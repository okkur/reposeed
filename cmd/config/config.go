/*
Copyright 2018 - The Reposeed authors
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

package config

type Config struct {
	Project struct {
		Name        string `yaml:"name" json:"name"`
		Description string `yaml:"description" json:"description"`
		State       string `yaml:"state,omitempty" json:"state,omitempty"`
		OneLiner    string `yaml:"oneLiner,omitempty" json:"oneLiner,omitempty"`
		Image       string `yaml:"image,omitempty" json:"image,omitempty"`
		Website     string `yaml:"website,omitempty" json:"website,omitempty"`
		Version     string `yaml:"version,omitempty" json:"version,omitempty"`
		MainLicense string `yaml:"mainLicense,omitempty" json:"mainLicense,omitempty"`
		DocsLicense string `yaml:"docsLicense,omitempty" json:"docsLicense,omitempty"`
	} `yaml:"project" json:"project"`
	Vision struct {
		Type     string   `yaml:"type" json:"type"`
		Items    []string `yaml:"items" json:"items"`
		Concept  string   `yaml:"concept" json:"concept"`
		Overview string   `yaml:"overview" json:"overview"`
		Aim      string   `yaml:"aim" json:"aim"`
	} `yaml:"vision" json:"vision"`
	Repo struct {
		Type string `yaml:"type" json:"type"`
		Link string `yaml:"link" json:"link"`
	} `yaml:"repo" json:"repo"`
	Copyright struct {
		Owner string `yaml:"owner" json:"owner"`
		Year  string `yaml:"year" json:"year"`
	} `yaml:"copyright" json:"copyright"`
	Cla struct {
		CopyrightHolder string `yaml:"copyrightHolder" json:"copyrightHolder"`
		Mail            string `yaml:"mail" json:"mail"`
	} `yaml:"cla" json:"cla"`
	Maintainers []struct {
		Name string `yaml:"name" json:"name"`
		Nick string `yaml:"nick,omitempty" json:"nick,omitempty"`
	} `yaml:"maintainers,omitempty" json:"maintainers,omitempty"`
	Emails struct {
		Security          string `yaml:"security" json:"security"`
		Coc               string `yaml:"coc" json:"coc"`
		CommercialSupport string `yaml:"commercialSupport,omitempty" json:"commercialSupport,omitempty"`
	} `yaml:"emails" json:"emails"`
	Badges []struct {
		Image string `yaml:"image,omitempty" json:"image,omitempty"`
		Link  string `yaml:"link,omitempty" json:"link,omitempty"`
		Alt   string `yaml:"alt,omitempty" json:"alt,omitempty"`
	} `yaml:"badges" json:"badges"`
	SupportLinks struct {
		Documentation   string `yaml:"documentation,omitempty" json:"documentation,omitempty"`
		Examples        string `yaml:"examples,omitempty" json:"examples,omitempty"`
		Troubleshooting string `yaml:"troubleshooting,omitempty" json:"troubleshooting,omitempty"`
	} `yaml:"supportLinks,omitempty" json:"supportLinks,omitempty"`
	Readme struct {
		UsageExample string `yaml:"usageExample,omitempty" json:"usageExample,omitempty"`
	} `yaml:"readme,omitempty" json:"readme,omitempty"`
	Docs struct {
		Development string `yaml:"development,omitempty" json:"development,omitempty"`
	} `yaml:"docs,omitempty" json:"docs,omitempty"`
	SupportPlatforms []struct {
		Service string `yaml:"service" json:"service"`
		Link    string `yaml:"link" json:"link"`
	} `yaml:"supportPlatforms,omitempty" json:"supportPlatforms,omitempty"`
	IssueTemplate struct {
		Questions []string `yaml:"questions" json:"questions"`
	} `yaml:"issueTemplate" json:"issueTemplate"`
	ContributionLinks struct {
		IssueTemplate string `yaml:"issueTemplate,omitempty" json:"issueTemplate,omitempty"`
		StarterIssues string `yaml:"starterIssues,omitempty" json:"starterIssues,omitempty"`
	} `yaml:"contributionLinks" json:"contributionLinks"`
	Reposeed struct {
		ConfigVersion string `yaml:"configVersion,omitempty" json:"configVersion,omitempty"`
	} `yaml:"reposeed" json:"reposeed"`
}
