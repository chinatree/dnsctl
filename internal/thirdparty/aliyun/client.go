package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

// Aliyun aliyun 实例
type Aliyun struct {
	Client *alidns.Client
}

// NewClient Aliyun default client
func NewClient(regionID, secretID, secretKey string) (*Aliyun, error) {
	client, err := alidns.NewClientWithAccessKey(regionID, secretID, secretKey)
	if err != nil {
		return nil, err
	}
	aliClient := &Aliyun{
		Client: client,
	}
	return aliClient, nil
}
