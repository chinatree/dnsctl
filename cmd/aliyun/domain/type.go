package domain

// Domains 域名列表
type Domains struct {
	Domains struct {
		Domain []Domain `json:"Domain"`
	} `json:"Domains"`
	PageNumber int
	PageSize   int
	RequestID  string
	TotalCount int
}

// Domain 域名信息
type Domain struct {
	AliDomain       bool   `json:"AliDomain"`
	CreateTime      string `json:"CreateTime"`
	CreateTimestamp int64  `json:"CreateTimestamp"`
	DNSServers      struct {
		DNSServer []string `json:"DnsServer"`
	} `json:"DnsServers"`
	DomainID        string `json:"DomainId"`
	DomainName      string `json:"DomainName"`
	GroupID         string `json:"GroupId"`
	GroupName       string `json:"GroupName"`
	PunyCode        string `json:"PunyCode"`
	RecordCount     int    `json:"RecordCount"`
	ResourceGroupID string `json:"ResourceGroupId"`
	Remark          string `json:"Remark"`
	Starmark        bool   `json:"Starmark"`
	VersionCode     string `json:"VersionCode"`
	VersionName     string `json:"VersionName"`
}

// MainDomain 主域名信息
type MainDomain struct {
	DomainLevel int    `json:"DomainLevel"`
	DomainName  string `json:"DomainName"`
	RR          string `json:"RR"`
}

// Logs 域名操作日志列表
type Logs struct {
	DomainLogs struct {
		DomainLog []Log `json:"DomainLog"`
	} `json:"DomainLogs"`
	PageNumber int    `json:"PageNumber"`
	PageSize   int    `json:"PageSize"`
	RequestID  string `json:"RequestID"`
	TotalCount int    `json:"TotalCount"`
}

// Log 域名操作日志信息
type Log struct {
	Action          string `json:"Action"`
	ActionTime      string `json:"ActionTime"`
	ActionTimestamp int64  `json:"ActionTimestamp"`
	DomainName      string `json:"DomainName"`
	Message         string `json:"Message"`
}
