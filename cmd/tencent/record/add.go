package record

import (
	"fmt"

	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/tencent"
	"github.com/chinatree/dnsctl/pkg/toolkits/convert"
)

var (
	addCmd = &cobra.Command{
		Use:     "add",
		Aliases: []string{"create"},
		Short:   "添加解析记录",
		Long:    "根据传入参数添加解析记录。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"record add --domain yuntree.com --rr www --type A --value 1.1.1.1",
		),
		PreRun: app.ParseArgs,
		Run:    addFunc,
	}
)

// addFunc executes the "record add" command.
func addFunc(cmd *cobra.Command, args []string) {
	domain := cmd.Flag("domain").Value.String()
	did, _ := cmd.Flags().GetUint64("domain-id")
	rr := cmd.Flag("rr").Value.String()
	_type := cmd.Flag("type").Value.String()
	ttl, _ := cmd.Flags().GetUint64("ttl")
	value := cmd.Flag("value").Value.String()
	status := cmd.Flag("status").Value.String()
	line := cmd.Flag("line").Value.String()
	lid := cmd.Flag("line-id").Value.String()
	mx, _ := cmd.Flags().GetUint64("mx")
	weight, _ := cmd.Flags().GetUint64("weight")

	if domain == "" || _type == "" || value == "" || line == "" {
		panic("Domain or Type or Value or RecordLine is mandatory for this action.")
	}

	resp, err := app.Client.AddRecord(
		domain, rr, _type, value, status, line, lid,
		ttl, mx, did, weight,
	)
	if err != nil {
		panic(err)
	}

	data, _ := convert.StringToJSONWithIndent(string(resp))
	fmt.Println(data)
}
