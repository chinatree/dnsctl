package record

import (
	"fmt"

	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/tencent"
	"github.com/chinatree/dnsctl/pkg/toolkits/convert"
)

var (
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "修改解析记录",
		Long:  "根据传入参数修改解析记录。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"record update --domain yuntree.com --id 666 --rr www --type A --value 1.1.1.1",
		),
		PreRun: app.ParseArgs,
		Run:    updateFunc,
	}
)

// updateFunc executes the "record update" command.
func updateFunc(cmd *cobra.Command, args []string) {
	domain := cmd.Flag("domain").Value.String()
	_type := cmd.Flag("type").Value.String()
	status := cmd.Flag("status").Value.String()
	line := cmd.Flag("line").Value.String()
	rr := cmd.Flag("rr").Value.String()
	value := cmd.Flag("value").Value.String()
	lid := cmd.Flag("line-id").Value.String()
	id, _ := cmd.Flags().GetUint64("id")
	ttl, _ := cmd.Flags().GetUint64("ttl")
	mx, _ := cmd.Flags().GetUint64("mx")
	did, _ := cmd.Flags().GetUint64("domain-id")
	weight, _ := cmd.Flags().GetUint64("weight")

	if domain == "" || id == 0 || _type == "" || value == "" || line == "" {
		panic("Domain or RecordID or Type or Value or RecordLine is mandatory for this action.")
	}

	resp, err := app.Client.UpdateRecord(
		domain, rr, _type, value, status, line, lid,
		id, ttl, mx, did, weight,
	)
	if err != nil {
		panic(err)
	}

	data, _ := convert.StringToJSONWithIndent(string(resp))
	fmt.Println(data)
}
