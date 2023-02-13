package models

const FirewallAccessProxyVirtualHostPath = "firewall/access-proxy-virtual-host/"

type FirewallAccessProxyVirtualHost struct {
	Host           *string `json:"host,omitempty"`
	HostType       *string `json:"host-type,omitempty"`
	Name           *string `json:"name,omitempty"`
	SslCertificate *string `json:"ssl-certificate,omitempty"`
}
