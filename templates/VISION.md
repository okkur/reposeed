# Vision
{{ .Project.Name }} is a {{ with .Vision.Type }}{{ . }}{{ else }}personal{{ end }} project. From our current viewpoint, its role is:
{{ .Vision.Overview }}

{{- with .Vision.Items }}
Our detailed items:
  {{ range . -}}
    * {{ . }}
{{ end -}}
{{- end -}}

{{ with .Vision.Concept }}
## Core Concept
{{ . -}}
{{ end -}}

{{ "" }}
{{ with .Vision.Aim }}
## General Aim
{{ . -}}
{{ end }}

## Communication
A lot of discussion about {{ .Project.Name }} happens within {{ .Repo.Type }} issues. Ideally, we will keep it that way until {{ .Project.Name }} becomes big enough that this turns into a problem. The advantages of this is that all documentation and issues are publicly searchable and easily linked to the source code.

## Contributions
The {{ .Project.Name }} project should always aim to devote a lot of time to making contributors feel like an important part of the community. We should strive to provide a place for anyone to be able to contribute. The aim of the core team should be to work on ways to showcase contributions, encourage adoption and show off great work. An example of where this worked is the [CocoaPods Quality Index](http://blog.cocoapods.org/CocoaPods.org-Two-point-Five/).
