package raw

type Actor struct {
	Typename       string           `json:"__typename"`
	ProfilePicture LargerProfilePic `json:"profile_picture"`
	ID             string           `json:"id"`
	URL            string           `json:"url,omitempty"`
	Name           string           `json:"name,omitempty"`
}
