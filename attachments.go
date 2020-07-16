package facebook

import (
	"encoding/json"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/tamboto2000/facebook/raw"
)

type Attachment struct {
	Type            string `json:"type,omitempty"`
	ID              string `json:"id,omitempty"`
	MediasetToken   string `json:"mediasetToken,omitempty"`
	URL             string `json:"url,omitempty"`
	MediaURL        string `json:"mediaUrl,omitempty"`
	Width           int    `json:"width,omitempty"`
	Height          int    `json:"height,omitempty"`
	Analysis        string `json:"analysis,omitempty"`
	IsPlayable      bool   `json:"isPlayable,omitempty"`
	IsLiveStreaming bool   `json:"isLiveStreaming,omitempty"`
	IsGamingVideo   bool   `json:"isGamingVideo,omitempty"`
	LiveViewerCount int    `json:"liveViewerCount,omitempty"`
	IsPremiere      bool   `json:"isPremiere,omitempty"`
}

type Attachments struct {
	UserID                         string `json:"userId"`
	Typename                       string `json:"typename"`
	MediasetToken                  string `json:"mediasetToken"`
	NodeID                         string `json:"nodeId"`
	Item                           *Story `json:"items"`
	CometPhotoRootQuery            string `json:"cometPhotoRootQuery"`
	CometVideoRootMediaViewerQuery string `json:"cometVideoRootMediaViewerQuery"`
	isEndOfList                    bool

	err error
	fb  *Facebook
}

func (att *Attachments) Error() error {
	return att.err
}

func (post *Story) Attachments() (*Attachments, error) {
	if len(post.PreviewAttachments) == 0 {
		return nil, nil
	}

	attachment := post.PreviewAttachments[0]
	att := &Attachments{
		UserID:              post.Author.ID,
		Typename:            attachment.Type,
		MediasetToken:       attachment.MediasetToken,
		NodeID:              attachment.ID,
		CometPhotoRootQuery: post.fb.CometPhotoRootQuery,
		fb:                  post.fb,
	}

	if err := att.sync(); err != nil {
		return nil, err
	}

	return att, nil
}

func (att *Attachments) Next() bool {
REITERATE:
	att.Item = nil
	if att.isEndOfList {
		return false
	}

	if err := att.sync(); err != nil {
		att.err = err
		return false
	}

	vars := map[string]interface{}{
		"UFI2CommentsProvider_commentsKey": "CometPhotoRootQuery",
		"containerIsFeedStory":             true,
		"containerIsLiveStory":             false,
		"containerIsTahoe":                 false,
		"containerIsWorkplace":             false,
		"feedLocation":                     "COMET_MEDIA_VIEWER",
		"feedbackSource":                   65,
		"isComet":                          true,
		"isMediaset":                       true,
		"loopMedia":                        false,
		"mediasetToken":                    att.MediasetToken,
		"nodeID":                           att.NodeID,
		"privacySelectorRenderLocation":    "COMET_MEDIA_VIEWER",
		"renderLocation":                   "permalink",
		"scale":                            1,
		"useDefaultActor":                  false,
	}

	var queryID string
	var apiName string
	if att.Typename == "photo" {
		queryID = att.CometPhotoRootQuery
		apiName = "CometPhotoRootQuery"
	}

	if att.Typename == "video" {
		queryID = att.CometVideoRootMediaViewerQuery
		apiName = "CometVideoRootMediaViewerQuery"
	}

	payloads, err := att.fb.doGraphQLRequest(vars, queryID, apiName, true)
	if err != nil {
		att.err = err
		return false
	}

	post := new(Story)
	for _, payload := range payloads {
		if payload.Data.CurrMedia != nil || payload.Data.Mediaset != nil {
			var media *raw.Media
			if payload.Data.CurrMedia != nil {
				media = payload.Data.CurrMedia
			}

			if payload.Data.Mediaset != nil {
				if payload.Data.Mediaset.CurrMedia != nil {
					media = payload.Data.Mediaset.CurrMedia
					if len(media.Edges) > 0 {
						media = media.Edges[0].Node
					}
				}
			}

			var fileURL string
			var width int
			var height int
			if media.Typename == "Photo" {
				fileURL = media.Image.URI
				width = int(media.Image.Width)
				height = int(media.Image.Height)
			}

			if media.Typename == "Video" {
				fileURL = media.PlayableURLQualityHD
				width = int(media.OriginalWidth)
				height = int(media.OriginalHeight)
			}

			var timestamp int64
			if media.CreatedTime > 0 {
				timestamp = media.CreatedTime
			}

			if media.CreationStory != nil {
				timestamp = media.CreationStory.CometSections.Timestamp.Story.CreationTime
			}

			date := time.Unix(timestamp, 0)
			dateStr := date.Format("2006-01-02 15:04:05")

			var reactionCount int
			if media.Feedback.Reactors != nil {
				reactionCount = int(media.Feedback.Reactors.Count)
			}
			post = &Story{
				ID:        media.ID,
				Timestamp: timestamp,
				CreatedAt: dateStr,
				Attachment: &Attachment{
					Type:            strings.ToLower(media.Typename),
					ID:              media.ID,
					URL:             fileURL,
					Height:          height,
					Width:           width,
					IsPlayable:      media.IsPlayable,
					IsLiveStreaming: media.IsLiveStreaming,
					IsGamingVideo:   media.IsGamingVideo,
					LiveViewerCount: int(media.LiveViewerCount),
					IsPremiere:      media.IsPremiere,
				},

				CanViewerShare: media.CanViewerShare,
				ReactionCount:  reactionCount,
				FeedbackID:     media.Feedback.ID,
			}
		}

		if payload.Data.Mediaset != nil {
			if payload.Data.Mediaset.NextMediaNoCursor != nil {
				att.NodeID = payload.Data.Mediaset.NextMediaNoCursor.ID
				if payload.Data.Mediaset.NextMediaNoCursor.Typename == "Video" {
					att.Typename = "video"
				} else {
					att.Typename = "photo"
				}
			} else {
				att.isEndOfList = true
			}
		}

		if payload.Data.URL != "" {
			post.URL = payload.Data.URL
			post.Attachment.MediaURL = post.URL
		}

		if payload.Data.CommentCount != nil {
			post.CommentCount = int(payload.Data.CommentCount.TotalCount)
		}

		if payload.Data.CanViewerComment {
			post.CanViewerComment = true
		}

		if payload.Data.TopLevelCommentListRenderer != nil {
			post.CommentsDisabled = payload.Data.TopLevelCommentListRenderer.Feedback.HaveCommentsBeenDisabled
		}

		if payload.Data.CometUFISummaryAndActionsRenderer != nil {
			renderer := payload.Data.CometUFISummaryAndActionsRenderer
			post.ReactionCount = int(renderer.Feedback.ReactionCount.Count)
			post.ShareCount = int(renderer.Feedback.ShareCount.Count)
			post.CanViewerComment = renderer.Feedback.CanViewerComment
			post.CanViewerReact = renderer.Feedback.CanViewerReact
			post.VideoViewCount = renderer.Feedback.VideoViewCount
		}
	}

	if post.ID == "" {
		goto REITERATE
	}

	post.fb = att.fb
	att.Item = post

	return true
}

