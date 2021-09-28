package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/chinatree/dnsctl/cmd/aliyun/domain"
	"github.com/chinatree/dnsctl/cmd/aliyun/group"
	"github.com/chinatree/dnsctl/cmd/aliyun/record"
	app "github.com/chinatree/dnsctl/internal/app/aliyun"
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
		Long:       fmt.Sprintf("%s is a simple command line client for aliyun dns.", app.Name),
		Short:      "aliyun dns command line client",
		SuggestFor: []string{app.Name},
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	rootCmd.AddCommand(
		domain.Cmd,
		group.Cmd,
		record.Cmd,
		versionCmd,
	)

	rootCmd.PersistentFlags().String("region-id", "default", "区域ID")
	rootCmd.PersistentFlags().String("access-key-id", "", "密钥ID")
	rootCmd.PersistentFlags().String("access-key-secret", "", "加密密钥")

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}

	os.Exit(0)
}
