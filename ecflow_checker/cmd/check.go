package cmd

import (
	"fmt"
	"github.com/perillaroc/ecflow-client-go/ecflow_checker/node_checker"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check the checklist",
	Long:  "check all tasks in config files.",
	Run: func(cmd *cobra.Command, args []string) {
		CheckCommand(checkConfigFilePath)
	},
}

var checkConfigFilePath string

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.PersistentFlags().StringVarP(
		&checkConfigFilePath, "config-path", "c", "", "config file path")
	checkCmd.MarkFlagRequired("config-path")
}

func CheckCommand(configPath string) {
	fmt.Printf("Config files directory: %s\n", configPath)
	var configFilePaths []string

	walkConfigDirectory := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			configFilePaths = append(configFilePaths, path)
		}
		return nil
	}

	err := filepath.Walk(configPath, walkConfigDirectory)

	if err != nil {
		log.Fatalf("load config files has error: %v\n", err)
	}

	for _, configFilePath := range configFilePaths {
		relFilePath, _ := filepath.Rel(configPath, configFilePath)
		fmt.Printf("-- config file: %s\n", relFilePath)
		node_checker.RunCheckTasks(configFilePath)
	}
}
