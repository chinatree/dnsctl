package record

import (
	"fmt"

	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/aliyun"
	"github.com/chinatree/dnsctl/pkg/toolkits/convert"
)

var (
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "添加解析记录",
		Long:  "根据传入参数添加解析记录。",
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
	rr := cmd.Flag("rr").Value.String()
	_type := cmd.Flag("type").Value.String()
	value := cmd.Flag("value").Value.String()
	ttl, _ := cmd.Flags().GetInt64("ttl")
	priority, _ := cmd.Flags().GetInt("priority")
	line := cmd.Flag("line").Value.String()

	if domain == "" || rr == "" || _type == "" || value == "" {
		panic("DomainName or RR or Type or Value is mandatory for this action.")
	}

	resp, err := app.Client.AddDomainRecord(
		domain, rr, _type, value, int(ttl), priority, line,
	)
	if err != nil {
		panic(err)
	}

	data, _ := convert.StringToJSONWithIndent(string(resp))
	fmt.Println(data)
}
