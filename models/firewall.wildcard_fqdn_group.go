package models

const FirewallWildcardFqdnGroupPath = "firewall.wildcard-fqdn/group/"

type FirewallWildcardFqdnGroup struct {
	Color      *int64                             `json:"color,omitempty"`
	Comment    *string                            `json:"comment,omitempty"`
	Member     *[]FirewallWildcardFqdnGroupMember `json:"member,omitempty"`
	Name       *string                            `json:"name,omitempty"`
	Uuid       *string                            `json:"uuid,omitempty"`
	Visibility *string                            `json:"visibility,omitempty"`
}

type FirewallWildcardFqdnGroupMember struct {
	Name *string `json:"name,omitempty"`
}
