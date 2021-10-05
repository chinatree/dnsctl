package domain

import (
	"fmt"

	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/tencent"
	"github.com/chinatree/dnsctl/pkg/toolkits/convert"
)

var (
	lockCmd = &cobra.Command{
		Use:   "lock",
		Short: "锁定域名",
		Long:  "根据传入参数锁定域名。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"domain lock -n yuntree.com --day 7",
		),
		PreRun: app.ParseArgs,
		Run:    lockFunc,
	}
)

// lockFunc executes the "domain lock" command.
func lockFunc(cmd *cobra.Command, args []string) {
	name := cmd.Flag("name").Value.String()
	id, _ := cmd.Flags().GetUint64("id")
	day, _ := cmd.Flags().GetUint64("day")

	if name == "" && id == 0 {
		panic("DomainName or DomainID is mandatory for this action.")
	}

	if day == 0 {
		panic("LockDays is mandatory for this action.")
	}

	resp, err := app.Client.LockDomain(name, id, day)
	if err != nil {
		panic(err)
	}

	data, _ := convert.StringToJSONWithIndent(string(resp))
	fmt.Println(data)
}
