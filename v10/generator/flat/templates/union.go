package templates

const UnionTemplate = `
import (
	"encoding/json"
)


type {{ .UnionEnumType }} int
const (
{{ range $i, $t := .ItemTypes -}}
	{{ if ne $i $.NullIndex -}}
	{{ $.UnionEnumType }}{{ .Name }} {{ $.UnionEnumType }} = {{ $i }}
        {{ end }}
{{ end -}}
)

type {{ .Name }} struct {
{{ range $i, $t := .ItemTypes -}}
	{{ .Name }} {{ .GoType }}
{{ end -}}
	UnionType {{ $.UnionEnumType }}
}

func {{ .ConstructorMethod }} *{{ .GoType }} {
	return &{{ .Name }}{}
} 

func (r {{ .GoType }}) MarshalJSON() ([]byte, error) {
	{{ if ne .NullIndex -1 }}
	if r == nil {
		return []byte("null"), nil
	}
	{{ end }}

	return json.Marshal(r.{{ (index .ItemTypes 1).Name }})
}


func (r {{ .GoType }}) UnmarshalJSON(data []byte) (error) {
	if string(data) != "null" {
		r.UnionType = {{ $.UnionEnumType }}{{ (index .ItemTypes 1).Name }}
	}
	return json.Unmarshal(data, &r.{{ (index .ItemTypes 1).Name }})
}
`
