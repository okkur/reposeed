package cmd

import (
	"os"

	templatesBox "github.com/okkur/reposeed/cmd/templates"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Create the sample config file",
	Example: "reposeed init ~/myproject",
	Run:     initHandler,
}

func initHandler(cmd *cobra.Command, args []string) {
	box := templatesBox.GetTemplates()
	seedString := box.String("seed-config.example.yaml")
	file, _ := os.Create(args[0] + "/.seed-config.yaml")
	defer file.Close()
	file.WriteString(seedString)
	os.Exit(1)
}
