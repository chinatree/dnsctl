package record

// DomainRecordList 解析记录列表返回体
type DomainRecordList struct {
	DomainRecords List `json:"DomainRecords"`
	PageNumber    int
	PageSize      int
	RequestID     string
	TotalCount    int
}

// List 解析记录列表
type List struct {
	Records []Record `json:"Record"`
}

// Record 解析记录信息
type Record struct {
	DomainName string `json:"DomainName"`
	Line       string `json:"Line"`
	Locked     bool   `json:"Locked"`
	RR         string `json:"RR"`
	RecordID   string `json:"RecordId"`
	Status     string `json:"Status"`
	TTL        int64  `json:"TTL"`
	Type       string `json:"Type"`
	Value      string `json:"Value"`
	Weight     int    `json:"Weight"`
	Remark     string `json:"Remark"`
}

// Logs 解析记录操作日志列表
type Logs struct {
	Logs struct {
		RecordLog []Log `json:"RecordLog"`
	} `json:"RecordLogs"`
	PageNumber int    `json:"PageNumber"`
	PageSize   int    `json:"PageSize"`
	RequestID  string `json:"RequestID"`
	TotalCount int    `json:"TotalCount"`
}

// Log 解析记录操作日志信息
type Log struct {
	Action          string `json:"Action"`
	ActionTime      string `json:"ActionTime"`
	ActionTimestamp int64  `json:"ActionTimestamp"`
	Message         string `json:"Message"`
}
