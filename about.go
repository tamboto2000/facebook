package facebook

import "github.com/tamboto2000/jsonextract"

// SyncAbout prepare Profile to fetch data collections on profile's about section
func (prof *Profile) SyncAbout() error {
	rawBody, err := prof.fb.getRequest("/"+prof.ID+"/about", nil)
	if err != nil {
		return err
	}

	// DELETE
	jsons, err := jsonextract.JSONFromBytes(rawBody)
	if err != nil {
		return err
	}
	jsonextract.SaveToPath(jsons, "raw_about.json")

	return nil
}
