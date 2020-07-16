package facebook

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"
	"sync"

	"github.com/tamboto2000/facebook/raw"
	"github.com/tamboto2000/htmltojson"
)

type UserVariables struct {
	UserID                string              `json:"userId"`
	Username              string              `json:"username"`
	AboutDocID            string              `json:"aboutDocID"`
	FriendsDocID          string              `json:"friendsDocId"`
	FriendsListDocID      string              `json:"friendsListDocId"`
	ReactionDialogQueryID string              `json:"reactionDialogQueryId"`
	CommentQueryID        string              `json:"commentQueryId"`
	TimelineQueryID       string              `json:"timelineQueryId"`
	UserSections          map[string]raw.Edge `json:"userSections"`
}

type User struct {
	ID                              string   `json:"id,omitempty"`
	Username                        string   `json:"username,omitempty"`
	Name                            string   `json:"name,omitempty"`
	URL                             string   `json:"url,omitempty"`
	ProfilePictURL                  string   `json:"profilePictURL,omitempty"`
	CoverPhotoURL                   string   `json:"coverPhotoUrl,omitempty"`
	Bio                             string   `json:"bio,omitempty"`
	Quote                           string   `json:"quote,omitempty"`
	WorkExperiences                 []Job    `json:"workExperiences,omitempty"`
	Educations                      []School `json:"educations,omitempty"`
	CurrentCity                     string   `json:"currentCity,omitempty"`
	Hometown                        string   `json:"hometown,omitempty"`
	Address                         string   `json:"address,omitempty"`
	District                        string   `json:"district,omitempty"`
	Phone                           string   `json:"phone,omitempty"`
	FBLink                          string   `json:"fbLink,omitempty"`
	Instagram                       string   `json:"instagram,omitempty"`
	Links                           []string `json:"links,omitempty"`
	DOB                             string   `json:"dob,omitempty"`
	Gender                          string   `json:"gender,omitempty"`
	SexualPreference                string   `json:"sexualPreference,omitempty"`
	Languages                       []string `json:"languages,omitempty"`
	Religion                        string   `json:"religion,omitempty"`
	RelationshipStatus              string   `json:"relationshipStatus,omitempty"`
	PoliticalView                   string   `json:"politicalView,omitempty"`
	TotalFriendCount                int      `json:"totalFriendCount,omitempty"`
	RecentFriendCount               int      `json:"recentFriendCount,omitempty"`
	FriendWithUpcomingBirthdayCount int      `json:"friendWithUpcomingBirthdayCount,omitempty"`
	CoworkerFriendCount             int      `json:"coworkerFriendCount,omitempty"`
	CollegeFriendCount              int      `json:"collegeFriendCount,omitempty"`
	HighSchoolFriendCount           int      `json:"highSchoolFriendCount,omitempty"`
	CurrentCityFriendCount          int      `json:"currentCityFriendCount,omitempty"`
	HometownFriendCount             int      `json:"hometownFriendCount,omitempty"`
	FollowerCount                   int      `json:"followerCount,omitempty"`
	FollowingCount                  int      `json:"followingCount,omitempty"`
	FamilyMember                    []Family `json:"familyMember,omitempty"`
	Events                          []Event  `json:"events,omitempty"`
	Places                          []Place  `json:"places,omitempty"`
	Contact                         *Contact `json:"contact,omitempty"`

	aboutDocID            string
	userSections          map[string]raw.Edge
	fb                    *Facebook
	reactionDialogQueryID string

	//can be used for listing friends list
	profileCometTopAppSectionQuery                       string
	profileCometAppCollectionListRendererPaginationQuery string
}

