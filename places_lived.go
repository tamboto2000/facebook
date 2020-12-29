package facebook

import (
	"errors"
	"net/url"

	"github.com/tamboto2000/jsonextract/v2"
)

// Place contains profile's place info
type Place struct {
	Name    string `json:"name"`
	URL     string `json:"url,omitempty"`
	PlaceIs string `json:"placeIs"`
	Icon    *Photo `json:"icon,omitempty"`
}

// SyncPlacesLived retrieve profile's places lived history
func (prof *Profile) SyncPlacesLived() error {
	var section *jsonextract.JSON
	for _, val := range prof.profileSections.KeyVal["edges"].Vals {
		node, ok := val.KeyVal["node"]
		if !ok {
			continue
		}

		if val, ok := node.KeyVal["section_type"]; ok && val.Val.(string) == SectionAbout {
			section = node
			break
		}
	}

	if section == nil {
		return errors.New("Important tokens for About section is not founs")
	}

	var coll *jsonextract.JSON
	for _, val := range section.KeyVal["all_collections"].KeyVal["nodes"].Vals {
		tabKey, ok := val.KeyVal["tab_key"]
		if !ok {
			continue
		}

		if tabKey.Val.(string) == "about_places" {
			coll = val
			break
		}
	}

	vars := prof.aboutSectionVars.KeyVal["variables"]
	vars.KeyVal["collectionToken"].Val = coll.KeyVal["id"].Val
	if err := vars.ReParse(); err != nil {
		return err
	}

	// fmt.Println(string(vars.Raw.Bytes()))

	reqBody := make(url.Values)
	reqBody.Set("fb_api_req_friendly_name", "ProfileCometAboutAppSectionQuery")
	reqBody.Set("variables", string(vars.Raw.Bytes()))
	reqBody.Set("doc_id", prof.aboutSectionVars.KeyVal["queryID"].Val.(string))
	rawBody, err := prof.fb.graphQlRequest(reqBody)
	if err != nil {
		return err
	}

	// DELETE
	// f, _ := os.Create("raw_work_education.json")
	// defer f.Close()
	// f.Write(rawBody)

	jsons, err := jsonextract.FromBytes(rawBody)
	if err != nil {
		return err
	}

	// DELETE
	// jsonextract.SaveToPath(jsons, "places_lived_bundle.json")

	for _, json := range jsons {
		if val, ok := json.KeyVal["label"]; ok {
			if val.Val.(string) == "ProfileCometAboutAppSectionQuery$defer$ProfileCometAboutAppSectionContent_appSection" {
				prof.About.PlacesLived = extractPlaceLived(json)
				break
			}
		}
	}

	return nil
}

func extractPlaceLived(json *jsonextract.JSON) []Place {
	places := make([]Place, 0)
	if val, ok := json.KeyVal["data"]; ok {
		if val, ok := val.KeyVal["activeCollections"]; ok {
			if val, ok := val.KeyVal["nodes"]; ok {
				for _, node := range val.Vals {
					if val, ok := node.KeyVal["style_renderer"]; ok {
						if val, ok := val.KeyVal["profile_field_sections"]; ok {
							for _, section := range val.Vals {
								if val, ok := section.KeyVal["profile_fields"]; ok {
									if val, ok := val.KeyVal["nodes"]; ok {
										for i, node := range val.Vals {
											if i == 0 {
												continue
											}

											place := Place{
												Name:    node.KeyVal["title"].KeyVal["text"].Val.(string),
												PlaceIs: node.KeyVal["field_type"].Val.(string),
											}

											// extract place url
											if val, ok := node.KeyVal["title"].KeyVal["ranges"]; ok {
												if len(val.Vals) > 0 {
													for _, rng := range val.Vals {
														if val, ok := rng.KeyVal["entity"]; ok {
															if val, ok := val.KeyVal["url"]; ok {
																place.URL = val.Val.(string)
															}
														}
													}
												}
											}

											// extract icon
											if val, ok := node.KeyVal["renderer"]; ok {
												if val, ok := val.KeyVal["field"]; ok {
													if val, ok := val.KeyVal["icon"]; ok {
														place.Icon = &Photo{
															Height: val.KeyVal["height"].Val.(int),
															Scale:  float64(val.KeyVal["scale"].Val.(int)),
															URI:    val.KeyVal["uri"].Val.(string),
															Width:  val.KeyVal["width"].Val.(int),
														}
													}
												}
											}

											places = append(places, place)
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

	return places
}
