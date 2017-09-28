<!-- This form is for bug reports and feature requests ONLY!

If you're looking for help check out [our support guidelines](/SUPPORT.md){{ with .SupportLinks.Troubleshooting }} and the [troubleshooting guide]({{ . }}).{{ else }}.{{ end }}
-->

**Is this a BUG REPORT or FEATURE REQUEST?**:
bug | feature


**What happened**:

**What you expected to happen**:

**How to reproduce it (as minimally and precisely as possible)**:


**Anything else we need to know?**:

**Environment**:
- {{ .Project.Name }} version:  {{ with .IssueTemplate.Questions }}{{ range . }}
- {{ . }}:  {{ end }}{{ end }}
- Others:
