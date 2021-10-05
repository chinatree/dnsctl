package record

import (
	"fmt"

	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/tencent"
	"github.com/chinatree/dnsctl/pkg/toolkits/convert"
)

var (
	statusCmd = &cobra.Command{
		Use:   "status",
		Short: "设置解析记录的状态",
		Long:  "根据传入参数设置解析记录的状态。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"record status --domain yuntree.com --id 666 --status ENABLE",
		),
		PreRun: app.ParseArgs,
		Run:    statusFunc,
	}
)

// statusFunc executes the "record status" command.
func statusFunc(cmd *cobra.Command, args []string) {
	domain := cmd.Flag("domain").Value.String()
	did, _ := cmd.Flags().GetUint64("domain-id")
	id, _ := cmd.Flags().GetUint64("id")
	status := cmd.Flag("status").Value.String()

	if domain == "" || id == 0 {
		panic("Domain or RecordID or Status is mandatory for this action.")
	}

	resp, err := app.Client.UpdateRecordStatus(domain, status, id, did)
	if err != nil {
		panic(err)
	}

	data, _ := convert.StringToJSONWithIndent(string(resp))
	fmt.Println(data)
}
