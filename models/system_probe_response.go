package models

const SystemProbeResponsePath = "system/probe-response/"

type SystemProbeResponse struct {
	HttpProbeValue *string `json:"http-probe-value,omitempty"`
	Mode           *string `json:"mode,omitempty"`
	Password       *string `json:"password,omitempty"`
	Port           *int64  `json:"port,omitempty"`
	SecurityMode   *string `json:"security-mode,omitempty"`
	Timeout        *int64  `json:"timeout,omitempty"`
	TtlMode        *string `json:"ttl-mode,omitempty"`
}
