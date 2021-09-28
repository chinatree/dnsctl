package domain

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var (
	// Cmd 命令集
	Cmd = &cobra.Command{
		Use:   "domain <subcommand>",
		Short: "域名管理",
	}
)

func init() {
	Cmd.AddCommand(
		listCmd,
		getCmd,
		addCmd,
		deleteCmd,
		logsCmd,
		remarkCmd,
		getMainCmd,
	)

	listCmd.Flags().StringP("keyword", "k", "", "关键字")
	listCmd.Flags().String("gid", "", "域名分组ID")
	listCmd.Flags().String("rgid", "", "资源组ID")
	listCmd.Flags().IntP("page", "p", 1, "当前页数，起始值为 1")
	listCmd.Flags().IntP("size", "s", 10, "分页查询时设置的每页行数，最大值100")
	listCmd.Flags().String("search-mode", "", "搜索模式，LIKE - 模糊搜索，EXACT - 精确搜索")
	listCmd.Flags().Bool("starmark", false, "是否查询域名星标")

	addCmd.Flags().StringP("name", "n", "", "域名名称")
	addCmd.Flags().String("gid", "", "域名分组ID")
	addCmd.Flags().String("rgid", "", "资源组ID")

	getCmd.Flags().StringP("name", "n", "", "域名名称")
	getCmd.Flags().Bool("detail", false, "是否显示细节属性")

	deleteCmd.Flags().StringP("name", "n", "", "域名名称")

	logsCmd.Flags().StringP("keyword", "k", "", "关键字")
	logsCmd.Flags().StringP("gid", "g", "", "域名分组ID")
	logsCmd.Flags().StringP("type", "t", "", "查询内容类型，domain - 域名，slavedns - 辅助DNS")
	logsCmd.Flags().StringP("lang", "l", "zh", "语言，en - 英文，zh - 中文")
	logsCmd.Flags().StringP(
		"start", "s", time.Now().Add(-1*time.Hour*24*7).Format("2006-01-02"),
		fmt.Sprintf("查询的开始时间，如 %s", time.Now().Format("2006-01-02")))
	logsCmd.Flags().StringP(
		"end", "e", time.Now().Format("2006-01-02"),
		fmt.Sprintf("查询的结束时间，如 %s", time.Now().Format("2006-01-02")))
	logsCmd.Flags().Int("page", 1, "当前页数，起始值为 1")
	logsCmd.Flags().Int("size", 20, "分页查询时设置的每页行数，最大值100")

	remarkCmd.Flags().StringP("name", "n", "", "域名名称")
	remarkCmd.Flags().StringP("remark", "r", "", "备注信息")

	getMainCmd.Flags().StringP("name", "n", "", "域名名称")
}
