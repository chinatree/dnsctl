package domain

type Domains struct {
	Domains struct {
		Domain []Domain `json:"Domain"`
	} `json:"Domains"`
	PageNumber int
	PageSize   int
	RequestID  string
	TotalCount int
}

type Domain struct {
	AliDomain       bool   `json:"AliDomain"`
	CreateTime      string `json:"CreateTime"`
	CreateTimestamp int64  `json:"CreateTimestamp"`
	DnsServers      struct {
		DnsServer []string `json:"DnsServer"`
	} `json:"DnsServers"`
	DomainId        string `json:"DomainId"`
	DomainName      string `json:"DomainName"`
	GroupId         string `json:"GroupId"`
	GroupName       string `json:"GroupName"`
	PunyCode        string `json:"PunyCode"`
	RecordCount     int    `json:"RecordCount"`
	ResourceGroupId string `json:"ResourceGroupId"`
	Remark          string `json:"Remark"`
	Starmark        bool   `json:"Starmark"`
	VersionCode     string `json:"VersionCode"`
	VersionName     string `json:"VersionName"`
}

type MainDomain struct {
	DomainLevel int    `json:"DomainLevel"`
	DomainName  string `json:"DomainName"`
	RR          string `json:"RR"`
}

type DomainLogs struct {
	DomainLogs struct {
		DomainLog []DomainLog `json:"DomainLog"`
	} `json:"DomainLogs"`
	PageNumber int    `json:"PageNumber"`
	PageSize   int    `json:"PageSize"`
	RequestID  string `json:"RequestID"`
	TotalCount int    `json:"TotalCount"`
}

type DomainLog struct {
	Action          string `json:"Action"`
	ActionTime      string `json:"ActionTime"`
	ActionTimestamp int64  `json:"ActionTimestamp"`
	DomainName      string `json:"DomainName"`
	Message         string `json:"Message"`
}
