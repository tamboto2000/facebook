package facebook

import (
	"encoding/base64"
	"strings"
	"time"
)

type Job struct {
	WorkPlaceName    string `json:"workPlaceName"`
	WorkPlaceFBLink  string `json:"workPlaceFBLink"`
	WorkPlaceLogoURL string `json:"workPlaceLogoURL"`
	Position         string `json:"position"`
	StartDate        string `json:"startDate"`
	EndDate          string `json:"endDate"`
	Location         string `json:"location"`
	Description      string `json:"description"`
}

type School struct {
	SchoolType    string `json:"schoolType"`
	SchoolName    string `json:"schoolName"`
	FieldOfStudy  string `json:"fieldOfStudy"`
	SchoolFBLink  string `json:"schoolFBLink"`
	SchoolLogoURL string `json:"schoolLogoURL"`
	StartDate     string `json:"startDate"`
	EndDate       string `json:"endDate"`
	ClassOf       string `json:"classOf"`
	Location      string `json:"location"`
}

func (user *User) SyncEducationAndWork() error {
	return user.reqWorkAndEducation()
}

func (user *User) reqWorkAndEducation() error {
	sectionKey, err := base64.StdEncoding.DecodeString(user.userSections["ABOUT"].Node.ID)
	if err != nil {
		return err
	}

	sectionKeyRaw := strings.Replace(string(sectionKey), "app_section:", "", 1)
	var collectionToken string
	for _, node := range user.userSections["ABOUT"].Node.AllCollections.Nodes {
		if strings.Contains(node.URL, "/about_work_and_education") {
			collectionToken = node.ID
			break
		}
	}

	vars := map[string]interface{}{
		"appSectionFeedKey":      "ProfileCometAppSectionFeed_timeline_nav_app_sections__" + sectionKeyRaw,
		"collectionToken":        collectionToken,
		"rawSectionToken":        sectionKeyRaw,
		"scale":                  1,
		"sectionToken":           user.userSections["ABOUT"].Node.ID,
		"userID":                 user.ID,
		"useIncrementalDelivery": true,
	}

	payloads, err := user.fb.doGraphQLRequest(vars, user.aboutDocID, "ProfileCometAboutAppSectionQuery", true)
	if err != nil {
		return err
	}

	//extract dates
	dateFormat1 := "January 2, 2006"
	dateFormat2 := "January 2006"
	dateFormat3 := "January 2"
	for _, payload := range payloads {
		if payload.Data.ActiveCollections.Nodes == nil || len(payload.Data.ActiveCollections.Nodes) == 0 {
			continue
		}

		for _, node := range payload.Data.ActiveCollections.Nodes {
			for _, section := range node.StyleRenderer.ProfileFieldSections {
				if section.SectionType == "work" {
					for _, field := range section.ProfileFields.Nodes {
						var position string
						var placeName string
						var startDate string
						var endDate string
						var location string
						var description string

						//extract place name
						if len(field.Renderer.Field.Title.Ranges) > 0 {
							offset := field.Renderer.Field.Title.Ranges[0].Offset
							length := field.Renderer.Field.Title.Ranges[0].Length
							text := strings.Split(field.Renderer.Field.Title.Text, "")

							for i := offset; i < offset+length; i++ {
								placeName += text[i]
							}

							for i := 0; i < int(offset)-3; i++ {
								position += text[i]
							}

							position = strings.TrimSpace(position)
						}

						for _, itemGroup := range field.Renderer.Field.ListItemGroups {
							for _, item := range itemGroup.ListItems {
								if strings.Contains(item.Text.Text, " - ") && item.HeadingType == "LOW" {
									stg1 := strings.Split(item.Text.Text, " - ")
									startdate, err := time.Parse(dateFormat1, stg1[0])
									if err == nil {
										startDate = startdate.Format(dateFormat1)
									} else {
										startdate, err := time.Parse(dateFormat2, stg1[0])
										if err == nil {
											startDate = startdate.Format(dateFormat2)
										} else {
											startdate, err := time.Parse(dateFormat3, stg1[0])
											if err == nil {
												startDate = startdate.Format(dateFormat3)
											}
										}
									}

									enddate, err := time.Parse(dateFormat1, stg1[1])
									if err == nil {
										endDate = enddate.Format(dateFormat1)
									} else {
										enddate, err := time.Parse(dateFormat2, stg1[1])
										if err == nil {
											endDate = enddate.Format(dateFormat2)
										} else {
											enddate, err := time.Parse(dateFormat3, stg1[1])
											if err == nil {
												endDate = enddate.Format(dateFormat3)
											}
										}
									}
								} else {
									if item.HeadingType == "LOW" {
										location = item.Text.Text
									}
								}

								if item.HeadingType == "MEDIUM" {
									description = item.Text.Text
								}
							}
						}

						var pageURL string
						if len(field.Renderer.Field.Title.Ranges) > 0 {
							pageURL = field.Renderer.Field.Title.Ranges[0].Entity.ProfileURL
						}

						user.WorkExperiences = append(user.WorkExperiences, Job{
							WorkPlaceName:    placeName,
							Position:         position,
							WorkPlaceFBLink:  pageURL,
							WorkPlaceLogoURL: field.Renderer.Field.Icon.URI,
							StartDate:        startDate,
							EndDate:          endDate,
							Location:         location,
							Description:      description,
						})
					}
				}
				if section.SectionType != "work" {
					for _, field := range section.ProfileFields.Nodes {
						var name string
						var fieldOfStudy string
						var pageURL string
						var startDate string
						var endDate string
						var classOf string
						//extract school name and field of study
						if len(field.Renderer.Field.Title.Ranges) > 0 {
							offset := field.Renderer.Field.Title.Ranges[0].Offset
							length := field.Renderer.Field.Title.Ranges[0].Length
							text := strings.Split(field.Renderer.Field.Title.Text, "")

							for i := offset; i < offset+length; i++ {
								name += text[i]
							}

							if section.SectionType == "college" {
								for i := 0; i < int(offset)-3; i++ {
									fieldOfStudy += text[i]
								}
							}

							pageURL = field.Renderer.Field.Title.Ranges[0].Entity.ProfileURL

							fieldOfStudy = strings.TrimSpace(fieldOfStudy)
						}

						if len(field.Renderer.Field.ListItemGroups) > 0 {
							for _, itemGroup := range field.Renderer.Field.ListItemGroups {
								for _, item := range itemGroup.ListItems {
									if !strings.Contains(item.Text.Text, " - ") && section.SectionType == "secondary_school" {
										classOf = item.Text.Text
										continue
									}

									stg1 := strings.Split(item.Text.Text, " - ")
									startdate, err := time.Parse(dateFormat1, stg1[0])
									if err == nil {
										startDate = startdate.Format(dateFormat1)
									} else {
										startdate, err := time.Parse(dateFormat2, stg1[0])
										if err == nil {
											startDate = startdate.Format(dateFormat2)
										} else {
											startdate, err := time.Parse(dateFormat3, stg1[0])
											if err == nil {
												startDate = startdate.Format(dateFormat3)
											}
										}
									}

									if len(stg1) > 1 {
										enddate, err := time.Parse(dateFormat1, stg1[1])
										if err == nil {
											endDate = enddate.Format(dateFormat1)
										} else {
											enddate, err := time.Parse(dateFormat2, stg1[1])
											if err == nil {
												endDate = enddate.Format(dateFormat2)
											} else {
												enddate, err := time.Parse(dateFormat3, stg1[1])
												if err == nil {
													endDate = enddate.Format(dateFormat3)
												}
											}
										}
									}
								}
							}
						}

						if name == "" {
							continue
						}

						user.Educations = append(user.Educations, School{
							SchoolName:    name,
							FieldOfStudy:  fieldOfStudy,
							SchoolFBLink:  pageURL,
							SchoolLogoURL: field.Renderer.Field.Icon.URI,
							StartDate:     startDate,
							EndDate:       endDate,
							ClassOf:       classOf,
						})
					}
				}
			}
		}
	}

	if len(user.WorkExperiences) == 0 {
		user.WorkExperiences = nil
	}

	if len(user.Educations) == 0 {
		user.Educations = nil
	}

	return nil
}
