package domain

// List 域名列表
type List struct {
	Response struct {
		// 列表页统计信息
		DomainCountInfo CountInfo `json:"DomainCountInfo"`
		// 域名列表
		DomainList []Domain `json:"DomainList"`
		RequestID  string   `json:"RequestId"`
	} `json:"Response"`
}

// CountInfo 域名的统计信息
type CountInfo struct {
	// 符合条件的域名数量
	DomainTotal uint64 `json:"DomainTotal"`
	// 用户可以查看的所有域名数量
	AllTotal uint64 `json:"AllTotal"`
	// 用户账号添加的域名数量
	MineTotal uint64 `json:"MineTotal"`
	// 共享给用户的域名数量
	ShareTotal uint64 `json:"ShareTotal"`
	// 付费域名数量
	VipTotal uint64 `json:"VipTotal"`
	// 暂停的域名数量
	PauseTotal uint64 `json:"PauseTotal"`
	// dns设置错误的域名数量
	ErrorTotal uint64 `json:"ErrorTotal"`
	// 锁定的域名数量
	LockTotal uint64 `json:"LockTotal"`
	// 封禁的域名数量
	SpamTotal uint64 `json:"SpamTotal"`
	// 30天内即将到期的域名数量
	VipExpire uint64 `json:"VipExpire"`
	// 分享给其它人的域名数量
	ShareOutTotal uint64 `json:"ShareOutTotal"`
	// 指定分组内的域名数量
	GroupTotal uint64 `json:"GroupTotal"`
}

// Domain 域名信息
type Domain struct {
	DNSStatus    string   `json:"DNSStatus"`
	EffectiveDNS []string `json:"EffectiveDNS"`
	ID           uint64   `json:"DomainId"`
	Name         string   `json:"Name"`
	RecordCount  uint64   `json:"RecordCount"`
	Grade        string   `json:"Grade"`
	GradeLevel   uint64   `json:"GradeLevel"`
	GradeTitle   string   `json:"GradeTitle"`
	GroupID      uint64   `json:"GroupId"`
	IsVip        string   `json:"IsVip"`
	Remark       string   `json:"Remark"`
	Status       string   `json:"Status"`
	TTL          uint64   `json:"TTL"`
	Owner        string   `json:"Owner"`
}

// InfoResp 域名信息返回体
type InfoResp struct {
	Response struct {
		DomainInfo Info   `json:"DomainInfo"`
		RequestID  string `json:"RequestId"`
	} `json:"Response"`
}

// Info 域名信息
type Info struct {
	DNSStatus    string   `json:"DNSStatus"`
	DomainID     uint64   `json:"DomainId"`
	DnspodNsList []string `json:"DnspodNsList"`
	Domain       string   `json:"Domain"`
	RecordCount  uint64   `json:"RecordCount"`
	Grade        string   `json:"Grade"`
	GradeLevel   uint64   `json:"GradeLevel"`
	GradeTitle   string   `json:"GradeTitle"`
	GroupID      uint64   `json:"GroupId"`
	IsVip        string   `json:"IsVip"`
	Remark       string   `json:"Remark"`
	Status       string   `json:"Status"`
	TTL          uint64   `json:"TTL"`
	Owner        string   `json:"Owner"`
}

// LogList 域名操作日志列表
type LogList struct {
	Response struct {
		LogList    []string `json:"LogList"`
		PageSize   uint64   `json:"PageSize"`
		RequestID  string   `json:"RequestId"`
		TotalCount uint64   `json:"TotalCount"`
	} `json:"Response"`
}
