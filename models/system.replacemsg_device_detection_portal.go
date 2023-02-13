package models

const SystemReplacemsgDeviceDetectionPortalPath = "system.replacemsg/device-detection-portal/"

type SystemReplacemsgDeviceDetectionPortal struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}
