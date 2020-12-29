package facebook

import (
	"strings"

	"github.com/araddon/dateparse"
	"github.com/tamboto2000/jsonextract/v2"
)

// Work contains Profile's work/occupation information
type Work struct {
	Title         string `json:"title,omitempty"`
	CompanyURL    string `json:"companyUrl,omitempty"`
	Description   string `json:"description,omitempty"`
	DateStart     string `json:"dateStart,omitempty"`
	DateStartUnix int64  `json:"dateStartUnix,omitempty"`
	DateEnd       string `json:"dateEnd,omitempty"`
	DateEndUnix   int64  `json:"dateEndUnix,omitempty"`
	CompanyIcon   *Photo `json:"companyIcon,omitempty"`
	Location      string `json:"location,omitempty"`
}

// Education contain education information
type Education struct {
	Title         string `json:"title,omitempty"`
	SchoolURL     string `json:"schoolUrl,omitempty"`
	DateStart     string `json:"dateStart,omitempty"`
	DateStartUnix int64  `json:"dateStartUnix,omitempty"`
	DateEnd       string `json:"dateEnd,omitempty"`
	DateEndUnix   int64  `json:"dateEndUnix,omitempty"`
	// This field contains degree, concentrations, and description, if any.
	// Why I can't separate these data to different fields? Unfortunately
	// data from Facebook GraphQL API is not specified all data to which data which,
	// no flag, no tag, no reference resource id, just spitting data according to view
	// position. But this only apply to college type, for high_school the description
	// can be assiggned on separate field
	OtherInfo   []string `json:"otherInfo,omitempty"`
	Description string   `json:"description,omitempty"`
	SchoolIcon  *Photo   `json:"schoolIcon,omitempty"`
	// Only assigned if type is high_school
	ClassOf string `json:"classOf,omitempty"`
	// college or high_school
	Type string `json:"type,omitempty"`
}

// SyncWorkAndEducation retrieve Profile's work/occupation history and education history
func (prof *Profile) SyncWorkAndEducation() error {
	jsons, err := prof.reqAboutCollection(workAndEducation)
	if err != nil {
		return err
	}

	// DELETE
	// jsonextract.SaveToPath(jsons, "work_education_bundle.json")

	for _, json := range jsons {
		val, ok := json.KeyVal["label"]
		if !ok {
			continue
		}

		if val.Val.(string) == "ProfileCometAboutAppSectionQuery$defer$ProfileCometAboutAppSectionContent_appSection" {
			works, educations := extractWorks(json)
			prof.About.WorkHistory = works
			prof.About.EducationHistory = educations
			break
		}
	}

	return nil
}

