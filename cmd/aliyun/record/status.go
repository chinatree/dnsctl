package record

import (
	"fmt"

	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/aliyun"
	"github.com/chinatree/dnsctl/pkg/toolkits/convert"
)

var (
	statusCmd = &cobra.Command{
		Use:   "status",
		Short: "设置解析记录状态",
		Long:  "根据传入参数设置解析记录状态。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"record status --id 721046463884277760 --status ENABLE",
		),
		PreRun: app.ParseArgs,
		Run:    statusFunc,
	}
)

// statusFunc executes the "record status" command.
func statusFunc(cmd *cobra.Command, args []string) {
	id := cmd.Flag("id").Value.String()
	status := cmd.Flag("status").Value.String()

	if id == "" || status == "" {
		panic("RecordID or Status is mandatory for this action.")
	}

	resp, err := app.Client.SetDomainRecordStatus(id, status)
	if err != nil {
		panic(err)
	}

	data, _ := convert.StringToJSONWithIndent(string(resp))
	fmt.Println(data)
}
