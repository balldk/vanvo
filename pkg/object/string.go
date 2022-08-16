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
func (s *String) At(index int) Object {
	if index < len(s.Value) {
		return &String{Value: string(s.Value[index])}
	}
	return IndexError
}
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
func (s *String) Equal(right Object) *Boolean {
	switch right := right.(type) {
	case *String:
		return Condition(s.Value == right.Value)
	default:
		return INCOMPARABLE
	}
}
func (s *String) Less(right Object) *Boolean {
	switch right := right.(type) {
	case *String:
		return Condition(s.Value < right.Value)
	default:
		return INCOMPARABLE
	}
}
