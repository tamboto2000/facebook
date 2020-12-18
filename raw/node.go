package raw

type Node struct {
	DisplayableCount *int        `json:"displayable_count"`
	Name             string      `json:"name"`
	SectionType      string      `json:"section_type"`
	TabKey           string      `json:"tab_key"`
	Tracking         string      `json:"tracking"`
	URL              string      `json:"url"`
	AllCollections   *Node       `json:"all_collections"`
	ID               string      `json:"id"`
	FirstStoryToShow interface{} `json:"first_story_to_show"`
}
