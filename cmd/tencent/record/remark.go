package record

import (
	"fmt"

	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/tencent"
	"github.com/chinatree/dnsctl/pkg/toolkits/convert"
)

var (
	remarkCmd = &cobra.Command{
		Use:   "remark",
		Short: "修改解析记录的备注",
		Long:  "根据传入参数修改解析记录的备注。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"record remark --domain yuntree.com --id 666 --remark 官网",
		),
		PreRun: app.ParseArgs,
		Run:    remarkFunc,
	}
)

// remarkFunc executes the "record remark" command.
func remarkFunc(cmd *cobra.Command, args []string) {
	domain := cmd.Flag("domain").Value.String()
	did, _ := cmd.Flags().GetUint64("domain-id")
	id, _ := cmd.Flags().GetUint64("id")
	remark := cmd.Flag("remark").Value.String()

	if domain == "" || id == 0 {
		panic("Domain or RecordID is mandatory for this action.")
	}

	resp, err := app.Client.UpdateRecordRemark(domain, remark, id, did)
	if err != nil {
		panic(err)
	}

	data, _ := convert.StringToJSONWithIndent(string(resp))
	fmt.Println(data)
}
