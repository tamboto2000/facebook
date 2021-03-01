package facebook

import (
	jsonenc "encoding/json"
	"errors"

	"github.com/tamboto2000/jsonextract/v3"
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

	aboutSectionVars  *ProfileSection
	friendSectionVars *jsonextract.JSON

	fb *Facebook
}

// Profile retrieve profile
func (fb *Facebook) Profile(user string) (*Profile, error) {
	resp, body, err := fb.getRequest("/"+user, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 404 {
		return nil, ErrUserNotFound
	}

	jsons, err := jsonextract.FromBytes(body)
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
		val, ok := json.Object()["require"]
		if !ok {
			return false
		}

		// try to confirm that this object contains the profile preview data
		if !findObj(val.Array(), func(json *jsonextract.JSON) bool {
			val, ok := json.Object()["__dr"]
			if !ok {
				return false
			}

			if val.Kind() == jsonextract.String && val.String() == "ProfileCometTabs_cometProfileTabs$normalization.graphql" {
				return true
			}

			return false
		}) {
			return false
		}

		if !findObj(val.Array(), func(json *jsonextract.JSON) bool {
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

			val, ok = val.Object()["user"]
			if !ok {
				return false
			}

			// find user id
			if val, ok := val.Object()["id"]; ok {
				prof.ID = val.String()
			} else {
				return false
			}

			// find name
			if val, ok := val.Object()["name"]; ok {
				prof.Name = val.String()
			} else {
				return false
			}

			// find alternate name (optional)
			if val, ok := val.Object()["alternate_name"]; ok {
				prof.AlternateName = val.String()
			}

			// find cover photo (optional)
			if val, ok := val.Object()["cover_photo"]; ok {
				if photo, ok := val.Object()["photo"]; ok {
					if img, ok := photo.Object()["image"]; ok {
						prof.CoverPhoto = &Photo{
							ID:     photo.Object()["id"].String(),
							URI:    img.Object()["uri"].String(),
							Width:  int(img.Object()["width"].Integer()),
							Height: int(img.Object()["height"].Integer()),
							URL:    photo.Object()["url"].String(),
						}
					}
				}
			}

			// find profile section and collection token
			if val, ok := val.Object()["profile_tabs"]; ok {
				if val, ok := val.Object()["profile_user"]; ok {
					if val, ok := val.Object()["timeline_nav_app_sections"]; ok {
						// iterate sections
						edges, ok := val.Object()["edges"]
						if ok {
							for _, edge := range edges.Array() {
								if node, ok := edge.Object()["node"]; ok {
									if sectionType, ok := node.Object()["section_type"]; ok {
										if allColls, ok := node.Object()["all_collections"]; ok {
											if nodes, ok := allColls.Object()["nodes"]; ok {
												colls := make([]Collection, 0)
												jsonenc.Unmarshal(nodes.Bytes(), &colls)

												// get about section
												if sectionType.String() == "ABOUT" {
													prof.aboutSectionVars = newProfileSection()
													prof.aboutSectionVars.collections = colls
												}
											}
										}
									}
								}
							}
						}
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
		val, ok := json.Object()["require"]
		if !ok {
			return false
		}

		if !findObj(val.Array(), func(json *jsonextract.JSON) bool {
			val, ok := json.Object()["rootView"]
			if !ok {
				return false
			}

			val, ok = val.Object()["props"]
			if !ok {
				return false
			}

			val, ok = val.Object()["userVanity"]
			if !ok {
				return false
			}

			prof.Username = val.String()

			return false
		}) {
			return false
		}

		return true
	})

	// find gender (optional)
	findObj(jsons, func(json *jsonextract.JSON) bool {
		val, ok := json.Object()["require"]
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

			if val, ok = val.Object()["gender"]; ok {
				prof.Gender = val.String()
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
		val, ok := json.Object()["require"]
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

			profPhoto, ok := val.Object()["profilePhoto"]
			if !ok {
				return false
			}

			profPicNormal, ok := val.Object()["profilePicNormal"]
			if !ok {
				return false
			}

			prof.ProfilePhoto = &Photo{
				ID:     profPhoto.Object()["id"].String(),
				URI:    profPicNormal.Object()["uri"].String(),
				Width:  int(profPhoto.Object()["viewer_image"].Object()["width"].Integer()),
				Height: int(profPhoto.Object()["viewer_image"].Object()["height"].Integer()),
				URL:    profPhoto.Object()["url"].String(),
			}

			return true
		}) {
			return true
		}

		return false
	})

	// find profile bio (optional)
	findObj(jsons, func(json *jsonextract.JSON) bool {
		val, ok := json.Object()["require"]
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

			val, ok = val.Object()["profile_intro_card"]
			if !ok {
				return false
			}

			val, ok = val.Object()["bio"]
			if !ok {
				return false
			}

			prof.BioText = val.Object()["text"].String()

			return true
		}) {
			return true
		}

		return false
	})

	return true
}
