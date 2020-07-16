package facebook

import (
	"encoding/base64"
	"strings"
)

type EventItem struct {
	Event string `json:"event"`
	URL   string `json:"url"`
}

type Event struct {
	Year   int         `json:"year"`
	Events []EventItem `json:"events"`
}

func (user *User) SyncLifeEvents() error {
	return user.reqLifeEvents()
}

func (user *User) reqLifeEvents() error {
	sectionKey, err := base64.StdEncoding.DecodeString(user.userSections["ABOUT"].Node.ID)
	if err != nil {
		return err
	}

	sectionKeyRaw := strings.Replace(string(sectionKey), "app_section:", "", 1)
	var collectionToken string
	for _, node := range user.userSections["ABOUT"].Node.AllCollections.Nodes {
		if strings.Contains(node.URL, "/about_life_events") {
			collectionToken = node.ID
			break
		}
	}

	payloads, err := user.fb.doGraphQLRequest(map[string]interface{}{
		"appSectionFeedKey":      "ProfileCometAppSectionFeed_timeline_nav_app_sections__" + sectionKeyRaw,
		"collectionToken":        collectionToken,
		"rawSectionToken":        sectionKeyRaw,
		"scale":                  1,
		"sectionToken":           user.userSections["ABOUT"].Node.ID,
		"userID":                 user.ID,
		"useIncrementalDelivery": true,
	}, user.aboutDocID, "ProfileCometAboutAppSectionQuery", true)
	if err != nil {
		return err
	}

	for _, payload := range payloads {
		for _, a := range payload.Data.ActiveCollections.Nodes {
			for _, b := range a.StyleRenderer.User.TimelineSections.Nodes {
				if b.Year == 0 {
					continue
				}

				event := Event{Year: int(b.Year)}
				for _, c := range b.YearOverview.Items.Nodes {
					event.Events = append(event.Events, EventItem{
						Event: c.Title.Text,
						URL:   c.URL,
					})
				}

				user.Events = append(user.Events, event)
			}
		}
	}

	return nil
}
