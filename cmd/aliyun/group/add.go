package group

import (
	"fmt"

	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/aliyun"
	"github.com/chinatree/dnsctl/pkg/toolkits/convert"
)

var (
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "添加域名分组",
		Long:  "根据传入参数添加域名分组。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"domain add --name 测试分组",
		),
		PreRun: app.ParseArgs,
		Run:    addFunc,
	}
)

// addFunc executes the "group add" command.
func addFunc(cmd *cobra.Command, args []string) {
	name := cmd.Flag("name").Value.String()

	if name == "" {
		panic("GroupName is mandatory for this action.")
	}

	resp, err := app.Client.AddDomainGroup(name)
	if err != nil {
		panic(err)
	}

	data, _ := convert.StringToJSONWithIndent(string(resp))
	fmt.Println(data)
}
