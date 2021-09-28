package record

import (
	"fmt"

	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/aliyun"
	"github.com/chinatree/dnsctl/pkg/toolkits/convert"
)

var (
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "修改解析记录",
		Long:  "根据传入参数修改解析记录。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"record update --id 721046463884277760 --rr www --type A --value 1.1.1.1",
		),
		PreRun: app.ParseArgs,
		Run:    updateFunc,
	}
)

// updateFunc executes the "record update" command.
func updateFunc(cmd *cobra.Command, args []string) {
	id := cmd.Flag("id").Value.String()
	rr := cmd.Flag("rr").Value.String()
	_type := cmd.Flag("type").Value.String()
	value := cmd.Flag("value").Value.String()
	ttl, _ := cmd.Flags().GetInt64("ttl")
	priority, _ := cmd.Flags().GetInt("priority")
	line := cmd.Flag("line").Value.String()

	if id == "" || rr == "" || _type == "" || value == "" {
		panic("RecordID or RR or Type or Value is mandatory for this action.")
	}

	resp, err := app.Client.UpdateDomainRecord(
		id, rr, _type, value, int(ttl), priority, line,
	)
	if err != nil {
		panic(err)
	}

	data, _ := convert.StringToJSONWithIndent(string(resp))
	fmt.Println(data)
}
