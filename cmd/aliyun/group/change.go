package group

import (
	"fmt"

	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/aliyun"
	"github.com/chinatree/dnsctl/pkg/toolkits/convert"
)

var (
	changeCmd = &cobra.Command{
		Use:   "change",
		Short: "更改域名分组",
		Long:  "根据传入参数将域名从原分组更换到新分组。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"domain update --domain yuntree.com --id fd31288b67e94dbdbab444e70937fb19",
		),
		PreRun: app.ParseArgs,
		Run:    changeFunc,
	}
)

// changeFunc executes the "group change" command.
func changeFunc(cmd *cobra.Command, args []string) {
	domain := cmd.Flag("domain").Value.String()
	id := cmd.Flag("id").Value.String()

	if domain == "" {
		panic("DomainName is mandatory for this action.")
	}

	resp, err := app.Client.ChangeDomainGroup(
		domain, id,
	)
	if err != nil {
		panic(err)
	}

	data, _ := convert.StringToJSONWithIndent(string(resp))
	fmt.Println(data)
}
