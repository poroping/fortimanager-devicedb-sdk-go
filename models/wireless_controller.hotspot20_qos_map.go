package models

const WirelessControllerHotspot20QosMapPath = "wireless-controller.hotspot20/qos-map/"

type WirelessControllerHotspot20QosMap struct {
	DscpExcept *[]WirelessControllerHotspot20QosMapDscpExcept `json:"dscp-except,omitempty"`
	DscpRange  *[]WirelessControllerHotspot20QosMapDscpRange  `json:"dscp-range,omitempty"`
	Name       *string                                        `json:"name,omitempty"`
}

type WirelessControllerHotspot20QosMapDscpExcept struct {
	Dscp  *int64 `json:"dscp,omitempty"`
	Index *int64 `json:"index,omitempty"`
	Up    *int64 `json:"up,omitempty"`
}

type WirelessControllerHotspot20QosMapDscpRange struct {
	High  *int64 `json:"high,omitempty"`
	Index *int64 `json:"index,omitempty"`
	Low   *int64 `json:"low,omitempty"`
	Up    *int64 `json:"up,omitempty"`
}
