package models

const WirelessControllerWtpGroupPath = "wireless-controller/wtp-group/"

type WirelessControllerWtpGroup struct {
	Name         *string                           `json:"name,omitempty"`
	PlatformType *string                           `json:"platform-type,omitempty"`
	Wtps         *[]WirelessControllerWtpGroupWtps `json:"wtps,omitempty"`
}

type WirelessControllerWtpGroupWtps struct {
	WtpId *string `json:"wtp-id,omitempty"`
}
