package facebook

import (
	"strings"

	"github.com/tamboto2000/jsonextract/v3"
)

// ContactAndBasicInfo contains profile's contact detail and basic info
type ContactAndBasicInfo struct {
	Phone       string       `json:"phone,omitempty"`
	Address     string       `json:"address,omitempty"`
	Websites    []string     `json:"websites,omitempty"`
	SocialLinks []SocialLink `json:"socialLinks,omitempty"`
	BasicInfo   *BasicInfo   `json:"basicInfo,omitempty"`
}

// SocialLink contains social media url and its type/social media platform name
type SocialLink struct {
	// can be an URL or screenname/username
	URLOrScreenname string `json:"urlOrScreenname"`
	Type            string `json:"type"`
}

// BasicInfo contains profile's about basic info, such as gender, birthday, sexual preferences, etc.
type BasicInfo struct {
	Gender         string `json:"gender,omitempty"`
	Birthday       string `json:"birthday,omitempty"`
	Languages      string `json:"languages,omitempty"`
	ReligiousViews string `json:"religiousViews,omitempty"`
	PoliticalViews string `json:"politicalViews,omitempty"`
	InterestedIn   string `json:"interestedIn,omitempty"`
}

// SyncContactAndBasicInfo retrieve profile's basic contact and info
func (about *About) SyncContactAndBasicInfo() error {
	jsons, err := about.profile.reqAboutCollection(aboutContactAndBasicInfo)
	if err != nil {
		return err
	}

	for _, json := range jsons {
		val, ok := json.Object()["label"]
		if !ok {
			continue
		}

		if val.String() == "ProfileCometAboutAppSectionQuery$defer$ProfileCometAboutAppSectionContent_appSection" {
			about.ContactAndBasicInfo = extractContactBasicInfo(json)
			break
		}
	}

	return nil
}

func extractContactBasicInfo(json *jsonextract.JSON) *ContactAndBasicInfo {
	contactBasic := new(ContactAndBasicInfo)
	if val, ok := json.Object()["data"]; ok {
		if val, ok := val.Object()["activeCollections"]; ok {
			if val, ok := val.Object()["nodes"]; ok {
				for _, node := range val.Array() {
					if val, ok := node.Object()["style_renderer"]; ok {
						if val, ok := val.Object()["profile_field_sections"]; ok {
							for _, section := range val.Array() {
								val, ok := section.Object()["field_section_type"]
								if !ok {
									continue
								}

								// extract phone number and address
								if val.String() == "about_contact_info" {
									if val, ok := section.Object()["profile_fields"]; ok {
										if val, ok := val.Object()["nodes"]; ok {
											for _, node := range val.Array() {
												val, ok := node.Object()["field_type"]
												if !ok {
													continue
												}

												// extract phone
												if val.String() == "other_phone" {
													if val, ok := node.Object()["title"]; ok {
														if val, ok := val.Object()["text"]; ok {
															contactBasic.Phone = val.String()
														}
													}
												}

												// extract address
												if val.String() == "address" {
													if val.String() == "other_phone" {
														if val, ok := node.Object()["title"]; ok {
															if val, ok := val.Object()["text"]; ok {
																contactBasic.Address = val.String()
															}
														}
													}
												}
											}
										}
									}
								}

								// extract websites and social links
								if val.String() == "websites_and_social_links" {
									if val, ok := section.Object()["profile_fields"]; ok {
										if val, ok := val.Object()["nodes"]; ok {
											for _, node := range val.Array() {
												val, ok := node.Object()["field_type"]
												if !ok {
													continue
												}

												// extract website
												if val.String() == "website" {
													if val, ok := node.Object()["title"]; ok {
														if val, ok := val.Object()["text"]; ok {
															contactBasic.Websites = append(contactBasic.Websites, val.String())
														}
													}
												}

												// extract social link
												if val.String() == "screenname" {
													socialLink := SocialLink{}
													if val, ok := node.Object()["title"]; ok {
														if val, ok := val.Object()["text"]; ok {
															socialLink.URLOrScreenname = val.String()
														}
													}

													// extract social media name
													if val, ok := node.Object()["list_item_groups"]; ok {
														for _, group := range val.Array() {
															if val, ok := group.Object()["list_items"]; ok {
																for _, item := range val.Array() {
																	if val, ok := item.Object()["text"]; ok {
																		socialLink.Type = val.Object()["text"].String()
																	}
																}
															}
														}
													}

													contactBasic.SocialLinks = append(contactBasic.SocialLinks, socialLink)
												}
											}
										}
									}
								}

								// extract basic infos
								if val.String() == "basic_info" {
									contactBasic.BasicInfo = new(BasicInfo)
									if val, ok := section.Object()["profile_fields"]; ok {
										if val, ok := val.Object()["nodes"]; ok {
											for _, node := range val.Array() {
												val, ok := node.Object()["field_type"]
												if !ok {
													continue
												}

												// extract gender
												if val.String() == "gender" {
													if val, ok := node.Object()["title"]; ok {
														if val, ok := val.Object()["text"]; ok {
															contactBasic.BasicInfo.Gender = strings.ToUpper(val.String())
														}
													}
												}

												// extract birthday
												if val.String() == "birthday" {
													if val, ok := node.Object()["title"]; ok {
														if val, ok := val.Object()["text"]; ok {
															if contactBasic.BasicInfo.Birthday != "" {
																contactBasic.BasicInfo.Birthday += ", "
															}

															contactBasic.BasicInfo.Birthday += val.String()
														}
													}
												}

												// extract languages
												if val.String() == "languages" {
													if val, ok := node.Object()["title"]; ok {
														if val, ok := val.Object()["text"]; ok {
															contactBasic.BasicInfo.Languages = val.String()
														}
													}
												}

												// extract relion views
												if val.String() == "religion" {
													if val, ok := node.Object()["title"]; ok {
														if val, ok := val.Object()["text"]; ok {
															contactBasic.BasicInfo.ReligiousViews = val.String()
														}
													}
												}

												// extract political views
												if val.String() == "politics" {
													if val, ok := node.Object()["title"]; ok {
														if val, ok := val.Object()["text"]; ok {
															contactBasic.BasicInfo.PoliticalViews = val.String()
														}
													}
												}

												// extract interested in
												if val.String() == "interested_in" {
													if val, ok := node.Object()["title"]; ok {
														if val, ok := val.Object()["text"]; ok {
															contactBasic.BasicInfo.InterestedIn = strings.ToUpper(val.String())
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
				}
			}
		}
	}

	return contactBasic
}
