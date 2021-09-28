package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

// GetDomains 获取域名列表
func (aliyun *Aliyun) GetDomains(keyword, gid, rgid, searchMode string, page, size int, starmark bool) ([]byte, error) {
	request := alidns.CreateDescribeDomainsRequest()
	request.KeyWord = keyword
	request.GroupId = gid
	request.ResourceGroupId = rgid
	request.PageNumber = requests.NewInteger(page)
	request.PageSize = requests.NewInteger(size)
	request.SearchMode = searchMode
	request.Starmark = requests.NewBoolean(starmark)
	response, err := aliyun.Client.DescribeDomains(request)
	if err != nil {
		return nil, err
	}
	return response.BaseResponse.GetHttpContentBytes(), nil
}

// GetDomain 获取域名信息
func (aliyun *Aliyun) GetDomain(name string, detail bool) ([]byte, error) {
	request := alidns.CreateDescribeDomainInfoRequest()
	request.DomainName = name
	request.NeedDetailAttributes = requests.NewBoolean(detail)
	response, err := aliyun.Client.DescribeDomainInfo(request)
	if err != nil {
		return nil, err
	}
	return response.BaseResponse.GetHttpContentBytes(), nil
}

// AddDomain 添加域名
func (aliyun *Aliyun) AddDomain(name, gid, rgid string) ([]byte, error) {
	request := alidns.CreateAddDomainRequest()
	request.DomainName = name
	request.GroupId = gid
	request.ResourceGroupId = rgid
	response, err := aliyun.Client.AddDomain(request)
	if err != nil {
		return nil, err
	}
	return response.BaseResponse.GetHttpContentBytes(), nil
}

// DeleteDomain 删除域名
func (aliyun *Aliyun) DeleteDomain(name string) ([]byte, error) {
	request := alidns.CreateDeleteDomainRequest()
	request.DomainName = name
	response, err := aliyun.Client.DeleteDomain(request)
	if err != nil {
		return nil, err
	}
	return response.BaseResponse.GetHttpContentBytes(), nil
}

// DomainLogs 获取域名操作日志
func (aliyun *Aliyun) DomainLogs(keyword, gid, start, end, _type, lang string, page, size int) ([]byte, error) {
	request := alidns.CreateDescribeDomainLogsRequest()
	request.KeyWord = keyword
	request.GroupId = gid
	request.StartDate = start
	request.EndDate = end
	request.Type = _type
	request.Lang = lang
	request.PageNumber = requests.NewInteger(page)
	request.PageSize = requests.NewInteger(size)
	response, err := aliyun.Client.DescribeDomainLogs(request)
	if err != nil {
		return nil, err
	}
	return response.BaseResponse.GetHttpContentBytes(), nil
}

// UpdateDomainRemark 修改域名的备注
func (aliyun *Aliyun) UpdateDomainRemark(name, remark string) ([]byte, error) {
	request := alidns.CreateUpdateDomainRemarkRequest()
	request.DomainName = name
	request.Remark = remark
	response, err := aliyun.Client.UpdateDomainRemark(request)
	if err != nil {
		return nil, err
	}
	return response.BaseResponse.GetHttpContentBytes(), nil
}

// GetMainDomainName 获取主域名名称
func (aliyun *Aliyun) GetMainDomainName(name string) ([]byte, error) {
	request := alidns.CreateGetMainDomainNameRequest()
	request.InputString = name
	response, err := aliyun.Client.GetMainDomainName(request)
	if err != nil {
		return nil, err
	}
	return response.BaseResponse.GetHttpContentBytes(), nil
}
