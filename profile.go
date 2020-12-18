package facebook

import (
	"github.com/tamboto2000/facebook/raw"
	"github.com/tamboto2000/jsonextract"
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

	ProfileSections *raw.TimelineNavAppSections `json:"profileSections,omitempty"`

	fb *Facebook
}

// Profile retrieve profile
func (fb *Facebook) Profile(user string) (*Profile, error) {
	profile := new(Profile)
	body, err := fb.getRequest("/"+user, nil)
	if err != nil {
		return nil, err
	}

	jsons, err := jsonextract.JSONFromBytes(body)
	if err != nil {
		return nil, err
	}

	// DELETE
	// jsonextract.Save(jsons)

	parser := newParser(jsons)
	parser.run(func(val interface{}) bool {
		item := val.(*raw.Item)
		if item.RootView != nil {
			if item.RootView.Props != nil {
				props := item.RootView.Props
				profile.ID = props.UserID
				profile.Username = props.UserVanity

				return true
			}
		}

		return false
	}, new(raw.Item), true, false)

	parser.reset()

	// Extract profile preview
	// Get profile name first
	parser.run(func(val interface{}) bool {
		item := val.(*raw.Item)
		composeProfile(item, profile)
		if profile.Name != "" {
			return true
		}

		return false
	}, new(raw.Item), true, false)

	// Reset parser and search for other data
	parser.reset()
	parser.run(func(val interface{}) bool {
		item := val.(*raw.Item)
		composeProfile(item, profile)

		return false
	}, new(raw.Item), true, false)

	return profile, nil
}

func composeProfile(item *raw.Item, prof *Profile) {
	if item.Require != nil {
		data := make([][]byte, 0)
		for _, d := range item.Require {
			data = append(data, d)
		}

		parser := newParser(data)
		parser.run(func(val interface{}) bool {
			item := val.(*raw.Item)
			composeProfile(item, prof)

			return false
		}, new(raw.Item), true, false)

		return
	}

	if item.Bbox != nil {
		if item.Bbox.Result != nil {
			if item.Bbox.Result.Data != nil {
				data := item.Bbox.Result.Data
				if prof.Name == "" && data.User != nil {
					if data.User.ID == prof.ID {
						// Extract name
						user := item.Bbox.Result.Data.User
						prof.Name = user.Name

						// Extract Alternate Name
						prof.AlternateName = user.AlternateName

						// Extract cover photo
						if user.CoverPhoto != nil {
							if user.CoverPhoto.Photo != nil {
								if user.CoverPhoto.Photo.Image != nil {
									photo := user.CoverPhoto.Photo
									prof.CoverPhoto = &Photo{
										ID:     photo.ID,
										URI:    photo.Image.URI,
										Width:  photo.Image.Width,
										Height: photo.Image.Height,
										URL:    photo.URL,
									}
								}
							}
						}

						if user.ProfileTabs != nil {
							if user.ProfileTabs.ProfileUser != nil {
								if user.ProfileTabs.ProfileUser.TimelineNavAppSections != nil {
									prof.ProfileSections = user.ProfileTabs.ProfileUser.TimelineNavAppSections
								}
							}
						}
					}

					return
				}

				// Extract gender
				if data.Gender != "" && prof.Gender == "" {
					prof.Gender = data.Gender
				}

				// Extract Profile Photo
				if data.ProfilePhoto != nil && data.ProfilePicNormal != nil && prof.ProfilePhoto == nil {
					prof.ProfilePhoto = &Photo{
						URL:    data.ProfilePhoto.URL,
						ID:     data.ProfilePhoto.ID,
						Height: data.ProfilePhoto.ViewerImage.Height,
						Width:  data.ProfilePhoto.ViewerImage.Width,
						URI:    data.ProfilePicNormal.URI,
					}
				}

				// Extract bio
				if data.ProfileIntroCard != nil && prof.BioText == "" {
					if data.ProfileIntroCard.Bio != nil {
						if data.ProfileIntroCard.Bio.Text != "" {
							prof.BioText = data.ProfileIntroCard.Bio.Text
						}
					}
				}
			}
		}
	}
}
