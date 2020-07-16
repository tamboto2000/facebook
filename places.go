package facebook

import (
	"encoding/base64"
	"strings"
)

type Place struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	URL     string `json:"url"`
	Name    string `json:"name"`
	IconURL string `json:"iconUrl"`
}

func (user *User) SyncPlaces() error {
	return user.reqPlaces()
}

func (user *User) reqPlaces() error {
	sectionKey, err := base64.StdEncoding.DecodeString(user.userSections["ABOUT"].Node.ID)
	if err != nil {
		return err
	}

	sectionKeyRaw := strings.Replace(string(sectionKey), "app_section:", "", 1)
	var collectionToken string
	for _, node := range user.userSections["ABOUT"].Node.AllCollections.Nodes {
		if strings.Contains(node.URL, "/about_places") {
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
			for _, b := range a.StyleRenderer.ProfileFieldSections {
				for _, c := range b.ProfileFields.Nodes {
					place := Place{
						Type: c.FieldType,
						Name: c.Title.Text,
					}

					if c.FieldType == "current_city" {
						user.CurrentCity = c.Title.Text
					}

					if len(c.Renderer.Field.Title.Ranges) > 0 {
						entity := c.Renderer.Field.Title.Ranges[0].Entity
						place.ID = entity.ID
						place.URL = entity.URL
					}

					place.IconURL = c.Renderer.Field.Icon.URI

					user.Places = append(user.Places, place)
				}
			}
		}
	}

	if len(user.Places) == 0 {
		user.Places = nil
	}

	return nil
}
