package facebook

import "github.com/tamboto2000/jsonextract/v2"

// NamePronunciation contains name pronunciation, include text and audio file URI
type NamePronunciation struct {
	Text     string `json:"text,omitempty"`
	AudioURI string `json:"audioURI,omitempty"`
}

// Nickname contains nickname and its type
type Nickname struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

// Details contains profile's detail information, such as "About You", nicknames, name pronunciation, etc.
type Details struct {
	About             string             `json:"about,omitempty"`
	NamePronunciation *NamePronunciation `json:"namePronunciation,omitempty"`
	Nicknames         []Nickname         `json:"nicknames,omitempty"`
	FavoriteQuotes    string             `json:"favoriteQuotes,omitempty"`
}

// SyncDetails retrieve profile's details info
func (prof *Profile) SyncDetails() error {
	jsons, err := prof.reqAboutCollection(details)
	if err != nil {
		return err
	}

	// DELETE
	// jsonextract.SaveToPath(jsons, "raw_details.json")

	for _, json := range jsons {
		val, ok := json.KeyVal["label"]
		if !ok {
			continue
		}

		if val.Val.(string) == "ProfileCometAboutAppSectionQuery$defer$ProfileCometAboutAppSectionContent_appSection" {
			details := extractDetails(json)
			prof.About.Details = details

			break
		}
	}

	return nil
}

func extractDetails(json *jsonextract.JSON) *Details {
	details := new(Details)
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

								// extract about
								if val.Val.(string) == "about_me" {
									if val, ok := section.KeyVal["profile_fields"]; ok {
										if val, ok := val.KeyVal["nodes"]; ok {
											for _, node := range val.Vals {
												if val, ok := node.KeyVal["renderer"]; ok {
													if val, ok := val.KeyVal["field"]; ok {
														if val, ok := val.KeyVal["text_content"]; ok {
															if val, ok := val.KeyVal["text"]; ok {
																details.About = val.Val.(string)
															}
														}
													}
												}
											}
										}
									}
								}

								// extract name pronunciation
								if val.Val.(string) == "name_pronunciation" {
									if val, ok := section.KeyVal["profile_fields"]; ok {
										if val, ok := val.KeyVal["nodes"]; ok {
											for _, node := range val.Vals {
												if val, ok := node.KeyVal["renderer"]; ok {
													pronun := new(NamePronunciation)
													if val, ok := val.KeyVal["audio_uri"]; ok {
														pronun.AudioURI = val.Val.(string)
													}

													if val, ok := val.KeyVal["text_content"]; ok {
														if val, ok := val.KeyVal["text"]; ok {
															pronun.Text = val.Val.(string)
														}
													}

													details.NamePronunciation = pronun
												}
											}
										}
									}
								}

								// extract nicknames
								if val.Val.(string) == "nicknames" {
									if val, ok := section.KeyVal["profile_fields"]; ok {
										if val, ok := val.KeyVal["nodes"]; ok {
											for i, node := range val.Vals {
												if i == 0 {
													continue
												}

												nickname := Nickname{
													Name: node.KeyVal["title"].KeyVal["text"].Val.(string),
												}

												if val, ok := node.KeyVal["list_item_groups"]; ok {
													for _, group := range val.Vals {
														if val, ok := group.KeyVal["list_items"]; ok {
															for _, item := range val.Vals {
																if val, ok := item.KeyVal["text"]; ok {
																	if val, ok := val.KeyVal["text"]; ok {
																		nickname.Type = val.Val.(string)
																	}
																}
															}
														}
													}
												}

												details.Nicknames = append(details.Nicknames, nickname)
											}
										}
									}
								}

								// extract quotes
								if val.Val.(string) == "favorite_quotes" {
									if val, ok := section.KeyVal["profile_fields"]; ok {
										if val, ok := val.KeyVal["nodes"]; ok {
											for _, node := range val.Vals {
												if val, ok := node.KeyVal["renderer"]; ok {
													if val, ok := val.KeyVal["field"]; ok {
														if val, ok := val.KeyVal["text_content"]; ok {
															if val, ok := val.KeyVal["text"]; ok {
																details.FavoriteQuotes = val.Val.(string)
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
