package object

import "fmt"

const (
	IntObj          = "Số Nguyên"
	RealObj         = "Số Thực"
	QuotientObj     = "Số Hữu Tỉ"
	BoolObj         = "Logic"
	NullObj         = "Rỗng"
	IncomparableObj = "Không thể so sánh được"
	CantOperateObj  = "Không thể thực hiện phép tính"
)

var (
	NULL         = &Null{}
	TRUE         = &Boolean{Value: true}
	FALSE        = &Boolean{Value: false}
	INCOMPARABLE = &Boolean{Value: false}
	CANT_OPERATE = &CantOperate{}
)

func NewInt(value int64) *Int {
	return &Int{Value: value}
}

type Int struct {
	Value int64
}

func (i *Int) Type() ObjectType { return IntObj }
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
		return CANT_OPERATE
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
		return CANT_OPERATE
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
		return CANT_OPERATE
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
		return CANT_OPERATE
	}
}
func (i *Int) Equal(right Object) *Boolean {
	switch right := right.(type) {
	case *Int:
		return Condition(i.Value == right.Value)
	case *Real:
		return Condition(i.ToReal().Value == right.Value)
	case *Quotient:
		return right.Equal(i)
	default:
		return INCOMPARABLE
	}
}
func (i *Int) Less(right Object) *Boolean {
	switch right := right.(type) {
	case *Int:
		return Condition(i.Value < right.Value)
	case *Real:
		return Condition(i.ToReal().Value < right.Value)
	case *Quotient:
		return Condition(i.ToReal().Value < right.ToReal().Value)
	default:
		return INCOMPARABLE
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

func (r *Real) Type() ObjectType { return RealObj }
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
		return CANT_OPERATE
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
		return CANT_OPERATE
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
		return CANT_OPERATE
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
		return CANT_OPERATE
	}
}
func (r *Real) Equal(right Object) *Boolean {
	switch right := right.(type) {
	case *Int:
		return right.Equal(r)
	case *Real:
		return Condition(r.Value == right.Value)
	case *Quotient:
		return right.Equal(r)
	default:
		return INCOMPARABLE
	}
}
func (r *Real) Less(right Object) *Boolean {
	switch right := right.(type) {
	case *Int:
		return Condition(r.Value < right.ToReal().Value)
	case *Real:
		return Condition(r.Value < right.Value)
	case *Quotient:
		return Condition(r.Value < right.ToReal().Value)
	default:
		return INCOMPARABLE
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

func (q *Quotient) Type() ObjectType { return QuotientObj }
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
		return CANT_OPERATE
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
		return CANT_OPERATE
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
		return CANT_OPERATE
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
		return CANT_OPERATE
	}
}
func (q *Quotient) Equal(right Object) *Boolean {
	switch right := right.(type) {
	case *Int:
		return q.Equal(right.ToQuotient())
	case *Real:
		return Condition(q.ToReal().Value == right.Value)
	case *Quotient:
		return Condition(q.Numer.Value == right.Numer.Value && q.Denom.Value == right.Denom.Value)
	default:
		return INCOMPARABLE
	}
}
func (q *Quotient) Less(right Object) *Boolean {
	switch right := right.(type) {
	case *Int:
		return q.Less(right.ToQuotient())
	case *Real:
		return Condition(q.ToReal().Value < right.Value)
	case *Quotient:
		return Condition(q.ToReal().Value < right.ToReal().Value)
	default:
		return INCOMPARABLE
	}
}

type Null struct{}

func (n *Null) Type() ObjectType { return NullObj }
func (n *Null) Display() string  { return "rỗng" }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BoolObj }
func (b *Boolean) Display() string {
	if b.Value {
		return "đúng"
	}
	return "sai"
}
func (b *Boolean) Not() *Boolean {
	return Condition(!b.Value)
}
func (b *Boolean) And(right *Boolean) *Boolean {
	return Condition(b.Value && right.Value)
}
func (b *Boolean) Or(right *Boolean) *Boolean {
	return Condition(b.Value || right.Value)
}

type CantOperate struct{}

func (*CantOperate) Type() ObjectType { return CantOperateObj }
func (*CantOperate) Display() string  { return CantOperateObj }

func Condition(condition bool) *Boolean {
	if condition {
		return TRUE
	}
	return FALSE
}
