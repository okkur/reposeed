# Support using {{ .Project.Name }}

Welcome to {{ .Project.Name }}! We use {{ .Repo.Type }} for tracking bugs and feature requests.
{{ if .SupportPlatforms }}
This isn't the right place to get support for using {{ .Project.Name }}. You can find user centric resources below.

If you're unsure where to post, use one of the [community platforms](#community-support).

If your particular issue turns out to be a bug, it will find its way from there.

{{ end -}}

There are helpful volunteers who may be able to help you.

If it happens that you know the solution to an existing bug, please first open the issue in order to keep track of it.
Afterwards open the relevant pull/merge request that potentially fixes it.

Please remember this is a community project and you are not entitled to free support.
Be kind to anyone helping out.

{{ with .Emails.CommercialSupport -}}
  {{ printf "For commercial support reach out to %s" . }}
{{- end -}}

{{ "" }}
{{ if .SupportPlatforms }}
## Community Support
The {{ .Project.Name }} community might be active on various platforms.

{{ with .SupportPlatforms -}}
  The current platforms used by the community:
  {{ range . -}}
    * [{{ .Service }}]({{ .Link }})
  {{ end -}}
{{ end }}

{{- end -}}

{{ "" }}
## Documentation
{{- with .SupportLinks.Documentation }}
  {{ printf "* [User Documentation](%s)" . }}
{{ end -}}
{{- with .SupportLinks.Troubleshooting }}
  {{ printf "* [Troubleshooting Guide](%s)" . }}
{{ end -}}
