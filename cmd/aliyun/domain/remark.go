package domain

import (
	"fmt"

	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/aliyun"
	"github.com/chinatree/dnsctl/pkg/toolkits/convert"
)

var (
	remarkCmd = &cobra.Command{
		Use:   "remark",
		Short: "修改域名的备注",
		Long:  "根据传入参数修改域名的备注。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"domain remark --name yuntree.com --remark 国际域名",
		),
		PreRun: app.ParseArgs,
		Run:    remarkFunc,
	}
)

// remarkFunc executes the "domain remark" command.
func remarkFunc(cmd *cobra.Command, args []string) {
	name := cmd.Flag("name").Value.String()
	remark := cmd.Flag("remark").Value.String()

	if name == "" {
		panic("DomainName is mandatory for this action.")
	}

	resp, err := app.Client.UpdateDomainRemark(name, remark)
	if err != nil {
		panic(err)
	}

	data, _ := convert.StringToJSONWithIndent(string(resp))
	fmt.Println(data)
}
