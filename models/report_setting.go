package models

const ReportSettingPath = "report/setting/"

type ReportSetting struct {
	Fortiview            *string `json:"fortiview,omitempty"`
	PdfReport            *string `json:"pdf-report,omitempty"`
	ReportSource         *string `json:"report-source,omitempty"`
	TopN                 *int64  `json:"top-n,omitempty"`
	WebBrowsingThreshold *int64  `json:"web-browsing-threshold,omitempty"`
}
