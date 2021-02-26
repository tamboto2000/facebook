package facebook

import (
	"errors"
	"net/url"
	"strings"

	"github.com/tamboto2000/jsonextract/v3"
)

// about Section's collections
const (
	aboutWorkAndEducation       = "about_work_and_education"
	aboutPlacesLived            = "about_places"
	aboutContactAndBasicInfo    = "about_contact_and_basic_info"
	aboutFamilyAndRelationships = "about_family_and_relationships"
	aboutDetails                = "about_details"
	aboutLifeEvents             = "about_life_events"
)

// About contains profile about section data
type About struct {
	WorkHistory            []Work                  `json:"workHistory,omitempty"`
	EducationHistory       []Education             `json:"educationHistory,omitempty"`
	PlacesLived            []Place                 `json:"placesLived,omitempty"`
	ContactAndBasicInfo    *ContactAndBasicInfo    `json:"contactAndBasicInfo,omitempty"`
	FamilyAndRelationships *FamilyAndRelationships `json:"familyAndRelationships,omitempty"`
	Details                *Details                `json:"details,omitempty"`
	LifeEvents             []LifeEvents            `json:"lifeEvents,omitempty"`
}

// SyncAbout fetch required tokens for requesting profile about data collections
func (prof *Profile) SyncAbout() error {
	var handle string
	if prof.Username != "" {
		handle = prof.Username
	} else {
		handle = prof.ID
	}

	_, rawBody, err := prof.fb.getRequest("/"+handle+"/about", nil)
	if err != nil {
		return err
	}

	jsons, err := jsonextract.FromBytes(rawBody)

	if err != nil {
		return err
	}	

	// find profile about section vars
	if !findObj(jsons, func(json *jsonextract.JSON) bool {
		obj := json.Object()
		val, ok := obj["require"]
		if !ok {
			return false
		}

		if val.Kind() != jsonextract.Array {
			return false
		}

		if findObj(val.Array(), func(json *jsonextract.JSON) bool {
			obj := json.Object()
			val, ok := obj["preloaderID"]
			if !ok {
				return false
			}

			if val.Kind() != jsonextract.String {
				return false
			}

			if strings.Contains(val.String(), "adp_ProfileCometAboutAppSectionQueryRelayPreloader") {
				if _, ok = obj["variables"]; ok {
					prof.aboutSectionVars = json
					return true
				}
			}

			return false
		}) {
			return true
		}

		return false
	}) {
		return errors.New("Important tokens for About section is not found")
	}

	prof.About = new(About)

	return nil
}

func (prof *Profile) reqAboutCollection(c string) ([]*jsonextract.JSON, error) {
	var section *jsonextract.JSON
	for _, val := range prof.profileSections.Object()["edges"].Array() {
		node, ok := val.Object()["node"]
		if !ok {
			continue
		}

		if val, ok := node.Object()["section_type"]; ok && val.String() == SectionAbout {
			section = node
			break
		}
	}

	if section == nil {
		return nil, errors.New("Important tokens for About section is not found")
	}

	var coll *jsonextract.JSON
	for _, val := range section.Object()["all_collections"].Object()["nodes"].Array() {
		tabKey, ok := val.Object()["tab_key"]
		if !ok {
			continue
		}

		if tabKey.String() == c {
			coll = val
			break
		}
	}

	vars := prof.aboutSectionVars.Object()["variables"]
	vars.Object()["collectionToken"].SetStr(coll.Object()["id"].String())

	reqBody := make(url.Values)
	reqBody.Set("fb_api_req_friendly_name", "ProfileCometAboutAppSectionQuery")
	reqBody.Set("variables", string(vars.Bytes()))
	reqBody.Set("doc_id", prof.aboutSectionVars.Object()["queryID"].String())
	rawBody, err := prof.fb.graphQlRequest(reqBody)
	if err != nil {
		return nil, err
	}

	return jsonextract.FromBytes(rawBody)
}
