package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gobuffalo/packr"
	"github.com/okkur/reposeed/cmd/config"
	templatesBox "github.com/okkur/reposeed/cmd/templates"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

var outputDir, conf string
var overwrite bool

const SupportedConfigVersion = "v1"

var rootCmd = &cobra.Command{
	Use:   "reposeed",
	Short: "Extensive repository base files up and running in an instant",
	Long: `Start with the base layer necessary to focus on your project and not on the repository.  
	Licensing, structure, documentation and more boilerplate to get you started from your first commit.`,
	Run: RootHandler,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outputDir, "output", "o", "", "Output directory")
	rootCmd.PersistentFlags().StringVarP(&conf, "conf", "c", ".seed-config.yaml", "Config file")
	rootCmd.PersistentFlags().BoolVarP(&overwrite, "overwrite", "w", false, "Force overwrite files")
}

func RootHandler(cmd *cobra.Command, args []string) {
	box := templatesBox.GetTemplates()

	config := parseConfig(conf)
	bl := make(map[string]bool)
	bl["seed-config.example.yaml"] = true
	configVersion := config.Reposeed.ConfigVersion
	templates := parseTemplates(box)

	if configVersion == SupportedConfigVersion {
		for _, templateName := range box.List() {
			file, _ := box.Open(templateName)
			fileStat, _ := file.Stat()
			if bl[fileStat.Name()] {
				log.Println(filepath.SkipDir)
				continue
			}

			if fileStat.IsDir() {
				continue
			}

			if strings.Contains(templateName, "partials/") {
				continue
			}

			err := generateFile(config, templates, templateName, overwrite)
			if err != nil {
				log.Println(err)
			}
		}
	} else {
		log.Fatalf("Invalid config version. Currently supported versions: v1")
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("Couldn't execute the command: %s", err.Error())
		os.Exit(1)
	}
}

// Reads the config file and returns the filled config struct
func parseConfig(path string) config.Config {
	var conf config.Config

	filename, _ := filepath.Abs(path)
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		panic(err)
	}
	return conf
}

// Reads all of the templates from the packr box and returns them
// as a *template.Template instance
func parseTemplates(box packr.Box) *template.Template {
	templatesName := box.List()
	templates := &template.Template{}
	for _, templateName := range templatesName {
		templateFile, err := box.Open(templateName)
		if err != nil {
			log.Fatalf("could not open the template file: %s", templateName)
		}
		defer templateFile.Close()
		templateContent := box.String(templateName)
		templates.New(templateName).Parse(templateContent)
	}
	return templates
}

// Generates the base files using the templates and config struct
func generateFile(config config.Config, templates *template.Template, newPath string, overwrite bool) error {
	if _, e := os.Stat(newPath); os.IsNotExist(e) {
		os.MkdirAll(filepath.Dir(newPath), os.ModePerm)
	}

	if !overwrite {
		if _, e := os.Stat(newPath); !os.IsNotExist(e) {
			return fmt.Errorf("file %s not overwritten", newPath)
		}
	}

	file, err := os.Create(newPath)
	if err != nil {
		return fmt.Errorf("unable to create file: %s", err)
	}
	defer file.Close()

	err = templates.Lookup(newPath).Execute(file, config)
	if err != nil {
		return fmt.Errorf("unable to parse template: %s", err)
	}

	return nil
}
