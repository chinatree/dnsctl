package domain

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
	logsCmd = &cobra.Command{
		Use:   "logs",
		Short: "获取域名的操作日志",
		Long:  "根据传入参数获取域名的操作日志。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"domain logs -n yuntree.com",
		),
		PreRun: app.ParseArgs,
		Run:    logsFunc,
	}
)

// logsFunc executes the "domain logs" command.
func logsFunc(cmd *cobra.Command, args []string) {
	name := cmd.Flag("name").Value.String()
	id, _ := cmd.Flags().GetUint64("id")
	offset, _ := cmd.Flags().GetUint64("offset")
	limit, _ := cmd.Flags().GetUint64("limit")

	if name == "" {
		panic("Domain is mandatory for this action.")
	}

	resp, err := app.Client.GetDomainLogs(name, id, offset, limit)
	if err != nil {
		panic(err)
	}

	// data, _ := convert.StringToJSONWithIndent(string(resp))
	// fmt.Println(data)

	// Parse Response
	var logs *LogList
	err = json.Unmarshal(resp, &logs)
	if err != nil {
		panic(err)
	}

	// data, _ := convert.StructToJSONWithIndent(logs)
	// fmt.Println(data)

	rows := make([][]string, 0)
	for idx, log := range logs.Response.LogList {
		row := make([]string, 0)
		row = append(row, strconv.Itoa(idx))
		row = append(row, log)

		rows = append(rows, row)
	}

	table := tablewriter.NewWriter(os.Stdout)
	tableHeader := []string{
		// "序号", "操作时间", "域名", "操作行为",
		"序号", "操作日志",
	}
	tableHeaderColor := make([]tablewriter.Colors, 0)
	for i := 0; i < len(tableHeader); i++ {
		tableHeaderColor = append(tableHeaderColor, tablewriter.Colors{tablewriter.Bold})
	}
	table.SetHeader(tableHeader)
	table.SetHeaderColor(tableHeaderColor...)
	table.SetColWidth(200)
	// table.SetRowLine(true)
	table.AppendBulk(rows)

	table.Render()
}
