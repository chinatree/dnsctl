package tencent

import (
	"errors"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

// Tencent aliyun 实例
type Tencent struct {
	Client *dnspod.Client
}

// NewClient Tencent default client
func NewClient(region, secretID, secretKey string) (*Tencent, error) {
	if secretID == "" || secretKey == "" {
		return nil, errors.New("SecretID or SecretKey is required.")
	}

	if region == "" {
		region = regions.Guangzhou
	}

	credential := common.NewCredential(secretID, secretKey)
	client, err := dnspod.NewClient(credential, region, profile.NewClientProfile())
	if err != nil {
		return nil, err
	}
	c := &Tencent{
		Client: client,
	}
	return c, nil
}
