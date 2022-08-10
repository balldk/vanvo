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
	IsCountable() bool
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
func (interval *IntInterval) IsCountable() bool { return true }
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
	Upper *Real
	Lower *Real
}

func (interval *RealInterval) IsCountable() bool { return false }
func (interval *RealInterval) Type() ObjectType  { return SetObj }
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

type UnionSet struct {
	Left       Set
	Right      Set
	leftLooped bool
}

func (set *UnionSet) Type() ObjectType { return SetObj }
func (set *UnionSet) Display() string {
	return set.Left.Display() + " + " + set.Right.Display()
}
func (set *UnionSet) IsCountable() bool {
	return set.Left.IsCountable() && set.Right.IsCountable()
}
func (set *UnionSet) Contain(obj Object) *Boolean {
	if set.Left.Contain(obj) == TRUE {
		return TRUE
	}
	return set.Right.Contain(obj)
}
func (set *UnionSet) StartIterate() {

	if left, isCountable := set.Left.(CountableSet); isCountable {
		left.BeginIterate()
	}
	if right, isCountable := set.Right.(CountableSet); isCountable {
		right.BeginIterate()
	}
}
func (set *UnionSet) NextElement() Object {
	if left, isCountable := set.Left.(CountableSet); isCountable {
		element := left.NextElement()
		if element == ENDLOOP {
			set.leftLooped = true
		}
	}
	return ENDLOOP
}
