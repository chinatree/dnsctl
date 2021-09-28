package record

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
		Short:   "删除解析记录",
		Long:    "根据传入参数删除解析记录。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"record delete --id 721046463884277760",
		),
		PreRun: app.ParseArgs,
		Run:    deleteFunc,
	}
)

// deleteFunc executes the "record delete" command.
func deleteFunc(cmd *cobra.Command, args []string) {
	id := cmd.Flag("id").Value.String()

	if id == "" {
		panic("RecordID is mandatory for this action.")
	}

	resp, err := app.Client.DeleteDomainRecord(id)
	if err != nil {
		panic(err)
	}

	data, _ := convert.StringToJSONWithIndent(string(resp))
	fmt.Println(data)
}
