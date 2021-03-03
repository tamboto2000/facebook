package facebook

import (
	"bytes"
	"compress/gzip"
	"net/http"
	"regexp"
	"time"

	"github.com/tamboto2000/jsonextract/v3"
)

func findObj(jsons []*jsonextract.JSON, onFound func(json *jsonextract.JSON) bool) bool {
	for _, json := range jsons {
		if json.Kind() == jsonextract.Object {
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

func extractDate(text string) (dates []time.Time, datesStr []string) {
	dPattern1 := `[a-zA-Z]+ \d{1,2}, \d{4}`
	dPattern2 := `[a-zA-Z]+ \d{4}`
	dPattern3 := `\d{4}`

	// try to get dates from dPattern1
	rgx, _ := regexp.Compile(dPattern1)
	if matchs := rgx.FindAllString(text, -1); len(matchs) > 0 {
		for _, match := range matchs {
			if date, err := time.Parse("January 02, 2006", match); err == nil {
				dates = append(dates, date)
				datesStr = append(datesStr, match)
			} else if date, err := time.Parse("January 2, 2006", match); err == nil {
				dates = append(dates, date)
				datesStr = append(datesStr, match)
			}
		}
	}

	// try to get dates from dPattern2
	if len(dates) < 2 && len(datesStr) < 2 {
		rgx, _ = regexp.Compile(dPattern2)
		if matchs := rgx.FindAllString(text, -1); len(matchs) > 0 {
			for _, match := range matchs {
				if date, err := time.Parse("January 2006", match); err == nil {
					dates = append(dates, date)
					datesStr = append(datesStr, match)
				}
			}
		}
	}

	// try to get dates from dPattern3
	if len(dates) < 2 && len(datesStr) < 2 {
		rgx, _ = regexp.Compile(dPattern3)
		if matchs := rgx.FindAllString(text, -1); len(matchs) > 0 {
			if len(dates) == 1 {
				if len(matchs) == 2 {
					match := matchs[len(matchs)-1]
					if date, err := time.Parse("2006", match); err == nil {
						dates = append(dates, date)
						datesStr = append(datesStr, match)
					}
				}
			}

			if len(dates) == 0 {
				for _, match := range matchs {
					if date, err := time.Parse("2006", match); err == nil {
						dates = append(dates, date)
						datesStr = append(datesStr, match)
					}
				}
			}

			// if len(dates) > 1 {
			// 	match := matchs[len(matchs)-1]
			// 	if date, err := time.Parse("2006", match); err == nil {
			// 		dates = append(dates, date)
			// 		datesStr = append(datesStr, match)
			// 	}
			// } else {
			// 	for _, match := range matchs {
			// 		if date, err := time.Parse("2006", match); err == nil {
			// 			dates = append(dates, date)
			// 			datesStr = append(datesStr, match)
			// 		}
			// 	}
			// }
		}
	}

	return dates, datesStr
}
