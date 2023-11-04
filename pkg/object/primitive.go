package object

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ALTree/bigfloat"
)

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
	NULL          = &Null{}
	TRUE          = &Boolean{Value: true}
	FALSE         = &Boolean{Value: false}
	INCOMPARABLE  = &Boolean{Value: false}
	ZERO_DIVISION = &Null{}
	CANT_OPERATE  = &CantOperate{}

	IntZero  = big.NewInt(0)
	IntOne   = big.NewInt(1)
	RealZero = big.NewFloat(0)
	RealOne  = big.NewFloat(1)
	Epsilon  = big.NewFloat(0.0000001)
)

type Number interface {
	Additive
	Subtractive
	Multiplicative
	Division
}

func NewInt(value *big.Int) *Int {
	return &Int{Value: value}
}

type Int struct {
	Value *big.Int
}

func (i *Int) Type() ObjectType { return IntObj }
func (i *Int) Display() string  { return fmt.Sprint(i.Value) }
func (i *Int) ToReal() *Real {
	return &Real{Value: new(big.Float).SetInt(i.Value)}
}
func (i *Int) ToComplex() *Complex {
	return NewComplex(i, NewInt(IntZero))
}
func (i *Int) ToQuotient() *Quotient {
	return NewQuotient(i.Value, IntOne)
}
func (i *Int) Add(right Object) Object {
	switch right := right.(type) {
	case *Int:
		return NewInt(new(big.Int).Add(i.Value, right.Value))
	case *Real:
		intVal := new(big.Float).SetInt(i.Value)
		return NewReal(new(big.Float).Add(intVal, right.Value))
	case *Quotient:
		return right.Add(i)
	case *Complex:
		return right.Add(i)
	default:
		return CANT_OPERATE
	}
}
func (i *Int) Subtract(right Object) Object {
	switch right := right.(type) {
	case *Int:
		return NewInt(new(big.Int).Sub(i.Value, right.Value))
	case *Real:
		intVal := new(big.Float).SetInt(i.Value)
		return NewReal(new(big.Float).Sub(intVal, right.Value))
	case *Quotient:
		return i.ToQuotient().Subtract(right)
	case *Complex:
		return i.ToComplex().Subtract(right)
	default:
		return CANT_OPERATE
	}
}
func (i *Int) Multiply(right Object) Object {
	switch right := right.(type) {
	case *Int:
		return NewInt(new(big.Int).Mul(i.Value, right.Value))
	case *Real:
		intVal := new(big.Float).SetInt(i.Value)
		return NewReal(new(big.Float).Mul(intVal, right.Value))
	case *Quotient:
		return right.Multiply(i)
	case *Complex:
		return right.Multiply(i)
	case *String:
		return right.Multiply(i)
	default:
		return CANT_OPERATE
	}
}
func (i *Int) Divide(right Object) Object {
	switch right := right.(type) {
	case *Int:
		if right.Value.Int64() == 0 {
			return ZERO_DIVISION
		}
		return NewQuotient(i.Value, right.Value)
	case *Real:
		rightVal, _ := right.Value.Float64()
		if rightVal == 0 {
			return ZERO_DIVISION
		}
		intVal := new(big.Float).SetInt(i.Value)
		return NewReal(new(big.Float).Quo(intVal, right.Value))
	case *Quotient:
		return i.ToQuotient().Divide(right)
	case *Complex:
		return i.ToComplex().Divide(right)
	default:
		return CANT_OPERATE
	}
}
func (i *Int) Mod(right Object) Object {
	switch right := right.(type) {
	case *Int:
		if right.Value.Int64() == 0 {
			return ZERO_DIVISION
		}
		return NewInt(new(big.Int).Mod(i.Value, right.Value))
	default:
		return CANT_OPERATE
	}
}
func (i *Int) Power(right Object) Object {
	switch right := right.(type) {
	case *Int:
		return NewInt(new(big.Int).Exp(i.Value, right.Value, nil))
	case *Real:
		intVal := new(big.Float).SetInt(i.Value)
		return NewReal(bigfloat.Pow(intVal, right.Value))
	case *Quotient:
		return NewReal(i.ToReal().Power(right.ToReal()).(*Real).Value)
	default:
		return CANT_OPERATE
	}
}
func (i *Int) Equal(right Object) *Boolean {
	switch right := right.(type) {
	case *Int:
		return Condition(i.Value.Cmp(right.Value) == 0)
	case *Real:
		return Condition(i.ToReal().Value.Cmp(right.Value) == 0)
	case *Quotient:
		return Condition(i.ToReal().Value.Cmp(right.ToReal().Value) == 0)
	default:
		return INCOMPARABLE
	}
}
func (i *Int) Less(right Object) *Boolean {
	switch right := right.(type) {
	case *Int:
		return Condition(i.Value.Cmp(right.Value) == -1)
	case *Real:
		return Condition(i.ToReal().Value.Cmp(right.Value) == -1)
	case *Quotient:
		return Condition(i.ToReal().Value.Cmp(right.ToReal().Value) == -1)
	default:
		return INCOMPARABLE
	}
}

