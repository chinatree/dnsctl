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
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "获取解析记录信息",
		Long:  "获取解析记录的详细信息。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"record get --id 721046463884277760",
		),
		PreRun: app.ParseArgs,
		Run:    getFunc,
	}
)

// getFunc executes the "record get" command.
func getFunc(cmd *cobra.Command, args []string) {
	id := cmd.Flag("id").Value.String()

	resp, err := app.Client.GetDomainRecord(id)
	if err != nil {
		panic(err)
	}

	// data, _ := convert.StringToJSONWithIndent(string(resp))
	// fmt.Println(data)

	// Parse Response
	var record *Record
	err = json.Unmarshal(resp, &record)
	if err != nil {
		panic(err)
	}

	setColor, _ := cmd.Flags().GetBool("set-color")
	rows := [][]string{
		{
			strconv.Itoa(0),
			record.RecordID,
			record.DomainName,
			record.RR,
			record.Type,
			record.Value,
			strconv.FormatInt(record.TTL, 10),
			strconv.Itoa(record.Weight),
			record.Line,
			statusM[record.Status],
			strconv.FormatBool(record.Locked),
		},
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
