package object

import (
	"fmt"
	"strings"
)

const (
	StringObj = "Chuỗi"
)

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return StringObj }
func (s *String) Display() string  { return fmt.Sprintf("\"%s\"", s.Value) }
func (s *String) Add(right Object) Object {
	switch right := right.(type) {
	case *String:
		return &String{Value: s.Value + right.Value}
	default:
		return CANT_OPERATE
	}
}
func (s *String) Multiply(right Object) Object {
	switch right := right.(type) {
	case *Int:
		return &String{Value: strings.Repeat(s.Value, int(right.Value.Int64()))}
	default:
		return CANT_OPERATE
	}
}
