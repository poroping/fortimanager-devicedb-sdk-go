package models

const VpnIpsecFecPath = "vpn.ipsec/fec/"

type VpnIpsecFec struct {
	Mappings *[]VpnIpsecFecMappings `json:"mappings,omitempty"`
	Name     *string                `json:"name,omitempty"`
}

type VpnIpsecFecMappings struct {
	BandwidthBiThreshold   *int64 `json:"bandwidth-bi-threshold,omitempty"`
	BandwidthDownThreshold *int64 `json:"bandwidth-down-threshold,omitempty"`
	BandwidthUpThreshold   *int64 `json:"bandwidth-up-threshold,omitempty"`
	Base                   *int64 `json:"base,omitempty"`
	LatencyThreshold       *int64 `json:"latency-threshold,omitempty"`
	PacketLossThreshold    *int64 `json:"packet-loss-threshold,omitempty"`
	Redundant              *int64 `json:"redundant,omitempty"`
	Seqno                  *int64 `json:"seqno,omitempty"`
}
