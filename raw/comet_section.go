package raw

type CometSections struct {
	AttachedStory       *Story               `json:"attached_story"`
	Message             CometSectionsMessage `json:"message"`
	MessageContainer    MessageContainer     `json:"message_container"`
	Content             Content              `json:"content"`
	ContextLayout       Story                `json:"context_layout"`
	Timestamp           Timestamp            `json:"timestamp"`
	Feedback            *Feedback            `json:"feedback"`
	AttachedStoryLayout AttachedStoryLayout  `json:"attached_story_layout"`
	ActorPhoto          ActorPhoto           `json:"actor_photo"`
	Audience             Audience    `json:"audience"`
}
