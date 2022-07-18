package object

type Int struct {
	Value int64
}

func (i *Int) Display() Any { return i.Value }

type Real struct {
	Value float64
}

func (r *Real) Display() Any { return r.Value }

type Null struct{}

func (n *Null) Display() Any { return "rỗng" }

type Boolean struct {
	Value bool
}

func (b *Boolean) Display() Any {
	if b.Value {
		return "đúng"
	}
	return "sai"
}