func (fb *Facebook) NewUser(username string) (*User, error) {
	username = strings.ToLower(username)
	username = strings.TrimSpace(username)
	user := new(User)

	user.Username = username
	user.userSections = make(map[string]raw.Edge)
	user.fb = fb
	if err := user.reqUserAboutPage(); err != nil {
		return nil, err
	}

	if user.ID == "" {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (fb *Facebook) UserFromVariables(v *UserVariables) (*User, error) {
	user := &User{
		ID:                    v.UserID,
		Username:              v.Username,
		aboutDocID:            v.AboutDocID,
		userSections:          v.UserSections,
		reactionDialogQueryID: v.ReactionDialogQueryID,
		fb:                    fb,
	}

	if err := user.reqUserAboutPage(); err != nil {
		return nil, err
	}

	if user.ID == "" {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (user *User) ExportVariables() *UserVariables {
	return &UserVariables{
		UserID:                user.ID,
		Username:              user.Username,
		AboutDocID:            user.aboutDocID,
		UserSections:          user.userSections,
		ReactionDialogQueryID: user.reactionDialogQueryID,
	}
}

func (user *User) reqUserAboutPage() error {
	urlParsed, _ := url.Parse(user.fb.RootURL + "/" + user.Username + "/about")
	resp, err := user.fb.doBasicGetRequest(urlParsed)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	buff, err := decompressResponseBody(resp)
	if err != nil {
		return err
	}

	if resp.StatusCode > 200 {
		if resp.StatusCode == 404 {
			return errors.New("user not found")
		}

		return errors.New("invalid session")
	}

	body := buff.Bytes()

	wg := new(sync.WaitGroup)
	wg.Add(4)

	//extract about section query ID
	jsons, err := extractJSONBytes(body)
	if err != nil {
		return err
	}

	go func() {
		user.extractAboutQueryID(jsons)
		wg.Done()
	}()

	//extract user sections
	go func() {
		user.extractUserSections(jsons)
		wg.Done()
	}()

	//extract user full name, username, user ID, profile pict, and cover image
	go func() {
		user.extractUserBasic(jsons)
		wg.Done()
	}()

	//extract friends query id
	rootNode, _ := htmltojson.ParseString(string(body))
	go func() {
		user.extractProfileCometTopAppSectionQuery(rootNode)
		wg.Done()
	}()

	wg.Wait()
	return nil
}

func (user *User) extractAboutQueryID(jsons [][]byte) {
	for _, str := range jsons {
		payload := new(raw.QueryID)
		if err := json.Unmarshal(str, payload); err != nil {
			continue
		}

		if strings.Contains(payload.PreloaderID, "adp_ProfileCometAboutAppSectionQueryRelayPreloader") {
			user.aboutDocID = payload.QueryID
			return
		}
	}
}

func (user *User) extractUserSections(jsons [][]byte) {
	for _, str := range jsons {
		payload := make(map[string][][]json.RawMessage)
		if err := json.Unmarshal(str, &payload); err != nil {
			continue
		}

		if _, exist := payload["require"]; !exist {
			continue
		}

		for _, a := range payload["require"] {
			for _, b := range a {
				c := make([]json.RawMessage, 0)
				if err := json.Unmarshal(b, &c); err != nil {
					continue
				}

				for _, d := range c {
					e := new(raw.RequireBbox)
					if err := json.Unmarshal(d, e); err != nil {
						continue
					}

					for _, f := range e.Bbox.Result.Data.TimelineNavAppSections.Edges {
						user.userSections[f.Node.SectionType] = f
						if f.Node.SectionType == "FRIENDS" {
							if f.Node.DisplayableCount > 0 {
								user.TotalFriendCount = f.Node.DisplayableCount
							}

							for _, g := range f.Node.AllCollections.Nodes {
								if g.Items != nil {
									if strings.Contains(g.URL, "/friends_all") {
										if g.Items.Count > 0 {
											user.TotalFriendCount = int(g.Items.Count)
										}

										continue
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func (user *User) extractUserBasic(jsons [][]byte) {
	for _, str := range jsons {
		payload := make(map[string][][]json.RawMessage)
		if err := json.Unmarshal(str, &payload); err != nil {
			continue
		}

		if _, exist := payload["require"]; !exist {
			continue
		}

		for _, a := range payload["require"] {
			for _, b := range a {
				c := make([]json.RawMessage, 0)
				if err := json.Unmarshal(b, &c); err != nil {
					continue
				}

				for _, d := range c {
					e := new(raw.RequireBbox)
					if err := json.Unmarshal(d, e); err != nil {
						continue
					}

					rawUser := e.Bbox.Result.Data.User
					if rawUser.Name != "" {
						user.Name = rawUser.Name
					}

					if rawUser.ID != "" {
						user.ID = rawUser.ID
					}

					if rawUser.CoverPhoto.Photo.Image.URI != "" {
						user.CoverPhotoURL = rawUser.CoverPhoto.Photo.Image.URI
					}

					if rawUser.ProfilePic160.URI != "" {
						user.ProfilePictURL = rawUser.ProfilePic160.URI
					}

					if user.ID != "" && user.Name != "" &&
						user.CoverPhotoURL != "" && user.ProfilePictURL != "" {
						return
					}
				}
			}
		}
	}
}

func (user *User) extractProfileCometTopAppSectionQuery(rootNode *htmltojson.Node) {
	nodes := htmltojson.SearchAllNode(
		"",
		"script",
		"",
		"",
		"",
		rootNode,
	)

	mutex := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	done := make(chan bool)
	isChanClosed := false
	counter := 0
	for i, node := range nodes {
		for _, attr := range node.Attr {
			if attr.Key == "src" {
				link := attr.Val
				wg.Add(1)
				counter++
				go func() {
					uri, _ := url.Parse(link)
					resp, err := user.fb.doBasicGetRequest(uri)
					if err != nil {
						wg.Done()
						return
					}

					buff, err := decompressResponseBody(resp)
					if err != nil {
						wg.Done()
						return
					}

					body := buff.String()
					mutex.Lock()
					if user.profileCometTopAppSectionQuery == "" {
						mutex.Unlock()
						substr := `__d("ProfileCometTopAppSectionQuery$Parameters"`
						if strings.Contains(body, substr) {
							stg1 := strings.Split(body, substr)
							stg1 = strings.Split(stg1[1], `{id:"`)
							stg1 = strings.Split(stg1[1], `"`)
							mutex.Lock()
							user.profileCometTopAppSectionQuery = stg1[0]
							mutex.Unlock()
						}
					} else {
						mutex.Unlock()
					}

					mutex.Lock()
					if user.profileCometAppCollectionListRendererPaginationQuery == "" {
						mutex.Unlock()
						substr := `__d("ProfileCometAppCollectionListRendererPaginationQuery.graphql"`
						if strings.Contains(body, substr) {
							stg1 := strings.Split(body, substr)
							stg1 = strings.Split(stg1[1], `{id:"`)
							stg1 = strings.Split(stg1[1], `"`)
							mutex.Lock()
							user.profileCometAppCollectionListRendererPaginationQuery = stg1[0]
							mutex.Unlock()
						}
					} else {
						mutex.Unlock()
					}

					mutex.Lock()
					if user.profileCometTopAppSectionQuery != "" && user.profileCometAppCollectionListRendererPaginationQuery != "" && !isChanClosed {
						mutex.Unlock()
						isChanClosed = true
						done <- true
					} else {
						mutex.Unlock()
					}

					wg.Done()
				}()

				break
			}
		}

		if counter == 10 || i == len(nodes)-1 {
			go func() {
				wg.Wait()
				if user.profileCometTopAppSectionQuery == "" || user.profileCometAppCollectionListRendererPaginationQuery == "" {
					done <- true
				}
			}()

			<-done
			counter = 0
			if user.profileCometTopAppSectionQuery != "" && user.profileCometAppCollectionListRendererPaginationQuery != "" {
				close(done)
				return
			}
		}
	}
}
