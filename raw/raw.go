// Package raw contains raw Facebook JSON data
package raw

type Edge struct {
	Node *Node `json:"node,omitempty"`
}

type Collection struct {
	Nodes []Node `json:"nodes,omitempty"`
}

type WemPrivateSharingBundle struct {
	PrivateSharingControlModelForUser *PrivateSharingControl `json:"private_sharing_control_model_for_user,omitempty"`
}

type PrivateSharingControl struct {
	PrivateSharingEnabled bool `json:"private_sharing_enabled,omitempty"`
}

type StoryBucket struct {
	Nodes []Node `json:"nodes,omitempty"`
}

type Extensions struct {
	IsFinal bool `json:"is_final,omitempty"`
}
