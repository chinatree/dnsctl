package domain

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/aliyun"
)

var (
	listCmd = &cobra.Command{
		Use:    "list",
		Short:  "获取域名列表",
		Long:   `根据传入参数查询该用户的域名列表。`,
		PreRun: app.ParseArgs,
		Run:    listFunc,
	}

	versionCodeM = map[string]string{
		"mianfei": "免费",
	}
)

// listFunc executes the "domain list" command.
func listFunc(cmd *cobra.Command, args []string) {
	keyword := cmd.Flag("keyword").Value.String()
	gid := cmd.Flag("gid").Value.String()
	rgid := cmd.Flag("rgid").Value.String()
	page, _ := cmd.Flags().GetInt("page")
	size, _ := cmd.Flags().GetInt("size")
	searchMode := cmd.Flag("search-mode").Value.String()
	starmark, _ := cmd.Flags().GetBool("starmark")

	resp, err := app.Client.GetDomains(
		keyword, gid, rgid, searchMode,
		page, size,
		starmark,
	)
	if err != nil {
		panic(err)
	}

	// data, _ := convert.StringToJSONWithIndent(string(resp))
	// fmt.Println(data)

	// Parse Response
	var domains *Domains
	err = json.Unmarshal(resp, &domains)
	if err != nil {
		panic(err)
	}

	// data, _ := convert.StructToJSONWithIndent(domains)
	// fmt.Println(data)

	rows := make([][]string, 0)
	for idx, domain := range domains.Domains.Domain {
		dnsServers := make([]string, 0)
		for _, dnsServer := range domain.DnsServers.DnsServer {
			dnsServers = append(dnsServers, dnsServer)
		}
		row := make([]string, 0)
		row = append(row, strconv.Itoa(idx))
		row = append(row, domain.DomainId)
		row = append(row, domain.DomainName)
		row = append(row, strings.Join(dnsServers, "\n"))
		row = append(row, strconv.FormatInt(int64(domain.RecordCount), 10))
		row = append(row, domain.GroupId)
		row = append(row, domain.GroupName)
		row = append(row, domain.ResourceGroupId)
		row = append(row, versionCodeM[domain.VersionCode])
		row = append(row, domain.Remark)

		rows = append(rows, row)
	}

	table := tablewriter.NewWriter(os.Stdout)
	tableHeader := []string{
		"序号", "域名ID", "域名名称", "DNS列表", "记录数", "域名分组ID", "域名分组名称",
		"资源组ID", "付费版本", "备注",
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
