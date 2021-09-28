package record

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/aliyun"
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
	listSubCmd = &cobra.Command{
		Use:   "listSub",
		Short: "获取子域名解析记录列表",
		Long:  "根据传入参数获取某个固定子域名的所有解析记录列表。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"record listSub --domain yuntree.com --sub-domain cdn.yuntree.com",
		),
		PreRun: app.ParseArgs,
		Run:    listSubFunc,
	}
)

// listFunc executes the "record list" command.
func listFunc(cmd *cobra.Command, args []string) {
	domain := cmd.Flag("domain").Value.String()
	_type := cmd.Flag("type").Value.String()
	status := cmd.Flag("status").Value.String()
	keyword := cmd.Flag("keyword").Value.String()
	rrKeyword := cmd.Flag("rr-keyword").Value.String()
	typeKeyword := cmd.Flag("type-keyword").Value.String()
	valueKeyword := cmd.Flag("value-keyword").Value.String()
	gid, _ := cmd.Flags().GetInt("gid")
	page, _ := cmd.Flags().GetInt("page")
	size, _ := cmd.Flags().GetInt("size")
	direction := cmd.Flag("direction").Value.String()

	if domain == "" {
		panic("DomainName is mandatory for this action.")
	}

	resp, err := app.Client.GetDomainRecords(
		domain, _type, status,
		keyword, rrKeyword, typeKeyword, valueKeyword,
		gid, page, size,
		direction,
	)
	if err != nil {
		panic(err)
	}

	// data, _ := convert.StringToJSONWithIndent(string(resp))
	// fmt.Println(data)

	// Parse Response
	var jsonData *DomainRecordList
	err = json.Unmarshal(resp, &jsonData)
	if err != nil {
		panic(err)
	}

	// data, _ := convert.StructToJSONWithIndent(jsonData)
	// fmt.Println(data)

	setColor, _ := cmd.Flags().GetBool("set-color")
	rows := make([][]string, 0)
	for idx, record := range jsonData.DomainRecords.Records {
		row := make([]string, 0)
		row = append(row, strconv.Itoa(idx))
		row = append(row, record.RecordId)
		row = append(row, record.DomainName)
		row = append(row, record.RR)
		row = append(row, record.Type)
		row = append(row, record.Value)
		row = append(row, strconv.FormatInt(record.TTL, 10))
		row = append(row, strconv.Itoa(record.Weight))
		row = append(row, record.Line)
		row = append(row, statusM[record.Status])
		row = append(row, strconv.FormatBool(record.Locked))
		row = append(row, record.Remark)

		rows = append(rows, row)
	}

	table := tablewriter.NewWriter(os.Stdout)
	tableHeader := []string{
		"序号", "记录ID", "域名名称", "主机记录", "记录类型", "记录值", "生存时间", "权重",
		"解析线路(isp)", "状态", "锁定状态", "备注",
	}
	tableHeaderColor := make([]tablewriter.Colors, 0)
	tableRowRedColor := make([]tablewriter.Colors, 0)
	tableRowGreenColor := make([]tablewriter.Colors, 0)
	for i := 0; i < len(tableHeader); i++ {
		tableHeaderColor = append(tableHeaderColor, tablewriter.Colors{tablewriter.Bold})
		if i == 9 {
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
			if row[9] == "暂停" {
				table.Rich(row, tableRowRedColor)
				continue
			} else if row[9] == "正常" {
				table.Rich(row, tableRowGreenColor)
				continue
			}
		}
		table.Append(row)
	}

	table.Render()
}

// listSubFunc executes the "record listSub" command.
func listSubFunc(cmd *cobra.Command, args []string) {
	domain := cmd.Flag("domain").Value.String()
	subDomain := cmd.Flag("sub-domain").Value.String()
	_type := cmd.Flag("type").Value.String()
	line := cmd.Flag("line").Value.String()
	page, _ := cmd.Flags().GetInt("page")
	size, _ := cmd.Flags().GetInt("size")

	if subDomain == "" {
		panic("SubDomain is mandatory for this action.")
	}

	resp, err := app.Client.GetDomainSubRecords(
		subDomain, domain, _type, line,
		page, size,
	)
	if err != nil {
		panic(err)
	}

	// data, _ := convert.StringToJSONWithIndent(string(resp))
	// fmt.Println(data)

	// Parse Response
	var jsonData *DomainRecordList
	err = json.Unmarshal(resp, &jsonData)
	if err != nil {
		panic(err)
	}

	// data, _ := convert.StructToJSONWithIndent(jsonData)
	// fmt.Println(data)

	setColor, _ := cmd.Flags().GetBool("set-color")
	rows := make([][]string, 0)
	for idx, record := range jsonData.DomainRecords.Records {
		row := make([]string, 0)
		row = append(row, strconv.Itoa(idx))
		row = append(row, record.RecordId)
		row = append(row, record.DomainName)
		row = append(row, record.RR)
		row = append(row, record.Type)
		row = append(row, record.Value)
		row = append(row, strconv.FormatInt(record.TTL, 10))
		row = append(row, strconv.Itoa(record.Weight))
		row = append(row, record.Line)
		row = append(row, statusM[record.Status])
		row = append(row, strconv.FormatBool(record.Locked))

		rows = append(rows, row)
	}

	table := tablewriter.NewWriter(os.Stdout)
	tableHeader := []string{
		"序号", "记录ID", "域名名称", "主机记录", "记录类型", "记录值", "生存时间", "权重",
		"解析线路(isp)", "状态", "锁定状态",
	}
	tableHeaderColor := make([]tablewriter.Colors, 0)
	tableRowRedColor := make([]tablewriter.Colors, 0)
	tableRowGreenColor := make([]tablewriter.Colors, 0)
	for i := 0; i < len(tableHeader); i++ {
		tableHeaderColor = append(tableHeaderColor, tablewriter.Colors{tablewriter.Bold})
		if i == 9 {
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
			if row[9] == "暂停" {
				table.Rich(row, tableRowRedColor)
				continue
			} else if row[9] == "正常" {
				table.Rich(row, tableRowGreenColor)
				continue
			}
		}
		table.Append(row)
	}

	table.Render()
}
