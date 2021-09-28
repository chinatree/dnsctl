package group

type DomainGroups struct {
	DomainGroups struct {
		DomainGroup []DomainGroup `json:"DomainGroup"`
	} `json:"DomainGroups"`
	PageNumber int    `json:"PageNumber"`
	PageSize   int    `json:"PageSize"`
	RequestID  string `json:"RequestID"`
	TotalCount int    `json:"TotalCount"`
}

type DomainGroup struct {
	DomainCount int64  `json:"DomainCount"`
	GroupId     string `json:"GroupId"`
	GroupName   string `json:"GroupName"`
}
