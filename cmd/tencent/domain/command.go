package domain

import (
	"github.com/spf13/cobra"
)

var (
	// Cmd 命令集
	Cmd = &cobra.Command{
		Use:   "domain <subcommand>",
		Short: "域名管理",
	}

	// statusM 域名状态MAP
	statusM = map[string]string{
		"ENABLE": "正常",
		"PAUSE":  "暂停",
		"SPAM":   "封禁",
		"LOCK":   "锁定",
	}

	// DNSStatusM  域名DNS状态MAP
	DNSStatusM = map[string]string{
		"":         "正常",
		"DNSERROR": "DNS 不正确",
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
		statusCmd,
		lockCmd,
		unlockCmd,
	)

	listCmd.Flags().StringP("keyword", "k", "", "关键字")
	listCmd.Flags().Int64P("gid", "g", 0, "域名分组ID")
	listCmd.Flags().StringP("type", "t", "", "域名分组类型。可取值为 ALL，MINE，SHARE，ISMARK，PAUSE，VIP，RECENT，SHARE_OUT")
	listCmd.Flags().Int64P("offset", "o", 0, "记录开始的偏移，起始值为 0")
	listCmd.Flags().Int64P("limit", "l", 20, "要获取的域名数量")
	listCmd.Flags().Bool("set-color", true, "设置颜色")

	getCmd.Flags().StringP("name", "n", "", "域名名称")
	getCmd.Flags().Uint64P("id", "i", 0, "域名ID，如果指定该参数，将忽略 --name 参数")
	getCmd.Flags().Bool("set-color", true, "设置颜色")

	addCmd.Flags().StringP("name", "n", "", "域名名称")
	addCmd.Flags().Uint64P("gid", "g", 0, "域名分组ID")
	addCmd.Flags().String("is-mark", "no", "是否星标域名，yes - 是，no - 否")

	deleteCmd.Flags().StringP("name", "n", "", "域名名称")
	deleteCmd.Flags().Uint64P("id", "i", 0, "域名ID，如果指定该参数，将忽略 --name 参数")

	logsCmd.Flags().StringP("name", "n", "", "域名名称")
	logsCmd.Flags().Uint64P("id", "i", 0, "域名ID，如果指定该参数，将忽略 --name 参数")
	logsCmd.Flags().Uint64P("offset", "o", 0, "记录开始的偏移，起始值为 0")
	logsCmd.Flags().Uint64P("limit", "l", 20, "要获取的域名数量")

	remarkCmd.Flags().StringP("name", "n", "", "域名名称")
	remarkCmd.Flags().Uint64P("id", "i", 0, "域名ID，如果指定该参数，将忽略 --name 参数")
	remarkCmd.Flags().StringP("remark", "r", "", "备注信息")

	statusCmd.Flags().StringP("name", "n", "", "域名名称")
	statusCmd.Flags().Uint64P("id", "i", 0, "域名ID，如果指定该参数，将忽略 --name 参数")
	statusCmd.Flags().StringP("status", "s", "", "域名状态，ENABLE - 启用解析、DISABLE - 暂停解析")

	lockCmd.Flags().StringP("name", "n", "", "域名名称")
	lockCmd.Flags().Uint64P("id", "i", 0, "域名ID，如果指定该参数，将忽略 --name 参数")
	lockCmd.Flags().Uint64P("day", "d", 0, "锁定的天数")

	unlockCmd.Flags().StringP("name", "n", "", "域名名称")
	unlockCmd.Flags().Uint64P("id", "i", 0, "域名ID，如果指定该参数，将忽略 --name 参数")
	unlockCmd.Flags().StringP("code", "c", "", "域名解锁码")
}
