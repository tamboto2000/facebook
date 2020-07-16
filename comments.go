package facebook

import (
	"errors"
	"strings"
	"time"

	"github.com/tamboto2000/facebook/raw"
)

type Comment struct {
	ID             string      `json:"id,omitempty"`
	LegacyFBID     string      `json:"legacyFbId,omitempty"`
	URL            string      `json:"url,omitempty"`
	Timestamp      int64       `json:"timestamp,omitempty"`
	CreatedAt      string      `json:"createdAt,omitempty"`
	Edited         bool        `json:"edited,omitempty"`
	Text           string      `json:"text,omitempty"`
	Attachment     *Attachment `json:"attachment,omitempty"`
	ReactionCount  int         `json:"reactionCount,omitempty"`
	Reactions      *Reactions  `json:"reactions,omitempty"`
	RepliesPreview []Comment   `json:"repliesPreview,omitempty"`
	Author         *User       `json:"author,omitempty"`
}

type Comments struct {
	FeedbackID     string     `json:"feedbackId"`
	Comments       []*Comment `json:"comments"`
	Before         string     `json:"before"`
	After          string     `json:"after"`
	CommentQueryID string     `json:"commentQueryId"`

	isEndOfList bool
	err         error
	fb          *Facebook
}

func (post *Story) Comments() *Comments {
	return &Comments{
		FeedbackID:     post.FeedbackID,
		CommentQueryID: post.fb.CommentQueryID,
		fb:             post.fb,
	}
}

func (comm *Comments) Next(count int) bool {
	comm.Comments = make([]*Comment, 0)
	if comm.isEndOfList {
		return false
	}

	if count > 50 {
		count = 50
	}

	if count <= 0 {
		count = 50
	}

	var vars map[string]interface{}
	vars = map[string]interface{}{
		"feedLocation":                     "TIMELINE",
		"feedbackID":                       comm.FeedbackID,
		"feedbackSource":                   0,
		"first":                            count,
		"includeHighlightedComments":       false,
		"includeNestedComments":            true,
		"isInitialFetch":                   false,
		"isPaginating":                     true,
		"isComet":                          true,
		"containerIsFeedStory":             true,
		"containerIsWorkplace":             false,
		"containerIsLiveStory":             false,
		"containerIsTahoe":                 false,
		"scale":                            1,
		"useDefaultActor":                  false,
		"UFI2CommentsProvider_commentsKey": "ProfileCometTimelineRoute",
	}

	if comm.After != "" {
		vars["after"] = comm.After
	}

	if comm.Before != "" {
		vars["before"] = comm.Before
	}

	payloads, err := comm.fb.doGraphQLRequest(
		vars,
		comm.CommentQueryID,
		"CometUFICommentsProviderPaginationQuery",
		false,
	)

	if err != nil {
		comm.err = err
		return false
	}

	if len(payloads) == 0 {
		return false
	}

	payload := payloads[0]
	if len(payload.Errors) > 0 {
		comm.err = errors.New(payload.Errors[0].Message)
		return false
	}

	edgeLen := len(payload.Data.Feedback.DisplayComments.Edges)
	if edgeLen == 0 {
		return false
	}

	comm.Comments = extractComments(payload.Data.Feedback.DisplayComments.Edges)
	if payload.Data.Feedback.DisplayComments.Edges[edgeLen-1].Cursor == "" {
		comm.isEndOfList = true
	} else {
		comm.After = payload.Data.Feedback.DisplayComments.Edges[edgeLen-1].Cursor
	}

	if !payload.Data.Feedback.DisplayComments.PageInfo.HasNextPage {
		comm.isEndOfList = true
	}

	return true
}

func (comm *Comments) Error() error {
	return comm.err
}

func (comm *Comments) IsEndOfList() bool {
	return comm.isEndOfList
}

func (comm *Comments) inject(fb *Facebook) {
	comm.fb = fb
}

func extractComments(edges []raw.Edge) []*Comment {
	comments := make([]*Comment, 0)
	for _, edge := range edges {
		datetime := time.Unix(edge.Node.CreatedTime, 0).Format("2006-01-02 15:04:05")
		edited := false
		if edge.Node.EditHistory.Count > 0 {
			edited = true
		}

		author := &User{
			ID:             edge.Node.Author.ID,
			Name:           edge.Node.Author.Name,
			ProfilePictURL: edge.Node.Author.ProfilePictureDepth0.URI,
			URL:            edge.Node.Author.URL,
		}

		comment := &Comment{
			ID:            edge.Node.ID,
			LegacyFBID:    edge.Node.LegacyFbid,
			URL:           edge.Node.URL,
			Timestamp:     edge.Node.CreatedTime,
			CreatedAt:     datetime,
			ReactionCount: int(edge.Node.Feedback.Reactors.Count),
			Edited:        edited,
			Author:        author,
		}

		if edge.Node.Body != nil {
			comment.Text = edge.Node.Body.Text
		}

		if len(edge.Node.Attachments) > 0 {
			att := edge.Node.Attachments[0]
			comment.Attachment = &Attachment{
				Type:       strings.ToLower(att.Media.Typename),
				ID:         att.Media.ID,
				URL:        att.Media.Image.URI,
				MediaURL:   att.URL,
				Width:      int(att.Media.Image.Width),
				Height:     int(att.Media.Image.Height),
				Analysis:   att.Media.AccessibilityCaption,
				IsPlayable: att.Media.IsPlayable,
			}
		}

		comments = append(comments, comment)
	}

	return comments
}
