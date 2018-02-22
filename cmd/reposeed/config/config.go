package config

type Config struct {
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
	Docs struct {
		Development string `yaml:"development,omitempty"`
	} `yaml:"docs,omitempty"`
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
