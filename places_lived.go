package facebook

import (
	"github.com/tamboto2000/jsonextract/v3"
)

// Place contains profile's place info
type Place struct {
	Name    string `json:"name"`
	URL     string `json:"url,omitempty"`
	PlaceIs string `json:"placeIs"`
	Icon    *Photo `json:"icon,omitempty"`
}

// SyncPlacesLived retrieve profile's places lived history
func (about *About) SyncPlacesLived() error {
	jsons, err := about.profile.reqAboutCollection(aboutPlacesLived)
	if err != nil {
		return err
	}

	for _, json := range jsons {
		if val, ok := json.Object()["label"]; ok {
			if val.String() == "ProfileCometAboutAppSectionQuery$defer$ProfileCometAboutAppSectionContent_appSection" {
				about.PlacesLived = extractPlaceLived(json)
				break
			}
		}
	}

	return nil
}

func extractPlaceLived(json *jsonextract.JSON) []Place {
	places := make([]Place, 0)
	if val, ok := json.Object()["data"]; ok {
		if val, ok := val.Object()["activeCollections"]; ok {
			if val, ok := val.Object()["nodes"]; ok {
				for _, node := range val.Array() {
					if val, ok := node.Object()["style_renderer"]; ok {
						if val, ok := val.Object()["profile_field_sections"]; ok {
							for _, section := range val.Array() {
								if val, ok := section.Object()["profile_fields"]; ok {
									if val, ok := val.Object()["nodes"]; ok {
										for i, node := range val.Array() {
											if i == 0 {
												continue
											}

											place := Place{
												Name:    node.Object()["title"].Object()["text"].String(),
												PlaceIs: node.Object()["field_type"].String(),
											}

											// extract place url
											if val, ok := node.Object()["title"].Object()["ranges"]; ok {
												if len(val.Array()) > 0 {
													for _, rng := range val.Array() {
														if val, ok := rng.Object()["entity"]; ok {
															if val, ok := val.Object()["url"]; ok {
																place.URL = val.String()
															}
														}
													}
												}
											}

											// extract icon
											if val, ok := node.Object()["renderer"]; ok {
												if val, ok := val.Object()["field"]; ok {
													if val, ok := val.Object()["icon"]; ok {
														place.Icon = &Photo{
															Height: int(val.Object()["height"].Integer()),
															URI:    val.Object()["uri"].String(),
															Width:  int(val.Object()["width"].Integer()),
														}

														scale := val.Object()["scale"]
														if scale.Kind() == jsonextract.Integer {
															place.Icon.Scale = float64(scale.Integer())
														} else {
															place.Icon.Scale = scale.Float()
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
