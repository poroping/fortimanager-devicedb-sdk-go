package models

type ModelDevice struct {
	Adom   *string `json:"adom,omitempty"`
	Device *Device `json:"device,omitempty"`
}

type Device struct {
	Name                  *string          `json:"name,omitempty"`
	Desc                  *string          `json:"desc,omitempty"`
	AdmUSR                *string          `json:"adm_usr,omitempty"`
	OSVer                 *string          `json:"os_ver,omitempty"`
	Version               *int64           `json:"version,omitempty"`
	OSType                *string          `json:"os_type,omitempty"`
	Mr                    *int64           `json:"mr,omitempty"`
	PlatformStr           *string          `json:"platform_str,omitempty"`
	PlatformID            *int64           `json:"platform_id,omitempty"`
	MgmtMode              *string          `json:"mgmt_mode,omitempty"`
	PreferImgVer          *string          `json:"prefer_img_ver,omitempty"`
	Patch                 *int64           `json:"patch,omitempty"`
	Build                 *int64           `json:"build,omitempty"`
	BranchPt              *int64           `json:"branch_pt,omitempty"`
	Sn                    *string          `json:"sn,omitempty"`
	Flags                 []string         `json:"flags,omitempty"`
	ExtraCommands         []FmgCmdbRequest `json:"extra commands,omitempty"`
	DeviceBlueprint       *string          `json:"device blueprint,omitempty"`
	AuthorizationTemplate *string          `json:"authorization template,omitempty"`
	DeviceAction          *string          `json:"device action,omitempty"`
	FazPerm               *int64           `json:"faz.perm,omitempty"`
	FazQuota              *int64           `json:"faz.quota,omitempty"`
}

type DelDevice struct {
	Adom   *string  `json:"adom,omitempty"`
	Device *string  `json:"device,omitempty"`
	Flags  []string `json:"flags,omitempty"`
}
