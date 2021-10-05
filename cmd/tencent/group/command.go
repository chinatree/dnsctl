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
		addCmd,
	)

	addCmd.Flags().StringP("name", "n", "", "域名分组名称")
}
