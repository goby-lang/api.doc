{{define "class"}}# {{.class.Name}}
{{ range .class.Methods }}
### {{.FnName}}

{{.Comment}}

[[source]({{$.class.Repo}}/tree/{{$.class.Commit}}/vm/{{$.class.Filename}}.go#L{{.FnLine}})]
{{ end }}
{{ end }}

