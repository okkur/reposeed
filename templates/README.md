<img src="{{ .Project.Image }}" width="100">

{{ .Project.OneLiner }}

{{ range .Badges }}
[![{{ .Image }}]][{{ .Link }}]
{{ end }}

----

# {{ .Project.Name }}
{{ .Project.Description }}

{{ if eq .Project.State "beta" }}NOTE: This is a beta release, we do not consider it completely production ready yet. Use at your own risk.{{ end }}
{{ if eq .Project.State "unstable" }}NOTE: This is a work-in-progress, we do not consider it production ready. Use at your own risk.{{ end }}

----

Copyright {{ .Copyright.Year }} {{ with .Copyright.Owner }}{{ . }}{{ else }} The {{ .Project.Name }} Authors{{ end }}
