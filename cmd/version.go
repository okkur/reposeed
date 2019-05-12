package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

const SupportedConfigVersion = "v1"

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
	reposeedVersion, err := ioutil.ReadFile("VERSION")
	if err != nil {
		log.Fatalf("Couldn't read the version file: %s", err.Error())
	}
	fmt.Printf("Reposeed version: %sConfig Version: %s\n", string(reposeedVersion), SupportedConfigVersion)
}
