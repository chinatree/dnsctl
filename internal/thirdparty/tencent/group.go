package tencent

import (
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

// AddGroup 添加域名分组
func (aliyun *Tencent) AddGroup(name string) ([]byte, error) {
	request := dnspod.NewCreateDomainGroupRequest()
	request.GroupName = &name
	response, err := aliyun.Client.CreateDomainGroup(request)
	if err != nil {
		return nil, err
	}
	return []byte(response.ToJsonString()), nil
}
