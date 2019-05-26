package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const SupportedConfigVersion = "v1"

var version = "dev"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Print the reposeed and config versions",
	Example: "reposeed version",
	Run:     versionHandler,
}

func versionHandler(cmd *cobra.Command, args []string) {
	fmt.Printf("Reposeed version: %s\nConfig Version: %s\n", version, SupportedConfigVersion)
}
