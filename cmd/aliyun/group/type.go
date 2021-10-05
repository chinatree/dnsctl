package group

// DomainGroups 域名分组列表
type DomainGroups struct {
	DomainGroups struct {
		DomainGroup []DomainGroup `json:"DomainGroup"`
	} `json:"DomainGroups"`
	PageNumber int    `json:"PageNumber"`
	PageSize   int    `json:"PageSize"`
	RequestID  string `json:"RequestID"`
	TotalCount int    `json:"TotalCount"`
}

// DomainGroup 域名分组信息
type DomainGroup struct {
	DomainCount int64  `json:"DomainCount"`
	GroupID     string `json:"GroupId"`
	GroupName   string `json:"GroupName"`
}
