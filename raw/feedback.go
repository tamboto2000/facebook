package raw

type Feedback struct {
	ID                                      string                   `json:"id"`
	Typename                                string                   `json:"__typename"`
	Story                                   *Story                   `json:"story"`
	CanViewerComment                        bool                     `json:"can_viewer_comment"`
	CanSeeVoiceSwitcher                     bool                     `json:"can_see_voice_switcher"`
	HaveCommentsBeenDisabled                bool                     `json:"have_comments_been_disabled"`
	IsViewerMuted                           bool                     `json:"is_viewer_muted"`
	Reactors                                *Reactors                `json:"reactors"`
	CanViewerReact                          bool                     `json:"can_viewer_react"`
	CommentCount                            *CommentCount            `json:"comment_count"`
	DisplayComments                         *FeedbackDisplayComments `json:"display_comments"`
	ReactionCount                           *ReactionCount           `json:"reaction_count"`
	ShareCount                              *ReactionCount           `json:"share_count"`
	CometUfiSummaryAndActionsRenderer       *Renderer                `json:"comet_ufi_summary_and_actions_renderer"`
	ViewerActor                             *ViewerActor             `json:"viewer_actor"`
	TopReactions                            *TopReactions            `json:"top_reactions"`
	CanViewerCommentWithFile                bool                     `json:"can_viewer_comment_with_file"`
	CanViewerCommentWithGIF                 bool                     `json:"can_viewer_comment_with_gif"`
	CanViewerCommentWithInsightPoint        bool                     `json:"can_viewer_comment_with_insight_point"`
	CanViewerCommentWithPhoto               bool                     `json:"can_viewer_comment_with_photo"`
	CanViewerCommentWithVideo               bool                     `json:"can_viewer_comment_with_video"`
	CanViewerCommentWithSticker             bool                     `json:"can_viewer_comment_with_sticker"`
	CanViewerCommentWithStars               bool                     `json:"can_viewer_comment_with_stars"`
	CanViewerCommentWithBotMention          bool                     `json:"can_viewer_comment_with_bot_mention"`
	CanViewerCommentWithMarker              bool                     `json:"can_viewer_comment_with_marker"`
	IsCommentMarkdownEligible               bool                     `json:"is_comment_markdown_eligible"`
	MentionsDatasourceJSConstructorArgsJSON string                   `json:"mentions_datasource_js_constructor_args_json"`
	CanViewerPinLiveComments                bool                     `json:"can_viewer_pin_live_comments"`
	DisplayCommentsCount                    *CommentCount            `json:"display_comments_count"`
	IsEligibleForRealTimeUpdates            bool                     `json:"is_eligible_for_real_time_updates"`
	SubscriptionTargetID                    string                   `json:"subscription_target_id"`
	CommentComposerPlaceholder              string                   `json:"comment_composer_placeholder"`
	ToplevelCommentCount                    *CommentCount            `json:"toplevel_comment_count"`
	ShareFbid                               string                   `json:"share_fbid"`
	DefaultCommentOrdering                  string                   `json:"default_comment_ordering"`
	UnfilteredCommentCount                  *CommentCount            `json:"unfiltered_comment_count"`
	RepliesListRenderer                     *Renderer                `json:"replies_list_renderer"`
	VideoViewCount                          int                      `json:"video_view_count"`
	TopLevelCommentListRenderer             *Renderer                `json:"top_level_comment_list_renderer"`
}
