package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/business_account_service/pkg/version"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   `version`,
	Short: "Prints business_account_service_cli version",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(version.VERSION)
		return nil
	},
}
