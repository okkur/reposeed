package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "reposeed",
	Short: "Extensive repository base files up and running in an instant",
	Long: `Start with the base layer necessary to focus on your project and not on the repository.  
	Licensing, structure, documentation and more boilerplate to get you started from your first commit.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Get the config and generate the files
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("Couldn't execute the command: %s", err.Error())
		os.Exit(1)
	}
}
