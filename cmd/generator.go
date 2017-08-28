/*
Copyright 2017 The Reposeed Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
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
	Vision struct {
		Type        string `yaml:"type"`
		Items        []string `yaml:"items"`
		Concept        string `yaml:"concept"`
		Aim        string `yaml:"aim"`
	} `yaml:"vision"`
	Repo struct {
		Type string `yaml:"type"`
		Link string `yaml:"link"`
	} `yaml:"repo"`
	Copyright struct {
		Owner string `yaml:"owner"`
		Year  string `yaml:"year"`
	} `yaml:"copyright"`
	Cla struct {
		CopyrightHolder string `yaml:"copyrightHolder"`
	} `yaml:"cla"`
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
		Questions []string `yaml:"questions"`
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

func generateFile(config config, path string, newPath string) {
	temp, err := template.ParseFiles(path)
	if err != nil {
		log.Println("Unable to parse file: ", err)
		return
	}

	if _, err := os.Stat(newPath); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(newPath), os.ModePerm)
	}

	file, err := os.Create(newPath)
	defer file.Close()
	if err != nil {
		log.Println("Unable to create file: ", err)
		return
	}

	err = temp.Execute(file, config)
	if err != nil {
		log.Print("Unable to parse template: ", err)
		return
	}
}
func main() {
	var tempDir string
	var conf string
	flag.StringVar(&tempDir, "input", "templates", "Template directory")
	flag.StringVar(&conf, "conf", ".seed-config.yaml", "Config file")
	flag.Parse()

	config := parseConfig(conf)
	tempDir, _ = filepath.Abs(tempDir)

	bl := make(map[string]bool)
	bl["seed-config.example.yaml"] = true

	filepath.Walk(tempDir, func(path string, info os.FileInfo, err error) error {
		if bl[info.Name()] {
			return filepath.SkipDir
		}
		if !info.IsDir() {
			newPath, _ := filepath.Rel(tempDir, path)
			generateFile(config, path, newPath)
		}
		return nil
	})

}
