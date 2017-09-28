{{ with .Project.Image }}<img src="{{ . }}" width="300">{{ else }} # {{ .Project.Name }}{{ end }}

{{ .Project.OneLiner }}

{{ range .Badges }}[![{{ .Alt }}]({{ .Image }})]({{ .Link }}) {{ end }}

----

# {{ .Project.Name }}
{{ .Project.Description }}

{{ if eq .Project.State "beta" }}NOTE: This is a beta release, we do not consider it completely production ready yet. Use at your own risk.{{ end }}
{{ if eq .Project.State "unstable" }}NOTE: This is a work-in-progress, we do not consider it production ready. Use at your own risk.{{ end }}

----

Copyright {{ .Copyright.Year }} {{ with .Copyright.Owner }}{{ . }}{{ else }} The {{ .Project.Name }} Authors{{ end }}
