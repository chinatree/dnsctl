package aliyun

import (
	"os"

	"github.com/chinatree/dnsctl/internal/thirdparty/aliyun"
	"github.com/spf13/cobra"
)

var (
	// Client Aliyun 实例
	Client *aliyun.Aliyun
)

// ParseArgs 从命令行参数获取 AccessKey
// 	如果命令行参数未指定，则从环境变量获取
// 	- ALIYUN_REGION_ID
// 	- ALIYUN_ACCESS_KEY_ID
// 	- ALIYUN_ACCESS_SECRET_ID
func ParseArgs(cmd *cobra.Command, args []string) {
	regionID := cmd.Flag("region-id").Value.String()
	accessKeyID := cmd.Flag("access-key-id").Value.String()
	accessKeySecret := cmd.Flag("access-key-secret").Value.String()
	if regionID == "" {
		regionID = os.Getenv("ALIYUN_REGION_ID")
	}
	if accessKeyID == "" {
		accessKeyID = os.Getenv("ALIYUN_ACCESS_KEY_ID")
	}
	if accessKeySecret == "" {
		accessKeySecret = os.Getenv("ALIYUN_ACCESS_SECRET_ID")
	}

	err := InitAliyun(regionID, accessKeyID, accessKeySecret)
	if err != nil {
		panic(err)
	}
}

// InitAliyun 初始化 Aliyun
func InitAliyun(regionID, accessKeyID, accessKeySecret string) error {
	// 初始化 AliClient
	client, err := aliyun.NewClient(regionID, accessKeyID, accessKeySecret)
	if err != nil {
		return err
	}

	Client = client
	return nil
}
