package object

import (
	"bytes"
	"math/big"
)

const (
	SetObj = "Tập Hợp"
)

var (
	ENDLOOP = &Null{}
)

type Set interface {
	Object
	Contain(Object) *Boolean
}

type CountableSet interface {
	Set
	BeginIterate()
	NextElement() Object
}

type IntInterval struct {
	Upper      *Real
	Lower      *Real
	curElement *Int
}

func (interval *IntInterval) Type() ObjectType { return SetObj }
func (interval *IntInterval) Display() string {
	var out bytes.Buffer
	out.WriteString("[")
	out.WriteString(interval.Lower.Display())
	out.WriteString("..")
	out.WriteString(interval.Upper.Display())
	out.WriteString("]")

	return out.String()
}
func (interval *IntInterval) Contain(obj Object) *Boolean {
	switch obj := obj.(type) {
	case *Int:
		cond1 := obj.Less(interval.Upper) == TRUE || obj.Equal(interval.Upper) == TRUE
		cond2 := interval.Lower.Less(obj) == TRUE || interval.Lower.Equal(obj) == TRUE
		return Condition(cond1 && cond2)
	default:
		return FALSE
	}
}
func (interval *IntInterval) BeginIterate() {
	interval.curElement = NewInt(new(big.Int))
	interval.Lower.Value.Int(interval.curElement.Value)
	if !interval.Lower.Value.IsInt() {
		interval.curElement.Value.Add(interval.curElement.Value, IntOne)
	}
}
func (interval *IntInterval) NextElement() Object {
	if interval.curElement == nil {
		interval.BeginIterate()
	}
	if interval.Upper.Less(interval.curElement) == TRUE {
		return ENDLOOP
	}
	defer interval.curElement.Value.Add(interval.curElement.Value, IntOne)
	val := new(big.Int).Set(interval.curElement.Value)
	return NewInt(val)
}

type RealInterval struct {
	Upper *Int
	Lower *Int
}

func (interval *RealInterval) Type() string { return SetObj }
func (interval *RealInterval) Display() string {
	var out bytes.Buffer
	out.WriteString("[")
	out.WriteString(interval.Lower.Display())
	out.WriteString(",")
	out.WriteString(interval.Upper.Display())
	out.WriteString("]")

	return out.String()
}
func (interval *RealInterval) Contain(obj Object) *Boolean {
	switch obj := obj.(type) {
	case Realness:
		real := obj.ToReal()
		cond1 := real.Less(interval.Upper) == TRUE || real.Equal(interval.Upper) == TRUE
		cond2 := interval.Lower.Less(real) == TRUE || interval.Lower.Equal(real) == TRUE
		return Condition(cond1 && cond2)
	default:
		return FALSE
	}
}
