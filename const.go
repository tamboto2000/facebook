package facebook

import "net/http"

const userAgent = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:73.0) Gecko/20100101 Firefox/73.0"

func useBaseHeader() http.Header {
	return http.Header{
		"User-Agent":                []string{userAgent},
		"Accept-Language":           []string{"en-US,en;q=0.5"},
		"Connection":                []string{"keep-alive"},
		"Upgrade-Insecure-Requests": []string{"1"},
		"Accept":                    []string{"*/*"},
		"Accept-Encoding":           []string{"gzip"},
	}
}
