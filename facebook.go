package facebook

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/tamboto2000/facebook/raw"
	"github.com/tamboto2000/htmltojson"
	"github.com/tamboto2000/random"
)

type injectable interface {
	inject(fb *Facebook)
}

type Facebook struct {
	Cookies                         []*http.Cookie `json:"cookies"`
	Tokens                          url.Values     `json:"tokens"`
	RootURL                         string         `json:"rootUrl"`
	CommentQueryID                  string         `json:"commentQueryId"`
	TimelineFeedQueryRelayPreloader string         `json:"timelineFeedQueryRelayPreloader"`
	TimelineFeedRefetchQuery        string         `json:"timelineFeedRefetchQuery"`
	ReactionsDialogQuery            string         `json:"reactionsDialogQuery"`
	CometPhotoRootQuery             string         `json:"cometPhotoRootQuery"`
	CometVideoRootMediaViewerQuery  string         `json:"cometVideoRootMediaViewerQuery"`

	mutex *sync.Mutex
}

func New() *Facebook {
	fb := &Facebook{
		mutex:  new(sync.Mutex),
		Tokens: make(url.Values),
	}
	return fb
}

func Import(i *Facebook) *Facebook {
	i.mutex = new(sync.Mutex)
	return i
}

func (fb *Facebook) Export() *Facebook {
	return &Facebook{
		Cookies:                         fb.Cookies,
		Tokens:                          fb.Tokens,
		RootURL:                         fb.RootURL,
		CommentQueryID:                  fb.CommentQueryID,
		TimelineFeedQueryRelayPreloader: fb.TimelineFeedQueryRelayPreloader,
		TimelineFeedRefetchQuery:        fb.TimelineFeedRefetchQuery,
		ReactionsDialogQuery:            fb.ReactionsDialogQuery,
		mutex:                           new(sync.Mutex),
	}
}

func (fb *Facebook) Inject(in injectable) {
	in.inject(fb)
}

func (fb *Facebook) SetCookie(cookie string) error {
	rawRequest := fmt.Sprintf("GET / HTTP/1.0\r\nCookie: %s\r\n\r\n", cookie)
	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(rawRequest)))
	if err != nil {
		return err
	}

	fb.Cookies = req.Cookies()
	return nil
}

func (fb *Facebook) Init() error {
	//try with www.facebook.com first
	rootURL := "https://www.facebook.com"

RETRY:
	uri, _ := url.Parse(rootURL + "/")
	resp, err := fb.doBasicGetRequest(uri)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	buff, err := decompressResponseBody(resp)
	body := buff.Bytes()
	resp.Body.Close()

	jsonStrs, err := extractJSONBytes(body)
	if err != nil {
		return err
	}

	for _, str := range jsonStrs {
		userIDInfo := new(raw.UserIDInfo)
		if err := json.Unmarshal(str, userIDInfo); err != nil {
			continue
		}

		if userIDInfo.UserID != "" {
			fb.Tokens.Set("__user", userIDInfo.UserID)
			fb.Tokens.Set("av", userIDInfo.UserID)
			break
		}
	}

	for _, str := range jsonStrs {
		token := new(raw.FBToken1)
		if err := json.Unmarshal(str, token); err != nil {
			continue
		}

		if token.SpinB != "" && token.SpinR != 0 && token.SpinT != 0 {
			fb.Tokens.Set("__rev", strconv.Itoa(token.ServerRevision))
			fb.Tokens.Set("__hsi", token.Hsi)
			fb.Tokens.Set("__spin_r", strconv.Itoa(token.SpinR))
			fb.Tokens.Set("__spin_b", token.SpinB)
			fb.Tokens.Set("__spin_t", strconv.Itoa(token.SpinT))
			break
		}
	}

	for _, str := range jsonStrs {
		token := new(raw.FBToken2)
		if err := json.Unmarshal(str, token); err != nil {
			continue
		}

		ajaxTokensFound := false
		for _, data1 := range token.Require {
			for _, data2 := range data1 {
				rawData := make([]json.RawMessage, 0)
				if err := json.Unmarshal(data2, &rawData); err != nil {
					continue
				}

				for _, data3 := range rawData {
					bbox := new(raw.RequireBbox)
					if err := json.Unmarshal(data3, bbox); err != nil {
						continue
					}

					loginData := bbox.Bbox.Result.Data.LoginData
					if loginData.Lsd.Value != "" && loginData.Jazoest.Value != "" {
						fb.Tokens.Set("fb_dtsg", loginData.Lsd.Value)
						fb.Tokens.Set("jazoest", loginData.Jazoest.Value)

						ajaxTokensFound = true
						break
					}
				}

				if ajaxTokensFound {
					break
				}
			}

			if ajaxTokensFound {
				break
			}
		}

		if ajaxTokensFound {
			break
		}
	}

	if fb.Tokens.Get("fb_dtsg") == "" || fb.Tokens.Get("__user") == "" {
		if rootURL == "https://web.facebook.com" {
			// DELETE
			f, _ := os.Create("index.html")
			f.Write(body)
			return nil
			return errors.New("invalid session cookie")
		}

		rootURL = "https://web.facebook.com"
		goto RETRY
	}

	fb.RootURL = rootURL

	fb.Tokens.Set("__a", "1")
	fb.Tokens.Set("__beoa", "1")
	fb.Tokens.Set("__pc", "EXP2:comet_pkg")
	fb.Tokens.Set("dpr", "1")
	fb.Tokens.Set("fb_api_caller_class", "RelayModern")

	//extract comment query id
	r := bytes.NewReader(body)
	rootNode, err := htmltojson.ParseFromReader(r)
	fb.extractCommentQueryID(rootNode)

	//extract timeline query ids
	uri, _ = url.Parse(fb.RootURL + "/" + fb.Tokens.Get("__user"))
	resp, err = fb.doBasicGetRequest(uri)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	buff, err = decompressResponseBody(resp)
	body = buff.Bytes()

	fb.extractTimelineFeedQueryID(body)
	fb.extractTimelineFeedRefetchQueryAndReactionsDialogQuery(body)

	//variables that needs to set dynamically
	//__req
	//fb_api_req_friendly_name
	//variables
	//doc_id

	return nil
}

