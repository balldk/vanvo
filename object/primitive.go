package object

import "fmt"

const (
	INT_OBJ  = "Số Nguyên"
	REAL_OBJ = "Số Thực"
	BOOL_OBJ = "Logic"
	NULL_OBJ = "Rỗng"
)

type Int struct {
	Value int64
}

func (i *Int) Type() ObjectType { return INT_OBJ }
func (i *Int) Display() string  { return fmt.Sprint(i.Value) }
func (i *Int) Add(right Object) Object {
	switch right := right.(type) {
	case *Int:
		return &Int{Value: i.Value + right.Value}
	case *Real:
		return &Real{Value: float64(i.Value) + right.Value}
	default:
		return nil
	}
}

type Real struct {
	Value float64
}

func (r *Real) Type() ObjectType { return REAL_OBJ }
func (r *Real) Display() string  { return fmt.Sprint(r.Value) }
func (r *Real) Add(right Object) Object {
	switch right := right.(type) {
	case *Int:
		return &Real{Value: r.Value + float64(right.Value)}
	case *Real:
		return &Real{Value: r.Value + right.Value}
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
