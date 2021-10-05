package domain

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/tencent"
)

var (
	getCmd = &cobra.Command{
		Use:     "get",
		Aliases: []string{"query"},
		Short:   "获取域名信息",
		Long:    `根据传入参数查询指定域名的信息。`,
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"domain get --name yuntree.com",
		),
		PreRun: app.ParseArgs,
		Run:    getFunc,
	}
)

// getFunc executes the "domain get" command.
func getFunc(cmd *cobra.Command, args []string) {
	name := cmd.Flag("name").Value.String()
	id, _ := cmd.Flags().GetUint64("id")

	if name == "" && id == 0 {
		panic("DomainName or DomainID is mandatory for this action.")
	}

	resp, err := app.Client.GetDomain(name, id)
	if err != nil {
		panic(err)
	}

	// data, _ := convert.StringToJSONWithIndent(string(resp))
	// fmt.Println(data)

	// Parse Response
	var info *InfoResp
	err = json.Unmarshal(resp, &info)
	if err != nil {
		panic(err)
	}

	// data, _ := convert.StructToJSONWithIndent(info)
	// fmt.Println(data)

	setColor, _ := cmd.Flags().GetBool("set-color")
	domain := info.Response.DomainInfo
	rows := [][]string{
		{
			strconv.Itoa(0),
			strconv.FormatUint(domain.DomainID, 10),
			domain.Domain,
			strings.Join(domain.DnspodNsList, "\n"),
			strconv.FormatUint(domain.RecordCount, 10),
			strconv.FormatUint(domain.GroupID, 10),
			strconv.FormatUint(domain.TTL, 10),
			statusM[strings.ToUpper(domain.Status)],
			DNSStatusM[strings.ToUpper(domain.DNSStatus)],
			domain.Owner,
			domain.GradeTitle,
			domain.IsVip,
			domain.Remark,
		},
	}

	table := tablewriter.NewWriter(os.Stdout)
	tableHeader := []string{
		"序号", "域名ID", "域名名称", "DNS列表", "记录数", "域名分组ID", "默认生存时间",
		"状态", "DNS状态", "域名所属帐号", "付费版本", "是否付费套餐", "备注",
	}
	tableHeaderColor := make([]tablewriter.Colors, 0)
	tableRowRedColor := make([]tablewriter.Colors, 0)
	tableRowGreenColor := make([]tablewriter.Colors, 0)
	for i := 0; i < len(tableHeader); i++ {
		tableHeaderColor = append(tableHeaderColor, tablewriter.Colors{tablewriter.Bold})
		if i == 7 || i == 8 {
			tableRowRedColor = append(tableRowRedColor,
				tablewriter.Colors{tablewriter.Normal, tablewriter.FgRedColor})
			tableRowGreenColor = append(tableRowGreenColor,
				tablewriter.Colors{tablewriter.Normal, tablewriter.FgGreenColor})
			continue
		}
		tableRowRedColor = append(tableRowRedColor, tablewriter.Colors{})
		tableRowGreenColor = append(tableRowGreenColor, tablewriter.Colors{})
	}
	table.SetHeader(tableHeader)
	table.SetHeaderColor(tableHeaderColor...)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetRowLine(true)
	for _, row := range rows {
		if setColor {
			if row[7] == "暂停" || row[7] == "锁定" || row[7] == "封禁" || row[8] == "DNS 不正确" {
				table.Rich(row, tableRowRedColor)
				continue
			} else {
				table.Rich(row, tableRowGreenColor)
				continue
			}
		}
		table.Append(row)
	}

	table.Render()
}
