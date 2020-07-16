package facebook

import (
	"encoding/base64"
	"net/url"
	"strings"
	"time"

	"github.com/tamboto2000/facebook/raw"
)

type Story struct {
	ID                 string        `json:"id"`
	URL                string        `json:"url"`
	Text               string        `json:"text,omitempty"`
	BackgroundImageURL string        `json:"backgroundImageUrl,omitempty"`
	Timestamp          int64         `json:"timestamp,omitempty"`
	CreatedAt          string        `json:"createdAt,omitempty"`
	AttachedStory      *Story        `json:"attachedStory,omitempty"`
	PreviewAttachments []*Attachment `json:"previewAttachments,omitempty"`
	Attachment         *Attachment   `json:"attachment,omitempty"`
	Reactions          *Reactions    `json:"reactions,omitempty"`
	CommentCount       int           `json:"commentCount,omitempty"`
	ReactionCount      int           `json:"reactionCount,omitempty"`
	ShareCount         int           `json:"shareCount,omitempty"`
	CanViewerComment   bool          `json:"canViewerComment,omitempty"`
	CanViewerReact     bool          `json:"canViewerReact,omitempty"`
	CanViewerShare     bool          `json:"canViewerShare,omitempty"`
	CommentsDisabled   bool          `json:"commentsDisabled,omitempty"`
	FeedbackID         string        `json:"feedbackID,omitempty"`
	PreviewComments    []*Comment    `json:"previewComments,omitempty"`
	Author             *User         `json:"author,omitempty"`
	Mentioned          []User        `json:"mentioned,omitempty"`
	VideoViewCount     int           `json:"videoViewCount,omitempty"`

	fb *Facebook
}

type Timeline struct {
	UserID                          string   `json:"userId"`
	Cursor                          string   `json:"cursor"`
	Items                           []*Story `json:"items"`
	TimelineFeedQueryRelayPreloader string   `json:"timelineFeedQueryRelayPreloader"`
	TimelineFeedRefetchQuery        string   `json:"timelineFeedRefetchQuery"`

	err            error
	fb             *Facebook
	isEndOfList    bool
	isFirstRequest bool
}

func (user *User) Timeline() *Timeline {
	return &Timeline{
		UserID:                          user.ID,
		TimelineFeedQueryRelayPreloader: user.fb.TimelineFeedQueryRelayPreloader,
		TimelineFeedRefetchQuery:        user.fb.TimelineFeedRefetchQuery,
		fb:                              user.fb,
	}
}

func (tl *Timeline) Sync() error {
	tl.isFirstRequest = true
	return nil
}

func (tl *Timeline) Error() error {
	return tl.err
}

func (tl *Timeline) IsEndOfList() bool {
	return tl.isEndOfList
}

func (tl *Timeline) Next(count int) bool {
	tl.Items = make([]*Story, 0)

	if tl.isEndOfList {
		return false
	}

	var apiName string
	var variables map[string]interface{}
	var docID string
	if tl.Cursor == "" {
		apiName = "ProfileCometTimelineFeedQuery"
		docID = tl.TimelineFeedQueryRelayPreloader
		variables = map[string]interface{}{
			"UFI2CommentsProvider_commentsKey": "ProfileCometTimelineRoute",
			"containerIsFeedStory":             false,
			"containerIsLiveStory":             false,
			"containerIsTahoe":                 false,
			"containerIsWorkplace":             false,
			"feedLocation":                     "TIMELINE",
			"feedbackSource":                   0,
			"isComet":                          true,
			"privacySelectorRenderLocation":    "COMET_STREAM",
			"renderLocation":                   "timeline",
			"scale":                            1,
			"useIncrementalDelivery":           true,
			"userID":                           tl.UserID,
		}
	} else {
		if count > 50 {
			count = 50
		}

		if count <= 0 {
			count = 50
		}

		count++

		apiName = "ProfileCometTimelineFeedRefetchQuery"
		docID = tl.TimelineFeedRefetchQuery
		variables = map[string]interface{}{
			"count":                            count,
			"cursor":                           tl.Cursor,
			"useIncrementalDelivery":           true,
			"renderLocation":                   "timeline",
			"scale":                            1,
			"privacySelectorRenderLocation":    "COMET_STREAM",
			"feedLocation":                     "TIMELINE",
			"feedbackSource":                   0,
			"useDefaultActor":                  false,
			"UFI2CommentsProvider_commentsKey": "ProfileCometTimelineRoute",
			"isComet":                          true,
			"containerIsFeedStory":             true,
			"containerIsLiveStory":             false,
			"containerIsTahoe":                 false,
			"containerIsWorkplace":             false,
			"id":                               tl.UserID,
		}

		tl.isFirstRequest = false
	}

	payloads, err := tl.fb.doGraphQLRequest(variables, docID, apiName, true)
	if err != nil {
		tl.err = err
		return false
	}

	rawStoryCounter := 0
	for _, payload := range payloads {
		rawStoryCounter++
		var node raw.Node

		if payload.Data.User.TimelineListFeedUnits.Edges != nil &&
			len(payload.Data.User.TimelineListFeedUnits.Edges) > 0 {
			node = payload.Data.User.TimelineListFeedUnits.Edges[0].Node
		} else if payload.Data.Node.CometSections != nil {
			node = payload.Data.Node
		} else {
			continue
		}

		tl.Cursor = payload.Data.Cursor
		post := extractStory(&node)
		if post == nil {
			continue
		}

		post.fb = tl.fb
		tl.Items = append(tl.Items, post)
	}

	if tl.Cursor == "" || len(tl.Items) == 0 {
		tl.isEndOfList = true
	} else if rawStoryCounter < count-1 && !tl.isFirstRequest {
		tl.isEndOfList = true
	}

	return true
}

