package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/chinatree/dnsctl/cmd/dnsctl/aliyun"
	"github.com/chinatree/dnsctl/cmd/dnsctl/tencent"
	app "github.com/chinatree/dnsctl/internal/app/dnsctl"
)

func main() {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "版本号",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s v%s\n", app.Name, app.Version)
		},
	}

	rootCmd := &cobra.Command{
		Use:        app.Name,
		Long:       fmt.Sprintf("%s is a simple command line client for cloud dns.", app.Name),
		Short:      "cloud dns command line client",
		SuggestFor: []string{app.Name},
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	rootCmd.AddCommand(
		tencent.Cmd,
		aliyun.Cmd,
		versionCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}

	os.Exit(0)
}
