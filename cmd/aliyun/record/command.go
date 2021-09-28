package record

import (
	"github.com/spf13/cobra"
)

var (
	// Cmd 命令集
	Cmd = &cobra.Command{
		Use:   "record <subcommand>",
		Short: "域名解析管理",
	}

	statusM = map[string]string{
		"ENABLE":  "正常",
		"DISABLE": "暂停",
	}
)

func init() {
	Cmd.AddCommand(
		listCmd,
		listSubCmd,
		getCmd,
		addCmd,
		updateCmd,
		deleteCmd,
		logsCmd,
		remarkCmd,
		statusCmd,
	)

	Cmd.PersistentFlags().StringP("domain", "d", "", "域名名称")

	listCmd.Flags().StringP("keyword", "k", "", "关键字")
	listCmd.Flags().String("rr-keyword", "", "主机记录的关键字，不区分大小写")
	listCmd.Flags().String("type-keyword", "", "解析类型的关键字，不区分大小写")
	listCmd.Flags().String("value-keyword", "", "记录值的关键字，不区分大小写")
	listCmd.Flags().IntP("gid", "g", 0, "域名分组ID")
	listCmd.Flags().StringP("type", "t", "", "记录类型")
	listCmd.Flags().IntP("page", "p", 1, "当前页数，起始值为 1")
	listCmd.Flags().IntP("size", "s", 10, "分页查询时设置的每页行数，最大值100")
	listCmd.Flags().String("direction", "", "排序方向，DESC - 降序、ASC - 升序")
	listCmd.Flags().String("status", "", "记录状态，ENABLE - 正常、DISABLE - 暂停")
	listCmd.Flags().Bool("set-color", true, "设置颜色")

	listSubCmd.Flags().String("sub-domain", "", "子域名名称")
	listSubCmd.Flags().StringP("type", "t", "", "记录类型")
	listSubCmd.Flags().StringP("line", "l", "default", "线路")
	listSubCmd.Flags().IntP("page", "p", 1, "当前页数，起始值为 1")
	listSubCmd.Flags().IntP("size", "s", 10, "分页查询时设置的每页行数，最大值100")
	listSubCmd.Flags().Bool("set-color", true, "设置颜色")

	getCmd.Flags().StringP("id", "i", "", "解析记录ID")
	getCmd.Flags().Bool("set-color", true, "设置颜色")

	addCmd.Flags().String("rr", "", "主机记录")
	addCmd.Flags().StringP("type", "t", "A", "记录类型")
	addCmd.Flags().StringP("value", "v", "", "记录值")
	addCmd.Flags().Int64("ttl", 600, "解析生效时间，默认为600秒(10分钟)。")
	addCmd.Flags().Int("priority", 1, "MX记录的优先级，取值范围：[1,50]。记录类型为MX记录时，此参数必需。")
	addCmd.Flags().StringP("line", "l", "default", "线路")

	updateCmd.Flags().StringP("id", "i", "", "解析记录ID")
	updateCmd.Flags().String("rr", "", "主机记录")
	updateCmd.Flags().StringP("type", "t", "A", "记录类型")
	updateCmd.Flags().StringP("value", "v", "", "记录值")
	updateCmd.Flags().Int64("ttl", 600, "解析生效时间，默认为600秒(10分钟)。")
	updateCmd.Flags().Int("priority", 1, "MX记录的优先级，取值范围：[1,50]。记录类型为MX记录时，此参数必需。")
	updateCmd.Flags().StringP("line", "l", "default", "线路")

	deleteCmd.Flags().StringP("id", "i", "", "解析记录ID")

	logsCmd.Flags().StringP("keyword", "k", "", "关键字")
	logsCmd.Flags().StringP("start", "s", "", "查询的开始时间")
	logsCmd.Flags().StringP("end", "e", "", "查询的结束时间")
	logsCmd.Flags().Int("page", 1, "当前页数，起始值为 1")
	logsCmd.Flags().Int("size", 20, "分页查询时设置的每页行数，最大值100")
	logsCmd.Flags().StringP("lang", "l", "zh", "语言，en - 英文， zh - 中文")

	remarkCmd.Flags().StringP("id", "i", "", "解析记录ID")
	remarkCmd.Flags().StringP("remark", "r", "", "备注信息")

	statusCmd.Flags().StringP("id", "i", "", "解析记录ID")
	statusCmd.Flags().StringP("status", "s", "", "记录状态，ENABLE - 启用解析、DISABLE - 暂停解析")
}
