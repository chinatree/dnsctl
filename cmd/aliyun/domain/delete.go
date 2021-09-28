package domain

import (
	"fmt"

	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/aliyun"
	"github.com/chinatree/dnsctl/pkg/toolkits/convert"
)

var (
	deleteCmd = &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del"},
		Short:   "删除域名",
		Long:    "根据传入参数删除域名。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"domain delete --name yuntree.com",
		),
		PreRun: app.ParseArgs,
		Run:    deleteFunc,
	}
)

// deleteFunc executes the "domain delete" command.
func deleteFunc(cmd *cobra.Command, args []string) {
	name := cmd.Flag("name").Value.String()

	if name == "" {
		panic("DomainName is mandatory for this action.")
	}

	resp, err := app.Client.DeleteDomain(name)
	if err != nil {
		panic(err)
	}

	data, _ := convert.StringToJSONWithIndent(string(resp))
	fmt.Println(data)
}
