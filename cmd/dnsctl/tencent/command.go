package tencent

import (
	"github.com/spf13/cobra"

	"github.com/chinatree/dnsctl/cmd/tencent/domain"
	"github.com/chinatree/dnsctl/cmd/tencent/group"
	"github.com/chinatree/dnsctl/cmd/tencent/record"
)

var (
	// Cmd 命令集
	Cmd = &cobra.Command{
		Use:   "tencent <subcommand>",
		Short: "腾讯云 云DNS",
	}
)

func init() {
	Cmd.AddCommand(
		domain.Cmd,
		group.Cmd,
		record.Cmd,
	)

	Cmd.PersistentFlags().String("region", "", "区域名称")
	Cmd.PersistentFlags().String("secret-id", "", "密钥ID")
	Cmd.PersistentFlags().String("secret-key", "", "加密密钥")
}
