/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	_ "os"

	_ "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	_ "github.com/spf13/viper"
)

// goGraphqlMicroserviceCmd represents the goGraphqlMicroservice command
var (
	// Used for flags.
	// cfgFile     string
	userLicense string

	goGraphqlMicroserviceCmd = &cobra.Command{
		Use:   "goGraphqlMicroservice",
		Short: "Generates a graphql golang microservice",
		Long: `This command skaffolds and generates a graphql golang microservice based on the proto file present in a directory. As a client
			it is imperative you first define a service interface via protocol buffers and provide this command with the path of the containing directory.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("goGraphqlMicroservice called")
		},
	}
)


func init() {
	rootCmd.AddCommand(goGraphqlMicroserviceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goGraphqlMicroserviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// goGraphqlMicroserviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
