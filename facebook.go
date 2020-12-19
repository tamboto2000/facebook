package facebook

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/tamboto2000/facebook/raw"
	"github.com/tamboto2000/jsonextract"
)

const (
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:83.0) Gecko/20100101 Firefox/83.0"
)

// Profile section types
const (
	About = "ABOUT"
)

// Collection types
const (
	WorkAndEducation = "about_work_and_education"
)

var rootURL = "https://www.facebook.com"

// Config saves Facebook client settings
type Config struct {
	FbDtsg   string
	Jazoest  string
	SiteData *SiteData
	Cookies  []*http.Cookie
}

// GraphQLVars contains GraphQL variables for request Facebook GraphQL API
type GraphQLVars struct {
	AppSectionFeedKey string `json:"appSectionFeedKey"`
	CollectionToken   string `json:"collectionToken"`
	RawSectionToken   string `json:"rawSectionToken"`
	Scale             string `json:"scale"`
	SectionToken      string `json:"sectionToken"`
	UserID            string `json:"userID"`
}

// Facebook is Facebook client
type Facebook struct {
	cookies  *sync.Map
	client   *http.Client
	FbDtsg   string
	Jazoest  string
	SiteData *SiteData
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
	cUser, ok := fb.cookies.Load("c_user")
	if !ok {
		return new(ErrInvalidSession)
	}

	userID := cUser.(*http.Cookie).Value

	body, err := fb.getRequest("/"+userID, nil)
	if err != nil {
		return err
	}

	jsons, err := jsonextract.JSONFromBytes(body)
	if err != nil {
		return err
	}

	// extract SiteData
	parser := newParser(jsons)
	parser.run(func(val interface{}) bool {
		std := val.(*SiteData)
		if std.SpinB != "" && std.SpinR > 0 && std.SpinT > 0 {
			fb.SiteData = std
			return true
		}

		return false
	}, new(SiteData), true, false)

	// extract fb_dtsg and jazoest
	parser.reset()
	parser.run(func(val interface{}) bool {
		item := val.(*raw.Item)
		if item.Require != nil {
			data := make([][]byte, 0)
			for _, jr := range item.Require {
				data = append(data, jr)
			}

			parser := newParser(data)
			parser.run(func(val interface{}) bool {
				item := val.(*raw.Item)
				if item.Bbox != nil {
					if item.Bbox.Result != nil {
						if item.Bbox.Result.Data != nil {
							if item.Bbox.Result.Data.LoginData != nil {
								if item.Bbox.Result.Data.LoginData.Lsd != nil && item.Bbox.Result.Data.LoginData.Jazoest != nil {
									fb.Jazoest = item.Bbox.Result.Data.LoginData.Jazoest.Value
									fb.FbDtsg = item.Bbox.Result.Data.LoginData.Lsd.Value

									return true
								}
							}
						}
					}
				}

				return false
			}, new(raw.Item), true, false)
		}

		return false
	}, new(raw.Item), true, false)

	if fb.SiteData != nil && fb.FbDtsg != "" && fb.Jazoest != "" {
		return nil
	}

	return new(ErrInvalidSession)
}

// Save saves session to ./fb_session.json
func (fb *Facebook) Save() error {
	return fb.save("./fb_session.json")
}

// SaveToPath saves session to path
func (fb *Facebook) SaveToPath(path string) error {
	return fb.save(path)
}

// save fb session to path
func (fb *Facebook) save(path string) error {
	cs := make([]*http.Cookie, 0)
	fb.cookies.Range(func(key, value interface{}) bool {
		cs = append(cs, value.(*http.Cookie))
		return true
	})

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	conf := Config{
		FbDtsg:   fb.FbDtsg,
		Jazoest:  fb.Jazoest,
		SiteData: fb.SiteData,
		Cookies:  cs,
	}
	return json.NewEncoder(f).Encode(conf)
}

// Load load cookies from ./fb_session.json and set to Facebook client
func (fb *Facebook) Load() error {
	conf, err := load("./fb_session.json")
	if err != nil {
		return err
	}

	for _, c := range conf.Cookies {
		fb.cookies.Store(c.Name, c)
	}

	fb.FbDtsg = conf.FbDtsg
	fb.Jazoest = conf.Jazoest
	fb.SiteData = conf.SiteData

	return nil
}

// LoadFromPath load cookies from path and set to Facebook client
func (fb *Facebook) LoadFromPath(path string) error {
	conf, err := load(path)
	if err != nil {
		return err
	}

	for _, c := range conf.Cookies {
		fb.cookies.Store(c.Name, c)
	}

	fb.FbDtsg = conf.FbDtsg
	fb.Jazoest = conf.Jazoest
	fb.SiteData = conf.SiteData

	return nil
}

// load fb session from path
func load(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	conf := new(Config)
	if err := json.NewDecoder(f).Decode(conf); err != nil {
		return nil, err
	}

	return conf, nil
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

func (fb *Facebook) getRequest(path string, query url.Values) ([]byte, error) {
	header := http.Header{
		"User-Agent":                {userAgent},
		"Accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"},
		"Accept-Language":           {"en-US,en;q=0.5"},
		"Accept-Encoding":           {"gzip"},
		"Connection":                {"keep-alive"},
		"Upgrade-Insecure-Requests": {"1"},
	}

	urlParsed, err := url.Parse(rootURL + path)
	if err != nil {
		return nil, err
	}

	if query != nil {
		urlParsed.RawQuery = query.Encode()
	}

	req, err := http.NewRequest("GET", urlParsed.String(), nil)
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

	req, err := http.NewRequest("POST", rootURL+"/api/graphql/", strings.NewReader(siteData.Encode()))
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

	return buff.Bytes(), nil
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
