package facebook

import (
	"github.com/tamboto2000/jsonextract/v3"
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
	jsons, err := prof.reqAboutCollection(aboutLifeEvents)
	if err != nil {
		return err
	}

	// DELETE
	// jsonextract.SaveToPath(jsons, "raw_life_events.json")
	for _, json := range jsons {
		val, ok := json.Object()["label"]
		if !ok {
			continue
		}

		if val.String() == "ProfileCometAboutAppSectionQuery$defer$ProfileCometAboutAppSectionContent_appSection" {
			prof.About.LifeEvents = extractLifeEvents(json)

			break
		}
	}

	return nil
}

func extractLifeEvents(json *jsonextract.JSON) []LifeEvents {
	events := make([]LifeEvents, 0)
	if val, ok := json.Object()["data"]; ok {
		if val, ok := val.Object()["activeCollections"]; ok {
			if val, ok := val.Object()["nodes"]; ok {
				for _, node := range val.Array() {
					if val, ok := node.Object()["style_renderer"]; ok {
						if val, ok := val.Object()["user"]; ok {
							if val, ok := val.Object()["timeline_sections"]; ok {
								if val, ok := val.Object()["nodes"]; ok {
									for i, node := range val.Array() {
										if i == 0 {
											continue
										}

										yearEvent := LifeEvents{Year: int(node.Object()["year"].Integer())}
										if val, ok := node.Object()["year_overview"]; ok {
											if val, ok := val.Object()["items"]; ok {
												if val, ok := val.Object()["nodes"]; ok {
													for _, node := range val.Array() {
														event := LifeEvent{
															Title: node.Object()["title"].Object()["text"].String(),
															URL:   node.Object()["url"].String(),
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
