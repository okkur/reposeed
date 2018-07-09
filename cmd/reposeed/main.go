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
	"text/template"

	"github.com/gobuffalo/packr"
	"github.com/okkur/reposeed/cmd/reposeed/config"
	templatesBox "github.com/okkur/reposeed/cmd/reposeed/templates"
	yaml "gopkg.in/yaml.v2"
)

const SupportedConfigVersion = "v1"

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

func main() {
	var outputDir, conf string
	var overwrite bool
	flag.StringVar(&outputDir, "output", "", "Output directory")
	flag.StringVar(&conf, "conf", ".seed-config.yaml", "Config file")
	flag.BoolVar(&overwrite, "overwrite", false, "Force overwrite files")

	flag.Parse()
	box := templatesBox.GetTemplates()

	// Commands
	if os.Args[1] == "init" {
		seedString := box.String("seed-config.example.yaml")
		file, _ := os.Create(os.Args[2] + "/.seed-config.yaml")
		defer file.Close()
		file.WriteString(seedString)
		os.Exit(1)
	}

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

			err := generateFile(config, templates, templateName, overwrite)
			if err != nil {
				log.Println(err)
			}
		}
	} else {
		log.Fatalf("Invalid config version. Currently supported versions: v1")
	}
}
