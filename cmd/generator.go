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
		State       string `yaml:"state,omitempty"`
		OneLiner    string `yaml:"oneLiner,omitempty"`
		Image       string `yaml:"image,omitempty"`
		Website     string `yaml:"website,omitempty"`
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
		 string `yaml:",omitempty"`
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

func generateFile(config config, path string, newPath string, overwrite bool) error {
	temp, err := template.ParseFiles(path)
	if err != nil {
		return fmt.Errorf("unable to parse file: %s", err)
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

	err = temp.Execute(file, config)
	if err != nil {
		return fmt.Errorf("unable to parse template: %s", err)
	}

	return nil
}

func main() {
	var tempDir, conf string
	var overwrite bool

	flag.StringVar(&tempDir, "input", "templates", "Template directory")
	flag.StringVar(&conf, "conf", ".seed-config.yaml", "Config file")
	flag.BoolVar(&overwrite, "overwrite", false, "Force overwrite files")

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
			err := generateFile(config, path, newPath, overwrite)
			if err != nil {
				log.Println(err)
			}
		}

		return nil
	})

}
