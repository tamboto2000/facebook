package facebook

import "github.com/tamboto2000/jsonextract/v2"

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
func (prof *Profile) SyncFamilyAndRelationships() error {
	jsons, err := prof.reqAboutCollection(familyAndRelationships)
	if err != nil {
		return err
	}

	// DELETE
	// jsonextract.SaveToPath(jsons, "raw_family_relationships.json")

	for _, json := range jsons {
		if val, ok := json.KeyVal["label"]; ok {
			if val.Val.(string) == "ProfileCometAboutAppSectionQuery$defer$ProfileCometAboutAppSectionContent_appSection" {
				members, relations := extractFamilyMember(json)
				prof.About.FamilyAndRelationships = &FamilyAndRelationships{
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

	if val, ok := json.KeyVal["data"]; ok {
		if val, ok := val.KeyVal["activeCollections"]; ok {
			if val, ok := val.KeyVal["nodes"]; ok {
				for _, node := range val.Vals {
					if val, ok := node.KeyVal["style_renderer"]; ok {
						if val, ok := val.KeyVal["profile_field_sections"]; ok {
							for _, section := range val.Vals {
								val, ok := section.KeyVal["field_section_type"]
								if !ok {
									continue
								}

								// extract relationships
								if val.Val.(string) == "relationship" {
									if val, ok := section.KeyVal["profile_fields"]; ok {
										if val, ok := val.KeyVal["nodes"]; ok {
											for _, node := range val.Vals {
												relation := Relationship{
													Profile: &Profile{
														Name: node.KeyVal["title"].KeyVal["text"].Val.(string),
													},
												}

												// extract user id and profile url
												if val, ok := node.KeyVal["title"]; ok {
													if val, ok := val.KeyVal["ranges"]; ok {
														for _, rng := range val.Vals {
															if val, ok := rng.KeyVal["entity"]; ok {
																relation.Profile.ID = val.KeyVal["id"].Val.(string)
																relation.URLToProfile = val.KeyVal["url"].Val.(string)
															}
														}
													}
												}

												// extract user profile photo (more like icon to be honest...)
												if val, ok := node.KeyVal["renderer"]; ok {
													if val, ok := val.KeyVal["field"]; ok {
														if val, ok := val.KeyVal["icon"]; ok {
															relation.Profile.ProfilePhoto = &Photo{
																Height: val.KeyVal["height"].Val.(int),
																Scale:  float64(val.KeyVal["scale"].Val.(int)),
																URI:    val.KeyVal["uri"].Val.(string),
																Width:  val.KeyVal["width"].Val.(int),
															}
														}
													}
												}

												// extract relation status
												if val, ok := node.KeyVal["list_item_groups"]; ok {
													for _, group := range val.Vals {
														if val, ok := group.KeyVal["list_items"]; ok {
															for _, item := range val.Vals {
																if val, ok := item.KeyVal["text"]; ok {
																	relation.Status = val.KeyVal["text"].Val.(string)
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
								if val.Val.(string) == "family" {
									if val, ok := section.KeyVal["profile_fields"]; ok {
										if val, ok := val.KeyVal["nodes"]; ok {
											for i, node := range val.Vals {
												if i == 0 {
													continue
												}

												family := FamilyMember{
													Profile: &Profile{
														Name: node.KeyVal["title"].KeyVal["text"].Val.(string),
													},
												}

												// extract user id and profile url
												if val, ok := node.KeyVal["title"]; ok {
													if val, ok := val.KeyVal["ranges"]; ok {
														for _, rng := range val.Vals {
															if val, ok := rng.KeyVal["entity"]; ok {
																family.Profile.ID = val.KeyVal["id"].Val.(string)
																family.URLToProfile = val.KeyVal["url"].Val.(string)
															}
														}
													}
												}

												// extract user profile photo (more like icon to be honest...)
												if val, ok := node.KeyVal["renderer"]; ok {
													if val, ok := val.KeyVal["field"]; ok {
														if val, ok := val.KeyVal["icon"]; ok {
															family.Profile.ProfilePhoto = &Photo{
																Height: val.KeyVal["height"].Val.(int),
																Scale:  float64(val.KeyVal["scale"].Val.(int)),
																URI:    val.KeyVal["uri"].Val.(string),
																Width:  val.KeyVal["width"].Val.(int),
															}
														}
													}
												}

												// extract member relation status
												if val, ok := node.KeyVal["list_item_groups"]; ok {
													for _, group := range val.Vals {
														if val, ok := group.KeyVal["list_items"]; ok {
															for _, item := range val.Vals {
																if val, ok := item.KeyVal["text"]; ok {
																	family.Relationship = val.KeyVal["text"].Val.(string)
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