type Realness interface {
	Number
	Order
	ToReal() *Real
}

func NewReal(value *big.Float) *Real {
	return &Real{Value: value}
}

type Real struct {
	Value *big.Float
}

func (r *Real) Type() ObjectType { return RealObj }
func (r *Real) Display() string  { return fmt.Sprint(r.Value) }
func (r *Real) ToReal() *Real {
	return r
}
func (r *Real) ToComplex() *Complex {
	return NewComplex(r, NewInt(IntZero))
}
func (r *Real) IsZero() bool {
	if val, _ := r.Value.Float32(); val == 0 {
		return true
	}
	return false
}
func (r *Real) Add(right Object) Object {
	switch right := right.(type) {
	case *Int:
		return right.Add(r)
	case *Real:
		return NewReal(new(big.Float).Add(r.Value, right.Value))
	case *Quotient:
		return right.Add(r)
	case *Complex:
		return right.Add(r)
	default:
		return CANT_OPERATE
	}
}
func (r *Real) Subtract(right Object) Object {
	switch right := right.(type) {
	case *Int:
		intVal := new(big.Float).SetInt(right.Value)
		return NewReal(new(big.Float).Sub(r.Value, intVal))
	case *Real:
		return NewReal(new(big.Float).Sub(r.Value, right.Value))
	case *Quotient:
		return NewReal(new(big.Float).Sub(r.Value, right.ToReal().Value))
	case *Complex:
		return r.ToComplex().Subtract(right)
	default:
		return CANT_OPERATE
	}
}
func (r *Real) Multiply(right Object) Object {
	switch right := right.(type) {
	case *Int:
		return right.Multiply(r)
	case *Real:
		return NewReal(new(big.Float).Mul(r.Value, right.Value))
	case *Quotient:
		return right.Multiply(r)
	case *Complex:
		return right.Multiply(r)
	default:
		return CANT_OPERATE
	}
}
func (r *Real) Divide(right Object) Object {
	switch right := right.(type) {
	case *Int:
		if right.Value.Cmp(IntZero) == 0 {
			return ZERO_DIVISION
		}
		return NewReal(new(big.Float).Quo(r.Value, right.ToReal().Value))
	case *Real:
		rightVal, _ := right.Value.Float64()
		if rightVal == 0 {
			return ZERO_DIVISION
		}
		return NewReal(new(big.Float).Quo(r.Value, right.Value))
	case *Quotient:
		if right.Value.Num().Cmp(IntZero) == 0 {
			return ZERO_DIVISION
		}
		return NewReal(new(big.Float).Quo(r.Value, right.ToReal().Value))
	case *Complex:
		return r.ToComplex().Divide(right)
	default:
		return CANT_OPERATE
	}
}
func (r *Real) Power(right Object) Object {
	switch right := right.(type) {
	case *Int:
		intVal := new(big.Float).SetInt(right.Value)
		return NewReal(bigfloat.Pow(r.Value, intVal))
	case *Real:
		return NewReal(bigfloat.Pow(r.Value, right.Value))
	case *Quotient:
		return NewReal(bigfloat.Pow(r.Value, right.ToReal().Value))
	default:
		return CANT_OPERATE
	}
}
func (r *Real) Sqrt() Number {
	val := r.Value
	if val.Cmp(RealZero) == -1 {
		val = new(big.Float).Abs(val)
		return NewComplex(NewInt(IntZero), NewReal(val))
	}
	return NewReal(new(big.Float).Sqrt(val))
}
func (r *Real) Cos() Number {
	real := r.ToReal().Value
	toSmallAngle(real)

	val, _ := real.Float64()
	val = math.Cos(val)
	if math.Abs(val) < 1e-15 {
		return NewReal(big.NewFloat(0))
	}

	bigVal := big.NewFloat(val)
	return NewReal(bigVal)
}
func (r *Real) Sin() Number {
	real := r.ToReal().Value
	toSmallAngle(real)

	val, _ := real.Float64()
	val = math.Sin(val)
	if math.Abs(val) < 1e-15 {
		return NewReal(big.NewFloat(0))
	}

	bigVal := big.NewFloat(val)
	return NewReal(bigVal)
}
func (r *Real) Tan() Number {
	real := r.ToReal().Value
	toSmallAngle(real)

	val, _ := real.Float64()
	val = math.Tan(val)
	if math.Abs(val) < 1e-15 {
		return NewReal(big.NewFloat(0))
	}

	bigVal := big.NewFloat(val)
	return NewReal(bigVal)
}
func (r *Real) NaturalLog() Number {
	return NewReal(bigfloat.Log(r.ToReal().Value))
}
func (r *Real) Round() *Real {
	i := new(big.Int)
	i, _ = new(big.Float).Add(r.Value, new(big.Float).Mul(big.NewFloat(0.5), big.NewFloat(float64(r.Value.Sign())))).Int(i)
	f := new(big.Float).SetInt(i)

	dif := new(big.Float).Sub(r.Value, f)
	dif = dif.Abs(dif)
	if dif.Cmp(Epsilon) == -1 {
		r.Value = f
	}
	return r
}
func (r *Real) Equal(right Object) *Boolean {
	switch right := right.(type) {
	case *Int:
		return right.Equal(r)
	case *Real:
		return Condition(r.Value.Cmp(right.Value) == 0)
	case *Quotient:
		return right.Equal(r)
	default:
		return INCOMPARABLE
	}
}
func (r *Real) Less(right Object) *Boolean {
	switch right := right.(type) {
	case *Int:
		return Condition(r.Value.Cmp(right.ToReal().Value) == -1)
	case *Real:
		return Condition(r.Value.Cmp(right.Value) == -1)
	case *Quotient:
		return Condition(r.Value.Cmp(right.ToReal().Value) == -1)
	default:
		return INCOMPARABLE
	}
}

