package raw

import "encoding/json"

type Item struct {
	Bbox                 *Bbox             `json:"__bbox,omitempty"`
	Require              []json.RawMessage `json:"require,omitempty"`
	ActorID              string            `json:"actorID,omitempty"`
	RootView             *View             `json:"rootView,omitempty"`
	TracePolicy          string            `json:"tracePolicy,omitempty"`
	Meta                 *Meta             `json:"meta,omitempty"`
	TimeSpentConfig      *TimeSpentConfig  `json:"timeSpentConfig,omitempty"`
	EntityKeyConfig      *EntityKeyConfig  `json:"entityKeyConfig,omitempty"`
	HostableView         *View             `json:"hostableView,omitempty"`
	ProductAttributionID string            `json:"productAttributionId,omitempty"`
	URL                  string            `json:"url,omitempty"`
	Params               *Params           `json:"params,omitempty"`
	RoutePath            string            `json:"routePath,omitempty"`
}

type EntityKeyConfig struct {
	EntityType *EntityID `json:"entity_type,omitempty"`
	EntityID   *EntityID `json:"entity_id,omitempty"`
	Section    *EntityID `json:"section,omitempty"`
}

type EntityID struct {
	Source string `json:"source,omitempty"`
	Value  string `json:"value,omitempty"`
}

type View struct {
	AllResources []EntryPoint `json:"allResources,omitempty"`
	Resource     *EntryPoint  `json:"resource,omitempty"`
	Props        *Props       `json:"props,omitempty"`
	EntryPoint   *EntryPoint  `json:"entryPoint,omitempty"`
}

type EntryPoint struct {
	DR string `json:"__dr,omitempty"`
}

type Props struct {
	ViewerID   string `json:"viewerID,omitempty"`
	UserVanity string `json:"userVanity,omitempty"`
	UserID     string `json:"userID,omitempty"`
}

type Meta struct {
	Title     string      `json:"title,omitempty"`
	Accessory interface{} `json:"accessory,omitempty"`
}

type Params struct {
	Vanity          string      `json:"vanity,omitempty"`
	Sk              interface{} `json:"sk,omitempty"`
	Viewas          interface{} `json:"viewas,omitempty"`
	BoostedAutoOpen interface{} `json:"boosted_auto_open,omitempty"`
	BoostPostID     interface{} `json:"boost_post_id,omitempty"`
	BoostID         interface{} `json:"boost_id,omitempty"`
	BoostRef        interface{} `json:"boost_ref,omitempty"`
	So              interface{} `json:"so,omitempty"`
}

type TimeSpentConfig struct {
	HasProfileSessionID bool `json:"has_profile_session_id,omitempty"`
}
