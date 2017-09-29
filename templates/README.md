{{ if .Project.Image }}
{{ if .Project.Website }}
<a href="{{ .Project.Website }}"><img src="{{ .Project.Image }}" width="500"/></a>
{{ else }}
<img src="{{ .Project.Image }}" width="500"/>
{{ end }}
{{ else }}
# {{ .Project.Name }}
{{ end }}

{{ .Project.OneLiner }}

{{ range .Badges }}[![{{ .Alt }}]({{ .Image }})]({{ .Link }}) {{ end }}

----

{{ if eq .Project.State "beta" }}**NOTE: This is a beta release, we do not consider it completely production ready yet. Use at your own risk.**{{ end }}
{{ if eq .Project.State "unstable" }}**NOTE: This is a work-in-progress, we do not consider it production ready. Use at your own risk.**{{ end }}

{{ if .Project.State }}
---
{{ end }}

# {{ .Project.Name }}
{{ .Project.Description }}

---

## Using {{ .Project.Name }}

{{ with .Readme.UsageExample }}{{ . }}{{ end }}

Take a look at our full [documentation]({{ .SupportLinks.Documentation }}).

## Helping out

<!-- contributing details -->

----

*Code is licensed under the [Apache License, Version 2.0](/LICENSE)*  
*Documentation is licensed under [Creative Commons BY 4.0](/docs/LICENSE)*  

*Illustrations, trademarks and third-party resources are owned by their respective party and are subject to different licensing.*

---

Copyright {{ .Copyright.Year }} - {{ with .Copyright.Owner }}{{ . }}{{ else }} The {{ .Project.Name }} Authors{{ end }}

[![Analytics](https://use.okkur.net/piwik.php?idsite=13rec=1)]()

{{ if .Tracking }}
[![Analytics]({{ .Tracking.Link }})]()
{{ end }}
