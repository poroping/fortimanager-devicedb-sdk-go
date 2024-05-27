package models

type InstallScriptPayload struct {
	Adom    *string  `json:"adom,omitempty"`
	Package *string  `json:"package,omitempty"`
	Pblock  *string  `json:"pblock,omitempty"`
	Script  *string  `json:"script,omitempty"`
	Flags   []string `json:"flags,omitempty"`
	Scope   []Scope  `json:"scope,omitempty"`
}

type InstallScriptResponse struct {
	Task *int64 `json:"task,omitempty"`
}
