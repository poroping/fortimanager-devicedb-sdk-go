package models

const LogFortianalyzerFilterPath = "log.fortianalyzer/filter/"

type LogFortianalyzerFilter struct {
	Anomaly          *string                            `json:"anomaly,omitempty"`
	DlpArchive       *string                            `json:"dlp-archive,omitempty"`
	Filter           *string                            `json:"filter,omitempty"`
	FilterType       *string                            `json:"filter-type,omitempty"`
	ForwardTraffic   *string                            `json:"forward-traffic,omitempty"`
	FreeStyle        *[]LogFortianalyzerFilterFreeStyle `json:"free-style,omitempty"`
	Gtp              *string                            `json:"gtp,omitempty"`
	LocalTraffic     *string                            `json:"local-traffic,omitempty"`
	MulticastTraffic *string                            `json:"multicast-traffic,omitempty"`
	Severity         *string                            `json:"severity,omitempty"`
	SnifferTraffic   *string                            `json:"sniffer-traffic,omitempty"`
	Voip             *string                            `json:"voip,omitempty"`
	ZtnaTraffic      *string                            `json:"ztna-traffic,omitempty"`
}

type LogFortianalyzerFilterFreeStyle struct {
	Category   *string `json:"category,omitempty"`
	Filter     *string `json:"filter,omitempty"`
	FilterType *string `json:"filter-type,omitempty"`
	Id         *int64  `json:"id,omitempty"`
}
