package facebook

import (
	"errors"
	"net/url"
	"strings"

	"github.com/tamboto2000/jsonextract/v2"
)

// about Section's collections
const (
	workAndEducation       = "about_work_and_education"
	placesLived            = "about_places"
	contactAndBasicInfo    = "about_contact_and_basic_info"
	familyAndRelationships = "about_family_and_relationships"
)

// About contains profile about section data
type About struct {
	WorkHistory            []Work                  `json:"workHistory,omitempty"`
	EducationHistory       []Education             `json:"educationHistories,omitempty"`
	PlacesLived            []Place                 `json:"placesLived,omitempty"`
	ContactAndBasicInfo    *ContactAndBasicInfo    `json:"contactAndBasicInfo,omitempty"`
	FamilyAndRelationships *FamilyAndRelationships `json:"familyAndRelationships,omitempty"`
}

// SyncAbout fetch required tokens for requesting profile about data collections
func (prof *Profile) SyncAbout() error {
	rawBody, err := prof.fb.getRequest("/"+prof.ID+"/about", nil)
	if err != nil {
		return err
	}

	jsons, err := jsonextract.FromBytesWithOpt(rawBody, jsonextract.Option{
		ParseObj:         true,
		ParseArray:       true,
		IgnoreEmptyObj:   true,
		IgnoreEmptyArray: true,
	})

	if err != nil {
		return err
	}

	// find profile about section vars
	if !findObj(jsons, func(json *jsonextract.JSON) bool {
		val, ok := json.KeyVal["require"]
		if !ok {
			return false
		}

		if findObj(val.Vals, func(json *jsonextract.JSON) bool {
			val, ok := json.KeyVal["preloaderID"]
			if !ok {
				return false
			}

			if strings.Contains(val.Val.(string), "adp_ProfileCometAboutAppSectionQueryRelayPreloader") {
				if _, ok = json.KeyVal["variables"]; ok {
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
	for _, val := range prof.profileSections.KeyVal["edges"].Vals {
		node, ok := val.KeyVal["node"]
		if !ok {
			continue
		}

		if val, ok := node.KeyVal["section_type"]; ok && val.Val.(string) == SectionAbout {
			section = node
			break
		}
	}

	if section == nil {
		return nil, errors.New("Important tokens for About section is not found")
	}

	var coll *jsonextract.JSON
	for _, val := range section.KeyVal["all_collections"].KeyVal["nodes"].Vals {
		tabKey, ok := val.KeyVal["tab_key"]
		if !ok {
			continue
		}

		if tabKey.Val.(string) == c {
			coll = val
			break
		}
	}

	vars := prof.aboutSectionVars.KeyVal["variables"]
	vars.KeyVal["collectionToken"].Val = coll.KeyVal["id"].Val
	if err := vars.ReParse(); err != nil {
		return nil, err
	}

	reqBody := make(url.Values)
	reqBody.Set("fb_api_req_friendly_name", "ProfileCometAboutAppSectionQuery")
	reqBody.Set("variables", string(vars.Raw.Bytes()))
	reqBody.Set("doc_id", prof.aboutSectionVars.KeyVal["queryID"].Val.(string))
	rawBody, err := prof.fb.graphQlRequest(reqBody)
	if err != nil {
		return nil, err
	}

	return jsonextract.FromBytes(rawBody)
}