func (fb *Facebook) doBasicGetRequest(uri *url.URL) (*http.Response, error) {
	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return nil, err
	}

	cl := new(http.Client)
	req.Header = useBaseHeader()
	fb.mutex.Lock()
	for _, c := range fb.Cookies {
		req.AddCookie(c)
	}

	fb.mutex.Unlock()
	resp, err := cl.Do(req)

	return resp, err
}

func (fb *Facebook) doGetRequest(uri *url.URL) (*http.Response, error) {
	fb.mutex.Lock()
	uri.RawQuery += "&" + fb.Tokens.Encode()
	fb.mutex.Unlock()

	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return nil, err
	}

	cl := new(http.Client)
	req.Header = useBaseHeader()
	fb.mutex.Lock()
	for _, c := range fb.Cookies {
		req.AddCookie(c)
	}

	fb.mutex.Unlock()
	resp, err := cl.Do(req)

	return resp, err
}

func (fb *Facebook) doGraphQLRequest(vr map[string]interface{}, docID, apiName string, isList bool) ([]raw.GraphQLData, error) {
	fb.mutex.Lock()
	urlParsed, _ := url.Parse(fb.RootURL + "/api/graphql/")
	q := fb.Tokens
	q.Add("fb_api_req_friendly_name", apiName)
	q.Add("__comet_req", "1")
	vrs, _ := json.Marshal(vr)
	q.Add("variables", string(vrs))
	q.Add("__req", random.RandStr(2))
	if docID != "" {
		q.Add("doc_id", docID)
	}

	req, err := http.NewRequest("POST", urlParsed.String(), strings.NewReader(q.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header = useBaseHeader()
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	cl := new(http.Client)
	for _, c := range fb.Cookies {
		req.AddCookie(c)
	}

	fb.mutex.Unlock()

	resp, err := cl.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	buff, err := decompressResponseBody(resp)
	if err != nil {
		return nil, err
	}

	payloads := make([]raw.GraphQLData, 0)
	if isList {
		jsons, err := extractJSONFromReader(buff)
		if err != nil {
			return nil, err
		}
		for _, str := range jsons {
			payload := new(raw.GraphQLData)
			if err := json.Unmarshal(str, payload); err == nil {
				// payload.Raw = str
				payloads = append(payloads, *payload)
			}
		}
	} else {
		payload := new(raw.GraphQLData)
		// body := buff.Bytes()
		if err := json.Unmarshal(buff.Bytes(), payload); err == nil {
			// payload.Raw = body
			payloads = append(payloads, *payload)
		} else {
			return nil, err
		}
	}

	return payloads, err
}

func (fb *Facebook) extractCommentQueryID(rootNode *htmltojson.Node) {
	nodes := htmltojson.SearchAllNode(
		"",
		"script",
		"",
		"",
		"",
		rootNode,
	)

	counter := 0
	wg := new(sync.WaitGroup)
	mutex := new(sync.Mutex)
	done := make(chan bool)
	for _, node := range nodes {
		for _, attr := range node.Attr {
			if attr.Key == "src" {
				wg.Add(1)
				counter++
				go func(link string) {
					uri, _ := url.Parse(link)
					resp, err := fb.doBasicGetRequest(uri)
					if err != nil {
						wg.Done()
						return
					}

					buff, err := decompressResponseBody(resp)
					defer resp.Body.Close()

					mutex.Lock()
					if fb.CommentQueryID != "" {
						mutex.Unlock()
						wg.Done()
						return
					}

					mutex.Unlock()

					body := buff.String()
					substr := `__d("CometUFICommentsProviderPaginationQuery.graphql"`
					if !strings.Contains(body, substr) {
						wg.Done()
						return
					}

					stg1 := strings.Split(body, substr)
					stg1 = strings.Split(stg1[1], `{id:"`)
					stg1 = strings.Split(stg1[1], `"`)
					mutex.Lock()
					fb.CommentQueryID = stg1[0]
					mutex.Unlock()
					wg.Done()
					done <- true
				}(attr.Val)

				break
			}
		}

		if counter == 10 {
			go func() {
				wg.Wait()
				if fb.CommentQueryID != "" {
					return
				}

				done <- true
			}()
			<-done
			if fb.CommentQueryID != "" {
				close(done)
				return
			}
		}
	}
}

func (fb *Facebook) extractTimelineFeedQueryID(body []byte) error {
	jsons, err := extractJSONBytes(body)
	if err != nil {
		return err
	}

	for _, rawData := range jsons {
		payload := new(raw.QueryID)
		if err := json.Unmarshal(rawData, payload); err != nil {
			continue
		}

		if !strings.Contains(payload.PreloaderID, "adp_ProfileCometTimelineFeedQueryRelayPreloader") {
			continue
		}

		fb.TimelineFeedQueryRelayPreloader = payload.QueryID
		return nil
	}

	return nil
}

func (fb *Facebook) extractTimelineFeedRefetchQueryAndReactionsDialogQuery(body []byte) error {
	rootNode, err := htmltojson.ParseString(string(body))
	if err != nil {
		return err
	}

	nodes := htmltojson.SearchAllNode(
		"",
		"script",
		"",
		"",
		"",
		rootNode,
	)

	wg := new(sync.WaitGroup)
	mx := new(sync.Mutex)
	done := make(chan bool)
	counter := 0
	isChanClosed := false
	for i, node := range nodes {
		for _, attr := range node.Attr {
			if attr.Key == "src" {
				wg.Add(1)
				link := attr.Val
				counter++
				go func() {
					uri, _ := url.Parse(link)
					resp, err := fb.doBasicGetRequest(uri)
					if err != nil {
						wg.Done()
						return
					}

					defer resp.Body.Close()

					mx.Lock()
					if fb.TimelineFeedRefetchQuery != "" {
						mx.Unlock()
						wg.Done()
						return
					}
					mx.Unlock()

					buffer, err := decompressResponseBody(resp)
					body := buffer.String()

					mx.Lock()
					if fb.TimelineFeedRefetchQuery == "" {
						mx.Unlock()
						substr := `__d("ProfileCometTimelineFeedRefetchQuery.graphql"`
						if strings.Contains(body, substr) {
							stg1 := strings.Split(body, substr)
							stg1 = strings.Split(stg1[1], `{id:"`)
							stg1 = strings.Split(stg1[1], `"`)
							mx.Lock()
							fb.TimelineFeedRefetchQuery = stg1[0]
							mx.Unlock()
						}
					} else {
						mx.Unlock()
					}

					mx.Lock()
					if fb.ReactionsDialogQuery == "" {
						mx.Unlock()
						substr := `__d("CometUFIReactionsDialogQuery$Parameters"`
						if strings.Contains(body, substr) {
							stg1 := strings.Split(body, substr)
							stg1 = strings.Split(stg1[1], `{id:"`)
							stg1 = strings.Split(stg1[1], `"`)
							mx.Lock()
							fb.ReactionsDialogQuery = stg1[0]
							mx.Unlock()
						}
					} else {
						mx.Unlock()
					}

					mx.Lock()
					if fb.TimelineFeedRefetchQuery != "" && fb.ReactionsDialogQuery != "" {
						if !isChanClosed {
							isChanClosed = true
							done <- true
						}
					}

					mx.Unlock()
					wg.Done()
				}()

				break
			}
		}

		if counter == 10 || i == len(nodes)-1 {
			counter = 0
			go func() {
				wg.Wait()
				if fb.TimelineFeedRefetchQuery == "" {
					done <- true
				}
			}()

			<-done
			if fb.TimelineFeedRefetchQuery != "" {
				close(done)
				return nil
			}
		}
	}

	return nil
}

func parseCookieStr(str string) []*http.Cookie {
	rawRequest := fmt.Sprintf("GET / HTTP/1.0\r\nCookie: %s\r\n\r\n", str)
	req, _ := http.ReadRequest(bufio.NewReader(strings.NewReader(rawRequest)))

	return req.Cookies()
}

func decompressResponseBody(resp *http.Response) (*bytes.Buffer, error) {
	buff := new(bytes.Buffer)
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}

		buff.ReadFrom(reader)
	default:
		buff.ReadFrom(resp.Body)
	}

	return buff, nil
}
