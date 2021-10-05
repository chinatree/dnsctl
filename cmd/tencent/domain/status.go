package domain

import (
	"fmt"

	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/tencent"
	"github.com/chinatree/dnsctl/pkg/toolkits/convert"
)

var (
	statusCmd = &cobra.Command{
		Use:   "status",
		Short: "设置域名的状态",
		Long:  "根据传入参数设置域名的状态。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"domain status -n yuntree.com --status ENABLE",
		),
		PreRun: app.ParseArgs,
		Run:    statusFunc,
	}
)

// statusFunc executes the "domain status" command.
func statusFunc(cmd *cobra.Command, args []string) {
	name := cmd.Flag("name").Value.String()
	id, _ := cmd.Flags().GetUint64("id")
	status := cmd.Flag("status").Value.String()

	if name == "" && id == 0 {
		panic("DomainName or DomainID is mandatory for this action.")
	}

	if status == "" {
		panic("Status is mandatory for this action.")
	}

	resp, err := app.Client.UpdateDomainStatus(name, status, id)
	if err != nil {
		panic(err)
	}

	data, _ := convert.StringToJSONWithIndent(string(resp))
	fmt.Println(data)
}
