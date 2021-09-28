package record

import (
	"fmt"

	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/aliyun"
	"github.com/chinatree/dnsctl/pkg/toolkits/convert"
)

var (
	remarkCmd = &cobra.Command{
		Use:   "remark",
		Short: "修改解析记录的备注",
		Long:  "根据传入参数修改解析记录的备注。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"record remark --id 721046463884277760 --remark 官网",
		),
		PreRun: app.ParseArgs,
		Run:    remarkFunc,
	}
)

// remarkFunc executes the "record remark" command.
func remarkFunc(cmd *cobra.Command, args []string) {
	id := cmd.Flag("id").Value.String()
	remark := cmd.Flag("remark").Value.String()

	if id == "" {
		panic("RecordID is mandatory for this action.")
	}

	resp, err := app.Client.UpdateDomainRecordRemark(id, remark)
	if err != nil {
		panic(err)
	}

	data, _ := convert.StringToJSONWithIndent(string(resp))
	fmt.Println(data)
}
