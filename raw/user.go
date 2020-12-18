package raw

type User struct {
	IsProfile                       string                      `json:"__isProfile"`
	Name                            string                      `json:"name"`
	ProfilePictureForStickyBar      *ProfilePictureForStickyBar `json:"profile_picture_for_sticky_bar"`
	IsEntity                        string                      `json:"__isEntity"`
	Typename                        string                      `json:"__typename"`
	URL                             string                      `json:"url"`
	IsViewerFriend                  bool                        `json:"is_viewer_friend"`
	IsRenderedProfile               string                      `json:"__isRenderedProfile"`
	CoverPhoto                      *CoverPhoto                 `json:"cover_photo"`
	ProfileActiveMessengerRoom      interface{}                 `json:"profile_active_messenger_room"`
	StoryBucket                     *StoryBucket                `json:"story_bucket"`
	ProfilePictureExpirationTime    interface{}                 `json:"profile_picture_expiration_time"`
	ProfileTabs                     *ProfileTabs                `json:"profile_tabs"`
	AlternateName                   string                      `json:"alternate_name"`
	UsernameForProfile              string                      `json:"username_for_profile"`
	ID                              string                      `json:"id"`
	IsVerified                      bool                        `json:"is_verified"`
	IsVisiblyMemorialized           bool                        `json:"is_visibly_memorialized"`
	ActiveFundraiserBadgeRenderer   interface{}                 `json:"active_fundraiser_badge_renderer"`
	WemPrivateSharingBundle         *WemPrivateSharingBundle    `json:"wem_private_sharing_bundle"`
	ProfilePlusWoodhengeCreatorInfo interface{}                 `json:"profile_plus_woodhenge_creator_info"`
	MentionsTabTooltipNuxText       string                      `json:"mentions_tab_tooltip_nux_text"`
	IsMemorialized                  bool                        `json:"is_memorialized"`
	MemorializedUserInfo            interface{}                 `json:"memorialized_user_info"`
	TimelineNavAppSections          *TimelineNavAppSections     `json:"timeline_nav_app_sections"`
}

type ProfileTabs struct {
	Typename                                            string                                      `json:"__typename"`
	ProfileUser                                         *User                                       `json:"profile_user"`
	ModuleOperationProfileCometTetraishEntityHeaderUser *ModuleProfileCometTetraishEntityHeaderUser `json:"__module_operation_ProfileCometTetraishEntityHeader_user"`
	ModuleComponentProfileCometTetraishEntityHeaderUser *ModuleProfileCometTetraishEntityHeaderUser `json:"__module_component_ProfileCometTetraishEntityHeader_user"`
}

type CoverPhoto struct {
	Focus *Focus `json:"focus"`
	Photo *Photo `json:"photo"`
}

type Focus struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type ProfilePictureForStickyBar struct {
	URI string `json:"uri"`
}

type ModuleProfileCometTetraishEntityHeaderUser struct {
	DR string `json:"__dr"`
}

type TimelineNavAppSections struct {
	Edges []Edge `json:"edges"`
}
