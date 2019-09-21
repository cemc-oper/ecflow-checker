package cmd

import (
	"fmt"
	"github.com/perillaroc/ecflow-checker/node_checker"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"time"
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
		&checkConfigFilePath, "config-path", "c",
		os.Getenv("ECFLOW_CHECKER_CONFIG_PATH"),
		"config file path, or use environment variable ECFLOW_CHECKER_CONFIG_PATH")
}

func CheckCommand(configPath string) {
	current := time.Now().UTC()
	fmt.Printf("Current time is %s\n", current.Format("2006-01-02 15:04 MST"))

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
