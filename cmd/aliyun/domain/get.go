package domain

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/aliyun"
)

var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "获取域名信息",
		Long:  "根据传入参数查询指定域名的信息。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"domain get --name yuntree.com",
		),
		PreRun: app.ParseArgs,
		Run:    getFunc,
	}
	getMainCmd = &cobra.Command{
		Use:   "getMain",
		Short: "获取主域名名称",
		Long:  "通过输入的参数，获取主域名名称。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"domain getMain --name cdn.yuntree.com",
		),
		PreRun: app.ParseArgs,
		Run:    getMainFunc,
	}
)

// getFunc executes the "domain get" command.
func getFunc(cmd *cobra.Command, args []string) {
	name := cmd.Flag("name").Value.String()
	detail, _ := cmd.Flags().GetBool("detail")

	if name == "" {
		panic("DomainName is mandatory for this action.")
	}

	resp, err := app.Client.GetDomain(name, detail)
	if err != nil {
		panic(err)
	}

	// data, _ := convert.StringToJSONWithIndent(string(resp))
	// fmt.Println(data)

	// Parse Response
	var domain *Domain
	err = json.Unmarshal(resp, &domain)
	if err != nil {
		panic(err)
	}

	dnsServers := make([]string, 0)
	for _, dnsServer := range domain.DnsServers.DnsServer {
		dnsServers = append(dnsServers, dnsServer)
	}
	rows := [][]string{
		{
			strconv.Itoa(0),
			domain.DomainId,
			domain.DomainName,
			strings.Join(dnsServers, "\n"),
			domain.GroupId,
			domain.GroupName,
			domain.ResourceGroupId,
			versionCodeM[domain.VersionCode],
			domain.Remark,
		},
	}

	table := tablewriter.NewWriter(os.Stdout)
	tableHeader := []string{
		"序号", "域名ID", "域名", "DNS列表", "域名分组ID", "域名分组名称", "资源组ID",
		"付费版本", "域名备注",
	}
	tableHeaderColor := make([]tablewriter.Colors, 0)
	for i := 0; i < len(tableHeader); i++ {
		tableHeaderColor = append(tableHeaderColor, tablewriter.Colors{tablewriter.Bold})
	}
	table.SetHeader(tableHeader)
	table.SetHeaderColor(tableHeaderColor...)
	table.AppendBulk(rows)

	table.Render()
}

// getMainFunc executes the "domain getMain" command.
func getMainFunc(cmd *cobra.Command, args []string) {
	name := cmd.Flag("name").Value.String()

	if name == "" {
		panic("DomainName is mandatory for this action.")
	}

	resp, err := app.Client.GetMainDomainName(name)
	if err != nil {
		panic(err)
	}

	// data, _ := convert.StringToJSONWithIndent(string(resp))
	// fmt.Println(data)

	// Parse Response
	var mainDomain *MainDomain
	err = json.Unmarshal(resp, &mainDomain)
	if err != nil {
		panic(err)
	}

	rows := [][]string{
		{
			strconv.Itoa(0),
			mainDomain.DomainName,
			mainDomain.RR,
			strconv.Itoa(mainDomain.DomainLevel),
		},
	}

	table := tablewriter.NewWriter(os.Stdout)
	tableHeader := []string{
		"序号", "域名", "主机记录信息", "级别",
	}
	tableHeaderColor := make([]tablewriter.Colors, 0)
	for i := 0; i < len(tableHeader); i++ {
		tableHeaderColor = append(tableHeaderColor, tablewriter.Colors{tablewriter.Bold})
	}
	table.SetHeader(tableHeader)
	table.SetHeaderColor(tableHeaderColor...)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.AppendBulk(rows)

	table.Render()
}
