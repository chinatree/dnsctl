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

	statusCodeM = map[uint64]string{
		1: "正常",
		0: "暂停",
	}
)

func init() {
	Cmd.AddCommand(
		listCmd,
		getCmd,
		addCmd,
		updateCmd,
		deleteCmd,
		remarkCmd,
		statusCmd,
	)

	Cmd.PersistentFlags().StringP("domain", "d", "", "域名名称")
	Cmd.PersistentFlags().Uint64("domain-id", 0, "域名ID。如果指定该参数，将忽略 --name 参数")

	listCmd.Flags().String("rr", "", "主机记录")
	listCmd.Flags().StringP("keyword", "k", "", "关键字")
	listCmd.Flags().Uint64P("gid", "g", 0, "域名分组ID")
	listCmd.Flags().StringP("type", "t", "", "记录类型")
	listCmd.Flags().Uint64P("offset", "o", 0, "记录开始的偏移，起始值为 0")
	listCmd.Flags().Uint64P("limit", "l", 20, "要获取的域名数量")
	listCmd.Flags().String("sort", "ASC", "排序方向，DESC - 降序、ASC - 升序")
	listCmd.Flags().String("sort-filed", "", "排序字段，支持 name, line, type, value, weight, mx, ttl, updated_on")
	listCmd.Flags().Bool("set-color", true, "设置颜色")

	getCmd.Flags().Uint64P("id", "i", 0, "解析记录ID")
	getCmd.Flags().Bool("set-color", true, "设置颜色")

	addCmd.Flags().String("rr", "", "主机记录。如果不传，默认为 @")
	addCmd.Flags().StringP("type", "t", "A", "记录类型")
	addCmd.Flags().StringP("value", "v", "", "记录值")
	addCmd.Flags().StringP("status", "s", "ENABLE", "记录状态，ENABLE - 启用解析、DISABLE - 暂停解析")
	addCmd.Flags().Uint64("ttl", 600, "解析生效时间(单位秒)，范围1-604800。默认为600秒(10分钟)")
	addCmd.Flags().StringP("line", "l", "默认", "线路")
	addCmd.Flags().Uint64("line-id", 0, "线路ID。如果指定该参数，将忽略 --line 参数。")
	addCmd.Flags().Uint64P("weight", "w", 0, "权重。仅企业 VIP 域名可用，0 表示关闭。")

	updateCmd.Flags().Uint64P("id", "i", 0, "解析记录ID")
	updateCmd.Flags().String("rr", "", "主机记录。如果不传，默认为 @")
	updateCmd.Flags().StringP("type", "t", "A", "记录类型")
	updateCmd.Flags().StringP("value", "v", "", "记录值")
	updateCmd.Flags().StringP("status", "s", "ENABLE", "记录状态，ENABLE - 启用解析、DISABLE - 暂停解析")
	updateCmd.Flags().Uint64("ttl", 600, "解析生效时间(单位秒)，范围1-604800。默认为600秒(10分钟)")
	updateCmd.Flags().StringP("line", "l", "默认", "线路")
	updateCmd.Flags().Uint64("line-id", 0, "线路ID。如果指定该参数，将忽略 --line 参数")
	updateCmd.Flags().Uint64P("weight", "w", 0, "权重。仅企业 VIP 域名可用，0 表示关闭")

	deleteCmd.Flags().Uint64P("id", "i", 0, "解析记录ID")

	remarkCmd.Flags().Uint64P("id", "i", 0, "解析记录ID")
	remarkCmd.Flags().StringP("remark", "r", "", "备注信息")

	statusCmd.Flags().Uint64P("id", "i", 0, "解析记录ID")
	statusCmd.Flags().StringP("status", "s", "", "记录状态，ENABLE - 启用解析、DISABLE - 暂停解析")
}
