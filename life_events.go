package facebook

import (
	"github.com/tamboto2000/jsonextract/v2"
)

// LifeEvent contains life event
type LifeEvent struct {
	Title string `json:"title,omitempty"`
	URL   string `json:"url,omitempty"`
}

// LifeEvents contains life events and the year the events happened
type LifeEvents struct {
	Year   int         `json:"year,omitempty"`
	Events []LifeEvent `json:"events,omitempty"`
}

// SyncLifeEvents retrieve profile's life events
func (prof *Profile) SyncLifeEvents() error {
	jsons, err := prof.reqAboutCollection(lifeEvents)
	if err != nil {
		return err
	}

	// DELETE
	// jsonextract.SaveToPath(jsons, "raw_life_events.json")
	for _, json := range jsons {
		val, ok := json.KeyVal["label"]
		if !ok {
			continue
		}

		if val.Val.(string) == "ProfileCometAboutAppSectionQuery$defer$ProfileCometAboutAppSectionContent_appSection" {
			prof.About.LifeEvents = extractLifeEvents(json)

			break
		}
	}

	return nil
}

func extractLifeEvents(json *jsonextract.JSON) []LifeEvents {
	events := make([]LifeEvents, 0)
	if val, ok := json.KeyVal["data"]; ok {
		if val, ok := val.KeyVal["activeCollections"]; ok {
			if val, ok := val.KeyVal["nodes"]; ok {
				for _, node := range val.Vals {
					if val, ok := node.KeyVal["style_renderer"]; ok {
						if val, ok := val.KeyVal["user"]; ok {
							if val, ok := val.KeyVal["timeline_sections"]; ok {
								if val, ok := val.KeyVal["nodes"]; ok {
									for i, node := range val.Vals {
										if i == 0 {
											continue
										}

										yearEvent := LifeEvents{Year: node.KeyVal["year"].Val.(int)}
										if val, ok := node.KeyVal["year_overview"]; ok {
											if val, ok := val.KeyVal["items"]; ok {
												if val, ok := val.KeyVal["nodes"]; ok {
													for _, node := range val.Vals {
														event := LifeEvent{
															Title: node.KeyVal["title"].KeyVal["text"].Val.(string),
															URL:   node.KeyVal["url"].Val.(string),
														}

														yearEvent.Events = append(yearEvent.Events, event)
													}
												}
											}
										}

										events = append(events, yearEvent)
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return events
}
