package facebook

const (
	sectionFriends = "FRIENDS"
	friendsAll     = "friends_all"
)

// Friends contains retrieved friend list and method for pagination
type Friends struct {
}

// // SyncFriends retrieve importants tokens to retrieve friend list
// func (prof *Profile) SyncFriends() error {
// 	_, rawBody, err := prof.fb.getRequest("/"+prof.ID+"/friends", nil)
// 	if err != nil {
// 		return err
// 	}

// 	jsons, err := jsonextract.FromBytes(rawBody)

// 	if err != nil {
// 		return err
// 	}

// 	// find friends tokens
// 	if !findObj(jsons, func(json *jsonextract.JSON) bool {
// 		if val, ok := json.Object()["require"]; ok {
// 			if findObj(val.Array(), func(json *jsonextract.JSON) bool {
// 				if val, ok := json.Object()["variables"]; ok {
// 					if _, ok := val.Object()["sectionToken"]; ok {
// 						prof.friendSectionVars = json
// 						return true
// 					}
// 				}

// 				return false
// 			}) {
// 				return true
// 			}
// 		}

// 		return false
// 	}) {
// 		return errors.New("Important tokens for Friends section is not found")
// 	}

// 	return nil
// }

// // AllFriends prepare for paging all friends list
// func (prof *Profile) AllFriends() (*Friends, error) {
// 	_, rawBody, err := prof.fb.getRequest("/"+prof.ID+"/friends_all", nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	jsons, err := jsonextract.FromBytes(rawBody)

// 	if err != nil {
// 		return nil, err
// 	}

// 	friends := new(Friends)
// 	jsons, err = prof.reqFriends(friendsAll, "")
// 	if err != nil {
// 		return nil, err
// 	}

// 	findKeyObj(jsons, "pageItems", func(parent, obj *jsonextract.JSON) bool {
// 		return true
// 	})

// 	return friends, nil
// }

// func (prof *Profile) reqFriends(c, cursor string) ([]*jsonextract.JSON, error) {
// 	var section *jsonextract.JSON
// 	for _, val := range prof.profileSections.Object()["edges"].Array() {
// 		node, ok := val.Object()["node"]
// 		if !ok {
// 			continue
// 		}

// 		if val, ok := node.Object()["section_type"]; ok && val.String() == sectionFriends {
// 			section = node
// 			break
// 		}
// 	}

// 	if section == nil {
// 		return nil, errors.New("Important tokens for Friends section is not found")
// 	}

// 	var coll *jsonextract.JSON
// 	for _, val := range section.Object()["all_collections"].Object()["nodes"].Array() {
// 		tabKey, ok := val.Object()["tab_key"]
// 		if !ok {
// 			continue
// 		}

// 		if tabKey.String() == c {
// 			coll = val
// 			break
// 		}
// 	}

// 	vars := prof.friendSectionVars.Object()["variables"]
// 	vars.Object()["collectionToken"].SetStr(coll.Object()["id"].String())

// 	var apiName string
// 	if cursor == "" {
// 		apiName = "ProfileCometTopAppSectionQuery"
// 	} else {
// 		apiName = "ProfileCometAppCollectionListRendererPaginationQuery"
// 		// alter the variables
// 	}

// 	reqBody := make(url.Values)
// 	reqBody.Set("fb_api_req_friendly_name", apiName)
// 	reqBody.Set("variables", string(vars.Bytes()))
// 	reqBody.Set("doc_id", prof.friendSectionVars.Object()["queryID"].String())
// 	_, rawBody, err := prof.fb.graphQlRequest(reqBody)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return jsonextract.FromBytes(rawBody)
// }
