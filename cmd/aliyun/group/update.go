package group

import (
	"fmt"

	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/aliyun"
	"github.com/chinatree/dnsctl/pkg/toolkits/convert"
)

var (
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "修改域名分组",
		Long:  "根据传入参数修改域名分组名称。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"domain update --id fd31288b67e94dbdbab444e70937fb19 --name 测试分组",
		),
		PreRun: app.ParseArgs,
		Run:    updateFunc,
	}
)

// updateFunc executes the "group update" command.
func updateFunc(cmd *cobra.Command, args []string) {
	id := cmd.Flag("id").Value.String()
	name := cmd.Flag("name").Value.String()

	if id == "" || name == "" {
		panic("GroupID and GroupName is mandatory for this action.")
	}

	resp, err := app.Client.UpdateDomainGroup(
		id, name,
	)
	if err != nil {
		panic(err)
	}

	data, _ := convert.StringToJSONWithIndent(string(resp))
	fmt.Println(data)
}
