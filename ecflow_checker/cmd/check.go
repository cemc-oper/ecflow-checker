package cmd

import (
	"github.com/perillaroc/ecflow-client-go/ecflow_checker/node_checker"
	"github.com/spf13/cobra"
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
}

func CheckCommand(configPath string) {
	node_checker.RunCheckTasks(configPath)
}
