package group

import (
	"github.com/spf13/cobra"
)

var (
	// Cmd 命令集
	Cmd = &cobra.Command{
		Use:   "group <subcommand>",
		Short: "域名分组管理",
	}
)

func init() {
	Cmd.AddCommand(
		listCmd,
		addCmd,
		updateCmd,
		deleteCmd,
		changeCmd,
	)

	listCmd.Flags().StringP("keyword", "k", "", "关键字")
	listCmd.Flags().IntP("page", "p", 1, "当前页数，起始值为 1")
	listCmd.Flags().IntP("size", "s", 10, "分页查询时设置的每页行数，最大值100")

	addCmd.Flags().StringP("name", "n", "", "域名分组名称")

	updateCmd.Flags().StringP("id", "i", "", "域名分组ID")
	updateCmd.Flags().StringP("name", "n", "", "域名分组名称")

	deleteCmd.Flags().StringP("id", "i", "", "域名分组ID")

	changeCmd.Flags().StringP("domain", "d", "", "域名名称")
	changeCmd.Flags().StringP("id", "i", "", "目标域名分组ID")
}
