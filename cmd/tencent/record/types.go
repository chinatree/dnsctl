package record

type RecordList struct {
	Response struct {
		RecordCountInfo RecordCountInfo `json:"RecordCountInfo"`
		RecordList      []Record        `json:"RecordList"`
		RequestID       string          `json:"RequestId"`
	} `json:"Response"`
}

type RecordCountInfo struct {
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

type Record struct {
	ID      uint64 `json:"RecordId"`
	Line    string `json:"Line"`
	LineId  string `json:"LineId"`
	Name    string `json:"Name"`
	Remark  string `json:"Remark"`
	Status  string `json:"Status"`
	TTL     uint64 `json:"TTL"`
	Type    string `json:"Type"`
	Updated string `json:"UpdatedOn"`
	Value   string `json:"Value"`
	Weight  uint64 `json:"Weight"`
}

type RecordInfoResp struct {
	Response struct {
		RecordInfo RecordInfo `json:"RecordInfo"`
		RequestID  string     `json:"RequestId"`
	} `json:"Response"`
}

type RecordInfo struct {
	DomainID uint64 `json:"DomainId"`
	Enabled  uint64 `json:"Enabled"`
	ID       uint64 `json:"Id"`
	Line     string `json:"RecordLine"`
	LineID   string `json:"RecordLineId"`
	RecordID uint64 `json:"RecordId"`
	Remark   string `json:"Remark"`
	RR       string `json:"SubDomain"`
	TTL      uint64 `json:"TTL"`
	Type     string `json:"RecordType"`
	Updated  string `json:"UpdatedOn"`
	Value    string `json:"Value"`
	Weight   uint64 `json:"Weight"`
}
