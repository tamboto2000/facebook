package facebook

import (
	"strings"

	"github.com/tamboto2000/jsonextract/v2"
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
func (prof *Profile) SyncContactAndBasicInfo() error {
	jsons, err := prof.reqAboutCollection(aboutContactAndBasicInfo)
	if err != nil {
		return err
	}

	// DELETE
	// jsonextract.SaveToPath(jsons, "raw_contact_basic_info.json")

	for _, json := range jsons {
		val, ok := json.KeyVal["label"]
		if !ok {
			continue
		}

		if val.Val.(string) == "ProfileCometAboutAppSectionQuery$defer$ProfileCometAboutAppSectionContent_appSection" {
			prof.About.ContactAndBasicInfo = extractContactBasicInfo(json)
			break
		}
	}

	return nil
}

func extractContactBasicInfo(json *jsonextract.JSON) *ContactAndBasicInfo {
	contactBasic := new(ContactAndBasicInfo)
	if val, ok := json.KeyVal["data"]; ok {
		if val, ok := val.KeyVal["activeCollections"]; ok {
			if val, ok := val.KeyVal["nodes"]; ok {
				for _, node := range val.Vals {
					if val, ok := node.KeyVal["style_renderer"]; ok {
						if val, ok := val.KeyVal["profile_field_sections"]; ok {
							for _, section := range val.Vals {
								val, ok := section.KeyVal["field_section_type"]
								if !ok {
									continue
								}

								// extract phone number and address
								if val.Val.(string) == "about_contact_info" {
									if val, ok := section.KeyVal["profile_fields"]; ok {
										if val, ok := val.KeyVal["nodes"]; ok {
											for _, node := range val.Vals {
												val, ok := node.KeyVal["field_type"]
												if !ok {
													continue
												}

												// extract phone
												if val.Val.(string) == "other_phone" {
													if val, ok := node.KeyVal["title"]; ok {
														if val, ok := val.KeyVal["text"]; ok {
															contactBasic.Phone = val.Val.(string)
														}
													}
												}

												// extract address
												if val.Val.(string) == "address" {
													if val.Val.(string) == "other_phone" {
														if val, ok := node.KeyVal["title"]; ok {
															if val, ok := val.KeyVal["text"]; ok {
																contactBasic.Address = val.Val.(string)
															}
														}
													}
												}
											}
										}
									}
								}

								// extract websites and social links
								if val.Val.(string) == "websites_and_social_links" {
									if val, ok := section.KeyVal["profile_fields"]; ok {
										if val, ok := val.KeyVal["nodes"]; ok {
											for _, node := range val.Vals {
												val, ok := node.KeyVal["field_type"]
												if !ok {
													continue
												}

												// extract website
												if val.Val.(string) == "website" {
													if val, ok := node.KeyVal["title"]; ok {
														if val, ok := val.KeyVal["text"]; ok {
															contactBasic.Websites = append(contactBasic.Websites, val.Val.(string))
														}
													}
												}

												// extract social link
												if val.Val.(string) == "screenname" {
													socialLink := SocialLink{}
													if val, ok := node.KeyVal["title"]; ok {
														if val, ok := val.KeyVal["text"]; ok {
															socialLink.URLOrScreenname = val.Val.(string)
														}
													}

													// extract social media name
													if val, ok := node.KeyVal["list_item_groups"]; ok {
														for _, group := range val.Vals {
															if val, ok := group.KeyVal["list_items"]; ok {
																for _, item := range val.Vals {
																	if val, ok := item.KeyVal["text"]; ok {
																		socialLink.Type = val.KeyVal["text"].Val.(string)
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
								if val.Val.(string) == "basic_info" {
									contactBasic.BasicInfo = new(BasicInfo)
									if val, ok := section.KeyVal["profile_fields"]; ok {
										if val, ok := val.KeyVal["nodes"]; ok {
											for _, node := range val.Vals {
												val, ok := node.KeyVal["field_type"]
												if !ok {
													continue
												}

												// extract gender
												if val.Val.(string) == "gender" {
													if val, ok := node.KeyVal["title"]; ok {
														if val, ok := val.KeyVal["text"]; ok {
															contactBasic.BasicInfo.Gender = strings.ToUpper(val.Val.(string))
														}
													}
												}

												// extract birthday
												if val.Val.(string) == "birthday" {
													if val, ok := node.KeyVal["title"]; ok {
														if val, ok := val.KeyVal["text"]; ok {
															if contactBasic.BasicInfo.Birthday != "" {
																contactBasic.BasicInfo.Birthday += ", "
															}

															contactBasic.BasicInfo.Birthday += val.Val.(string)
														}
													}
												}

												// extract languages
												if val.Val.(string) == "languages" {
													if val, ok := node.KeyVal["title"]; ok {
														if val, ok := val.KeyVal["text"]; ok {
															contactBasic.BasicInfo.Languages = val.Val.(string)
														}
													}
												}

												// extract relion views
												if val.Val.(string) == "religion" {
													if val, ok := node.KeyVal["title"]; ok {
														if val, ok := val.KeyVal["text"]; ok {
															contactBasic.BasicInfo.ReligiousViews = val.Val.(string)
														}
													}
												}

												// extract political views
												if val.Val.(string) == "politics" {
													if val, ok := node.KeyVal["title"]; ok {
														if val, ok := val.KeyVal["text"]; ok {
															contactBasic.BasicInfo.PoliticalViews = val.Val.(string)
														}
													}
												}

												// extract interested in
												if val.Val.(string) == "interested_in" {
													if val, ok := node.KeyVal["title"]; ok {
														if val, ok := val.KeyVal["text"]; ok {
															contactBasic.BasicInfo.InterestedIn = strings.ToUpper(val.Val.(string))
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
