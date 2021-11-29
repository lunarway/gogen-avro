package flat

import (
	"fmt"

	avro "github.com/actgardner/gogen-avro/v10/schema"
)

type Checker struct {

}

func NewChecker() *Checker {
	return &Checker{}
}

func (c *Checker) Check(def avro.Node) error {

	switch d := def.(type) {
	case *avro.UnionField:
		if len(d.AvroTypes()) != 2 {
			return fmt.Errorf("unions are not generally supported, only to make nullable fields. Found a union with %d sub types", len(d.AvroTypes()))
		}
		if d.AvroTypes()[0].Name() != "Null" {
			return fmt.Errorf("unions are not generally supported, only to make nullable fields. First sub type was not a null it was a %s", d.AvroTypes()[0].Name())
		}
	}

	for _, child := range def.Children() {
		if err := c.Check(child); err != nil {
			return err
		}
	}

	return nil
}