func (tl *Timeline) inject(fb *Facebook) {
	tl.fb = fb
}

func extractStory(node *raw.Node) *Story {
	story := new(Story)
	//copy raw story for easy data extraction
	rawStory := node.CometSections.Content.Story

	//get URL
	if node.CometSections.ContextLayout.Story.CometSections.Timestamp.Story != nil {
		story.URL = node.CometSections.ContextLayout.Story.CometSections.Timestamp.Story.URL
	} else {
		return nil
	}

	//get id
	story.ID = storyIDFromURL(story.URL)

	//get message/text/caption
	if rawStory.CometSections.Message.Story != nil {
		story.Text = rawStory.CometSections.Message.Story.Message.Text
	}

	//get timestamp
	if node.CometSections.ContextLayout.Story.CometSections.Timestamp.Story != nil {
		timestamp := node.CometSections.ContextLayout.Story.CometSections.Timestamp.Story.CreationTime
		date := time.Unix(timestamp, 0)
		dateStr := date.Format("2006-01-02 15:04:05")
		story.Timestamp = timestamp
		story.CreatedAt = dateStr
	}

	//get attachments
	story.PreviewAttachments = extractAttachments(rawStory)
	if len(story.PreviewAttachments) == 1 {
		story.Attachment = story.PreviewAttachments[0]
		story.PreviewAttachments = nil
	}

	//get reaction count
	story.ReactionCount = int(node.CometSections.Feedback.Story.FeedbackContext.FeedbackTargetWithContext.CometUfiSummaryAndActionsRenderer.Feedback.ReactionCount.Count)

	//get comment count
	story.CommentCount = int(node.CometSections.Feedback.Story.FeedbackContext.FeedbackTargetWithContext.CommentCount.TotalCount)

	//get share count
	story.ShareCount = int(node.CometSections.Feedback.Story.FeedbackContext.FeedbackTargetWithContext.CometUfiSummaryAndActionsRenderer.Feedback.ShareCount.Count)

	story.CanViewerComment = node.CometSections.Feedback.Story.FeedbackContext.FeedbackTargetWithContext.CanViewerComment
	story.CanViewerReact = node.CometSections.Feedback.Story.FeedbackContext.FeedbackTargetWithContext.CometUfiSummaryAndActionsRenderer.Feedback.CanViewerReact
	//can viewer share
	if len(node.CometSections.Content.Story.Attachments) > 0 {
		if node.CometSections.Content.Story.Attachments[0].StyleTypeRenderer.Attachment != nil {
			if node.CometSections.Content.Story.Attachments[0].StyleTypeRenderer.Attachment.Media != nil {
				story.CanViewerShare = node.CometSections.Content.Story.Attachments[0].StyleTypeRenderer.Attachment.Media.CanViewerShare
			}
		}
	}

	story.CommentsDisabled = node.CometSections.Feedback.Story.FeedbackContext.FeedbackTargetWithContext.TopLevelCommentListRenderer.Feedback.HaveCommentsBeenDisabled

	//create feedbcak id
	story.FeedbackID = base64.StdEncoding.EncodeToString([]byte("feedback:" + story.ID))

	//get preview comments
	rawComments := node.CometSections.Feedback.Story.FeedbackContext.FeedbackTargetWithContext.DisplayComments.Edges
	story.PreviewComments = extractComments(rawComments)

	//get author
	rawActor := node.CometSections.ContextLayout.Story.CometSections.ActorPhoto.Story.Actors[0]
	story.Author = &User{
		ID:             rawActor.ID,
		URL:            rawActor.URL,
		Name:           rawActor.Name,
		ProfilePictURL: rawActor.ProfilePicture.URI,
	}

	//get attached story, if any
	if node.AttachedStory != nil {
		story.AttachedStory = extractAttachedStory(node)
	}

	return story
}

