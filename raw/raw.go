// Package raw contains raw Facebook JSON data
package raw

type Edge struct {
	Node *Node `json:"node"`
}

type Collection struct {
	Nodes []Node `json:"nodes"`
}

type WemPrivateSharingBundle struct {
	PrivateSharingControlModelForUser *PrivateSharingControl `json:"private_sharing_control_model_for_user"`
}

type PrivateSharingControl struct {
	PrivateSharingEnabled bool `json:"private_sharing_enabled"`
}

type StoryBucket struct {
	Nodes []Node `json:"nodes"`
}

type Extensions struct {
	IsFinal bool `json:"is_final"`
}
