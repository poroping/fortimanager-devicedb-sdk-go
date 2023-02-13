package models

const VpnIpsecPhase1Path = "vpn.ipsec/phase1/"

type VpnIpsecPhase1 struct {
	AcctVerify                *string                           `json:"acct-verify,omitempty"`
	AddGwRoute                *string                           `json:"add-gw-route,omitempty"`
	AddRoute                  *string                           `json:"add-route,omitempty"`
	AssignIp                  *string                           `json:"assign-ip,omitempty"`
	AssignIpFrom              *string                           `json:"assign-ip-from,omitempty"`
	Authmethod                *string                           `json:"authmethod,omitempty"`
	AuthmethodRemote          *string                           `json:"authmethod-remote,omitempty"`
	Authpasswd                *string                           `json:"authpasswd,omitempty"`
	Authusr                   *string                           `json:"authusr,omitempty"`
	Authusrgrp                *string                           `json:"authusrgrp,omitempty"`
	AutoNegotiate             *string                           `json:"auto-negotiate,omitempty"`
	BackupGateway             *[]VpnIpsecPhase1BackupGateway    `json:"backup-gateway,omitempty"`
	Banner                    *string                           `json:"banner,omitempty"`
	CertIdValidation          *string                           `json:"cert-id-validation,omitempty"`
	Certificate               *[]VpnIpsecPhase1Certificate      `json:"certificate,omitempty"`
	ChildlessIke              *string                           `json:"childless-ike,omitempty"`
	ClientAutoNegotiate       *string                           `json:"client-auto-negotiate,omitempty"`
	ClientKeepAlive           *string                           `json:"client-keep-alive,omitempty"`
	Comments                  *string                           `json:"comments,omitempty"`
	DhcpRaGiaddr              *string                           `json:"dhcp-ra-giaddr,omitempty"`
	Dhcp6RaLinkaddr           *string                           `json:"dhcp6-ra-linkaddr,omitempty"`
	Dhgrp                     *string                           `json:"dhgrp,omitempty"`
	DigitalSignatureAuth      *string                           `json:"digital-signature-auth,omitempty"`
	Distance                  *int64                            `json:"distance,omitempty"`
	DnsMode                   *string                           `json:"dns-mode,omitempty"`
	Domain                    *string                           `json:"domain,omitempty"`
	Dpd                       *string                           `json:"dpd,omitempty"`
	DpdRetrycount             *int64                            `json:"dpd-retrycount,omitempty"`
	DpdRetryinterval          *string                           `json:"dpd-retryinterval,omitempty"`
	Eap                       *string                           `json:"eap,omitempty"`
	EapExcludePeergrp         *string                           `json:"eap-exclude-peergrp,omitempty"`
	EapIdentity               *string                           `json:"eap-identity,omitempty"`
	EnforceUniqueId           *string                           `json:"enforce-unique-id,omitempty"`
	Esn                       *string                           `json:"esn,omitempty"`
	FecBase                   *int64                            `json:"fec-base,omitempty"`
	FecCodec                  *string                           `json:"fec-codec,omitempty"`
	FecEgress                 *string                           `json:"fec-egress,omitempty"`
	FecHealthCheck            *string                           `json:"fec-health-check,omitempty"`
	FecIngress                *string                           `json:"fec-ingress,omitempty"`
	FecMappingProfile         *string                           `json:"fec-mapping-profile,omitempty"`
	FecReceiveTimeout         *int64                            `json:"fec-receive-timeout,omitempty"`
	FecRedundant              *int64                            `json:"fec-redundant,omitempty"`
	FecSendTimeout            *int64                            `json:"fec-send-timeout,omitempty"`
	ForticlientEnforcement    *string                           `json:"forticlient-enforcement,omitempty"`
	Fragmentation             *string                           `json:"fragmentation,omitempty"`
	FragmentationMtu          *int64                            `json:"fragmentation-mtu,omitempty"`
	GroupAuthentication       *string                           `json:"group-authentication,omitempty"`
	GroupAuthenticationSecret *string                           `json:"group-authentication-secret,omitempty"`
	HaSyncEspSeqno            *string                           `json:"ha-sync-esp-seqno,omitempty"`
	IdleTimeout               *string                           `json:"idle-timeout,omitempty"`
	IdleTimeoutinterval       *int64                            `json:"idle-timeoutinterval,omitempty"`
	IkeVersion                *string                           `json:"ike-version,omitempty"`
	IncludeLocalLan           *string                           `json:"include-local-lan,omitempty"`
	Interface                 *string                           `json:"interface,omitempty"`
	IpDelayInterval           *int64                            `json:"ip-delay-interval,omitempty"`
	Ipv4DnsServer1            *string                           `json:"ipv4-dns-server1,omitempty"`
	Ipv4DnsServer2            *string                           `json:"ipv4-dns-server2,omitempty"`
	Ipv4DnsServer3            *string                           `json:"ipv4-dns-server3,omitempty"`
	Ipv4EndIp                 *string                           `json:"ipv4-end-ip,omitempty"`
	Ipv4ExcludeRange          *[]VpnIpsecPhase1Ipv4ExcludeRange `json:"ipv4-exclude-range,omitempty"`
	Ipv4Name                  *string                           `json:"ipv4-name,omitempty"`
	Ipv4Netmask               *string                           `json:"ipv4-netmask,omitempty"`
	Ipv4SplitExclude          *string                           `json:"ipv4-split-exclude,omitempty"`
	Ipv4SplitInclude          *string                           `json:"ipv4-split-include,omitempty"`
	Ipv4StartIp               *string                           `json:"ipv4-start-ip,omitempty"`
	Ipv4WinsServer1           *string                           `json:"ipv4-wins-server1,omitempty"`
	Ipv4WinsServer2           *string                           `json:"ipv4-wins-server2,omitempty"`
	Ipv6DnsServer1            *string                           `json:"ipv6-dns-server1,omitempty"`
	Ipv6DnsServer2            *string                           `json:"ipv6-dns-server2,omitempty"`
	Ipv6DnsServer3            *string                           `json:"ipv6-dns-server3,omitempty"`
	Ipv6EndIp                 *string                           `json:"ipv6-end-ip,omitempty"`
	Ipv6ExcludeRange          *[]VpnIpsecPhase1Ipv6ExcludeRange `json:"ipv6-exclude-range,omitempty"`
	Ipv6Name                  *string                           `json:"ipv6-name,omitempty"`
	Ipv6Prefix                *int64                            `json:"ipv6-prefix,omitempty"`
	Ipv6SplitExclude          *string                           `json:"ipv6-split-exclude,omitempty"`
	Ipv6SplitInclude          *string                           `json:"ipv6-split-include,omitempty"`
	Ipv6StartIp               *string                           `json:"ipv6-start-ip,omitempty"`
	Keepalive                 *int64                            `json:"keepalive,omitempty"`
	Keylife                   *int64                            `json:"keylife,omitempty"`
	LocalGw                   *string                           `json:"local-gw,omitempty"`
	Localid                   *string                           `json:"localid,omitempty"`
	LocalidType               *string                           `json:"localid-type,omitempty"`
	LoopbackAsymroute         *string                           `json:"loopback-asymroute,omitempty"`
	MeshSelectorType          *string                           `json:"mesh-selector-type,omitempty"`
	Mode                      *string                           `json:"mode,omitempty"`
	ModeCfg                   *string                           `json:"mode-cfg,omitempty"`
	Name                      *string                           `json:"name,omitempty"`
	Nattraversal              *string                           `json:"nattraversal,omitempty"`
	NegotiateTimeout          *int64                            `json:"negotiate-timeout,omitempty"`
	NetworkId                 *int64                            `json:"network-id,omitempty"`
	NetworkOverlay            *string                           `json:"network-overlay,omitempty"`
	NpuOffload                *string                           `json:"npu-offload,omitempty"`
	Peer                      *string                           `json:"peer,omitempty"`
	Peergrp                   *string                           `json:"peergrp,omitempty"`
	Peerid                    *string                           `json:"peerid,omitempty"`
	Peertype                  *string                           `json:"peertype,omitempty"`
	Ppk                       *string                           `json:"ppk,omitempty"`
	PpkIdentity               *string                           `json:"ppk-identity,omitempty"`
	PpkSecret                 *string                           `json:"ppk-secret,omitempty"`
	Priority                  *int64                            `json:"priority,omitempty"`
	Proposal                  *string                           `json:"proposal,omitempty"`
	Psksecret                 *string                           `json:"psksecret,omitempty"`
	PsksecretRemote           *string                           `json:"psksecret-remote,omitempty"`
	Reauth                    *string                           `json:"reauth,omitempty"`
	Rekey                     *string                           `json:"rekey,omitempty"`
	RemoteGw                  *string                           `json:"remote-gw,omitempty"`
	RemotegwDdns              *string                           `json:"remotegw-ddns,omitempty"`
	RsaSignatureFormat        *string                           `json:"rsa-signature-format,omitempty"`
	SavePassword              *string                           `json:"save-password,omitempty"`
	SendCertChain             *string                           `json:"send-cert-chain,omitempty"`
	SignatureHashAlg          *string                           `json:"signature-hash-alg,omitempty"`
	SplitIncludeService       *string                           `json:"split-include-service,omitempty"`
	SuiteB                    *string                           `json:"suite-b,omitempty"`
	Type                      *string                           `json:"type,omitempty"`
	UnitySupport              *string                           `json:"unity-support,omitempty"`
	Usrgrp                    *string                           `json:"usrgrp,omitempty"`
	WizardType                *string                           `json:"wizard-type,omitempty"`
	Xauthtype                 *string                           `json:"xauthtype,omitempty"`
}

type VpnIpsecPhase1BackupGateway struct {
	Address *string `json:"address,omitempty"`
}

type VpnIpsecPhase1Certificate struct {
	Name *string `json:"name,omitempty"`
}

type VpnIpsecPhase1Ipv4ExcludeRange struct {
	EndIp   *string `json:"end-ip,omitempty"`
	Id      *int64  `json:"id,omitempty"`
	StartIp *string `json:"start-ip,omitempty"`
}

type VpnIpsecPhase1Ipv6ExcludeRange struct {
	EndIp   *string `json:"end-ip,omitempty"`
	Id      *int64  `json:"id,omitempty"`
	StartIp *string `json:"start-ip,omitempty"`
}
