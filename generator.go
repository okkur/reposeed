package main

import (
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Project struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
	} `yaml:"project"`
	SupportPlatforms []struct {
		Service string `yaml:"service"`
		Link    string `yaml:"link"`
	} `yaml:"supportPlatforms"`
	SupportLinks struct {
		Documentation string `yaml:"documentation,omitempty"`
		Examples    string `yaml:"examples,omitempty"`
		Troubleshooting    string `yaml:"troubleshooting,omitempty"`
	} `yaml:"supportLinks,omitempty"`
	ContributionLinks struct {
		IssueTemplate string `yaml:"issueTemplate,omitempty"`
		StarterIssues string `yaml:"starterIssues,omitempty"`
	} `yaml:"contributionLinks"`
	Emails struct {
		CommercialSupport string `yaml:"commercialSupport,omitempty"`
		Security          string `yaml:"security"`
		Coc               string `yaml:"coc"`
	} `yaml:"emails"`
	Copyright struct {
		Owner string `yaml:"owner"`
		Year  string `yaml:"year"`
	} `yaml:"copyright"`
	Maintainers []struct {
		Name string `yaml:"name"`
		Nick string `yaml:"nick"`
	} `yaml:"maintainers"`
	Repo struct {
		Type string `yaml:"type"`
		Link string `yaml:"link"`
	} `yaml:"repo"`
}

func parseConfig(path string) Config {
	var config Config

	filename, _ := filepath.Abs(path)
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}
	return config
}

func generateFile(path string, newPath string) {
	temp, err := template.ParseFiles(path)
	if err != nil {
		log.Println("Unable to parse file: ", err)
		return
	}

	file, err := os.Create(newPath)
	defer file.Close()
	if err != nil {
		log.Println("Unable to create file: ", err)
		return
	}

	config := parseConfig(".seed-config.yaml")

	err = temp.Execute(file, config)
	if err != nil {
		log.Print("Unable to parse template: ", err)
		return
	}
}
func main() {
	dir := "templates/"
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.Name() == "seed-config.example.yaml" {
			return filepath.SkipDir
		}
		if !info.IsDir() {
			newPath := filepath.Base(path)
			generateFile(path, newPath)
		}
		return nil
	})

}
