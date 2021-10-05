package aliyun

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/chinatree/dnsctl/internal/thirdparty/tencent"
)

var (
	// Client Aliyun 实例
	Client *tencent.Tencent
)

// ParseArgs 从命令行参数获取 AccessKey
// 	如果命令行参数未指定，则从环境变量获取
// 	- TENCENT_REGION
// 	- TENCENT_SECRET_ID
// 	- TENCENT_SECRET_KEY
func ParseArgs(cmd *cobra.Command, args []string) {
	region := cmd.Flag("region").Value.String()
	secretID := cmd.Flag("secret-id").Value.String()
	secretKey := cmd.Flag("secret-key").Value.String()
	if region == "" {
		region = os.Getenv("TENCENT_REGION")
	}
	if secretID == "" {
		secretID = os.Getenv("TENCENT_SECRET_ID")
	}
	if secretKey == "" {
		secretKey = os.Getenv("TENCENT_SECRET_KEY")
	}

	err := InitTencent(region, secretID, secretKey)
	if err != nil {
		panic(err)
	}
}

// InitTencent 初始化 Tencent
func InitTencent(regionID, secretID, secretKey string) error {
	// 初始化 AliClient
	client, err := tencent.NewClient(regionID, secretID, secretKey)
	if err != nil {
		return err
	}

	Client = client
	return nil
}
