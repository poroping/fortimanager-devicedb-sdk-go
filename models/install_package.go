package models

type InstallPackagePayload struct {
	Adom  *string  `json:"adom,omitempty"`
	Pkg   *string  `json:"pkg,omitempty"`
	Flags []string `json:"flags,omitempty"`
	Scope []Scope  `json:"scope,omitempty"`
}

type InstallPackageResponse struct {
	Task *int64 `json:"task,omitempty"`
}
