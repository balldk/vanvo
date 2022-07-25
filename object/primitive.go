package object

import "fmt"

const (
	INT_OBJ      = "Số Nguyên"
	REAL_OBJ     = "Số Thực"
	QUOTIENT_OBJ = "Số Hữu Tỉ"
	BOOL_OBJ     = "Logic"
	NULL_OBJ     = "Rỗng"
)

func NewInt(value int64) *Int {
	return &Int{Value: value}
}

type Int struct {
	Value int64
}

func (i *Int) Type() ObjectType { return INT_OBJ }
func (i *Int) Display() string  { return fmt.Sprint(i.Value) }
func (i *Int) ToReal() *Real {
	return &Real{Value: float64(i.Value)}
}
func (i *Int) ToQuotient() *Quotient {
	return NewQuotient(i, NewInt(1))
}
func (i *Int) Add(right Object) Object {
	switch right := right.(type) {
	case *Int:
		return NewInt(i.Value + right.Value)
	case *Real:
		return NewReal(float64(i.Value) + right.Value)
	case *Quotient:
		return right.Add(i)
	default:
		return nil
	}
}
func (i *Int) Subtract(right Object) Object {
	switch right := right.(type) {
	case *Int:
		return NewInt(i.Value - right.Value)
	case *Real:
		return NewReal(float64(i.Value) - right.Value)
	case *Quotient:
		return i.ToQuotient().Subtract(right)
	default:
		return nil
	}
}
func (i *Int) Multiply(right Object) Object {
	switch right := right.(type) {
	case *Int:
		return NewInt(i.Value * right.Value)
	case *Real:
		return NewReal(float64(i.Value) * right.Value)
	case *Quotient:
		return right.Multiply(i)
	default:
		return nil
	}
}
func (i *Int) Divide(right Object) Object {
	switch right := right.(type) {
	case *Int:
		return NewQuotient(i, right)
	case *Real:
		return NewReal(float64(i.Value) / right.Value)
	case *Quotient:
		return i.ToQuotient().Divide(right)
	default:
		return nil
	}
}

type Realness interface {
	ToReal() *Real
}

func NewReal(value float64) *Real {
	return &Real{Value: value}
}

type Real struct {
	Value float64
}

func (r *Real) Type() ObjectType { return REAL_OBJ }
func (r *Real) Display() string  { return fmt.Sprint(r.Value) }
func (r *Real) Add(right Object) Object {
	switch right := right.(type) {
	case *Int:
		return right.Add(r)
	case *Real:
		return NewReal(r.Value + right.Value)
	case *Quotient:
		return right.Add(r)
	default:
		return nil
	}
}
func (r *Real) Subtract(right Object) Object {
	switch right := right.(type) {
	case *Int:
		return NewReal(r.Value - float64(right.Value))
	case *Real:
		return NewReal(r.Value - right.Value)
	case *Quotient:
		return NewReal(r.Value - right.ToReal().Value)
	default:
		return nil
	}
}
func (r *Real) Multiply(right Object) Object {
	switch right := right.(type) {
	case *Int:
		return right.Multiply(r)
	case *Real:
		return NewReal(r.Value * right.Value)
	case *Quotient:
		return right.Multiply(r)
	default:
		return nil
	}
}
func (r *Real) Divide(right Object) Object {
	switch right := right.(type) {
	case *Int:
		return NewReal(r.Value / float64(right.Value))
	case *Real:
		return NewReal(r.Value / right.Value)
	case *Quotient:
		return NewReal(r.Value / right.ToReal().Value)
	default:
		return nil
	}
}

func NewQuotient(numer, denom *Int) *Quotient {
	if denom.Value == 0 {
		return nil
	}
	q := &Quotient{Numer: numer, Denom: denom}

	d := gcd(numer.Value, denom.Value)
	q.Numer.Value = numer.Value / d
	q.Denom.Value = denom.Value / d
	if q.Denom.Value < 0 {
		q.Numer.Value *= -1
		q.Denom.Value *= -1
	}

	return q
}

type Quotient struct {
	Numer *Int
	Denom *Int
}

func (q *Quotient) Type() ObjectType { return QUOTIENT_OBJ }
func (q *Quotient) Display() string {
	if q.Denom.Value == 1 {
		return q.Numer.Display()
	}
	return fmt.Sprintf("(%d/%d)", q.Numer.Value, q.Denom.Value)
}
func (q *Quotient) ToReal() *Real {
	return NewReal(float64(q.Numer.Value) / float64(q.Denom.Value))
}
func (q *Quotient) Inverse() *Quotient {
	return NewQuotient(q.Denom, q.Numer)
}
func (q *Quotient) Add(right Object) Object {
	switch right := right.(type) {
	case *Int:
		return q.Add(right.ToQuotient())
	case *Real:
		return q.ToReal().Add(right)
	case *Quotient:
		numer := NewInt(q.Numer.Value*right.Denom.Value + right.Numer.Value*q.Denom.Value)
		return NewQuotient(numer, q.Denom.Multiply(right.Denom).(*Int))
	default:
		return nil
	}
}
func (q *Quotient) Subtract(right Object) Object {
	switch right := right.(type) {
	case *Int:
		return q.Subtract(right.ToQuotient())
	case *Real:
		return q.ToReal().Subtract(right)
	case *Quotient:
		numer := NewInt(q.Numer.Value*right.Denom.Value - right.Numer.Value*q.Denom.Value)
		return NewQuotient(numer, NewInt(q.Denom.Value*right.Denom.Value))
	default:
		return nil
	}
}
func (q *Quotient) Multiply(right Object) Object {
	switch right := right.(type) {
	case *Int:
		return q.Multiply(right.ToQuotient())
	case *Real:
		return q.ToReal().Multiply(right)
	case *Quotient:
		return NewQuotient(NewInt(q.Numer.Value*right.Numer.Value), NewInt(q.Denom.Value*right.Denom.Value))
	default:
		return nil
	}
}
func (q *Quotient) Divide(right Object) Object {
	switch right := right.(type) {
	case *Int:
		return q.Divide(right.ToQuotient())
	case *Real:
		return q.ToReal().Divide(right)
	case *Quotient:
		return q.Multiply(right.Inverse())
	default:
		return nil
	}
}

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Display() string  { return "rỗng" }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOL_OBJ }
func (b *Boolean) Display() string {
	if b.Value {
		return "đúng"
	}
	return "sai"
}