func (att *Attachments) sync() error {
	required := "CometPhotoRoot.react"
	if att.Typename == "video" {
		required = "CometVideoRoot.react"
	}

	if att.Typename == "photo" && att.fb.CometPhotoRootQuery != "" {
		att.CometPhotoRootQuery = att.fb.CometPhotoRootQuery
		return nil
	}

	if att.Typename == "video" && att.fb.CometVideoRootMediaViewerQuery != "" {
		att.CometVideoRootMediaViewerQuery = att.fb.CometVideoRootMediaViewerQuery
		return nil
	}

	var mediaURL string
	if att.Typename == "photo" {
		mediaURL = "/photo/?fbid=" + att.NodeID + "&&set=" + att.MediasetToken
	}

	if att.Typename == "video" {
		mediaURL = "/" + att.UserID + "/videos/" + att.MediasetToken + "/" + att.NodeID + "/"
	}

	queryID := ""
	uri, _ := url.Parse(att.fb.RootURL + "/ajax/bulk-route-definitions/")
	q := uri.Query()
	q.Add("route_urls[0]", mediaURL)
	uri.RawQuery = q.Encode()
	resp, err := att.fb.doGetRequest(uri)
	if err != nil {
		return err
	}

	buff, err := decompressResponseBody(resp)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body := strings.Replace(buff.String(), "for (;;);", "", 1)
	routeDef := new(raw.RouteDefinitions)
	if err = json.Unmarshal([]byte(body), routeDef); err != nil {
		return err
	}

	wg := new(sync.WaitGroup)
	mx := new(sync.Mutex)
	done := make(chan bool)
	counter := 0
	for key, val := range routeDef.Payload.SrPayload.Bootloadable {
		if key == required {
			for i, r := range val.R {
				if routeDef.Payload.SrPayload.ResourceMap[r].Type != "js" {
					continue
				}

				link := routeDef.Payload.SrPayload.ResourceMap[r].Src
				wg.Add(1)
				counter++
				go func() {
					uri, _ := url.Parse(link)
					resp, err := att.fb.doBasicGetRequest(uri)
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
					defer resp.Body.Close()

					mx.Lock()
					if queryID != "" {
						mx.Unlock()
						wg.Done()
						return
					}

					mx.Unlock()

					substr := `__d("CometPhotoRootQuery.graphql"`
					if att.Typename == "video" {
						substr = `__d("CometVideoRootMediaViewerQuery.graphql"`
					}

					if !strings.Contains(body, substr) {
						wg.Done()
						return
					}

					stg1 := strings.Split(body, substr)
					stg1 = strings.Split(stg1[1], `{id:"`)
					stg1 = strings.Split(stg1[1], `"`)
					mx.Lock()
					queryID = stg1[0]
					if att.Typename == "photo" {
						att.fb.CometPhotoRootQuery = queryID
						att.CometPhotoRootQuery = queryID
					}

					if att.Typename == "video" {
						att.fb.CometVideoRootMediaViewerQuery = queryID
						att.CometVideoRootMediaViewerQuery = queryID
					}

					mx.Unlock()
					wg.Done()
					done <- true
				}()

				if counter == 5 || i == len(val.R)-1 {
					go func() {
						wg.Wait()
						if queryID == "" {
							done <- true
							return
						}

						close(done)
					}()

					<-done
					counter = 0
					if queryID != "" {
						break
					}
				}
			}

			break
		}
	}

	return nil
}

func (att *Attachments) IsEndOfList() bool {
	return att.isEndOfList
}

func (att *Attachments) inject(fb *Facebook) {
	att.fb = fb
}
