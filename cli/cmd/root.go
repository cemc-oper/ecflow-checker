package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "ecflow_checker",
	Short: "ecflow_checker is a check list tool for ecFlow server.",
	Long: `ecflow_checker
A checklist tool for ecflow server. 
It checks node status and node variables according to time triggers.`,

	//Run: func(cmd *cobra.Command, args []string) {
	//},
}

func init() {
	var magenta = color.New(color.FgMagenta).SprintfFunc()
	fmt.Printf("%s\n",
		magenta("This is an experimental tool which is under development."))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
