package facebook

import (
	"github.com/tamboto2000/jsonextract/v3"
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
	Title     string `json:"title,omitempty"`
	SchoolURL string `json:"schoolUrl,omitempty"`
	// start and end date can't be determined
	DateRange     []string `json:"dateRange,omitempty"`
	DateRangeUnix []int64  `json:"dateRangeUnix,omitempty"`
	OtherInfo     []string `json:"otherInfo,omitempty"`
	Description   string   `json:"description,omitempty"`
	SchoolIcon    *Photo   `json:"schoolIcon,omitempty"`
	// college or secondary_school
	Type string `json:"type,omitempty"`
}

// SyncWorkAndEducation retrieve Profile's work/occupation history and education history
func (about *About) SyncWorkAndEducation() error {
	jsons, err := about.profile.reqAboutCollection(aboutWorkAndEducation)
	if err != nil {
		return err
	}

	for _, json := range jsons {
		val, ok := json.Object()["label"]
		if !ok {
			continue
		}

		if val.String() == "ProfileCometAboutAppSectionQuery$defer$ProfileCometAboutAppSectionContent_appSection" {
			works, educations := extractWorks(json)
			about.WorkHistory = works
			about.EducationHistory = educations
			break
		}
	}

	return nil
}

func extractWorks(json *jsonextract.JSON) ([]Work, []Education) {
	works := make([]Work, 0)
	educations := make([]Education, 0)
	val, ok := json.Object()["data"]
	if !ok {
		return nil, nil
	}

	val, ok = val.Object()["activeCollections"]
	if !ok {
		return nil, nil
	}

	val, ok = val.Object()["nodes"]
	if !ok {
		return nil, nil
	}

	for _, node := range val.Array() {
		val, ok := node.Object()["style_renderer"]
		if !ok {
			continue
		}

		val, ok = val.Object()["profile_field_sections"]
		if !ok {
			continue
		}

		for _, section := range val.Array() {
			val, ok := section.Object()["field_section_type"]
			if !ok {
				continue
			}

			if val.String() == "work" {
				// start parsing work history
				val, ok := section.Object()["profile_fields"]
				if !ok {
					return nil, nil
				}

				val, ok = val.Object()["nodes"]
				if !ok {
					return nil, nil
				}

				for _, node := range val.Array() {
					if node.Object()["field_type"].String() != "work" {
						continue
					}

					// create Work and assign Title
					work := Work{
						Title: node.Object()["title"].Object()["text"].String(),
					}

					// find company url
					if val, ok := node.Object()["title"].Object()["ranges"]; ok {
						for _, rng := range val.Array() {
							val, ok := rng.Object()["entity"]
							if !ok {
								continue
							}

							val, ok = val.Object()["url"]
							if !ok {
								continue
							}

							work.CompanyURL = val.String()
						}
					}

					// find company icon
					if val, ok := node.Object()["renderer"]; ok {
						if val, ok := val.Object()["field"]; ok {
							if val, ok := val.Object()["icon"]; ok {
								work.CompanyIcon = &Photo{
									Height: int(val.Object()["height"].Integer()),
									URI:    val.Object()["uri"].String(),
									Width:  int(val.Object()["width"].Integer()),
								}

								scale := val.Object()["scale"]
								if scale.Kind() == jsonextract.Integer {
									work.CompanyIcon.Scale = float64(scale.Integer())
								} else {
									work.CompanyIcon.Scale = scale.Float()
								}
							}
						}
					}

					// find description
					if val, ok := node.Object()["list_item_groups"]; ok {
						for _, itemGroup := range val.Array() {
							if val, ok := itemGroup.Object()["list_items"]; ok {
								for _, item := range val.Array() {
									val, ok := item.Object()["heading_type"]
									if !ok {
										continue
									}

									// description
									if val.String() == "MEDIUM" {
										if val, ok := item.Object()["text"]; ok {
											work.Description = val.Object()["text"].String()
										}
									}

									// can be date range or location
									if val.String() == "LOW" {
										if val, ok := item.Object()["text"]; ok {
											// find date range with regex
											// if non found, then this must be location, hopefully...
											text := val.Object()["text"].String()
											dates, datesStr := extractDate(text)

											if len(dates) == 0 && len(datesStr) == 0 {
												work.Location = text
											} else {
												if len(dates) == 2 {
													var dateStartIdx int
													var dateEndIdx int

													if dates[0].Unix() < dates[1].Unix() {
														dateStartIdx = 0
														dateEndIdx = 1
													} else {
														dateStartIdx = 1
														dateEndIdx = 0
													}

													work.DateStart = datesStr[dateStartIdx]
													work.DateStartUnix = dates[dateStartIdx].Unix()

													work.DateEnd = datesStr[dateEndIdx]
													work.DateEndUnix = dates[dateEndIdx].Unix()
												} else {
													work.DateStart = datesStr[0]
													work.DateStartUnix = dates[0].Unix()
												}
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

			if val.String() == "college" {
				if val, ok := section.Object()["profile_fields"]; ok {
					if val, ok := val.Object()["nodes"]; ok {
						for _, node := range val.Array() {
							if node.Object()["field_type"].String() != "education" {
								continue
							}

							education := Education{
								Title: node.Object()["title"].Object()["text"].String(),
								Type:  "college",
							}

							// find school url
							if val, ok := node.Object()["title"].Object()["ranges"]; ok {
								for _, rng := range val.Array() {
									if val, ok := rng.Object()["entity"]; ok {
										if val, ok := val.Object()["url"]; ok {
											education.SchoolURL = val.String()
										}
									}
								}
							}

							// find school icon
							if val, ok := node.Object()["renderer"]; ok {
								if val, ok := val.Object()["field"]; ok {
									if val, ok := val.Object()["icon"]; ok {
										education.SchoolIcon = &Photo{
											Height: int(val.Object()["height"].Integer()),
											URI:    val.Object()["uri"].String(),
											Width:  int(val.Object()["width"].Integer()),
										}

										scale := val.Object()["scale"]
										if scale.Kind() == jsonextract.Integer {
											education.SchoolIcon.Scale = float64(scale.Integer())
										} else {
											education.SchoolIcon.Scale = scale.Float()
										}
									}
								}
							}

							// find date range, degree, concentrations, and description
							if val, ok := node.Object()["list_item_groups"]; ok {
								for _, val := range val.Array() {
									if val, ok := val.Object()["list_items"]; ok {
										for _, val := range val.Array() {
											heading, ok := val.Object()["heading_type"]
											if !ok {
												continue
											}

											headStr := heading.String()

											// extract date range
											if headStr == "LOW" {
												text := val.Object()["text"].Object()["text"].String()
												dates, datesStr := extractDate(text)
												education.DateRange = datesStr
												for _, date := range dates {
													education.DateRangeUnix = append(education.DateRangeUnix, date.Unix())
												}
											}

											// extract other info, like degree, concentrations, and description
											if headStr == "MEDIUM" {
												education.OtherInfo = append(education.OtherInfo, val.Object()["text"].Object()["text"].String())
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

			if val.String() == "secondary_school" {
				if val, ok := section.Object()["profile_fields"]; ok {
					if val, ok := val.Object()["nodes"]; ok {
						for _, node := range val.Array() {
							if node.Object()["field_type"].String() != "education" {
								continue
							}

							education := Education{
								Title: node.Object()["title"].Object()["text"].String(),
								Type:  "secondary_school",
							}

							// find school url
							if val, ok := node.Object()["title"].Object()["ranges"]; ok {
								for _, rng := range val.Array() {
									if val, ok := rng.Object()["entity"]; ok {
										if val, ok := val.Object()["url"]; ok {
											education.SchoolURL = val.String()
										}
									}
								}
							}

							// find school icon
							if val, ok := node.Object()["renderer"]; ok {
								if val, ok := val.Object()["field"]; ok {
									if val, ok := val.Object()["icon"]; ok {
										education.SchoolIcon = &Photo{
											Height: int(val.Object()["height"].Integer()),
											URI:    val.Object()["uri"].String(),
											Width:  int(val.Object()["width"].Integer()),
										}

										scale := val.Object()["scale"]
										if scale.Kind() == jsonextract.Integer {
											education.SchoolIcon.Scale = float64(scale.Integer())
										} else {
											education.SchoolIcon.Scale = scale.Float()
										}
									}
								}
							}

							// find date range, degree, concentrations, and description
							if val, ok := node.Object()["list_item_groups"]; ok {
								for _, val := range val.Array() {
									if val, ok := val.Object()["list_items"]; ok {
										for _, val := range val.Array() {
											heading, ok := val.Object()["heading_type"]
											if !ok {
												continue
											}

											headStr := heading.String()

											// extract date range
											if headStr == "LOW" {
												text := val.Object()["text"].Object()["text"].String()
												dates, datesStr := extractDate(text)
												education.DateRange = datesStr
												for _, date := range dates {
													education.DateRangeUnix = append(education.DateRangeUnix, date.Unix())
												}
											}

											// extract other info, like degree, concentrations, and description
											if headStr == "MEDIUM" {
												education.Description = val.Object()["text"].Object()["text"].String()
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
