{{- if .Project.Image -}}
  {{- if .Project.Website -}}
    {{- printf "<a href='%s'><img src='%s' width='500'/></a>" .Project.Website .Project.Image -}}
  {{- else -}}
    {{- printf "<img src='%s' width='500'/>" .Project.Image -}}
  {{- end -}}
{{- else -}}
# {{ .Project.Name }}
{{ end }}

{{ .Project.OneLiner }}

{{ range .Badges }}
  {{- printf " [![%s](%s)](%s)" .Alt .Image .Link -}}
{{ end }}

{{ if eq .Project.State "beta" }}
  {{- printf "**NOTE: This is a beta release, we do not consider it completely production ready yet. Use at your own risk.**" -}}
{{ end }}

{{ if eq .Project.State "unstable" -}}
  {{- printf "**NOTE: This is a work-in-progress, we do not consider it production ready. Use at your own risk.**" -}}
{{ end }}

{{- printf "# %s" .Project.Name }}
{{ .Project.Description -}}

{{ "" }}
{{ printf "## Using %s" .Project.Name }}
{{ with .Readme.UsageExample -}}
  {{- . -}}
{{- end -}}

{{- with .SupportLinks.Documentation -}}
Take a look at our full [documentation]({{ . }}).
{{- end }}

## Support
For detailed information on support options see our [support guide](/SUPPORT.md).

## Helping out
Best place to start is our [contribution guide](/CONTRIBUTING.md).

----

*Code is licensed under the [Apache License, Version 2.0](/LICENSE)*  
*Documentation is licensed under [Creative Commons BY-SA 4.0](/docs/LICENSE)*  

*Illustrations, trademarks and third-party resources are owned by their respective party and are subject to different licensing.*

---

{{ if .Copyright.Owner -}}
  {{- $owner := .Copyright.Owner -}}
  {{ printf "Copyright %s - %s" .Copyright.Year $owner }}
{{- else -}}
  {{- $owner := "output" | printf "The %s Authors" .Project.Name -}}
  {{ printf "Copyright %s - %s" .Copyright.Year $owner }}
{{ end }}
