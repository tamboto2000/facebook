package facebook

type Image struct {
	ID         string `json:"id"`
	URL        string `json:"url" csv:"url"`
	PostURL    string `json:"postUrl"`
	Width      int    `json:"width" csv:"width"`
	Height     int    `json:"height" csv:"height"`
	Analysis   string `json:"analysis" csv:"analysis"`
	IsPlayable bool   `json:"isPlayable"`
}

type Video struct {
	ID           string `json:"id"`
	URL          string `json:"url"`
	PostURL      string `json:"postUrl"`
	IsLiveStream bool   `json:"isLiveStream"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
}

type Reactions struct {
	Like  int `json:"like" csv:"react_like_count"`
	Love  int `json:"love" csv:"react_love_count"`
	Haha  int `json:"haha" csv:"react_haha_count"`
	Wow   int `json:"wow" csv:"react_wow_count"`
	Sad   int `json:"sad" csv:"react_sad_count"`
	Angry int `json:"angry" csv:"react_angry_count"`
}

type Reaction struct {
	Type     string `json:"type"`
	Reactors []User `json:"reactors"`
}

func (p *Story) SyncReactions() error {
	payloads, err := p.fb.doGraphQLRequest(map[string]interface{}{
		"feedbackTargetID": p.FeedbackID,
		"reactionType":     "NONE",
		"scale":            1,
	}, p.fb.ReactionsDialogQuery, "CometUFIReactionsDialogQuery", false)

	if err != nil {
		return err
	}

	payload := payloads[0]
	p.Reactions = new(Reactions)
	if payload.Data.Node.TopReactions != nil {
		for _, sum := range payload.Data.Node.TopReactions.Summary {
			if sum.Reaction.ReactionType == "LIKE" {
				p.Reactions.Like = int(sum.ReactionCount)
			}

			if sum.Reaction.ReactionType == "LOVE" {
				p.Reactions.Love = int(sum.ReactionCount)
			}

			if sum.Reaction.ReactionType == "HAHA" {
				p.Reactions.Haha = int(sum.ReactionCount)
			}

			if sum.Reaction.ReactionType == "WOW" {
				p.Reactions.Wow = int(sum.ReactionCount)
			}

			if sum.Reaction.ReactionType == "SORRY" {
				p.Reactions.Sad = int(sum.ReactionCount)
			}

			if sum.Reaction.ReactionType == "ANGER" {
				p.Reactions.Angry = int(sum.ReactionCount)
			}
		}
	}

	return nil
}
