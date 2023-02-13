package models

const VpnL2tpPath = "vpn/l2tp/"

type VpnL2tp struct {
	Compress        *string `json:"compress,omitempty"`
	Eip             *string `json:"eip,omitempty"`
	EnforceIpsec    *string `json:"enforce-ipsec,omitempty"`
	HelloInterval   *int64  `json:"hello-interval,omitempty"`
	LcpEchoInterval *int64  `json:"lcp-echo-interval,omitempty"`
	LcpMaxEchoFails *int64  `json:"lcp-max-echo-fails,omitempty"`
	Sip             *string `json:"sip,omitempty"`
	Status          *string `json:"status,omitempty"`
	Usrgrp          *string `json:"usrgrp,omitempty"`
}
