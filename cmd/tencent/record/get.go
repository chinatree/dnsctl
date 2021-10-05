package record

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/tencent"
)

var (
	getCmd = &cobra.Command{
		Use:     "get",
		Aliases: []string{"query"},
		Short:   "获取解析记录信息",
		Long:    "获取解析记录的详细信息。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"record get --id 666",
		),
		PreRun: app.ParseArgs,
		Run:    getFunc,
	}
)

// getFunc executes the "record get" command.
func getFunc(cmd *cobra.Command, args []string) {
	domain := cmd.Flag("domain").Value.String()
	did, _ := cmd.Flags().GetUint64("domain-id")
	id, _ := cmd.Flags().GetUint64("id")

	if domain == "" || id == 0 {
		panic("DomainName or RecordID is mandatory for this action.")
	}

	resp, err := app.Client.GetRecord(domain, did, id)
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
	record := info.Response.RecordInfo
	rows := [][]string{
		{
			strconv.Itoa(0),
			strconv.FormatUint(record.ID, 10),
			record.RR,
			record.Type,
			record.Value,
			strconv.FormatUint(record.TTL, 10),
			strconv.FormatUint(record.Weight, 10),
			record.Line,
			statusCodeM[record.Enabled],
			record.Updated,
			record.Remark,
		},
	}

	table := tablewriter.NewWriter(os.Stdout)
	tableHeader := []string{
		"序号", "记录ID", "主机记录", "记录类型", "记录值", "生存时间", "权重",
		"解析线路(isp)", "状态", "更新时间", "备注",
	}
	tableHeaderColor := make([]tablewriter.Colors, 0)
	tableRowRedColor := make([]tablewriter.Colors, 0)
	tableRowGreenColor := make([]tablewriter.Colors, 0)
	for i := 0; i < len(tableHeader); i++ {
		tableHeaderColor = append(tableHeaderColor, tablewriter.Colors{tablewriter.Bold})
		if i == 8 {
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
	for _, row := range rows {
		if setColor {
			if row[8] == "暂停" {
				table.Rich(row, tableRowRedColor)
				continue
			} else if row[8] == "正常" {
				table.Rich(row, tableRowGreenColor)
				continue
			}
		}
		table.Append(row)
	}

	table.Render()
}
