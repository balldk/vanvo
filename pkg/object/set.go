package object

import (
	"bytes"
)

const (
	SetObj = "Tập Hợp"
)

type IterateCallback func(Object)

type Set interface {
	Object
	Contain(Object) *Boolean
	IsCountable() bool
}

type CountableSet interface {
	Set
	Iterate(IterateCallback)
}

type List struct {
	Data []Object
}

func (list *List) Type() ObjectType { return SetObj }
func (list *List) Display() string {
	var out bytes.Buffer
	out.WriteString("{")

	for ind, each := range list.Data {
		out.WriteString(each.Display())
		if ind != len(list.Data)-1 {
			out.WriteString(",")
		}
	}
	out.WriteString("}")

	return out.String()
}
func (list *List) IsCountable() bool { return true }
func (list *List) Contain(obj Object) *Boolean {
	if obj, ok := obj.(Equal); ok {
		for _, each := range list.Data {
			if obj.Equal(each).Value {
				return TRUE
			}
		}
	}
	return FALSE
}
func (list *List) Iterate(callback IterateCallback) {
	for _, each := range list.Data {
		callback(each)
	}
}

type IntInterval struct {
	Upper Realness
	Lower Realness
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
func (interval *IntInterval) Iterate(callback IterateCallback) {
	element := interval.Lower
	for element.Less(interval.Upper) == TRUE || element.Equal(interval.Upper) == TRUE {
		callback(element)
		element = element.Add(NewInt(IntOne)).(Realness)
	}
}

type RealInterval struct {
	Upper Realness
	Lower Realness
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
func (interval *RealInterval) Less(right Object) *Boolean {
	switch right := right.(type) {
	case *IntInterval:
		cond1 := interval.Lower.Less(right.Lower).Not()
		cond2 := interval.Upper.Less(right.Upper).Or(interval.Upper.Equal(right.Upper))
		return cond1.And(cond2)
	default:
		return INCOMPARABLE
	}
}

type UnionSet struct {
	Left  Set
	Right Set
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
func (set *UnionSet) Iterate(callback IterateCallback) {
	if left, isCountable := set.Left.(CountableSet); isCountable {
		left.Iterate(callback)
	}
	if right, isCountable := set.Right.(CountableSet); isCountable {
		right.Iterate(callback)
	}
}

type DiffSet struct {
	Left  Set
	Right Set
}

func (set *DiffSet) Type() ObjectType { return SetObj }
func (set *DiffSet) Display() string {
	return set.Left.Display() + " + " + set.Right.Display()
}
func (set *DiffSet) IsCountable() bool {
	return set.Left.IsCountable()
}
func (set *DiffSet) Contain(obj Object) *Boolean {
	if set.Left.Contain(obj) == FALSE {
		return FALSE
	}
	return set.Right.Contain(obj).Not()
}
func (set *DiffSet) Iterate(callback IterateCallback) {
	if left, isCountable := set.Left.(CountableSet); isCountable {
		left.Iterate(func(element Object) {
			if !set.Right.Contain(element).Value {
				callback(element)
			}
		})
	}
}
