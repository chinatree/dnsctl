package group

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
		Short:   "删除域名分组",
		Long:    "根据传入参数删除域名分组名称。",
		PreRun:  app.ParseArgs,
		Run:     deleteFunc,
	}
)

// deleteFunc executes the "group delete" command.
func deleteFunc(cmd *cobra.Command, args []string) {
	id := cmd.Flag("id").Value.String()

	if id == "" {
		panic("GroupID is mandatory for this action.")
	}

	resp, err := app.Client.DeleteDomainGroup(id)
	if err != nil {
		panic(err)
	}

	data, _ := convert.StringToJSONWithIndent(string(resp))
	fmt.Println(data)
}
