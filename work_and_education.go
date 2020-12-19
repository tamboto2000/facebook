package facebook

// Work contains Profile's work/occupation history
type Work struct {
}

// WorkAndEducation retrieve Profile's work/occupation history and education history
func (prof *Profile) WorkAndEducation() ([]Work, error) {
	// var node *raw.Node
	// for _, edge := range prof.ProfileSections.Edges {
	// 	if edge.Node != nil {
	// 		if edge.Node.SectionType == About {
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

	// rawSectionToken, err := base64.StdEncoding.DecodeString(node.ID)
	// if err != nil {
	// 	return nil, err
	// }

	// rawSectionToken = bytes.ReplaceAll(rawSectionToken, []byte("app_section:"), []byte{})

	// graphQlVars := GraphQLVars{
	// 	AppSectionFeedKey: "ProfileCometAppSectionFeed_timeline_nav_app_sections__" + string(rawSectionToken),
	// 	CollectionToken:   collection.ID,
	// 	RawSectionToken:   string(rawSectionToken),
	// 	Scale:             fmt.Sprintf("%g", prof.Variables.Scale),
	// 	SectionToken:      node.ID,
	// 	UserID:            prof.ID,
	// }

	// graphQlbyte, err := json.Marshal(graphQlVars)
	// if err != nil {
	// 	return nil, err
	// }

	// reqBody := make(url.Values)
	// reqBody.Set("fb_api_req_friendly_name", "ProfileCometAboutAppSectionQuery")

	return nil, nil
}
