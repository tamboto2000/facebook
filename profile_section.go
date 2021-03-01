package facebook

import "sync"

// ProfileSection contains tokens for profile sections data, like About, Friends, etc.
type ProfileSection struct {
	ActorID     string     `json:"actorID,omitempty"`
	PreloaderID string     `json:"preloaderID,omitempty"`
	QueryID     string     `json:"queryID,omitempty"`
	Variables   *Variables `json:"variables,omitempty"`
	collections []Collection

	mutex *sync.Mutex
}

type Variables struct {
	AppSectionFeedKey string  `json:"appSectionFeedKey,omitempty"`
	CollectionToken   string  `json:"collectionToken,omitempty"`
	RawSectionToken   string  `json:"rawSectionToken,omitempty"`
	Scale             float64 `json:"scale,omitempty"`
	SectionToken      string  `json:"sectionToken,omitempty"`
	UserID            string  `json:"userID,omitempty"`
}

// Collection contains token for section data
type Collection struct {
	TabKey string `json:"tab_key,omitempty"`
	ID     string `json:"id,omitempty"`
}

func newProfileSection() *ProfileSection {
	return &ProfileSection{mutex: new(sync.Mutex)}
}

func (p *ProfileSection) getColl(tabKey string) Collection {
	p.mutex.Lock()
	var coll Collection

	for _, c := range p.collections {
		if c.TabKey == tabKey {
			coll = c
			break
		}
	}

	p.mutex.Unlock()

	return coll
}

func (p *ProfileSection) getVariables() Variables {
	p.mutex.Lock()
	vars := *p.Variables
	p.mutex.Unlock()

	return vars
}
