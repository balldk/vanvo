package object

type Set interface {
	Object
	contain(Object) bool
}
