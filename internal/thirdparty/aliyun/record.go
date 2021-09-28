package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

// GetDomainRecords 获取解析记录列表
func (aliyun *Aliyun) GetDomainRecords(domain, _type, status, keyword, rrKeyword, typeKeyword, valueKeyword string, gid, page, size int, direction string) ([]byte, error) {
	request := alidns.CreateDescribeDomainRecordsRequest()
	request.DomainName = domain
	request.Type = _type
	request.Status = status
	request.KeyWord = keyword
	request.RRKeyWord = rrKeyword
	request.TypeKeyWord = typeKeyword
	request.ValueKeyWord = valueKeyword
	// request.GroupId = requests.NewInteger(gid)
	request.PageNumber = requests.NewInteger(page)
	request.PageSize = requests.NewInteger(size)
	request.Direction = direction
	response, err := aliyun.Client.DescribeDomainRecords(request)
	if err != nil {
		return nil, err
	}
	return response.BaseResponse.GetHttpContentBytes(), nil
}

// GetDomainSubRecords 获取子域名解析记录列表
func (aliyun *Aliyun) GetDomainSubRecords(subDomain, domain, _type, line string, page, size int) ([]byte, error) {
	request := alidns.CreateDescribeSubDomainRecordsRequest()
	request.SubDomain = subDomain
	request.DomainName = domain
	request.Type = _type
	request.Line = line
	request.PageNumber = requests.NewInteger(page)
	request.PageSize = requests.NewInteger(size)
	response, err := aliyun.Client.DescribeSubDomainRecords(request)
	if err != nil {
		return nil, err
	}
	return response.BaseResponse.GetHttpContentBytes(), nil
}

// GetDomainRecord 获取解析记录信息
func (aliyun *Aliyun) GetDomainRecord(id string) ([]byte, error) {
	request := alidns.CreateDescribeDomainRecordInfoRequest()
	request.RecordId = id
	response, err := aliyun.Client.DescribeDomainRecordInfo(request)
	if err != nil {
		return nil, err
	}
	return response.BaseResponse.GetHttpContentBytes(), nil
}

// AddDomainRecord 添加解析记录
func (aliyun *Aliyun) AddDomainRecord(domain, rr, _type, value string, ttl, priority int, line string) ([]byte, error) {
	request := alidns.CreateAddDomainRecordRequest()
	request.DomainName = domain
	request.RR = rr
	request.Type = _type
	request.Value = value
	request.Priority = requests.NewInteger(priority)
	request.TTL = requests.NewInteger(ttl)
	request.Line = line
	response, err := aliyun.Client.AddDomainRecord(request)
	if err != nil {
		return nil, err
	}
	return response.BaseResponse.GetHttpContentBytes(), nil
}

// UpdateDomainRecord 修改解析记录
func (aliyun *Aliyun) UpdateDomainRecord(id, rr, _type, value string, ttl, priority int, line string) ([]byte, error) {
	request := alidns.CreateUpdateDomainRecordRequest()
	request.RecordId = id
	request.RR = rr
	request.Type = _type
	request.Value = value
	request.Priority = requests.NewInteger(priority)
	request.TTL = requests.NewInteger(ttl)
	request.Line = line
	response, err := aliyun.Client.UpdateDomainRecord(request)
	if err != nil {
		return nil, err
	}
	return response.BaseResponse.GetHttpContentBytes(), nil
}

// DeleteDomainRecord 删除解析记录
func (aliyun *Aliyun) DeleteDomainRecord(id string) ([]byte, error) {
	request := alidns.CreateDeleteDomainRecordRequest()
	request.RecordId = id
	response, err := aliyun.Client.DeleteDomainRecord(request)
	if err != nil {
		return nil, err
	}
	return response.BaseResponse.GetHttpContentBytes(), nil
}

// GetDomainRecordLogs 获取解析记录操作日志
func (aliyun *Aliyun) GetDomainRecordLogs(domain, keyword, lang, start, end string, page, size int) ([]byte, error) {
	request := alidns.CreateDescribeRecordLogsRequest()
	request.DomainName = domain
	request.KeyWord = keyword
	request.Lang = lang
	request.StartDate = start
	request.EndDate = end
	request.PageNumber = requests.NewInteger(page)
	request.PageSize = requests.NewInteger(size)
	response, err := aliyun.Client.DescribeRecordLogs(request)
	if err != nil {
		return nil, err
	}
	return response.BaseResponse.GetHttpContentBytes(), nil
}

// UpdateDomainRecordRemark 修改解析记录的备注
func (aliyun *Aliyun) UpdateDomainRecordRemark(id, remark string) ([]byte, error) {
	request := alidns.CreateUpdateDomainRecordRemarkRequest()
	request.RecordId = id
	request.Remark = remark
	response, err := aliyun.Client.UpdateDomainRecordRemark(request)
	if err != nil {
		return nil, err
	}
	return response.BaseResponse.GetHttpContentBytes(), nil
}

// SetDomainRecordStatus 设置解析记录状态
func (aliyun *Aliyun) SetDomainRecordStatus(id, status string) ([]byte, error) {
	request := alidns.CreateSetDomainRecordStatusRequest()
	request.RecordId = id
	request.Status = status
	response, err := aliyun.Client.SetDomainRecordStatus(request)
	if err != nil {
		return nil, err
	}
	return response.BaseResponse.GetHttpContentBytes(), nil
}
