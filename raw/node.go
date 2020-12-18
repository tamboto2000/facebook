package raw

type Node struct {
	Nodes            []Node      `json:"nodes,omitempty"`
	DisplayableCount *int        `json:"displayable_count,omitempty"`
	Name             string      `json:"name,omitempty"`
	SectionType      string      `json:"section_type,omitempty"`
	TabKey           string      `json:"tab_key,omitempty"`
	Tracking         string      `json:"tracking,omitempty"`
	URL              string      `json:"url,omitempty"`
	AllCollections   *Node       `json:"all_collections,omitempty"`
	ID               string      `json:"id,omitempty"`
	FirstStoryToShow interface{} `json:"first_story_to_show,omitempty"`
}
