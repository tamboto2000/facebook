package facebook

import (
	"bytes"
	"compress/gzip"
	"net/http"

	"github.com/tamboto2000/jsonextract/v3"
)

func findObj(jsons []*jsonextract.JSON, onFound func(json *jsonextract.JSON) bool) bool {
	for _, json := range jsons {
		if json.Kind() == jsonextract.Object {
			// DELETE
			// if strings.Contains(string(json.Raw.Bytes()), "fb_dtsg") {
			// 	fmt.Println("fb_dtsg found")
			// 	fmt.Println(string(json.Raw.Bytes()))
			// }

			if onFound(json) {
				return true
			}
		}

		if json.Kind() == jsonextract.Array {
			if findObj(json.Array(), onFound) {
				return true
			}
		}
	}

	return false
}

func findKeyObj(jsons []*jsonextract.JSON, key string, onFound func(parent *jsonextract.JSON, obj *jsonextract.JSON) bool) bool {
	for _, json := range jsons {
		if json.Kind() == jsonextract.Object {
			if val, ok := json.Object()[key]; ok {
				if val.Kind() == jsonextract.Object {
					if onFound(json, val) {
						return true
					}
				}
			}

			for _, val := range json.Object() {
				if findKeyObj([]*jsonextract.JSON{val}, key, onFound) {
					return true
				}
			}
		}

		if json.Kind() == jsonextract.Array {
			if findKeyObj(json.Array(), key, onFound) {
				return true
			}
		}
	}

	return false
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
