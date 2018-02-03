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
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/erbesharat/reposeed/templates"
	"gopkg.in/yaml.v2"
)

type config struct {
	Project struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		State       string `yaml:"state,omitempty"`
		OneLiner    string `yaml:"oneLiner,omitempty"`
		Image       string `yaml:"image,omitempty"`
		Website     string `yaml:"website,omitempty"`
		Version     string `yaml:"version,omitempty"`
		MainLicense string `yaml:"mainLicense,omitempty"`
		DocsLicense string `yaml:"docsLicense,omitempty"`
	} `yaml:"project"`
	Vision struct {
		Type     string   `yaml:"type"`
		Items    []string `yaml:"items"`
		Concept  string   `yaml:"concept"`
		Overview string   `yaml:"overview"`
		Aim      string   `yaml:"aim"`
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
	Badges []struct {
		Image string `yaml:"image,omitempty"`
		Link  string `yaml:"link,omitempty"`
		Alt   string `yaml:"alt,omitempty"`
	} `yaml:"badges"`
	SupportLinks struct {
		Documentation   string `yaml:"documentation,omitempty"`
		Examples        string `yaml:"examples,omitempty"`
		Troubleshooting string `yaml:"troubleshooting,omitempty"`
	} `yaml:"supportLinks,omitempty"`
	Readme struct {
		UsageExample string `yaml:"usageExample,omitempty"`
	} `yaml:"readme,omitempty"`
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

func generateFile(config config, fileContent []byte, newPath string, overwrite bool) error {

	// Create a temporary file based on fileContent
	tmpfile, err := ioutil.TempFile("", "template")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	if _, err := tmpfile.Write(fileContent); err != nil {
		log.Fatal(err)
	}
	if _, e := os.Stat(newPath); os.IsNotExist(e) {
		os.MkdirAll(filepath.Dir(newPath), os.ModePerm)
	}

	if !overwrite {
		if _, e := os.Stat(newPath); !os.IsNotExist(e) {
			return fmt.Errorf("file %s not overwritten", newPath)
		}
	}

	file, err := os.Create(newPath)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("unable to create file: %s", err)
	}

	temp, err := template.ParseFiles(tmpfile.Name())
	if err != nil {
		return fmt.Errorf("unable to parse file: %s", err)
	}

	err = temp.Execute(file, config)
	if err != nil {
		return fmt.Errorf("unable to parse template: %s", err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func main() {
	var outputDir, conf string
	var overwrite bool
	flag.StringVar(&outputDir, "output", "", "Output directory")
	flag.StringVar(&conf, "conf", ".seed-config.yaml", "Config file")
	flag.BoolVar(&overwrite, "overwrite", false, "Force overwrite files")

	flag.Parse()

	config := parseConfig(conf)
	templatesName, _ := templates.AssetDir("templates")
	bl := make(map[string]bool)
	bl["seed-config.example.yaml"] = true

	for _, templateName := range templatesName {
		file, _ := templates.AssetInfo("templates/" + templateName)
		fileContent := templates.MustAsset("templates/" + templateName)
		if bl[file.Name()] {
			log.Println(filepath.SkipDir)
		}
		if !file.IsDir() {
			var newPath string
			if outputDir != "" {
				fileName := strings.Split(file.Name(), "/")
				newPath = filepath.Join(outputDir, fileName[1])
			}
			err := generateFile(config, fileContent, newPath, overwrite)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
