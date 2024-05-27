package models

type InstallDevicePayload struct {
	Adom           *string  `json:"adom,omitempty"`
	DevRevComments *string  `json:"dev_rev_comments,omitempty"`
	Flags          []string `json:"flags,omitempty"`
	Scope          []Scope  `json:"scope,omitempty"`
}

type InstallDeviceResponse struct {
	Task *int64 `json:"task,omitempty"`
}
