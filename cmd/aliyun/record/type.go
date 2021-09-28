package record

type DomainRecordList struct {
	DomainRecords RecordList `json:"DomainRecords"`
	PageNumber    int
	PageSize      int
	RequestID     string
	TotalCount    int
}

type RecordList struct {
	Records []Record `json:"Record"`
}

type Record struct {
	DomainName string `json:"DomainName"`
	Line       string `json:"Line"`
	Locked     bool   `json:"Locked"`
	RR         string `json:"RR"`
	RecordId   string `json:"RecordId"`
	Status     string `json:"Status"`
	TTL        int64  `json:"TTL"`
	Type       string `json:"Type"`
	Value      string `json:"Value"`
	Weight     int    `json:"Weight"`
	Remark     string `json:"Remark"`
}

type RecordLogs struct {
	RecordLogs struct {
		RecordLog []RecordLog `json:"RecordLog"`
	} `json:"RecordLogs"`
	PageNumber int    `json:"PageNumber"`
	PageSize   int    `json:"PageSize"`
	RequestID  string `json:"RequestID"`
	TotalCount int    `json:"TotalCount"`
}

type RecordLog struct {
	Action          string `json:"Action"`
	ActionTime      string `json:"ActionTime"`
	ActionTimestamp int64  `json:"ActionTimestamp"`
	Message         string `json:"Message"`
}
