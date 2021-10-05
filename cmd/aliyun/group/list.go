package group

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/aliyun"
)

var (
	listCmd = &cobra.Command{
		Use:    "list",
		Short:  "获取域名分组列表",
		Long:   "根据传入参数获取所有分组列表。",
		PreRun: app.ParseArgs,
		Run:    listFunc,
	}

	versionCodeM = map[string]string{
		"mianfei": "免费",
	}
)

// listFunc executes the "group list" command.
func listFunc(cmd *cobra.Command, args []string) {
	keyword := cmd.Flag("keyword").Value.String()
	page, _ := cmd.Flags().GetInt("page")
	size, _ := cmd.Flags().GetInt("size")

	resp, err := app.Client.GetDomainGroups(
		keyword, page, size,
	)
	if err != nil {
		panic(err)
	}

	// data, _ := convert.StringToJSONWithIndent(string(resp))
	// fmt.Println(data)

	// Parse Response
	var groups *DomainGroups
	err = json.Unmarshal(resp, &groups)
	if err != nil {
		panic(err)
	}

	// data, _ := convert.StructToJSONWithIndent(groups)
	// fmt.Println(data)

	rows := make([][]string, 0)
	for idx, group := range groups.DomainGroups.DomainGroup {
		row := make([]string, 0)
		row = append(row, strconv.Itoa(idx))
		row = append(row, group.GroupID)
		row = append(row, group.GroupName)
		row = append(row, strconv.FormatInt(group.DomainCount, 10))

		rows = append(rows, row)
	}

	table := tablewriter.NewWriter(os.Stdout)
	tableHeader := []string{
		"序号", "域名分组ID", "域名分组名称", "域名分组数量",
	}
	tableHeaderColor := make([]tablewriter.Colors, 0)
	for i := 0; i < len(tableHeader); i++ {
		tableHeaderColor = append(tableHeaderColor, tablewriter.Colors{tablewriter.Bold})
	}
	table.SetHeader(tableHeader)
	table.SetHeaderColor(tableHeaderColor...)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetRowLine(true)
	table.AppendBulk(rows)

	table.Render()
}
