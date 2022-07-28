package object

const (
	IMPLY_OBJ = "Giá trị trả về"
)

type Imply struct {
	Value Object
}

func (r *Imply) Type() ObjectType { return IMPLY_OBJ }
func (r *Imply) Display() string  { return r.Value.Display() }
