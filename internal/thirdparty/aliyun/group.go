package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

// GetDomainGroups 获取域名分组列表
func (aliyun *Aliyun) GetDomainGroups(keyword string, page, size int) ([]byte, error) {
	request := alidns.CreateDescribeDomainGroupsRequest()
	request.KeyWord = keyword
	request.PageNumber = requests.NewInteger(page)
	request.PageSize = requests.NewInteger(size)
	response, err := aliyun.Client.DescribeDomainGroups(request)
	if err != nil {
		return nil, err
	}
	return response.BaseResponse.GetHttpContentBytes(), nil
}

// AddDomainGroup 添加域名分组
func (aliyun *Aliyun) AddDomainGroup(name string) ([]byte, error) {
	request := alidns.CreateAddDomainGroupRequest()
	request.GroupName = name
	response, err := aliyun.Client.AddDomainGroup(request)
	if err != nil {
		return nil, err
	}
	return response.BaseResponse.GetHttpContentBytes(), nil
}

// UpdateDomainGroup 修改域名分组
func (aliyun *Aliyun) UpdateDomainGroup(id, name string) ([]byte, error) {
	request := alidns.CreateUpdateDomainGroupRequest()
	request.GroupId = id
	request.GroupName = name
	response, err := aliyun.Client.UpdateDomainGroup(request)
	if err != nil {
		return nil, err
	}
	return response.BaseResponse.GetHttpContentBytes(), nil
}

// DeleteDomainGroup 删除域名分组
func (aliyun *Aliyun) DeleteDomainGroup(id string) ([]byte, error) {
	request := alidns.CreateDeleteDomainGroupRequest()
	request.GroupId = id
	response, err := aliyun.Client.DeleteDomainGroup(request)
	if err != nil {
		return nil, err
	}
	return response.BaseResponse.GetHttpContentBytes(), nil
}

// ChangeDomainGroup 更改域名分组
func (aliyun *Aliyun) ChangeDomainGroup(domain, id string) ([]byte, error) {
	request := alidns.CreateChangeDomainGroupRequest()
	request.DomainName = domain
	request.GroupId = id
	response, err := aliyun.Client.ChangeDomainGroup(request)
	if err != nil {
		return nil, err
	}
	return response.BaseResponse.GetHttpContentBytes(), nil
}
