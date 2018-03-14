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

	"github.com/okkur/reposeed/cmd/reposeed/config"
	templates "github.com/okkur/reposeed/cmd/reposeed/templates"
	yaml "gopkg.in/yaml.v2"
)

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

func generateFile(config config.Config, fileContent []byte, newPath string, overwrite bool) error {

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
	box := templates.GetTemplates()

	// Commands
	if os.Args[1] == "init" {
		seedString := box.String("seed-config.example.yaml")
		file, _ := os.Create(os.Args[2] + "/.seed-config.yaml")
		defer file.Close()
		file.WriteString(seedString)
		os.Exit(1)
	}

	config := parseConfig(conf)
	templatesName := box.List()
	bl := make(map[string]bool)
	bl["seed-config.example.yaml"] = true

	for _, templateName := range templatesName {
		file, _ := box.Open(templateName)
		fileStat, _ := file.Stat()
		fileContent := box.Bytes(templateName)
		if bl[fileStat.Name()] {
			log.Println(filepath.SkipDir)
		}

		if !fileStat.IsDir() {
			err := generateFile(config, fileContent, templateName, overwrite)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
