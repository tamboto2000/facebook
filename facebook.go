package facebook

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/tamboto2000/jsonextract/v3"
)

const (
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:83.0) Gecko/20100101 Firefox/83.0"
)

// Profile section types
const (
	SectionAbout = "ABOUT"
)

// Facebook is Facebook client
type Facebook struct {
	cookies  *sync.Map
	client   *http.Client
	FbDtsg   string
	Jazoest  string
	SiteData *jsonextract.JSON

	host string
}

// New initialize Facebook client
func New() *Facebook {
	return &Facebook{
		cookies: new(sync.Map),
		client:  http.DefaultClient,
	}
}

// SetCookieStr set cookies from string
func (fb *Facebook) SetCookieStr(cookie string) {
	req := new(http.Request)
	req.Header = http.Header{}
	req.Header.Add("Cookie", cookie)
	cookies := req.Cookies()
	for _, c := range cookies {
		fb.cookies.Store(c.Name, c)
	}
}

// Init initialize Facebook client
func (fb *Facebook) Init() error {
	fb.host = "https://www.facebook.com"

RETRY:
	cUser, ok := fb.cookies.Load("c_user")
	if !ok {
		return ErrInvalidSession
	}

	userID := cUser.(*http.Cookie).Value

	_, body, err := fb.getRequest("/"+userID, nil)
	if err != nil {
		return err
	}

	jsons, err := jsonextract.FromBytes(body)
	if err != nil {
		return err
	}

	if !findObj(jsons, func(json *jsonextract.JSON) bool {
		obj := json.Object()
		val, ok := obj["require"]
		if !ok {
			return false
		}

		if findObj(val.Array(), func(json *jsonextract.JSON) bool {
			val, ok := json.Object()["__bbox"]
			if !ok {
				return false
			}

			val, ok = val.Object()["result"]
			if !ok {
				return false
			}

			val, ok = val.Object()["data"]
			if !ok {
				return false
			}

			val, ok = val.Object()["login_data"]
			if !ok {
				return false
			}

			if val, ok := val.Object()["lsd"]; ok {
				if val, ok := val.Object()["value"]; ok {
					fb.FbDtsg = val.String()
				} else {
					return false
				}
			} else {
				return false
			}

			if val, ok := val.Object()["jazoest"]; ok {
				if val, ok := val.Object()["value"]; ok {
					fb.Jazoest = val.String()
				} else {
					return false
				}
			} else {
				return false
			}

			return true
		}) {
			return true
		}

		return false
	}) {
		if fb.host == "https://www.facebook.com" {
			fb.host = "https://web.facebook.com"
			goto RETRY
		}

		return ErrInvalidSession
	}

	// find site data
	if !findObj(jsons, func(json *jsonextract.JSON) bool {
		val, ok := json.Object()["define"]
		if !ok {
			return false
		}

		if val.Kind() == jsonextract.Array {
			if findObj(val.Array(), func(json *jsonextract.JSON) bool {
				if _, ok := json.Object()["__spin_r"]; ok {
					fb.SiteData = json
					return true
				}

				return false
			}) {
				return true
			}
		}

		return false
	}) {
		return errors.New("SiteData tokens not found")
	}

	return nil
}

// merge cookies
func (fb *Facebook) mergeCookie(newC []*http.Cookie) {
	for _, c := range newC {
		_, ok := fb.cookies.Load(c.Name)
		if ok {
			fb.cookies.Delete(c.Name)
		}

		fb.cookies.Store(c.Name, c)
	}
}

func (fb *Facebook) getRequest(path string, query url.Values) (*http.Response, []byte, error) {
	header := http.Header{
		"User-Agent":                {userAgent},
		"Accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"},
		"Accept-Language":           {"en-US,en;q=0.5"},
		"Accept-Encoding":           {"gzip"},
		"Connection":                {"keep-alive"},
		"Upgrade-Insecure-Requests": {"1"},
	}

	urlParsed, err := url.Parse(fb.host + path)
	if err != nil {
		return nil, nil, err
	}

	if query != nil {
		urlParsed.RawQuery = query.Encode()
	}

	req, err := http.NewRequest("GET", urlParsed.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header = header
	fb.cookies.Range(func(key, val interface{}) bool {
		req.AddCookie(val.(*http.Cookie))
		return true
	})

	resp, err := fb.client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()

	buff, err := decompressResponseBody(resp)
	if err != nil {
		return nil, nil, err
	}

	fb.mergeCookie(resp.Cookies())

	return resp, buff.Bytes(), nil
}

func (fb *Facebook) graphQlRequest(body url.Values) ([]byte, error) {
	header := http.Header{}
	header.Set("Accept", "*/*")
	header.Set("Accept-Encoding", "gzip")
	header.Set("Accept-Language", "en-US,en;q=0.5")
	header.Set("Content-Type", "application/x-www-form-urlencoded")
	header.Set("Connection", "keep-alive")
	header.Set("Upgrade-Insecure-Requests", "1")

	siteData := fb.prepSiteData(fb.SiteData)
	for k, v := range body {
		siteData[k] = v
	}

	req, err := http.NewRequest("POST", fb.host+"/api/graphql/", strings.NewReader(siteData.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header = header
	fb.cookies.Range(func(key, val interface{}) bool {
		req.AddCookie(val.(*http.Cookie))
		return true
	})

	resp, err := fb.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	buff, err := decompressResponseBody(resp)
	if err != nil {
		return nil, err
	}

	fb.mergeCookie(resp.Cookies())

	return buff.Bytes(), nil
}
