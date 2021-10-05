package record

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
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "获取域名解析列表",
		Long:  "根据传入参数获取指定主域名的所有解析记录列表。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"record list --domain yuntree.com",
		),
		PreRun: app.ParseArgs,
		Run:    listFunc,
	}
)

// listFunc executes the "record list" command.
func listFunc(cmd *cobra.Command, args []string) {
	domain := cmd.Flag("domain").Value.String()
	rr := cmd.Flag("rr").Value.String()
	_type := cmd.Flag("type").Value.String()
	keyword := cmd.Flag("keyword").Value.String()
	sort := cmd.Flag("sort").Value.String()
	sortField := cmd.Flag("sort-filed").Value.String()
	did, _ := cmd.Flags().GetUint64("domain-id")
	gid, _ := cmd.Flags().GetUint64("gid")
	offset, _ := cmd.Flags().GetUint64("offset")
	limit, _ := cmd.Flags().GetUint64("limit")

	if domain == "" {
		panic("DomainName is mandatory for this action.")
	}

	resp, err := app.Client.GetRecordList(
		domain, rr, keyword, _type, sort, sortField,
		did, gid, offset, limit,
	)
	if err != nil {
		panic(err)
	}

	// data, _ := convert.StringToJSONWithIndent(string(resp))
	// fmt.Println(data)

	// Parse Response
	var records *Records
	err = json.Unmarshal(resp, &records)
	if err != nil {
		panic(err)
	}

	// data, _ := convert.StructToJSONWithIndent(jsonData)
	// fmt.Println(data)

	setColor, _ := cmd.Flags().GetBool("set-color")
	rows := make([][]string, 0)
	for idx, record := range records.Response.RecordList {
		row := make([]string, 0)
		row = append(row, strconv.Itoa(idx))
		row = append(row, strconv.FormatUint(record.ID, 10))
		row = append(row, record.Name)
		row = append(row, record.Type)
		row = append(row, record.Value)
		row = append(row, strconv.FormatUint(record.TTL, 10))
		row = append(row, strconv.FormatUint(record.Weight, 10))
		row = append(row, record.Line)
		row = append(row, statusM[strings.ToUpper(record.Status)])
		row = append(row, record.Updated)
		row = append(row, record.Remark)

		rows = append(rows, row)
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
