# Support using {{ .Project.Name }}

Welcome to {{ .Project.Name }}! We use {{ .Repo.Type }} for tracking bugs and feature requests.
This isn't the right place to get support for using {{ .Project.Name }}, but below you can find user centric resources.

If you're unsure where to post, use one of the [community platforms](#community-support):

There are helpful volunteers who may be able to help you. If your particular issue turns out to be a bug, it will find its way from there.

If it happens that you know the solution to an existing bug, please first open the issue in order to keep track of it.
Afterwards open the relevant pull/merge request that potentially fixes it.

Please remember this is a community project and you are not entitled to free support. Be kind to anyone helping out.
{{ with .Emails.CommercialSupport }}For commercial support reach out to {{ . }}.{{ end }}

## Community Support

The {{ .Project.Name }} community might be active on various platforms.

{{ with .SupportPlatforms }}The current platforms used by the community:
{{ range . }}* [{{ .Service }}]({{ .Link }})
{{ end }}{{ end }}

## Documentation

{{ with .SupportLinks.Documentation }}* [User Documentation]({{ . }}){{ end }}
{{ with .SupportLinks.Troubleshooting }}* [Troubleshooting Guide]({{ . }}){{ end }}