func extractAttachedStory(node *raw.Node) *Story {
	story := new(Story)
	//get url
	story.URL = node.AttachedStory.CometSections.ContextLayout.Story.CometSections.Timestamp.Story.URL

	//get id
	story.ID = storyIDFromURL(story.URL)

	//get text
	story.Text = node.CometSections.Content.Story.CometSections.AttachedStory.Story.AttachedStory.CometSections.AttachedStoryLayout.Story.Message.Text

	//get timestamp
	timestamp := node.AttachedStory.CometSections.ContextLayout.Story.CometSections.Timestamp.Story.CreationTime
	date := time.Unix(timestamp, 0)
	dateStr := date.Format("2006-01-02 15:04:05")
	story.Timestamp = timestamp
	story.CreatedAt = dateStr

	//get preview attachments
	story.PreviewAttachments = extractAttachments(node.CometSections.Content.Story.CometSections.AttachedStory.Story.AttachedStory.CometSections.AttachedStoryLayout.Story)
	if len(story.PreviewAttachments) == 1 {
		story.Attachment = story.PreviewAttachments[0]
		story.PreviewAttachments = nil
	}

	//get author
	rawActor := node.AttachedStory.CometSections.ContextLayout.Story.CometSections.ActorPhoto.Story.Actors[0]
	story.Author = &User{
		ID:             rawActor.ID,
		URL:            rawActor.URL,
		Name:           rawActor.Name,
		ProfilePictURL: rawActor.ProfilePicture.URI,
	}

	//create feedback id
	story.FeedbackID = base64.StdEncoding.EncodeToString([]byte("feedback:" + story.ID))

	return story
}

func extractAttachments(story *raw.Story) []*Attachment {
	attachments := make([]*Attachment, 0)
	if len(story.Attachments) == 0 {
		return nil
	}

	rawAttachments := story.Attachments[0].StyleTypeRenderer.Attachment
	if rawAttachments == nil {
		return nil
	}

	if rawAttachments.AllSubattachments != nil {
		for _, att := range rawAttachments.AllSubattachments.Nodes {
			media := att.Media
			if media.Typename == "Photo" {
				attachments = append(attachments, &Attachment{
					Type:          "photo",
					ID:            media.ID,
					MediasetToken: rawAttachments.MediasetToken,
					URL:           media.Image.URI,
					MediaURL:      att.URL,
					Width:         int(media.Image.Width),
					Height:        int(media.Image.Height),
					Analysis:      media.AccessibilityCaption,
					IsPlayable:    media.IsPlayable,
				})
			}

			if media.Typename == "Video" {
				attachments = append(attachments, &Attachment{
					Type:            "video",
					ID:              media.ID,
					MediasetToken:   rawAttachments.MediasetToken,
					URL:             media.PlayableURL,
					MediaURL:        media.PermalinkURL,
					Width:           int(media.Width),
					Height:          int(media.Height),
					IsPlayable:      true,
					IsLiveStreaming: media.IsLiveStreaming,
					IsGamingVideo:   media.IsGamingVideo,
					LiveViewerCount: int(media.LiveViewerCount),
					IsPremiere:      media.IsPremiere,
				})
			}
		}
	}

	if len(attachments) > 0 {
		return attachments
	}

	if rawAttachments.Media != nil {
		media := rawAttachments.Media
		if media.Typename == "Photo" {
			var image *raw.Image
			if media.PhotoImage != nil {
				image = media.PhotoImage
			}

			if rawAttachments.Media.Image != nil {
				image = media.Image
			}

			attachments = append(attachments, &Attachment{
				Type:       "photo",
				ID:         media.ID,
				URL:        image.URI,
				MediaURL:   media.URL,
				Width:      int(image.Width),
				Height:     int(image.Height),
				Analysis:   media.AccessibilityCaption,
				IsPlayable: media.IsPlayable,
			})

			return attachments
		}

		if media.Typename == "Video" {
			attachments = append(attachments, &Attachment{
				Type:            "video",
				ID:              media.ID,
				URL:             media.PlayableURL,
				MediaURL:        media.PermalinkURL,
				Width:           int(media.Width),
				Height:          int(media.Height),
				IsPlayable:      true,
				IsLiveStreaming: media.IsLiveStreaming,
				IsGamingVideo:   media.IsGamingVideo,
				LiveViewerCount: int(media.LiveViewerCount),
				IsPremiere:      media.IsPremiere,
			})

			return attachments
		}
	}

	return attachments
}

func storyIDFromURL(uri string) string {
	if strings.Contains(uri, "story_fbid") {
		parsed, _ := url.Parse(uri)
		return parsed.Query().Get("story_fbid")
	}

	stg1 := strings.Split(uri, "/")
	for i, frag := range stg1 {
		if frag == "posts" {
			return stg1[i+1]
		}
	}

	return ""
}
