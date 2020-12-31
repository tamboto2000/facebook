package facebook

import (
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
	jsons, err := prof.reqAboutCollection(aboutPlacesLived)
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
