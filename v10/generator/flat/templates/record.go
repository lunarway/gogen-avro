package templates

const RecordTemplate = `

{{ if containsDateTimeField . }}
	import "time"
{{ end }}

{{ if ne .Doc "" }}// {{ .Doc}}{{ end }}
type {{ .Name }} struct {
{{ range $i, $field := .Fields -}}
	{{ if ne $field.Doc "" }}// {{ $field.Doc }}{{ end }}
	{{ if ne $field.Tags "" -}}
		{{ $field.GoName }} {{ $field.Type.GoType }} ` + "`{{ $field.Tags }}`" + `
	{{ else -}}
		{{ $field.GoName }} {{ $field.Type.GoType }}
	{{ end -}}
{{ end }}
}

func (r *{{ .GoType }}) ApplyDefaults() {
{{ range $i, $field := .Fields -}}
	{{ if eq $field.Type.WrapperType "types.Record" -}}
		r.{{ $field.GoName }}.ApplyDefaults()
	{{ else if .HasDefault -}}
		{{ $.DefaultForField $field }}
	{{ end -}}
{{ end -}}
}
`
