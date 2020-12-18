package raw

import "encoding/json"

type Item struct {
	Bbox                 *Bbox             `json:"__bbox"`
	Require              []json.RawMessage `json:"require"`
	ActorID              string            `json:"actorID"`
	RootView             *View             `json:"rootView"`
	TracePolicy          string            `json:"tracePolicy"`
	Meta                 *Meta             `json:"meta"`
	TimeSpentConfig      *TimeSpentConfig  `json:"timeSpentConfig"`
	EntityKeyConfig      *EntityKeyConfig  `json:"entityKeyConfig"`
	HostableView         *View             `json:"hostableView"`
	ProductAttributionID string            `json:"productAttributionId"`
	URL                  string            `json:"url"`
	Params               *Params           `json:"params"`
	RoutePath            string            `json:"routePath"`
}

type EntityKeyConfig struct {
	EntityType *EntityID `json:"entity_type"`
	EntityID   *EntityID `json:"entity_id"`
	Section    *EntityID `json:"section"`
}

type EntityID struct {
	Source string `json:"source"`
	Value  string `json:"value"`
}

type View struct {
	AllResources []EntryPoint `json:"allResources"`
	Resource     *EntryPoint  `json:"resource"`
	Props        *Props       `json:"props"`
	EntryPoint   *EntryPoint  `json:"entryPoint"`
}

type EntryPoint struct {
	DR string `json:"__dr"`
}

type Props struct {
	ViewerID   string `json:"viewerID"`
	UserVanity string `json:"userVanity"`
	UserID     string `json:"userID"`
}

type Meta struct {
	Title     string      `json:"title"`
	Accessory interface{} `json:"accessory"`
}

type Params struct {
	Vanity          string      `json:"vanity"`
	Sk              interface{} `json:"sk"`
	Viewas          interface{} `json:"viewas"`
	BoostedAutoOpen interface{} `json:"boosted_auto_open"`
	BoostPostID     interface{} `json:"boost_post_id"`
	BoostID         interface{} `json:"boost_id"`
	BoostRef        interface{} `json:"boost_ref"`
	So              interface{} `json:"so"`
}

type TimeSpentConfig struct {
	HasProfileSessionID bool `json:"has_profile_session_id"`
}
