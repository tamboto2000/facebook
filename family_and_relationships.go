package facebook

import "github.com/tamboto2000/jsonextract/v3"

// FamilyMember contains information about a family member, including some info about its profile
type FamilyMember struct {
	*Profile
	URLToProfile string `json:"urlToProfile,omitempty"`
	// Relationship status, like mother, mother-in-law, or maybe, the one that tempting a lot of young men, step-sister
	Relationship string `json:"relationship"`
}

// Relationship contains relationship with someone, including some info about its profile
type Relationship struct {
	*Profile
	URLToProfile string `json:"urlToProfile,omitempty"`
	// Relationship status like girlfriend, wife, or even divorced...
	Status string `json:"status"`
}

// FamilyAndRelationships contains profile's family and relationships info
type FamilyAndRelationships struct {
	FamilyMembers []FamilyMember `json:"familyMembers,omitempty"`
	Relationships []Relationship `json:"relationships,omitempty"`
}

// SyncFamilyAndRelationships retrieve family and relationships
func (about *About) SyncFamilyAndRelationships() error {
	jsons, err := about.profile.reqAboutCollection(aboutFamilyAndRelationships)
	if err != nil {
		return err
	}

	for _, json := range jsons {
		if val, ok := json.Object()["label"]; ok {
			if val.String() == "ProfileCometAboutAppSectionQuery$defer$ProfileCometAboutAppSectionContent_appSection" {
				members, relations := extractFamilyMember(json)
				about.FamilyAndRelationships = &FamilyAndRelationships{
					FamilyMembers: members,
					Relationships: relations,
				}

				break
			}
		}
	}

	return nil
}

func extractFamilyMember(json *jsonextract.JSON) ([]FamilyMember, []Relationship) {
	members := make([]FamilyMember, 0)
	relations := make([]Relationship, 0)

	if val, ok := json.Object()["data"]; ok {
		if val, ok := val.Object()["activeCollections"]; ok {
			if val, ok := val.Object()["nodes"]; ok {
				for _, node := range val.Array() {
					if val, ok := node.Object()["style_renderer"]; ok {
						if val, ok := val.Object()["profile_field_sections"]; ok {
							for _, section := range val.Array() {
								val, ok := section.Object()["field_section_type"]
								if !ok {
									continue
								}

								// extract relationships
								if val.String() == "relationship" {
									if val, ok := section.Object()["profile_fields"]; ok {
										if val, ok := val.Object()["nodes"]; ok {
											for _, node := range val.Array() {
												relation := Relationship{
													Profile: &Profile{
														Name: node.Object()["title"].Object()["text"].String(),
													},
												}

												// extract user id and profile url
												if val, ok := node.Object()["title"]; ok {
													if val, ok := val.Object()["ranges"]; ok {
														for _, rng := range val.Array() {
															if val, ok := rng.Object()["entity"]; ok {
																relation.Profile.ID = val.Object()["id"].String()
																relation.URLToProfile = val.Object()["url"].String()
															}
														}
													}
												}

												// extract user profile photo (more like icon to be honest...)
												if val, ok := node.Object()["renderer"]; ok {
													if val, ok := val.Object()["field"]; ok {
														if val, ok := val.Object()["icon"]; ok {
															relation.Profile.ProfilePhoto = &Photo{
																Height: int(val.Object()["height"].Integer()),
																URI:    val.Object()["uri"].String(),
																Width:  int(val.Object()["width"].Integer()),
															}

															scale := val.Object()["scale"]
															if scale.Kind() == jsonextract.Integer {
																relation.Profile.ProfilePhoto.Scale = float64(scale.Integer())
															} else {
																relation.Profile.ProfilePhoto.Scale = scale.Float()
															}
														}
													}
												}

												// extract relation status
												if val, ok := node.Object()["list_item_groups"]; ok {
													for _, group := range val.Array() {
														if val, ok := group.Object()["list_items"]; ok {
															for _, item := range val.Array() {
																if val, ok := item.Object()["text"]; ok {
																	relation.Status = val.Object()["text"].String()
																}
															}
														}
													}
												}

												relations = append(relations, relation)
											}
										}
									}
								}

								// extract family members
								if val.String() == "family" {
									if val, ok := section.Object()["profile_fields"]; ok {
										if val, ok := val.Object()["nodes"]; ok {
											for i, node := range val.Array() {
												if i == 0 {
													continue
												}

												family := FamilyMember{
													Profile: &Profile{
														Name: node.Object()["title"].Object()["text"].String(),
													},
												}

												// extract user id and profile url
												if val, ok := node.Object()["title"]; ok {
													if val, ok := val.Object()["ranges"]; ok {
														for _, rng := range val.Array() {
															if val, ok := rng.Object()["entity"]; ok {
																family.Profile.ID = val.Object()["id"].String()
																family.URLToProfile = val.Object()["url"].String()
															}
														}
													}
												}

												// extract user profile photo (more like icon to be honest...)
												if val, ok := node.Object()["renderer"]; ok {
													if val, ok := val.Object()["field"]; ok {
														if val, ok := val.Object()["icon"]; ok {
															family.Profile.ProfilePhoto = &Photo{
																Height: int(val.Object()["height"].Integer()),
																URI:    val.Object()["uri"].String(),
																Width:  int(val.Object()["width"].Integer()),
															}

															scale := val.Object()["scale"]
															if scale.Kind() == jsonextract.Integer {
																family.Profile.ProfilePhoto.Scale = float64(scale.Integer())
															} else {
																family.Profile.ProfilePhoto.Scale = scale.Float()
															}
														}
													}
												}

												// extract member relation status
												if val, ok := node.Object()["list_item_groups"]; ok {
													for _, group := range val.Array() {
														if val, ok := group.Object()["list_items"]; ok {
															for _, item := range val.Array() {
																if val, ok := item.Object()["text"]; ok {
																	family.Relationship = val.Object()["text"].String()
																}
															}
														}
													}
												}

												members = append(members, family)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return members, relations
}
