<!--
This form is for bug reports and feature requests ONLY!  
If you're looking for help check out [our support guidelines](/SUPPORT.md)
{{- with .SupportLinks.Troubleshooting -}}
  {{- printf " and the [troubleshooting guide](%s)." . -}}
{{ else }}
  {{- printf "." -}}
{{ end }}
-->
**Bug report**

**What happened**:

**What you expected to happen**:

**How to reproduce it (as minimally and precisely as possible)**:

**Anything else we need to know?**:

**Environment**:
{{ printf "- %s version:  " .Project.Name }}
{{ with .IssueTemplate.Questions -}}
{{- range . -}}
  {{ printf "- %s:  " . }}
{{ end -}}
{{ end -}}
{{ printf "- Others:" }}
