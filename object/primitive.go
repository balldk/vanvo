package object

import "fmt"

type Int struct {
	Value int64
}

func (i *Int) Display() string { return fmt.Sprint(i.Value) }

type Real struct {
	Value float64
}

func (r *Real) Display() string { return fmt.Sprint(r.Value) }

type Null struct{}

func (n *Null) Display() string { return "rỗng" }

type Boolean struct {
	Value bool
}

func (b *Boolean) Display() string {
	if b.Value {
		return "đúng"
	}
	return "sai"
}
