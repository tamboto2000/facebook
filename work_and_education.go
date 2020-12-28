package facebook

import (
	"github.com/tamboto2000/facebook/raw"
)

// Work contains Profile's work/occupation history
type Work struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Icon        *Photo `json:"icon,omitempty"`
}

// SyncWorkAndEducation retrieve Profile's work/occupation history and education history
func (prof *Profile) SyncWorkAndEducation() error {
	// var node *raw.Node
	// for _, edge := range prof.ProfileSections.Edges {
	// 	if edge.Node != nil {
	// 		if edge.Node.SectionType == SectionAbout {
	// 			node = edge.Node
	// 			break
	// 		}
	// 	}
	// }

	// var collection raw.Node
	// for _, col := range node.AllCollections.Nodes {
	// 	if col.TabKey == WorkAndEducation {
	// 		collection = col
	// 		break
	// 	}
	// }

	// vars := *prof.AboutSectionVars.Variables
	// vars.CollectionToken = collection.ID

	// varsByts, err := json.Marshal(vars)
	// if err != nil {
	// 	return err
	// }

	// reqBody := make(url.Values)
	// reqBody.Set("fb_api_req_friendly_name", "ProfileCometAboutAppSectionQuery")
	// reqBody.Set("variables", string(varsByts))
	// reqBody.Set("doc_id", prof.AboutSectionVars.QueryID)
	// rawBody, err := prof.fb.graphQlRequest(reqBody)
	// if err != nil {
	// 	return err
	// }

	// // DELETE
	// f, _ := os.Create("raw_work_education.json")
	// defer f.Close()
	// f.Write(rawBody)

	// jsons, err := jsonextract.JSONFromBytes(rawBody)
	// if err != nil {
	// 	return err
	// }

	// // DELETE
	// jsonextract.SaveToPath(jsons, "work_education_bundle.json")

	// for i, frag := range jsons {
	// 	item := new(raw.Item)
	// 	if err := json.Unmarshal(frag, item); err == nil {
	// 		// DELETE
	// 		f, _ := os.Create("work_education_try_success_" + strconv.Itoa(i) + ".json")
	// 		defer f.Close()
	// 		f.Write(frag)

	// 		works := extractWorks(item)
	// 		if len(works) > 0 {
	// 			prof.About.WorkHistory = append(prof.About.WorkHistory, works...)
	// 		}

	// 		// DELETE ELSE BLOCK
	// 	} else {
	// 		// DELETE
	// 		f, _ := os.Create("work_education_try_fail_" + strconv.Itoa(i) + ".json")
	// 		defer f.Close()
	// 		f.Write(frag)
	// 	}
	// }

	return nil
}

func extractWorks(item *raw.Item) []Work {
	works := make([]Work, 0)
	if item.Data != nil {
		if item.Data.ActiveCollections != nil {
			if item.Data.ActiveCollections.Nodes != nil {
				for _, node := range item.Data.ActiveCollections.Nodes {
					if node.StyleRenderer != nil {
						if node.StyleRenderer.ProfileFieldSections != nil {
							for _, section := range node.StyleRenderer.ProfileFieldSections {
								if section.FieldSectionType == "work" {
									if section.ProfileFields != nil {
										if section.ProfileFields.Nodes != nil {
											for _, node := range section.ProfileFields.Nodes {
												// extract title
												work := Work{Title: node.Title.Text}

												// if node.Renderer != nil {
												// 	renderer := node.Renderer

												// 	if renderer.Field
												// }

												works = append(works, work)
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

	return works
}
