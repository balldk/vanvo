package object

import (
	"fmt"
	"math/big"
)

const (
	ComplexObj = "Số Phức"
)

func NewComplex(real Realness, imagine Realness) *Complex {
	fmt.Println(real.Type(), imagine.Type())
	return &Complex{Real: real, Imagine: imagine}
}

type Complex struct {
	Real    Realness
	Imagine Realness
}

func (c *Complex) Type() ObjectType { return ComplexObj }
func (c *Complex) Display() string {
	s := ""
	real, _ := c.Real.ToReal().Value.Float32()
	imagine, _ := c.Imagine.ToReal().Value.Float32()

	if real != 0 {
		s += c.Real.Display()
	}
	if real != 0 && imagine != 0 {
		if c.Imagine.ToReal().Value.Cmp(RealZero) == -1 {
			s += " - "
		} else {
			s += " + "
		}
	}
	if imagine != 0 {
		if imagine < 0 && real == 0 {
			s += "-"
			s += NewInt(big.NewInt(-1)).Multiply(c.Imagine).Display()
		}
		if imagine != 1 && imagine > 0 {
			s += c.Imagine.Display()
		}
		if imagine != -1 && imagine < 0 {
			s += NewInt(big.NewInt(-1)).Multiply(c.Imagine).Display()
		}
		s += "i"
	}
	if real == 0 && imagine == 0 {
		s = "0"
	}

	return s
}
func (c *Complex) Module() *Real {
	a := c.Real.Multiply(c.Real).(*Real)
	b := c.Imagine.Multiply(c.Imagine).(*Real)
	return a.Add(b).(*Real)
}
func (c *Complex) Add(right Object) Object {
	switch right := right.(type) {
	case Realness:
		real := right.Add(c.Real).(Realness)
		return NewComplex(real, c.Imagine)
	case *Complex:
		real := c.Real.Add(right.Real).(Realness)
		imagine := c.Imagine.Add(right.Imagine).(Realness)
		return NewComplex(real, imagine)
	default:
		return CANT_OPERATE
	}
}
func (c *Complex) Subtract(right Object) Object {
	switch right := right.(type) {
	case Realness:
		real := c.Real.Subtract(right).(Realness)
		return NewComplex(real, c.Imagine)
	case *Complex:
		real := c.Real.Subtract(right.Real).(Realness)
		imagine := c.Imagine.Subtract(right.Imagine).(Realness)
		return NewComplex(real, imagine)
	default:
		return CANT_OPERATE
	}
}
func (c *Complex) Multiply(right Object) Object {
	switch right := right.(type) {
	case Realness:
		real, ok1 := right.Multiply(c.Real).(Realness)
		imagine, ok2 := right.Multiply(c.Imagine).(Realness)

		if ok1 && ok2 {
			return NewComplex(real, imagine)
		}
		return CANT_OPERATE
	case *Complex:
		a, ok1 := c.Real.Multiply(right.Real).(Realness)
		b, ok2 := c.Imagine.Multiply(right.Imagine).(Realness)
		e, ok3 := c.Real.Multiply(right.Imagine).(Realness)
		f, ok4 := c.Imagine.Multiply(right.Real).(Realness)

		if ok1 && ok2 && ok3 && ok4 {
			real := a.Subtract(b).(Realness)
			imagine := e.Add(f).(Realness)
			return NewComplex(real, imagine)
		}
		return CANT_OPERATE
	default:
		return CANT_OPERATE
	}
}
func (c *Complex) Divide(right Object) Object {
	switch right := right.(type) {
	case Realness:
		real := c.Real.Divide(right).(Realness)
		imagine := c.Imagine.Divide(right).(Realness)
		return NewComplex(real, imagine)
	case *Complex:
		if right.Real.ToReal().IsZero() && right.Imagine.ToReal().IsZero() {
			return ZERO_DIVISION
		}

		a, ok1 := c.Real.Multiply(right.Real).(Realness)
		b, ok2 := c.Imagine.Multiply(right.Imagine).(Realness)
		e, ok3 := c.Imagine.Multiply(right.Real).(Realness)
		f, ok4 := c.Real.Multiply(right.Imagine).(Realness)

		if ok1 && ok2 && ok3 && ok4 {
			real := a.Add(b).(Realness).Divide(right.Module()).(Realness)
			imagine := e.Subtract(f).(Realness).Divide(right.Module()).(Realness)
			return NewComplex(real, imagine)
		}
		return CANT_OPERATE
	default:
		return CANT_OPERATE
	}
}
