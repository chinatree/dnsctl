package aliyun

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/chinatree/dnsctl/internal/thirdparty/aliyun"
)

var (
	// Client Aliyun 实例
	Client *aliyun.Aliyun
)

// ParseArgs 从命令行参数获取 AccessKey
// 	如果命令行参数未指定，则从环境变量获取
// 	- ALIYUN_REGION_ID
// 	- ALIYUN_SECRET_ID
// 	- ALIYUN_SECRET_KEY
func ParseArgs(cmd *cobra.Command, args []string) {
	regionID := cmd.Flag("region-id").Value.String()
	secretID := cmd.Flag("secret-id").Value.String()
	secretKey := cmd.Flag("secret-key").Value.String()
	if regionID == "" {
		regionID = os.Getenv("ALIYUN_REGION_ID")
	}
	if secretID == "" {
		secretID = os.Getenv("ALIYUN_SECRET_ID")
	}
	if secretKey == "" {
		secretKey = os.Getenv("ALIYUN_SECRET_KEY")
	}

	err := InitAliyun(regionID, secretID, secretKey)
	if err != nil {
		panic(err)
	}
}

// InitAliyun 初始化 Aliyun
func InitAliyun(regionID, secretID, secretKey string) error {
	// 初始化 AliClient
	client, err := aliyun.NewClient(regionID, secretID, secretKey)
	if err != nil {
		return err
	}

	Client = client
	return nil
}
