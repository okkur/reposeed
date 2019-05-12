package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var outputDir, conf string
var overwrite bool

var rootCmd = &cobra.Command{
	Use:   "reposeed",
	Short: "Extensive repository base files up and running in an instant",
	Long: `Start with the base layer necessary to focus on your project and not on the repository.  
	Licensing, structure, documentation and more boilerplate to get you started from your first commit.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Get the config and generate the files
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outputDir, "output", "o", "", "Output directory")
	rootCmd.PersistentFlags().StringVarP(&conf, "conf", "c", ".seed-config.yaml", "Config file")
	rootCmd.PersistentFlags().BoolVarP(&overwrite, "overwrite", "ow", false, "Force overwrite files")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("Couldn't execute the command: %s", err.Error())
		os.Exit(1)
	}
}