func NewQuotient(numer, denom *big.Int) *Quotient {
	if denom.Cmp(IntZero) == 0 {
		return nil
	}
	q := &Quotient{Value: new(big.Rat).SetFrac(numer, denom)}
	return q
}

type Quotient struct {
	Value *big.Rat
}

func (q *Quotient) Type() ObjectType { return QuotientObj }
func (q *Quotient) Display() string {
	if q.Value.Denom().Cmp(IntOne) == 0 {
		return q.Value.Num().String()
	}
	return q.Value.String()
}
func (q *Quotient) ToReal() *Real {
	val, _ := q.Value.Float64()
	return NewReal(big.NewFloat(val))
}
func (q *Quotient) ToComplex() *Complex {
	return NewComplex(q, NewInt(IntZero))
}
func (q *Quotient) Inverse() *Quotient {
	return &Quotient{Value: new(big.Rat).Inv(q.Value)}
}
func (q *Quotient) Add(right Object) Object {
	switch right := right.(type) {
	case *Int:
		return q.Add(right.ToQuotient())
	case *Real:
		return q.ToReal().Add(right)
	case *Quotient:
		return &Quotient{Value: new(big.Rat).Add(q.Value, right.Value)}
	case *Complex:
		return right.Add(q)
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
		return &Quotient{Value: new(big.Rat).Sub(q.Value, right.Value)}
	case *Complex:
		return q.ToComplex().Subtract(right)
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
		return &Quotient{Value: new(big.Rat).Mul(q.Value, right.Value)}
	case *Complex:
		return right.Multiply(q)
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
		if right.Value.Denom().Cmp(IntZero) == 0 {
			return ZERO_DIVISION
		}
		return &Quotient{Value: new(big.Rat).Quo(q.Value, right.Value)}
	case *Complex:
		return q.ToComplex().Divide(right)
	default:
		return CANT_OPERATE
	}
}
func (q *Quotient) Power(right Object) Object {
	switch right := right.(type) {
	case *Int:
		numer := new(big.Int).Exp(q.Value.Num(), right.Value, nil)
		denom := new(big.Int).Exp(q.Value.Denom(), right.Value, nil)
		return NewQuotient(numer, denom)
	case *Real:
		return q.ToReal().Power(right)
	case *Quotient:
		return q.ToReal().Power(right.ToReal())
	default:
		return CANT_OPERATE
	}
}
func (q *Quotient) Equal(right Object) *Boolean {
	switch right := right.(type) {
	case *Int:
		return q.Equal(right.ToQuotient())
	case *Real:
		return q.ToReal().Equal(right)
	case *Quotient:
		numerEqual := q.Value.Num().Cmp(right.Value.Num()) == 0
		denomEqual := q.Value.Denom().Cmp(right.Value.Denom()) == 0
		return Condition(numerEqual && denomEqual)
	default:
		return INCOMPARABLE
	}
}
func (q *Quotient) Less(right Object) *Boolean {
	switch right := right.(type) {
	case *Int:
		return q.Less(right.ToQuotient())
	case *Real:
		return q.ToReal().Less(right)
	case *Quotient:
		return q.ToReal().Less(right.ToReal())
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
func (b *Boolean) Equal(right Object) *Boolean {
	switch right := right.(type) {
	case *Boolean:
		return Condition(b.Value == right.Value)
	case *Int:
		isFalse := b == FALSE && right.Value.Cmp(IntZero) == 0
		isTrue := b == TRUE && right.Value.Cmp(IntOne) == 0
		return Condition(isFalse || isTrue)
	default:
		return FALSE
	}
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
