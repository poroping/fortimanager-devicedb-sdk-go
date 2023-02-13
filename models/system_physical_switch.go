package models

const SystemPhysicalSwitchPath = "system/physical-switch/"

type SystemPhysicalSwitch struct {
	AgeEnable *string `json:"age-enable,omitempty"`
	AgeVal    *int64  `json:"age-val,omitempty"`
	Name      *string `json:"name,omitempty"`
}
