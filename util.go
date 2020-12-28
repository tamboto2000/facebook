package facebook

import (
	"github.com/tamboto2000/jsonextract/v2"
)

func findObj(jsons []*jsonextract.JSON, onFound func(json *jsonextract.JSON) bool) bool {
	for _, json := range jsons {
		if json.Kind == jsonextract.Object {
			// DELETE
			// if strings.Contains(string(json.Raw.Bytes()), "fb_dtsg") {
			// 	fmt.Println("fb_dtsg found")
			// 	fmt.Println(string(json.Raw.Bytes()))
			// }

			if onFound(json) {
				return true
			}
		}

		if json.Kind == jsonextract.Array {
			if findObj(json.Vals, onFound) {
				return true
			}
		}
	}

	return false
}
