package raw

type User struct {
	IsProfile                       string                      `json:"__isProfile,omitempty"`
	Name                            string                      `json:"name,omitempty"`
	ProfilePictureForStickyBar      *ProfilePictureForStickyBar `json:"profile_picture_for_sticky_bar,omitempty"`
	IsEntity                        string                      `json:"__isEntity,omitempty"`
	Typename                        string                      `json:"__typename,omitempty"`
	URL                             string                      `json:"url,omitempty"`
	IsViewerFriend                  bool                        `json:"is_viewer_friend,omitempty"`
	IsRenderedProfile               string                      `json:"__isRenderedProfile,omitempty"`
	CoverPhoto                      *CoverPhoto                 `json:"cover_photo,omitempty"`
	ProfileActiveMessengerRoom      interface{}                 `json:"profile_active_messenger_room,omitempty"`
	StoryBucket                     *StoryBucket                `json:"story_bucket,omitempty"`
	ProfilePictureExpirationTime    interface{}                 `json:"profile_picture_expiration_time,omitempty"`
	ProfileTabs                     *ProfileTabs                `json:"profile_tabs,omitempty"`
	AlternateName                   string                      `json:"alternate_name,omitempty"`
	UsernameForProfile              string                      `json:"username_for_profile,omitempty"`
	ID                              string                      `json:"id,omitempty"`
	IsVerified                      bool                        `json:"is_verified,omitempty"`
	IsVisiblyMemorialized           bool                        `json:"is_visibly_memorialized,omitempty"`
	ActiveFundraiserBadgeRenderer   interface{}                 `json:"active_fundraiser_badge_renderer,omitempty"`
	WemPrivateSharingBundle         *WemPrivateSharingBundle    `json:"wem_private_sharing_bundle,omitempty"`
	ProfilePlusWoodhengeCreatorInfo interface{}                 `json:"profile_plus_woodhenge_creator_info,omitempty"`
	MentionsTabTooltipNuxText       string                      `json:"mentions_tab_tooltip_nux_text,omitempty"`
	IsMemorialized                  bool                        `json:"is_memorialized,omitempty"`
	MemorializedUserInfo            interface{}                 `json:"memorialized_user_info,omitempty"`
	TimelineNavAppSections          *TimelineNavAppSections     `json:"timeline_nav_app_sections,omitempty"`
}

type ProfileTabs struct {
	Typename                                            string                                      `json:"__typename,omitempty"`
	ProfileUser                                         *User                                       `json:"profile_user,omitempty"`
	ModuleOperationProfileCometTetraishEntityHeaderUser *ModuleProfileCometTetraishEntityHeaderUser `json:"__module_operation_ProfileCometTetraishEntityHeader_user,omitempty"`
	ModuleComponentProfileCometTetraishEntityHeaderUser *ModuleProfileCometTetraishEntityHeaderUser `json:"__module_component_ProfileCometTetraishEntityHeader_user,omitempty"`
}

type CoverPhoto struct {
	Focus *Focus `json:"focus,omitempty"`
	Photo *Photo `json:"photo,omitempty"`
}

type Focus struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
}

type ProfilePictureForStickyBar struct {
	URI string `json:"uri,omitempty"`
}

type ModuleProfileCometTetraishEntityHeaderUser struct {
	DR string `json:"__dr,omitempty"`
}

type TimelineNavAppSections struct {
	Edges []Edge `json:"edges,omitempty"`
}
