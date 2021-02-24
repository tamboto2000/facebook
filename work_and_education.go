package facebook

import (
	"strings"

	"github.com/araddon/dateparse"
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
	jsons, err := prof.reqAboutCollection(aboutWorkAndEducation)
	if err != nil {
		return err
	}

	// DELETE
	// jsonextract.SaveToPath(jsons, "work_education_bundle.json")

	for _, json := range jsons {
		val, ok := json.Object()["label"]
		if !ok {
			continue
		}

		if val.String() == "ProfileCometAboutAppSectionQuery$defer$ProfileCometAboutAppSectionContent_appSection" {
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

				for i, node := range val.Array() {
					// skip the first index because the first index is a button for add new work history
					if i == 0 {
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

							val, ok = rng.Object()["url"]
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
											// if strings contains " - ", this must be date range, otherwise a location
											if strings.Contains(val.Object()["text"].String(), " - ") {
												split := strings.Split(val.Object()["text"].String(), " - ")
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
												work.Location = val.Object()["text"].String()
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
						for i, node := range val.Array() {
							if i == 0 {
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
												split := strings.Split(val.Object()["text"].Object()["text"].String(), " - ")
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
						for i, node := range val.Array() {
							if i == 0 {
								continue
							}

							education := Education{
								Title: node.Object()["title"].Object()["text"].String(),
								Type:  "high_school",
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
												education.ClassOf = val.Object()["text"].Object()["text"].String()
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
