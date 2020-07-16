package facebook

import (
	"encoding/base64"
	"net/url"
	"strings"
)

type Contact struct {
	Phones     []string `json:"phones"`
	Emails     []string `json:"emails"`
	Websites   []string `json:"websites"`
	GitHubs    []string `json:"gitHubs"`
	Linkedins  []string `json:"linkedins"`
	Youtubes   []string `json:"youtubes"`
	Twitters   []string `json:"twitters"`
	Instagrams []string `json:"instagrams"`
	Addresses  []string `json:"addresses"`
}

func (user *User) SyncContact() error {
	user.Contact = new(Contact)
	return user.reqContactAndBasicInfo()
}

func (user *User) reqContactAndBasicInfo() error {
	sectionKey, err := base64.StdEncoding.DecodeString(user.userSections["ABOUT"].Node.ID)
	if err != nil {
		return err
	}

	sectionKeyRaw := strings.Replace(string(sectionKey), "app_section:", "", 1)
	var collectionToken string
	for _, node := range user.userSections["ABOUT"].Node.AllCollections.Nodes {
		if strings.Contains(node.URL, "/about_contact_and_basic_info") {
			collectionToken = node.ID
			break
		}
	}

	payloads, err := user.fb.doGraphQLRequest(map[string]interface{}{
		"appSectionFeedKey":      "ProfileCometAppSectionFeed_timeline_nav_app_sections__" + sectionKeyRaw,
		"collectionToken":        collectionToken,
		"rawSectionToken":        sectionKeyRaw,
		"scale":                  1,
		"sectionToken":           user.userSections["ABOUT"].Node.ID,
		"userID":                 user.ID,
		"useIncrementalDelivery": true,
	}, user.aboutDocID, "ProfileCometAboutAppSectionQuery", true)
	if err != nil {
		return err
	}

	for _, payload := range payloads {
		for _, a := range payload.Data.ActiveCollections.Nodes {
			for _, b := range a.StyleRenderer.ProfileFieldSections {
				for _, c := range b.ProfileFields.Nodes {
					//phone numbers
					if c.FieldType == "other_phone" {
						user.Contact.Phones = append(user.Contact.Phones, c.Title.Text)
						continue
					}

					//address
					if c.FieldType == "address" {
						user.Contact.Addresses = append(user.Contact.Addresses, c.Title.Text)
						continue
					}

					//websites and social medias
					if b.SectionType == "websites_and_social_links" && c.FieldType == "screenname" {
						if strings.Contains(c.LinkURL, "github.com") {
							parsed, _ := url.Parse(c.LinkURL)
							user.Contact.GitHubs = append(user.Contact.GitHubs, parsed.Query().Get("u"))
							continue
						}

						if strings.Contains(c.LinkURL, "linkedin.com") {
							parsed, _ := url.Parse(c.LinkURL)
							user.Contact.Linkedins = append(user.Contact.Linkedins, parsed.Query().Get("u"))
							continue
						}

						if strings.Contains(c.LinkURL, "youtube.com") {
							parsed, _ := url.Parse(c.LinkURL)
							user.Contact.Youtubes = append(user.Contact.Youtubes, parsed.Query().Get("u"))
							continue
						}

						if strings.Contains(c.LinkURL, "twitter.com") {
							parsed, _ := url.Parse(c.LinkURL)
							user.Contact.Twitters = append(user.Contact.Twitters, parsed.Query().Get("u"))
							continue
						}

						if strings.Contains(c.LinkURL, "instagram.com") {
							parsed, _ := url.Parse(c.LinkURL)
							user.Contact.Instagrams = append(user.Contact.Instagrams, parsed.Query().Get("u"))
							continue
						}
					}

					if c.FieldType == "website" {
						parsed, _ := url.Parse(c.LinkURL)
						user.Contact.Websites = append(user.Contact.Websites, parsed.Query().Get("u"))
						continue
					}

					//basic info
					//gender
					if c.FieldType == "gender" {
						user.Gender = c.Title.Text
						continue
					}

					if c.FieldType == "birthday" {
						if user.DOB != "" {
							user.DOB += ", " + c.Title.Text
						} else {
							user.DOB = c.Title.Text
						}

						continue
					}

					if c.FieldType == "languages" {
						splited := strings.Split(c.Title.Text, ", ")
						for i, lang := range splited {
							if i == len(splited)-1 {
								stg1 := strings.Split(lang, " and ")
								user.Languages = append(user.Languages, stg1[0])
								if len(stg1) > 1 {
									user.Languages = append(user.Languages, stg1[1])
								}
							} else {
								user.Languages = append(user.Languages, lang)
							}
						}

						continue
					}

					if c.FieldType == "religion" {
						user.Religion = c.Title.Text
						continue
					}

					if c.FieldType == "politics" {
						user.PoliticalView = c.Title.Text
						continue
					}

					if c.FieldType == "interested_in" {
						user.SexualPreference = c.Title.Text
						continue
					}
				}
			}
		}
	}

	return nil
}
