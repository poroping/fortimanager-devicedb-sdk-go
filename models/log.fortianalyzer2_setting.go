package models

const LogFortianalyzer2SettingPath = "log.fortianalyzer2/setting/"

type LogFortianalyzer2Setting struct {
	AccessConfig              *string                           `json:"access-config,omitempty"`
	Certificate               *string                           `json:"certificate,omitempty"`
	CertificateVerification   *string                           `json:"certificate-verification,omitempty"`
	ConnTimeout               *int64                            `json:"conn-timeout,omitempty"`
	EncAlgorithm              *string                           `json:"enc-algorithm,omitempty"`
	HmacAlgorithm             *string                           `json:"hmac-algorithm,omitempty"`
	Interface                 *string                           `json:"interface,omitempty"`
	InterfaceSelectMethod     *string                           `json:"interface-select-method,omitempty"`
	IpsArchive                *string                           `json:"ips-archive,omitempty"`
	MaxLogRate                *int64                            `json:"max-log-rate,omitempty"`
	MonitorFailureRetryPeriod *int64                            `json:"monitor-failure-retry-period,omitempty"`
	MonitorKeepalivePeriod    *int64                            `json:"monitor-keepalive-period,omitempty"`
	PresharedKey              *string                           `json:"preshared-key,omitempty"`
	Priority                  *string                           `json:"priority,omitempty"`
	Reliable                  *string                           `json:"reliable,omitempty"`
	Serial                    *[]LogFortianalyzer2SettingSerial `json:"serial,omitempty"`
	Server                    *string                           `json:"server,omitempty"`
	SourceIp                  *string                           `json:"source-ip,omitempty"`
	SslMinProtoVersion        *string                           `json:"ssl-min-proto-version,omitempty"`
	Status                    *string                           `json:"status,omitempty"`
	UploadDay                 *string                           `json:"upload-day,omitempty"`
	UploadInterval            *string                           `json:"upload-interval,omitempty"`
	UploadOption              *string                           `json:"upload-option,omitempty"`
	UploadTime                *string                           `json:"upload-time,omitempty"`
}

type LogFortianalyzer2SettingSerial struct {
	Name *string `json:"name,omitempty"`
}
