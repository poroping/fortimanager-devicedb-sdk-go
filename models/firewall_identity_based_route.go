package models

const FirewallIdentityBasedRoutePath = "firewall/identity-based-route/"

type FirewallIdentityBasedRoute struct {
	Comments *string                           `json:"comments,omitempty"`
	Name     *string                           `json:"name,omitempty"`
	Rule     *[]FirewallIdentityBasedRouteRule `json:"rule,omitempty"`
}

type FirewallIdentityBasedRouteRule struct {
	Device  *string                                 `json:"device,omitempty"`
	Gateway *string                                 `json:"gateway,omitempty"`
	Groups  *[]FirewallIdentityBasedRouteRuleGroups `json:"groups,omitempty"`
	Id      *int64                                  `json:"id,omitempty"`
}

type FirewallIdentityBasedRouteRuleGroups struct {
	Name *string `json:"name,omitempty"`
}
