package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Long:  "show version.",
	Run: func(cmd *cobra.Command, args []string) {
		VersionCommand()
	},
}

var (
	Version   = "Unknown version"
	BuildTime = "Unknown build time"
	GitCommit = "Unknown GitCommit"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

func VersionCommand() {
	fmt.Printf("Version %s (%s)\n", Version, GitCommit)
	fmt.Printf("Build at %s\n", BuildTime)
}
