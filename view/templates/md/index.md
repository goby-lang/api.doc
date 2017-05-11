{{define "index"}}# All Classes
{{range .classes}}
- [{{.Name}}](classes/{{.Filename}}.md)
{{end}}
{{end}}
