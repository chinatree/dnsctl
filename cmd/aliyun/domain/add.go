package domain

import (
	"fmt"

	"github.com/spf13/cobra"

	app "github.com/chinatree/dnsctl/internal/app/aliyun"
	"github.com/chinatree/dnsctl/pkg/toolkits/convert"
)

var (
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "添加域名",
		Long:  "根据传入参数添加域名。",
		Example: fmt.Sprintf(
			"%s %s", app.Name,
			"domain add --name yuntree.com",
		),
		PreRun: app.ParseArgs,
		Run:    addFunc,
	}
)

// addFunc executes the "domain add" command.
func addFunc(cmd *cobra.Command, args []string) {
	name := cmd.Flag("name").Value.String()
	gid := cmd.Flag("gid").Value.String()
	rgid := cmd.Flag("rgid").Value.String()

	if name == "" {
		panic("DomainName is mandatory for this action.")
	}

	resp, err := app.Client.AddDomain(name, gid, rgid)
	if err != nil {
		panic(err)
	}

	data, _ := convert.StringToJSONWithIndent(string(resp))
	fmt.Println(data)
}
