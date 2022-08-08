package object

import (
	"fmt"
	"math/big"
)

const (
	ComplexObj = "Số Phức"
)

func NewComplex(real *Real, imagine *Real) *Complex {
	return &Complex{Real: real, Imagine: imagine}
}

type Complex struct {
	Real    *Real
	Imagine *Real
}

func (c *Complex) Type() ObjectType { return ComplexObj }
func (c *Complex) Display() string {
	s := ""
	real, _ := c.Real.Value.Float32()
	imagine, _ := c.Imagine.Value.Float32()

	if real != 0 {
		s += fmt.Sprint(c.Real.Value)
	}
	if real != 0 && imagine != 0 {
		if c.Imagine.Value.Cmp(RealZero) == -1 {
			s += " - "
		} else {
			s += " + "
		}
	}
	if imagine != 0 {
		if imagine < 0 && real == 0 {
			s += "-"
		}
		if imagine != 1 && imagine != -1 {
			s += fmt.Sprint(new(big.Float).Abs(c.Imagine.Value))
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
		real := right.ToReal().Add(c.Real).(*Real)
		return NewComplex(real, c.Imagine)
	case *Complex:
		real := c.Real.Add(right.Real).(*Real)
		imagine := c.Imagine.Add(right.Imagine).(*Real)
		return NewComplex(real, imagine)
	default:
		return CANT_OPERATE
	}
}
func (c *Complex) Subtract(right Object) Object {
	switch right := right.(type) {
	case Realness:
		real := c.Real.Subtract(right.ToReal()).(*Real)
		return NewComplex(real, c.Imagine)
	case *Complex:
		real := c.Real.Subtract(right.Real).(*Real)
		imagine := c.Imagine.Subtract(right.Imagine).(*Real)
		return NewComplex(real, imagine)
	default:
		return CANT_OPERATE
	}
}
func (c *Complex) Multiply(right Object) Object {
	switch right := right.(type) {
	case Realness:
		r := right.ToReal()
		real := r.Multiply(c.Real).(*Real)
		imagine := r.Multiply(c.Imagine).(*Real)
		return NewComplex(real, imagine)
	case *Complex:
		a := c.Real.Multiply(right.Real).(*Real)
		b := c.Imagine.Multiply(right.Imagine).(*Real)
		e := c.Real.Multiply(right.Imagine).(*Real)
		f := c.Imagine.Multiply(right.Real).(*Real)

		real := a.Subtract(b).(*Real)
		imagine := e.Add(f).(*Real)
		return NewComplex(real, imagine)
	default:
		return CANT_OPERATE
	}
}
func (c *Complex) Divide(right Object) Object {
	switch right := right.(type) {
	case Realness:
		r := right.ToReal()
		real := c.Real.Divide(r).(*Real)
		imagine := c.Imagine.Divide(r).(*Real)
		return NewComplex(real, imagine)
	case *Complex:
		if right.Real.IsZero() && right.Imagine.IsZero() {
			return ZERO_DIVISION
		}

		a := c.Real.Multiply(right.Real).(*Real)
		b := c.Imagine.Multiply(right.Imagine).(*Real)
		e := c.Imagine.Multiply(right.Real).(*Real)
		f := c.Real.Multiply(right.Imagine).(*Real)

		real := a.Add(b).(*Real).Divide(right.Module()).(*Real)
		imagine := e.Subtract(f).(*Real).Divide(right.Module()).(*Real)
		return NewComplex(real, imagine)
	default:
		return CANT_OPERATE
	}
}
