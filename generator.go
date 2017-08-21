package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type config struct {
	Project struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
	} `yaml:"project"`
	Repo struct {
		Type string `yaml:"type"`
		Link string `yaml:"link"`
	} `yaml:"repo"`
	Copyright struct {
		Owner string `yaml:"owner"`
		Year  string `yaml:"year"`
	} `yaml:"copyright"`
	Maintainers []struct {
		Name string `yaml:"name"`
		Nick string `yaml:"nick,omitempty"`
	} `yaml:"maintainers,omitempty"`
	Emails struct {
		CommercialSupport string `yaml:"commercialSupport,omitempty"`
		Security          string `yaml:"security"`
		Coc               string `yaml:"coc"`
	} `yaml:"emails"`
	SupportLinks struct {
		Documentation   string `yaml:"documentation,omitempty"`
		Examples        string `yaml:"examples,omitempty"`
		Troubleshooting string `yaml:"troubleshooting,omitempty"`
	} `yaml:"supportLinks,omitempty"`
	SupportPlatforms []struct {
		Service string `yaml:"service"`
		Link    string `yaml:"link"`
	} `yaml:"supportPlatforms,omitempty"`
	IssueTemplate struct {
		Questions []struct {
			Question string `yaml:"question"`
		} `yaml:"questions"`
	} `yaml:"issueTemplate"`
	ContributionLinks struct {
		IssueTemplate string `yaml:"issueTemplate,omitempty"`
		StarterIssues string `yaml:"starterIssues,omitempty"`
	} `yaml:"contributionLinks"`
}

func parseConfig(path string) config {
	var conf config

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
	bl := make(map[string]bool)
	bl["seed-config.example.yaml"] = true
	bl[".gitlab"] = true

	dir := "templates/"
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if bl[info.Name()] {
			return filepath.SkipDir
		}
		if !info.IsDir() {
			newPath := filepath.Base(path)
			generateFile(path, newPath)
		}
		return nil
	})

}
