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
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

const blackspaceRepo string = "BlackspacePlatform"
const repoUrl string = "https://github.com/stefanprodan/podinfo.git"
const repoName string = "podinfo"
const defaultService string = "default"
var podInfoStrToRemove = []string{
	"podinfo", "Podinfo","PODINFO",
}

var (
	serviceName string
	pkgDirectories = []string{
	"pkg/database",
	"pkg/database_test",
	"pkg/models",
	"pkg/middleware",
	"pkg/metrics",
	"pkg/logging",
	}

	svcDirectories = []string{
		"models",
		"alerts",
	}

	svcFiles = []string{
		"models/common.proto",
		"models/schema.proto",
	}

	pkgFiles = []string{
		"pkg/database/database.go",
		"pkg/database/doc.go",
		"pkg/database_test/database_test.go",
		"pkg/middleware/middleware.go",
		"pkg/middleware/middleware_test.go",
		"pkg/metrics/metrics.go",
		"pkg/metrics/metrics_test.go",
	}
)

func RemoveAllReferencesFromFile(path string, from []string, to string) {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		for _, searchStr := range from {
			if strings.Contains(line, searchStr) {
				lines[i] = strings.ReplaceAll(line, searchStr, to)
			}
		}
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(path, []byte(output), 0644)
		if err != nil {
			fmt.Errorf(err.Error())
		}
	}
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func WalkAndUpdate(serviceName string){
	searchDir := "./" + serviceName

	fileList := make([]string, 0)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return err
	})

	if e != nil {
		panic(e)
	}

	for _, file := range fileList {
		if fileExists(file){
			// remove all references of podinfo in file
			RemoveAllReferencesFromFile(file, podInfoStrToRemove, serviceName)
			// remove all references of stefanprodan
			var stefanSearchStr = []string{ "stefanprodan", "Stefan Prodan"}
			RemoveAllReferencesFromFile(file, stefanSearchStr, blackspaceRepo)
		}
	}
}


func RenameFiles(serviceName string) {
	dir := "./" + serviceName
	file, err := os.Open(dir)
	if err != nil {
		fmt.Errorf("failed opening directory: %s", err)
	}
	defer file.Close()

	list, err := file.Readdirnames(0) // 0 to read all files and folders
	if err != nil {
		fmt.Errorf("failed reading directory: %s", err)
	}
	re := regexp.MustCompile("[^A-Za-z]")
	for _, name := range list {
		if strings.Contains(name, repoName) {
			newName := re.ReplaceAllString(name, serviceName)
			err := os.Rename(filepath.Join(dir, name), filepath.Join(dir, newName))
			if err != nil {
				fmt.Errorf("error renaming file: %s", err)
				continue
			}
		}
		fmt.Println("File names have been changed")
	}
}

func GenerateDirectories(serviceName string, directories []string) {
	for _, directory := range directories {
		var path = "./" + serviceName + "/" + directory
		exec.Command("mkdir", path).Run()
	}
}

func GenerateFiles(serviceName string, files []string) {
	for _, directory := range files {
		var path = "./" + serviceName + "/" + directory
		exec.Command("touch", path).Run()
	}
}

func SkaffoldService(serviceName string){
	fmt.Println("skaffolding service")
	// create pkg specific directories
	GenerateDirectories(serviceName, pkgDirectories)
	fmt.Println("sucessfully created pkg specific directories")

	GenerateDirectories(serviceName, svcDirectories)
	fmt.Println("sucessfully created svc specific directories")

	GenerateFiles(serviceName, svcFiles)
	fmt.Println("sucessfully created svc specific files")

	GenerateFiles(serviceName, pkgFiles)
	fmt.Println("sucessfully created pkg specific files")
}

func RenameRepo(serviceName string){
	fmt.Println("renaming cloned repository to " + serviceName)
	// change the name of the repository to specified input argument
	exec.Command("mv", "-f", repoName, serviceName).Run()
}

func CloneMicroservicesRepoTemplate(serviceName string){
	fmt.Println("cloning repository " + repoUrl)
	// change the name of the repository to specified input argument
	exec.Command("git", "clone", repoUrl).Run()
}

func SetupService(serviceName string){
	CloneMicroservicesRepoTemplate(serviceName)
	RenameRepo(serviceName)
	SkaffoldService(serviceName)
	WalkAndUpdate(serviceName)
	RenameFiles(serviceName)
}

func GenerateService(cmd *cobra.Command, args []string)  error {
	if len(args) < 1 {
		// return fmt.Errorf("service name is required!")
		args = append(args, defaultService)
	}

	serviceName := args[0]
	SetupService(serviceName)

	os.Exit(1)
	return nil
}

// goRestMicroserviceCmd represents the goRestMicroservice command
var goRestMicroserviceCmd = &cobra.Command{
	Use:   "goRestMicroservice",
	Short: "Generates a rest based golang microservice",
	Long: `This command skaffolds and generates a rest golang microservice.`,
	RunE: GenerateService,
}

func init() {
	goRestMicroserviceCmd.Flags().StringVarP(&serviceName, "servicename", "s", "test", "Micro Service name (required)")
	goRestMicroserviceCmd.MarkFlagRequired("servicename")
	rootCmd.AddCommand(goRestMicroserviceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goRestMicroserviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// goRestMicroserviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
