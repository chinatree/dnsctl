package tencent

import (
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

// GetDomainList 获取域名列表
func (sdk *Tencent) GetDomainList(
	keyword, _type string, gid, offset, limit int64) ([]byte, error) {
	request := dnspod.NewDescribeDomainListRequest()
	request.Keyword = &keyword
	request.GroupId = &gid
	request.Type = &_type
	request.Offset = &offset
	request.Limit = &limit
	response, err := sdk.Client.DescribeDomainList(request)
	if err != nil {
		return nil, err
	}
	return []byte(response.ToJsonString()), nil
}

// GetDomain 获取域名信息
func (sdk *Tencent) GetDomain(name string, id uint64) ([]byte, error) {
	request := dnspod.NewDescribeDomainRequest()
	request.Domain = &name
	request.DomainId = &id
	response, err := sdk.Client.DescribeDomain(request)
	if err != nil {
		return nil, err
	}
	return []byte(response.ToJsonString()), nil
}

// AddDomain 添加域名
func (sdk *Tencent) AddDomain(name, isMark string, gid uint64) ([]byte, error) {
	request := dnspod.NewCreateDomainRequest()
	request.Domain = &name
	request.GroupId = &gid
	request.IsMark = &isMark
	response, err := sdk.Client.CreateDomain(request)
	if err != nil {
		return nil, err
	}
	return []byte(response.ToJsonString()), nil
}

// DeleteDomain 删除域名
func (sdk *Tencent) DeleteDomain(name string, id uint64) ([]byte, error) {
	request := dnspod.NewDeleteDomainRequest()
	request.Domain = &name
	request.DomainId = &id
	response, err := sdk.Client.DeleteDomain(request)
	if err != nil {
		return nil, err
	}
	return []byte(response.ToJsonString()), nil
}

// GetDomainLogs 获取域名操作日志
func (sdk *Tencent) GetDomainLogs(
	name string, id, offset, limit uint64) ([]byte, error) {
	request := dnspod.NewDescribeDomainLogListRequest()
	request.Domain = &name
	request.DomainId = &id
	request.Offset = &offset
	request.Limit = &limit
	response, err := sdk.Client.DescribeDomainLogList(request)
	if err != nil {
		return nil, err
	}
	return []byte(response.ToJsonString()), nil
}

// UpdateDomainRemark 修改域名的备注
func (sdk *Tencent) UpdateDomainRemark(name, remark string, id uint64) ([]byte, error) {
	request := dnspod.NewModifyDomainRemarkRequest()
	request.Domain = &name
	request.DomainId = &id
	request.Remark = &remark
	response, err := sdk.Client.ModifyDomainRemark(request)
	if err != nil {
		return nil, err
	}
	return []byte(response.ToJsonString()), nil
}

// UpdateDomainStatus 修改域名状态
func (sdk *Tencent) UpdateDomainStatus(name, status string, id uint64) ([]byte, error) {
	request := dnspod.NewModifyDomainStatusRequest()
	request.Domain = &name
	request.DomainId = &id
	request.Status = &status
	response, err := sdk.Client.ModifyDomainStatus(request)
	if err != nil {
		return nil, err
	}
	return []byte(response.ToJsonString()), nil
}

// LockDomain 锁定域名
func (sdk *Tencent) LockDomain(name string, id, day uint64) ([]byte, error) {
	request := dnspod.NewModifyDomainLockRequest()
	request.Domain = &name
	request.DomainId = &id
	request.LockDays = &day
	response, err := sdk.Client.ModifyDomainLock(request)
	if err != nil {
		return nil, err
	}
	return []byte(response.ToJsonString()), nil
}

// UnlockDomain 域名锁定解锁
func (sdk *Tencent) UnlockDomain(name, code string, id uint64) ([]byte, error) {
	request := dnspod.NewModifyDomainUnlockRequest()
	request.Domain = &name
	request.DomainId = &id
	request.LockCode = &code
	response, err := sdk.Client.ModifyDomainUnlock(request)
	if err != nil {
		return nil, err
	}
	return []byte(response.ToJsonString()), nil
}
