// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v8/compiler"
	"github.com/actgardner/gogen-avro/v8/vm"
	"github.com/actgardner/gogen-avro/v8/vm/types"
)

type UnionStringLongIntFloatDoubleNullBoolTypeEnum int

const (
	UnionStringLongIntFloatDoubleNullBoolTypeEnumString UnionStringLongIntFloatDoubleNullBoolTypeEnum = 0

	UnionStringLongIntFloatDoubleNullBoolTypeEnumLong UnionStringLongIntFloatDoubleNullBoolTypeEnum = 1

	UnionStringLongIntFloatDoubleNullBoolTypeEnumInt UnionStringLongIntFloatDoubleNullBoolTypeEnum = 2

	UnionStringLongIntFloatDoubleNullBoolTypeEnumFloat UnionStringLongIntFloatDoubleNullBoolTypeEnum = 3

	UnionStringLongIntFloatDoubleNullBoolTypeEnumDouble UnionStringLongIntFloatDoubleNullBoolTypeEnum = 4

	UnionStringLongIntFloatDoubleNullBoolTypeEnumBool UnionStringLongIntFloatDoubleNullBoolTypeEnum = 6
)

type UnionStringLongIntFloatDoubleNullBool struct {
	String    string
	Long      int64
	Int       int32
	Float     float32
	Double    float64
	Null      *types.NullVal
	Bool      bool
	UnionType UnionStringLongIntFloatDoubleNullBoolTypeEnum
}

func writeUnionStringLongIntFloatDoubleNullBool(r *UnionStringLongIntFloatDoubleNullBool, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(5, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionStringLongIntFloatDoubleNullBoolTypeEnumString:
		return vm.WriteString(r.String, w)
	case UnionStringLongIntFloatDoubleNullBoolTypeEnumLong:
		return vm.WriteLong(r.Long, w)
	case UnionStringLongIntFloatDoubleNullBoolTypeEnumInt:
		return vm.WriteInt(r.Int, w)
	case UnionStringLongIntFloatDoubleNullBoolTypeEnumFloat:
		return vm.WriteFloat(r.Float, w)
	case UnionStringLongIntFloatDoubleNullBoolTypeEnumDouble:
		return vm.WriteDouble(r.Double, w)
	case UnionStringLongIntFloatDoubleNullBoolTypeEnumBool:
		return vm.WriteBool(r.Bool, w)
	}
	return fmt.Errorf("invalid value for *UnionStringLongIntFloatDoubleNullBool")
}

func NewUnionStringLongIntFloatDoubleNullBool() *UnionStringLongIntFloatDoubleNullBool {
	return &UnionStringLongIntFloatDoubleNullBool{}
}

func (r *UnionStringLongIntFloatDoubleNullBool) Serialize(w io.Writer) error {
	return writeUnionStringLongIntFloatDoubleNullBool(r, w)
}

func DeserializeUnionStringLongIntFloatDoubleNullBool(r io.Reader) (*UnionStringLongIntFloatDoubleNullBool, error) {
	t := NewUnionStringLongIntFloatDoubleNullBool()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func DeserializeUnionStringLongIntFloatDoubleNullBoolFromSchema(r io.Reader, schema string) (*UnionStringLongIntFloatDoubleNullBool, error) {
	t := NewUnionStringLongIntFloatDoubleNullBool()
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func (r *UnionStringLongIntFloatDoubleNullBool) Schema() string {
	return "[\"string\",\"long\",\"int\",\"float\",\"double\",\"null\",\"boolean\"]"
}

func (_ *UnionStringLongIntFloatDoubleNullBool) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionStringLongIntFloatDoubleNullBool) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionStringLongIntFloatDoubleNullBool) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionStringLongIntFloatDoubleNullBool) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionStringLongIntFloatDoubleNullBool) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionStringLongIntFloatDoubleNullBool) SetString(v string)  { panic("Unsupported operation") }
func (r *UnionStringLongIntFloatDoubleNullBool) SetLong(v int64) {
	r.UnionType = (UnionStringLongIntFloatDoubleNullBoolTypeEnum)(v)
}
func (r *UnionStringLongIntFloatDoubleNullBool) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.String{Target: (&r.String)}
	case 1:
		return &types.Long{Target: (&r.Long)}
	case 2:
		return &types.Int{Target: (&r.Int)}
	case 3:
		return &types.Float{Target: (&r.Float)}
	case 4:
		return &types.Double{Target: (&r.Double)}
	case 5:
		return r.Null
	case 6:
		return &types.Boolean{Target: (&r.Bool)}
	}
	panic("Unknown field index")
}
func (_ *UnionStringLongIntFloatDoubleNullBool) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionStringLongIntFloatDoubleNullBool) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionStringLongIntFloatDoubleNullBool) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionStringLongIntFloatDoubleNullBool) AppendArray() types.Field {
	panic("Unsupported operation")
}
func (_ *UnionStringLongIntFloatDoubleNullBool) Finalize() {}

func (r *UnionStringLongIntFloatDoubleNullBool) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	switch r.UnionType {
	case UnionStringLongIntFloatDoubleNullBoolTypeEnumString:
		return json.Marshal(map[string]interface{}{"string": r.String})
	case UnionStringLongIntFloatDoubleNullBoolTypeEnumLong:
		return json.Marshal(map[string]interface{}{"long": r.Long})
	case UnionStringLongIntFloatDoubleNullBoolTypeEnumInt:
		return json.Marshal(map[string]interface{}{"int": r.Int})
	case UnionStringLongIntFloatDoubleNullBoolTypeEnumFloat:
		return json.Marshal(map[string]interface{}{"float": r.Float})
	case UnionStringLongIntFloatDoubleNullBoolTypeEnumDouble:
		return json.Marshal(map[string]interface{}{"double": r.Double})
	case UnionStringLongIntFloatDoubleNullBoolTypeEnumBool:
		return json.Marshal(map[string]interface{}{"boolean": r.Bool})
	}
	return nil, fmt.Errorf("invalid value for *UnionStringLongIntFloatDoubleNullBool")
}

func (r *UnionStringLongIntFloatDoubleNullBool) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["string"]; ok {
		r.UnionType = 0
		return json.Unmarshal([]byte(value), &r.String)
	}
	if value, ok := fields["long"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Long)
	}
	if value, ok := fields["int"]; ok {
		r.UnionType = 2
		return json.Unmarshal([]byte(value), &r.Int)
	}
	if value, ok := fields["float"]; ok {
		r.UnionType = 3
		return json.Unmarshal([]byte(value), &r.Float)
	}
	if value, ok := fields["double"]; ok {
		r.UnionType = 4
		return json.Unmarshal([]byte(value), &r.Double)
	}
	if value, ok := fields["boolean"]; ok {
		r.UnionType = 6
		return json.Unmarshal([]byte(value), &r.Bool)
	}
	return fmt.Errorf("invalid value for *UnionStringLongIntFloatDoubleNullBool")
}
