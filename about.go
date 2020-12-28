package facebook

// About Section's collections
const (
	WorkAndEducation = "about_work_and_education"
)

type About struct {
	WorkHistory []Work `json:"workHistory,omitempty"`
}

// SyncAbout prepare Profile to fetch data collections on profile's about section
func (prof *Profile) SyncAbout() error {
	// rawBody, err := prof.fb.getRequest("/"+prof.ID+"/about", nil)
	// if err != nil {
	// 	return err
	// }

	// jsons, err := jsonextract.JSONFromBytes(rawBody)
	// if err != nil {
	// 	return err
	// }

	// // // DELETE
	// // jsonextract.SaveToPath(jsons, "raw_about.json")

	// parser := newParser(jsons)
	// parser.run(func(val interface{}) bool {
	// 	item := val.(*raw.Item)
	// 	if item.Variables != nil && strings.Contains(item.PreloaderID, "adp_ProfileCometAboutAppSectionQueryRelayPreloader") {
	// 		prof.AboutSectionVars = item
	// 	}

	// 	return false
	// }, new(raw.Item), true, false)

	// prof.About = new(About)

	return nil
}
