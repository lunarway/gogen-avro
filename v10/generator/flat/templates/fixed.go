package templates

const FixedTemplate = `
import (
	"io"
	"encoding/json"

	"github.com/actgardner/gogen-avro/v10/vm/types"
	"github.com/actgardner/gogen-avro/v10/util"
)

func {{ .SerializerMethod }}(r {{ .GoType }}, w io.Writer) error {
	_, err := w.Write(r[:])
	return err
}

type {{ .GoType }} [{{ .SizeBytes }}]byte

func (b *{{ .GoType }}) UnmarshalJSON(data []byte) (error) {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	codepoints := util.DecodeByteString(s)
	copy((*b)[:], codepoints)
	return nil
}

func (b {{ .GoType }}) MarshalJSON() ([]byte, error) {
	return []byte(util.EncodeByteString(b[:])), nil
}

`
