package aliyun

import (
	"github.com/spf13/cobra"

	"github.com/chinatree/dnsctl/cmd/aliyun/domain"
	"github.com/chinatree/dnsctl/cmd/aliyun/group"
	"github.com/chinatree/dnsctl/cmd/aliyun/record"
)

var (
	// Cmd 命令集
	Cmd = &cobra.Command{
		Use:   "aliyun <subcommand>",
		Short: "阿里云 云DNS",
	}
)

func init() {
	Cmd.AddCommand(
		domain.Cmd,
		group.Cmd,
		record.Cmd,
	)

	Cmd.PersistentFlags().String("region-id", "default", "区域ID")
	Cmd.PersistentFlags().String("secret-id", "", "密钥ID")
	Cmd.PersistentFlags().String("secret-key", "", "加密密钥")
}