func extractWorks(json *jsonextract.JSON) ([]Work, []Education) {
	works := make([]Work, 0)
	educations := make([]Education, 0)
	val, ok := json.KeyVal["data"]
	if !ok {
		return nil, nil
	}

	val, ok = val.KeyVal["activeCollections"]
	if !ok {
		return nil, nil
	}

	val, ok = val.KeyVal["nodes"]
	if !ok {
		return nil, nil
	}

	for _, node := range val.Vals {
		val, ok := node.KeyVal["style_renderer"]
		if !ok {
			continue
		}

		val, ok = val.KeyVal["profile_field_sections"]
		if !ok {
			continue
		}

		for _, section := range val.Vals {
			val, ok := section.KeyVal["field_section_type"]
			if !ok {
				continue
			}

			if val.Val.(string) == "work" {
				// start parsing work history
				val, ok := section.KeyVal["profile_fields"]
				if !ok {
					return nil, nil
				}

				val, ok = val.KeyVal["nodes"]
				if !ok {
					return nil, nil
				}

				for i, node := range val.Vals {
					// skip the first index because the first index is a button for add new work history
					if i == 0 {
						continue
					}

					// create Work and assign Title
					work := Work{
						Title: node.KeyVal["title"].KeyVal["text"].Val.(string),
					}

					// find company url
					if val, ok := node.KeyVal["title"].KeyVal["ranges"]; ok {
						for _, rng := range val.Vals {
							val, ok := rng.KeyVal["entity"]
							if !ok {
								continue
							}

							val, ok = rng.KeyVal["url"]
							if !ok {
								continue
							}

							work.CompanyURL = val.Val.(string)
						}
					}

					// find company icon
					if val, ok := node.KeyVal["renderer"]; ok {
						if val, ok := val.KeyVal["field"]; ok {
							if val, ok := val.KeyVal["icon"]; ok {
								work.CompanyIcon = &Photo{
									Height: val.KeyVal["height"].Val.(int),
									Scale:  val.KeyVal["scale"].Val.(float64),
									URI:    val.KeyVal["uri"].Val.(string),
									Width:  val.KeyVal["width"].Val.(int),
								}
							}
						}
					}

					// find description
					if val, ok := node.KeyVal["list_item_groups"]; ok {
						for _, itemGroup := range val.Vals {
							if val, ok := itemGroup.KeyVal["list_items"]; ok {
								for _, item := range val.Vals {
									val, ok := item.KeyVal["heading_type"]
									if !ok {
										continue
									}

									// description
									if val.Val.(string) == "MEDIUM" {
										if val, ok := item.KeyVal["text"]; ok {
											work.Description = val.KeyVal["text"].Val.(string)
										}
									}

									// can be date range or location
									if val.Val.(string) == "LOW" {
										if val, ok := item.KeyVal["text"]; ok {
											// if strings contains " - ", this must be date range, otherwise a location
											if strings.Contains(val.KeyVal["text"].Val.(string), " - ") {
												split := strings.Split(val.KeyVal["text"].Val.(string), " - ")
												if len(split) > 1 {
													date1, err := dateparse.ParseAny(split[0])
													if err == nil {
														work.DateStart = split[0]
														work.DateStartUnix = date1.Unix()
													}

													date2, err := dateparse.ParseAny(split[1])
													if err == nil {
														work.DateEnd = split[1]
														work.DateEndUnix = date2.Unix()
													}
												}

											} else {
												work.Location = val.KeyVal["text"].Val.(string)
											}
										}
									}
								}
							}
						}
					}

					works = append(works, work)
				}
			}

			if val.Val.(string) == "college" {
				if val, ok := section.KeyVal["profile_fields"]; ok {
					if val, ok := val.KeyVal["nodes"]; ok {
						for i, node := range val.Vals {
							if i == 0 {
								continue
							}

							education := Education{
								Title: node.KeyVal["title"].KeyVal["text"].Val.(string),
								Type:  "college",
							}

							// find school url
							if val, ok := node.KeyVal["title"].KeyVal["ranges"]; ok {
								for _, rng := range val.Vals {
									if val, ok := rng.KeyVal["entity"]; ok {
										if val, ok := val.KeyVal["url"]; ok {
											education.SchoolURL = val.Val.(string)
										}
									}
								}
							}

							// find school icon
							if val, ok := node.KeyVal["renderer"]; ok {
								if val, ok := val.KeyVal["field"]; ok {
									if val, ok := val.KeyVal["icon"]; ok {
										education.SchoolIcon = &Photo{
											Height: val.KeyVal["height"].Val.(int),
											Scale:  float64(val.KeyVal["scale"].Val.(int)),
											URI:    val.KeyVal["uri"].Val.(string),
											Width:  val.KeyVal["width"].Val.(int),
										}
									}
								}
							}

							// find date range, degree, concentrations, and description
							if val, ok := node.KeyVal["list_item_groups"]; ok {
								for _, val := range val.Vals {
									if val, ok := val.KeyVal["list_items"]; ok {
										for _, val := range val.Vals {
											heading, ok := val.KeyVal["heading_type"]
											if !ok {
												continue
											}

											headStr := heading.Val.(string)

											// extract date range
											if headStr == "LOW" {
												split := strings.Split(val.KeyVal["text"].KeyVal["text"].Val.(string), " - ")
												if len(split) > 1 {
													date1, err := dateparse.ParseAny(split[0])
													if err == nil {
														education.DateStart = split[0]
														education.DateStartUnix = date1.Unix()
													}

													date2, err := dateparse.ParseAny(split[1])
													if err == nil {
														education.DateEnd = split[1]
														education.DateEndUnix = date2.Unix()
													}
												}
											}

											// extract other info, like degree, concentrations, and description
											if headStr == "MEDIUM" {
												education.OtherInfo = append(education.OtherInfo, val.KeyVal["text"].KeyVal["text"].Val.(string))
											}
										}
									}
								}
							}

							educations = append(educations, education)
						}
					}
				}
			}

			if val.Val.(string) == "secondary_school" {
				if val, ok := section.KeyVal["profile_fields"]; ok {
					if val, ok := val.KeyVal["nodes"]; ok {
						for i, node := range val.Vals {
							if i == 0 {
								continue
							}

							education := Education{
								Title: node.KeyVal["title"].KeyVal["text"].Val.(string),
								Type:  "high_school",
							}

							// find school url
							if val, ok := node.KeyVal["title"].KeyVal["ranges"]; ok {
								for _, rng := range val.Vals {
									if val, ok := rng.KeyVal["entity"]; ok {
										if val, ok := val.KeyVal["url"]; ok {
											education.SchoolURL = val.Val.(string)
										}
									}
								}
							}

							// find school icon
							if val, ok := node.KeyVal["renderer"]; ok {
								if val, ok := val.KeyVal["field"]; ok {
									if val, ok := val.KeyVal["icon"]; ok {
										education.SchoolIcon = &Photo{
											Height: val.KeyVal["height"].Val.(int),
											Scale:  float64(val.KeyVal["scale"].Val.(int)),
											URI:    val.KeyVal["uri"].Val.(string),
											Width:  val.KeyVal["width"].Val.(int),
										}
									}
								}
							}

							// find date range, degree, concentrations, and description
							if val, ok := node.KeyVal["list_item_groups"]; ok {
								for _, val := range val.Vals {
									if val, ok := val.KeyVal["list_items"]; ok {
										for _, val := range val.Vals {
											heading, ok := val.KeyVal["heading_type"]
											if !ok {
												continue
											}

											headStr := heading.Val.(string)

											// extract date range
											if headStr == "LOW" {
												education.ClassOf = val.KeyVal["text"].KeyVal["text"].Val.(string)
											}

											// extract other info, like degree, concentrations, and description
											if headStr == "MEDIUM" {
												education.Description = val.KeyVal["text"].KeyVal["text"].Val.(string)
											}
										}
									}
								}
							}

							educations = append(educations, education)
						}
					}
				}
			}
		}
	}

	return works, educations
}
