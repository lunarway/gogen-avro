package templates

const EnumTemplate = `
import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
)

var _ = fmt.Printf

{{ if ne .Doc "" }}// {{ .Doc}}{{ end }}  
type {{ .GoType }} int32

const (
{{ range $i, $symbol := .Symbols -}}
	{{ $.SymbolName $symbol }} {{ $.GoType }} = {{ $i }}
{{ end -}}
)

func (e {{ .GoType  }}) String() string {
	switch e {
	{{ range $i, $symbol := .Symbols -}}
	case {{ $.SymbolName $symbol }}:
		return {{ printf "%q" $symbol }}
	{{ end -}}
	}
	return "unknown"
}

func {{ .SerializerMethod }}(r {{ .GoType }}, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func {{ .FromStringMethod }}(raw string) (r {{ .GoType }}, err error) {
	switch raw {
{{ range $i, $symbol := .Symbols -}}
	case {{ printf "%q" $symbol }}:
		return {{ $.SymbolName $symbol }}, nil
{{ end -}}
	}

{{ if ne .Default "" }} 
	return {{ $.SymbolName .Default }}, nil
{{ else }}
	return -1, fmt.Errorf("invalid value for {{ $.GoType }}: '%s'", raw)
{{ end }}
}

func (b {{ .GoType }}) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *{{ .GoType }}) UnmarshalJSON(data []byte) (error) {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := {{ .FromStringMethod }}(stringVal)
	*b = val
	return err
}

`
