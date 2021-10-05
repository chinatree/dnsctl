package domain

import (
	"fmt"

	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/tencent"
	"github.com/chinatree/dnsctl/pkg/toolkits/convert"
)

var (
	unlockCmd = &cobra.Command{
		Use:   "unlock",
		Short: "域名锁定解锁",
		Long:  "根据传入参数解锁域名。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"domain unlock -n yuntree.com --code xxxx",
		),
		PreRun: app.ParseArgs,
		Run:    unlockFunc,
	}
)

// unlockFunc executes the "domain unlock" command.
func unlockFunc(cmd *cobra.Command, args []string) {
	name := cmd.Flag("name").Value.String()
	id, _ := cmd.Flags().GetUint64("id")
	code := cmd.Flag("code").Value.String()

	if name == "" && id == 0 {
		panic("DomainName or DomainID is mandatory for this action.")
	}

	if code == "" {
		panic("LockCode is mandatory for this action.")
	}

	resp, err := app.Client.UnlockDomain(name, code, id)
	if err != nil {
		panic(err)
	}

	data, _ := convert.StringToJSONWithIndent(string(resp))
	fmt.Println(data)
}
