package facebook

import (
	"encoding/base64"
	"strings"
)

type Family struct {
	User     *User  `json:"user"`
	Relation string `json:"relation"`
}

func (user *User) SyncFamilyAndRelationships() error {
	return user.reqFamilyRelationships()
}

func (user *User) reqFamilyRelationships() error {
	sectionKey, err := base64.StdEncoding.DecodeString(user.userSections["ABOUT"].Node.ID)
	if err != nil {
		return err
	}

	sectionKeyRaw := strings.Replace(string(sectionKey), "app_section:", "", 1)
	var collectionToken string
	for _, node := range user.userSections["ABOUT"].Node.AllCollections.Nodes {
		if strings.Contains(node.URL, "/about_family_and_relationships") {
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
					if c.FieldType == "relationship" {
						family := Family{}
						if c.Renderer.Field.Title.Ranges != nil && len(c.Renderer.Field.Title.Ranges) > 0 {
							newUser := new(User)
							newUser.Name = c.Title.Text
							rawUser := c.Renderer.Field.Title.Ranges[0].Entity
							newUser.ID = rawUser.ID

							//extract username
							stg1 := strings.Split(c.Renderer.Field.Title.Ranges[0].Entity.ProfileURL, "/")
							newUser.Username = stg1[len(stg1)-1]

							family.User = newUser
						}

						if len(c.Renderer.Field.ListItemGroups) > 0 {
							for _, itemGroup := range c.Renderer.Field.ListItemGroups {
								for _, item := range itemGroup.ListItems {
									if strings.Contains(item.Text.Text, "Married") {
										user.RelationshipStatus = "married"
										family.Relation = "married"
									}
								}
							}
						} else {
							if strings.Contains(c.Renderer.Field.TextContent.Text, "Married") {
								user.RelationshipStatus = "married"
							}
						}

						if family.User != nil {
							user.FamilyMember = append(user.FamilyMember, family)
						}
					}

					if c.FieldType == "family" {
						family := Family{}
						newUser := new(User)
						newUser.Name = c.Renderer.Field.Title.Text
						if c.Renderer.Field.Title.Ranges != nil && len(c.Renderer.Field.Title.Ranges) > 0 {
							rawUser := c.Renderer.Field.Title.Ranges[0].Entity
							newUser.ID = rawUser.ID

							//extract username
							stg1 := strings.Split(rawUser.ProfileURL, "/")
							newUser.Username = stg1[len(stg1)-1]
							newUser.ProfilePictURL = c.Renderer.Field.Icon.URI
						}

						for _, itemGroup := range c.Renderer.Field.ListItemGroups {
							for _, item := range itemGroup.ListItems {
								family.Relation = strings.ToLower(item.Text.Text)
							}
						}

						family.User = newUser
						user.FamilyMember = append(user.FamilyMember, family)
					}
				}
			}
		}
	}

	if len(user.FamilyMember) == 0 {
		user.FamilyMember = nil
	}

	return nil
}
