package facebook

import (
	"encoding/base64"
	"strings"
)

func (user *User) SyncDetails() error {
	return user.reqDetails()
}

func (user *User) reqDetails() error {
	sectionKey, err := base64.StdEncoding.DecodeString(user.userSections["ABOUT"].Node.ID)
	if err != nil {
		return err
	}

	sectionKeyRaw := strings.Replace(string(sectionKey), "app_section:", "", 1)
	var collectionToken string
	for _, node := range user.userSections["ABOUT"].Node.AllCollections.Nodes {
		if strings.Contains(node.URL, "/about_details") {
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
					if c.FieldType == "about_me" {
						user.Bio = c.Renderer.Field.TextContent.Text
						continue
					}

					if c.FieldType == "quotes" {
						user.Quote = c.Renderer.Field.TextContent.Text
						continue
					}
				}
			}
		}
	}

	return nil
}
