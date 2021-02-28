package facebook

import "github.com/tamboto2000/jsonextract/v3"

// NamePronunciation contains name pronunciation, include text and audio file URI
type NamePronunciation struct {
	Text     string `json:"text,omitempty"`
	AudioURI string `json:"audioURI,omitempty"`
}

// OtherName contains user name and its type
type OtherName struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

// Details contains profile's detail information, such as "About You", nicknames, name pronunciation, etc.
type Details struct {
	About             string             `json:"about,omitempty"`
	NamePronunciation *NamePronunciation `json:"namePronunciation,omitempty"`
	OtherNames        []OtherName        `json:"otherNames,omitempty"`
	FavoriteQuotes    string             `json:"favoriteQuotes,omitempty"`
}

// SyncDetails retrieve profile's details info
func (about *About) SyncDetails() error {
	jsons, err := about.profile.reqAboutCollection(aboutDetails)
	if err != nil {
		return err
	}

	for _, json := range jsons {
		val, ok := json.Object()["label"]
		if !ok {
			continue
		}

		if val.String() == "ProfileCometAboutAppSectionQuery$defer$ProfileCometAboutAppSectionContent_appSection" {
			details := extractDetails(json)
			about.Details = details

			break
		}
	}

	return nil
}

func extractDetails(json *jsonextract.JSON) *Details {
	details := new(Details)
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

								// extract about
								if val.String() == "about_me" {
									if val, ok := section.Object()["profile_fields"]; ok {
										if val, ok := val.Object()["nodes"]; ok {
											for _, node := range val.Array() {
												if val, ok := node.Object()["renderer"]; ok {
													if val, ok := val.Object()["field"]; ok {
														if val, ok := val.Object()["text_content"]; ok {
															if val, ok := val.Object()["text"]; ok {
																details.About = val.String()
															}
														}
													}
												}
											}
										}
									}
								}

								// extract name pronunciation
								if val.String() == "name_pronunciation" {
									if val, ok := section.Object()["profile_fields"]; ok {
										if val, ok := val.Object()["nodes"]; ok {
											for _, node := range val.Array() {
												if val, ok := node.Object()["renderer"]; ok {
													pronun := new(NamePronunciation)
													if val, ok := val.Object()["audio_uri"]; ok {
														pronun.AudioURI = val.String()
													}

													if val, ok := val.Object()["text_content"]; ok {
														if val, ok := val.Object()["text"]; ok {
															pronun.Text = val.String()
														}
													}

													details.NamePronunciation = pronun
												}
											}
										}
									}
								}

								// extract nicknames
								if val.String() == "nicknames" {
									if val, ok := section.Object()["profile_fields"]; ok {
										if val, ok := val.Object()["nodes"]; ok {
											for i, node := range val.Array() {
												if i == 0 {
													continue
												}

												nickname := OtherName{
													Name: node.Object()["title"].Object()["text"].String(),
												}

												if val, ok := node.Object()["list_item_groups"]; ok {
													for _, group := range val.Array() {
														if val, ok := group.Object()["list_items"]; ok {
															for _, item := range val.Array() {
																if val, ok := item.Object()["text"]; ok {
																	if val, ok := val.Object()["text"]; ok {
																		nickname.Type = val.String()
																	}
																}
															}
														}
													}
												}

												details.OtherNames = append(details.OtherNames, nickname)
											}
										}
									}
								}

								// extract quotes
								if val.String() == "favorite_quotes" {
									if val, ok := section.Object()["profile_fields"]; ok {
										if val, ok := val.Object()["nodes"]; ok {
											for _, node := range val.Array() {
												if val, ok := node.Object()["renderer"]; ok {
													if val, ok := val.Object()["field"]; ok {
														if val, ok := val.Object()["text_content"]; ok {
															if val, ok := val.Object()["text"]; ok {
																details.FavoriteQuotes = val.String()
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
	}

	return details
}
