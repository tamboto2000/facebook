package facebook

import (
	"encoding/base64"
	"errors"
	"strings"

	"github.com/tamboto2000/facebook/raw"
)

const (
	AllFriends         = "friends_all"
	RecentFriends      = "friends_recent"
	WorkFriends        = "friends_work"
	HighSchoolFriends  = "friends_high_school"
	CurrentCityFriends = "friends_current_city"
	Following          = "following"
	Followers          = "followers"
)

type Friends struct {
	UserID                                               string  `json:"userId"`
	ProfileCometTopAppSectionQuery                       string  `json:"profileCometTopAppSectionQuery"`
	ProfileCometAppCollectionListRendererPaginationQuery string  `json:"profileCometAppCollectionListRendererPaginationQuery"`
	Users                                                []*User `json:"users"`
	CollectionToken                                      string  `json:"collectionToken"`
	SectionToken                                         string  `json:"sectionToken"`
	RawSectionKey                                        string  `json:"rawSectionKey"`
	Cursor                                               string  `json:"cursor"`

	isEndOfList bool
	fb          *Facebook
	err         error
}

func (user *User) Friends(coll string) (*Friends, error) {
	if coll == "" {
		coll = AllFriends
	}

	if coll != AllFriends && coll != RecentFriends &&
		coll != WorkFriends && coll != HighSchoolFriends &&
		coll != CurrentCityFriends {
		return nil, errors.New("collection not valid")
	}

	var collToken string
	section := user.userSections["FRIENDS"]
	for _, node := range section.Node.AllCollections.Nodes {
		if strings.Contains(node.URL, coll) {
			collToken = node.ID
			break
		}
	}

	sectionKey, err := base64.StdEncoding.DecodeString(user.userSections["FRIENDS"].Node.ID)
	if err != nil {
		return nil, err
	}

	sectionKeyRaw := strings.Replace(string(sectionKey), "app_section:", "", 1)

	return &Friends{
		UserID:                         user.ID,
		CollectionToken:                collToken,
		SectionToken:                   section.Node.ID,
		RawSectionKey:                  sectionKeyRaw,
		ProfileCometTopAppSectionQuery: user.profileCometTopAppSectionQuery,
		ProfileCometAppCollectionListRendererPaginationQuery: user.profileCometAppCollectionListRendererPaginationQuery,
		fb: user.fb,
	}, nil
}

func (fr *Friends) IsEndOfList() bool {
	return fr.isEndOfList
}

func (fr *Friends) Next(count int) bool {
	fr.Users = make([]*User, 0)

	if fr.isEndOfList {
		return false
	}

	if count > 50 {
		count = 50
	}

	if count <= 0 {
		count = 50
	}

	docID := fr.ProfileCometTopAppSectionQuery
	apiName := "ProfileCometTopAppSectionQuery"
	vars := map[string]interface{}{
		"collectionToken": fr.CollectionToken,
		"scale":           1,
		"sectionToken":    fr.SectionToken,
		"userID":          fr.UserID,
	}

	if fr.Cursor != "" {
		vars = map[string]interface{}{
			"count":  count,
			"cursor": fr.Cursor,
			"search": nil,
			"scale":  1,
			"id":     fr.CollectionToken,
		}

		vars["cursor"] = fr.Cursor
		vars["count"] = count
		docID = fr.ProfileCometAppCollectionListRendererPaginationQuery
		apiName = "ProfileCometAppCollectionListRendererPaginationQuery"
	}

	payloads, err := fr.fb.doGraphQLRequest(vars, docID, apiName, false)
	if err != nil {
		fr.err = err
		return false
	}

	payload := payloads[0]
	var items *raw.Items
	//get items
	if fr.Cursor == "" {
		if len(payload.Data.User.TimelineNavAppSections.Nodes[0].AllCollections.Nodes) == 0 {
			fr.isEndOfList = true
			return false
		}

		items = payload.Data.User.TimelineNavAppSections.Nodes[0].AllCollections.Nodes[0].StyleRenderer.Collection.Items
	} else {
		items = payload.Data.Node.Items
	}
	//get cursor
	pageInfo := items.PageInfo
	if pageInfo.HasNextPage {
		fr.Cursor = pageInfo.EndCursor
	} else {
		fr.isEndOfList = true
	}

	//extract users
	for _, edge := range items.Edges {
		user := &User{
			ID:             edge.Node.Node.ID,
			Name:           edge.Node.Title.Text,
			URL:            edge.Node.URL,
			ProfilePictURL: edge.Node.Image.URI,
		}

		fr.Users = append(fr.Users, user)
	}

	return true
}

func (fr *Friends) Error() error {
	return fr.err
}

func (fr *Friends) inject(fb *Facebook) {
	fr.fb = fb
}
