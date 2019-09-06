package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	templatesBox "go.okkur.org/reposeed/cmd/templates"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Create the sample config file",
	Example: "reposeed init\nreposeed init -o ~/myproject",
	Run:     initHandler,
}

func initHandler(cmd *cobra.Command, args []string) {
	if outputDir == "" {
		outputDir, _ = os.Getwd()
	}
	if _, err := os.Stat(outputDir + "/.seed-config.yaml"); err == nil {
		log.Fatal(".seed-config.yaml already exists in the current directory. You can overwrite it using the -w flag")
	}

	box := templatesBox.GetTemplates()
	seedString := box.String("seed-config.example.yaml")
	file, _ := os.Create(outputDir + "/.seed-config.yaml")
	defer file.Close()
	file.WriteString(seedString)
	os.Exit(1)
}
