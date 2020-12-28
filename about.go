package facebook

import (
	"errors"
	"strings"

	"github.com/tamboto2000/jsonextract/v2"
)

// About Section's collections
const (
	WorkAndEducation = "about_work_and_education"
)

// About contains profile about section data
type About struct {
	WorkHistory []Work `json:"workHistory,omitempty"`
}

// SyncAbout fetch required tokens for requesting profile about data collections
func (prof *Profile) SyncAbout() error {
	rawBody, err := prof.fb.getRequest("/"+prof.ID+"/about", nil)
	if err != nil {
		return err
	}

	jsons, err := jsonextract.FromBytesWithOpt(rawBody, jsonextract.Option{
		ParseObj:         true,
		ParseArray:       true,
		IgnoreEmptyObj:   true,
		IgnoreEmptyArray: true,
	})

	if err != nil {
		return err
	}

	// DELETE
	// jsonextract.SaveToPath(jsons, "raw_about.json")

	// find profile about section vars
	if !findObj(jsons, func(json *jsonextract.JSON) bool {
		val, ok := json.KeyVal["require"]
		if !ok {
			return false
		}

		if findObj(val.Vals, func(json *jsonextract.JSON) bool {
			val, ok := json.KeyVal["preloaderID"]
			if !ok {
				return false
			}

			if strings.Contains(val.Val.(string), "adp_ProfileCometAboutAppSectionQueryRelayPreloader") {
				if _, ok = json.KeyVal["variables"]; ok {
					prof.aboutSectionVars = json
					return true
				}
			}

			return false
		}) {
			return true
		}

		return false
	}) {
		return errors.New("Important tokens for About section is not found")
	}

	return nil
}
