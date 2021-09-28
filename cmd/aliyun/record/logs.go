package record

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/aliyun"
)

var (
	logsCmd = &cobra.Command{
		Use:   "logs",
		Short: "获取解析记录操作日志",
		Long:  "根据传入参数获取域名的解析操作日志。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"record logs --domain yuntree.com --lang zh",
		),
		PreRun: app.ParseArgs,
		Run:    logsFunc,
	}
)

// logsFunc executes the "record logs" command.
func logsFunc(cmd *cobra.Command, args []string) {
	domain := cmd.Flag("domain").Value.String()
	keyword := cmd.Flag("keyword").Value.String()
	lang := cmd.Flag("lang").Value.String()
	start := cmd.Flag("start").Value.String()
	end := cmd.Flag("end").Value.String()
	page, _ := cmd.Flags().GetInt("page")
	size, _ := cmd.Flags().GetInt("size")

	if domain == "" {
		panic("DomainName is mandatory for this action.")
	}

	resp, err := app.Client.GetDomainRecordLogs(
		domain, keyword, lang,
		start,
		end,
		page,
		size,
	)
	if err != nil {
		panic(err)
	}

	// data, _ := convert.StringToJSONWithIndent(string(resp))
	// fmt.Println(data)

	// Parse Response
	var logs *RecordLogs
	err = json.Unmarshal(resp, &logs)
	if err != nil {
		panic(err)
	}

	// data, _ := convert.StructToJSONWithIndent(logs)
	// fmt.Println(data)

	rows := make([][]string, 0)
	for idx, log := range logs.RecordLogs.RecordLog {
		row := make([]string, 0)
		row = append(row, strconv.Itoa(idx))
		row = append(row, time.Unix(log.ActionTimestamp/1e3, 0).Format("2006-01-02 15:04:05"))
		row = append(row, log.Message)

		rows = append(rows, row)
	}

	table := tablewriter.NewWriter(os.Stdout)
	tableHeader := []string{
		"序号", "操作时间", "操作行为",
	}
	tableHeaderColor := make([]tablewriter.Colors, 0)
	for i := 0; i < len(tableHeader); i++ {
		tableHeaderColor = append(tableHeaderColor, tablewriter.Colors{tablewriter.Bold})
	}
	table.SetHeader(tableHeader)
	table.SetHeaderColor(tableHeaderColor...)
	table.SetColWidth(200)
	table.SetRowLine(true)
	table.AppendBulk(rows)

	table.Render()
}
