package domain

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/tencent"
)

var (
	listCmd = &cobra.Command{
		Use:    "list",
		Short:  "获取域名列表",
		Long:   `根据传入参数查询该用户的域名列表。`,
		PreRun: app.ParseArgs,
		Run:    listFunc,
	}
)

// listFunc executes the "domain list" command.
func listFunc(cmd *cobra.Command, args []string) {
	keyword := cmd.Flag("keyword").Value.String()
	gid, _ := cmd.Flags().GetInt64("gid")
	_type := cmd.Flag("type").Value.String()
	offset, _ := cmd.Flags().GetInt64("offset")
	limit, _ := cmd.Flags().GetInt64("limit")

	resp, err := app.Client.GetDomainList(keyword, _type, gid, offset, limit)
	if err != nil {
		panic(err)
	}

	// data, _ := convert.StringToJSONWithIndent(string(resp))
	// fmt.Println(data)

	// Parse Response
	var domains *List
	err = json.Unmarshal(resp, &domains)
	if err != nil {
		panic(err)
	}

	// data, _ := convert.StructToJSONWithIndent(domains)
	// fmt.Println(data)

	setColor, _ := cmd.Flags().GetBool("set-color")
	rows := make([][]string, 0)
	for idx, domain := range domains.Response.DomainList {
		row := make([]string, 0)
		row = append(row, strconv.Itoa(idx))
		row = append(row, strconv.FormatUint(domain.ID, 10))
		row = append(row, domain.Name)
		row = append(row, strings.Join(domain.EffectiveDNS, "\n"))
		row = append(row, strconv.FormatUint(domain.RecordCount, 10))
		row = append(row, strconv.FormatUint(domain.GroupID, 10))
		row = append(row, strconv.FormatUint(domain.TTL, 10))
		row = append(row, statusM[strings.ToUpper(domain.Status)])
		row = append(row, DNSStatusM[strings.ToUpper(domain.DNSStatus)])
		row = append(row, domain.Owner)
		row = append(row, domain.GradeTitle)
		row = append(row, domain.IsVip)
		row = append(row, domain.Remark)

		rows = append(rows, row)
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
