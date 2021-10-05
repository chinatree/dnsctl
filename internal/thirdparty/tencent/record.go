package tencent

import (
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

// GetRecordList 获取解析记录列表
func (sdk *Tencent) GetRecordList(
	domain, rr, keyword, _type, sort, sortField string,
	did, gid, offset, limit uint64) ([]byte, error) {
	request := dnspod.NewDescribeRecordListRequest()
	request.Domain = &domain
	request.Subdomain = &rr
	request.DomainId = &did
	request.Keyword = &keyword
	request.GroupId = &gid
	request.RecordType = &_type
	request.Offset = &offset
	request.Limit = &limit
	request.SortType = &sort
	request.SortField = &sortField
	response, err := sdk.Client.DescribeRecordList(request)
	if err != nil {
		return nil, err
	}
	return []byte(response.ToJsonString()), nil
}

// GetRecord 获取解析记录信息
func (sdk *Tencent) GetRecord(domain string, did, id uint64) ([]byte, error) {
	request := dnspod.NewDescribeRecordRequest()
	request.Domain = &domain
	request.DomainId = &did
	request.RecordId = &id
	response, err := sdk.Client.DescribeRecord(request)
	if err != nil {
		return nil, err
	}
	return []byte(response.ToJsonString()), nil
}

// AddRecord 修改解析记录
func (sdk *Tencent) AddRecord(
	domain, rr, _type, value, status, line, lid string,
	ttl, mx, did, weight uint64) ([]byte, error) {
	request := dnspod.NewCreateRecordRequest()
	request.Domain = &domain
	request.DomainId = &did
	request.SubDomain = &rr
	request.RecordType = &_type
	request.Value = &value
	request.Status = &status
	request.RecordLine = &line
	request.RecordLineId = &lid
	request.TTL = &ttl
	request.MX = &mx
	request.Weight = &weight
	response, err := sdk.Client.CreateRecord(request)
	if err != nil {
		return nil, err
	}
	return []byte(response.ToJsonString()), nil
}

// UpdateRecord 修改解析记录
func (sdk *Tencent) UpdateRecord(
	domain, rr, _type, value, status, line, lid string,
	id, ttl, mx, did, weight uint64) ([]byte, error) {
	request := dnspod.NewModifyRecordRequest()
	request.Domain = &domain
	request.DomainId = &did
	request.RecordId = &id
	request.SubDomain = &rr
	request.RecordType = &_type
	request.Value = &value
	request.Status = &status
	request.RecordLine = &line
	request.RecordLineId = &lid
	request.TTL = &ttl
	request.MX = &mx
	request.Weight = &weight
	response, err := sdk.Client.ModifyRecord(request)
	if err != nil {
		return nil, err
	}
	return []byte(response.ToJsonString()), nil
}

// DeleteRecord 删除解析记录
func (sdk *Tencent) DeleteRecord(domain string, id, did uint64) ([]byte, error) {
	request := dnspod.NewDeleteRecordRequest()
	request.Domain = &domain
	request.DomainId = &did
	request.RecordId = &id
	response, err := sdk.Client.DeleteRecord(request)
	if err != nil {
		return nil, err
	}
	return []byte(response.ToJsonString()), nil
}

// UpdateRecordRemark 修改解析记录的备注
func (sdk *Tencent) UpdateRecordRemark(domain, remark string,
	id, did uint64) ([]byte, error) {
	request := dnspod.NewModifyRecordRemarkRequest()
	request.Domain = &domain
	request.DomainId = &did
	request.RecordId = &id
	request.Remark = &remark
	response, err := sdk.Client.ModifyRecordRemark(request)
	if err != nil {
		return nil, err
	}
	return []byte(response.ToJsonString()), nil
}

// UpdateRecordStatus 修改解析记录的备注
func (sdk *Tencent) UpdateRecordStatus(domain, status string,
	id, did uint64) ([]byte, error) {
	request := dnspod.NewModifyRecordStatusRequest()
	request.Domain = &domain
	request.DomainId = &did
	request.RecordId = &id
	request.Status = &status
	response, err := sdk.Client.ModifyRecordStatus(request)
	if err != nil {
		return nil, err
	}
	return []byte(response.ToJsonString()), nil
}
