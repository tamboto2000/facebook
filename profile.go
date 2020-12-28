package facebook

import (
	"errors"

	"github.com/tamboto2000/jsonextract/v2"
)

// Profile represent Facebook profile
type Profile struct {
	ID                    string `json:"id,omitempty"`
	Username              string `json:"username,omitempty"`
	Name                  string `json:"name,omitempty"`
	AlternateName         string `json:"alternateName,omitempty"`
	IsVerified            bool   `json:"isVerified,omitempty"`
	IsVisiblyMemorialized bool   `json:"is_visibly_memorialized,omitempty"`
	CoverPhoto            *Photo `json:"coverPhoto,omitempty"`
	ProfilePhoto          *Photo `json:"profilePhoto,omitempty"`
	Gender                string `json:"gender,omitempty"`
	BioText               string `json:"bioText,omitempty"`
	About                 *About `json:"about,omitempty"`

	profileSections  *jsonextract.JSON `json:"-"`
	variables        *jsonextract.JSON `json:"-"`
	aboutSectionVars *jsonextract.JSON `json:"-"`

	fb *Facebook
}

// Profile retrieve profile
func (fb *Facebook) Profile(user string) (*Profile, error) {
	body, err := fb.getRequest("/"+user, nil)
	if err != nil {
		return nil, err
	}

	jsons, err := jsonextract.FromBytesWithOpt(body, jsonextract.Option{
		ParseObj:         true,
		ParseArray:       true,
		IgnoreEmptyArray: true,
		IgnoreEmptyObj:   true,
	})
	if err != nil {
		return nil, err
	}

	prof := &Profile{fb: fb}
	if !composeProfile(jsons, prof) {
		return nil, errors.New("required data is missing")
	}

	return prof, nil
}

func composeProfile(jsons []*jsonextract.JSON, prof *Profile) bool {
	if !findObj(jsons, func(json *jsonextract.JSON) bool {
		val, ok := json.KeyVal["require"]
		if !ok {
			return false
		}

		// try to confirm that this object contains the profile preview data
		if !findObj(val.Vals, func(json *jsonextract.JSON) bool {
			val, ok := json.KeyVal["__dr"]
			if !ok {
				return false
			}

			if val.Kind == jsonextract.String && val.Val.(string) == "ProfileCometTabs_cometProfileTabs$normalization.graphql" {
				return true
			}

			return false
		}) {
			return false
		}

		if !findObj(val.Vals, func(json *jsonextract.JSON) bool {
			val, ok := json.KeyVal["__bbox"]
			if !ok {
				return false
			}

			// find profile variables (optional)
			if val, ok := val.KeyVal["variables"]; ok {
				prof.variables = val
			}

			val, ok = val.KeyVal["result"]
			if !ok {
				return false
			}

			val, ok = val.KeyVal["data"]
			if !ok {
				return false
			}

			val, ok = val.KeyVal["user"]
			if !ok {
				return false
			}

			// find user id
			if val, ok := val.KeyVal["id"]; ok {
				prof.ID = val.Val.(string)
			} else {
				return false
			}

			// find name
			if val, ok := val.KeyVal["name"]; ok {
				prof.Name = val.Val.(string)
			} else {
				return false
			}

			// find alternate name (optional)
			if val, ok := val.KeyVal["alternate_name"]; ok {
				prof.AlternateName = val.Val.(string)
			}

			// find cover photo (optional)
			if val, ok := val.KeyVal["cover_photo"]; ok {
				if photo, ok := val.KeyVal["photo"]; ok {
					if img, ok := photo.KeyVal["image"]; ok {
						prof.CoverPhoto = &Photo{
							ID:     photo.KeyVal["id"].Val.(string),
							URI:    img.KeyVal["uri"].Val.(string),
							Width:  img.KeyVal["width"].Val.(int),
							Height: img.KeyVal["height"].Val.(int),
							URL:    photo.KeyVal["url"].Val.(string),
						}
					}
				}
			}

			// find profile section and collection token
			if val, ok := val.KeyVal["profile_tabs"]; ok {
				if val, ok := val.KeyVal["profile_user"]; ok {
					if val, ok := val.KeyVal["timeline_nav_app_sections"]; ok {
						prof.profileSections = val
					}
				}
			}

			return true
		}) {
			return false
		}

		return true
	}) {
		return false
	}

	// find username (optional, username can be private)
	findObj(jsons, func(json *jsonextract.JSON) bool {
		val, ok := json.KeyVal["require"]
		if !ok {
			return false
		}

		if !findObj(val.Vals, func(json *jsonextract.JSON) bool {
			val, ok := json.KeyVal["rootView"]
			if !ok {
				return false
			}

			val, ok = val.KeyVal["props"]
			if !ok {
				return false
			}

			val, ok = val.KeyVal["userVanity"]
			if !ok {
				return false
			}

			prof.Username = val.Val.(string)

			return false
		}) {
			return false
		}

		return true
	})

	// find gender (optional)
	findObj(jsons, func(json *jsonextract.JSON) bool {
		val, ok := json.KeyVal["require"]
		if !ok {
			return false
		}

		if findObj(val.Vals, func(json *jsonextract.JSON) bool {
			val, ok := json.KeyVal["__bbox"]
			if !ok {
				return false
			}

			val, ok = val.KeyVal["result"]
			if !ok {
				return false
			}

			val, ok = val.KeyVal["data"]
			if !ok {
				return false
			}

			if val, ok = val.KeyVal["gender"]; ok {
				prof.Gender = val.Val.(string)
				return true
			}

			return false
		}) {
			return true
		}

		return false
	})

	// find profile photo (optional)
	findObj(jsons, func(json *jsonextract.JSON) bool {
		val, ok := json.KeyVal["require"]
		if !ok {
			return false
		}

		if findObj(val.Vals, func(json *jsonextract.JSON) bool {
			val, ok := json.KeyVal["__bbox"]
			if !ok {
				return false
			}

			val, ok = val.KeyVal["result"]
			if !ok {
				return false
			}

			val, ok = val.KeyVal["data"]
			if !ok {
				return false
			}

			profPhoto, ok := val.KeyVal["profilePhoto"]
			if !ok {
				return false
			}

			profPicNormal, ok := val.KeyVal["profilePicNormal"]
			if !ok {
				return false
			}

			prof.ProfilePhoto = &Photo{
				ID:     profPhoto.KeyVal["id"].Val.(string),
				URI:    profPicNormal.KeyVal["uri"].Val.(string),
				Width:  profPhoto.KeyVal["viewer_image"].KeyVal["width"].Val.(int),
				Height: profPhoto.KeyVal["viewer_image"].KeyVal["height"].Val.(int),
				URL:    profPhoto.KeyVal["url"].Val.(string),
			}

			return true
		}) {
			return true
		}

		return false
	})

	// find profile bio (optional)
	findObj(jsons, func(json *jsonextract.JSON) bool {
		val, ok := json.KeyVal["require"]
		if !ok {
			return false
		}

		if findObj(val.Vals, func(json *jsonextract.JSON) bool {
			val, ok := json.KeyVal["__bbox"]
			if !ok {
				return false
			}

			val, ok = val.KeyVal["result"]
			if !ok {
				return false
			}

			val, ok = val.KeyVal["data"]
			if !ok {
				return false
			}

			val, ok = val.KeyVal["profile_intro_card"]
			if !ok {
				return false
			}

			val, ok = val.KeyVal["bio"]
			if !ok {
				return false
			}

			prof.BioText = val.KeyVal["text"].Val.(string)

			return true
		}) {
			return true
		}

		return false
	})

	return true
}
