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

{{ if eq .Project.MainLicense "creativeCommonsBy" }}
  {{- printf "*Code is licensed under the [%s](/LICENSE)*  " (print "Creative Commons BY 4.0") }}
{{ else if eq .Project.MainLicense "creativeCommonsSa" }}
  {{- printf "*Code is licensed under the [%s](/LICENSE)*  " (print "Creative Commons SA 4.0") }}
{{ else if eq .Project.MainLicense "creativeCommonsBySa" }}
  {{- printf "*Code is licensed under the [%s](/LICENSE)*  " (print "Creative Commons BY-SA 4.0") }}
{{ else if eq .Project.MainLicense "creativeCommonsByNcSa" }}
  {{- printf "*Code is licensed under the [%s](/LICENSE)*  " (print "Creative Commons BY-NC-SA 4.0") }}
{{ else if eq .Project.MainLicense "creativeCommons0" }}
  {{- printf "*Code is licensed under the [%s](/LICENSE)*  " (print "Creative Commons CC0 1.0") }}
{{ else }}
  {{- printf "*Code is licensed under the [%s](/LICENSE)*  " (print "Apache License, Version 2.0") }}
{{ end -}}

{{ if eq .Project.DocsLicense "apache2" }}
  {{- printf "*Documentation/examples are licensed under [%s](/docs/LICENSE)* " (print "Apache License, Version 2.0") }}
{{ else if eq .Project.MainLicense "creativeCommonsBy" }}
  {{- printf "*Documentation/examples are licensed under [%s](/docs/LICENSE)*  " (print "Creative Commons BY 4.0") }}
{{ else if eq .Project.MainLicense "creativeCommonsSa" }}
  {{- printf "*Documentation/examples are licensed under [%s](/docs/LICENSE)*  " (print "Creative Commons SA 4.0") }}
{{ else if eq .Project.MainLicense "creativeCommonsByNcSa" }}
  {{- printf "*Documentation/examples are licensed under [%s](/docs/LICENSE)*  " (print "Creative Commons BY-NC-SA 4.0") }}
{{ else if eq .Project.MainLicense "creativeCommons0" }}
  {{- printf "*Documentation/examples are licensed under [%s](/docs/LICENSE)*  " (print "Creative Commons CC0 1.0") }}
{{ else }}
  {{- printf "*Documentation/examples are licensed under [%s](/docs/LICENSE)*  " (print "Creative Commons BY-SA 4.0") }}
{{ end -}}

*Illustrations, trademarks and third-party resources are owned by their respective party and are subject to different licensing.*

---

{{ if .Copyright.Owner -}}
  {{ printf "Copyright %s - %s" .Copyright.Year .Copyright.Owner }}
{{- else -}}
  {{ printf "Copyright %s - %s" .Copyright.Year (printf "The %s Authors" .Project.Name) }}
{{ end }}
