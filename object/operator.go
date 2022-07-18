package object

type Additive interface {
	Object
	Add(Object) Object
}

type Subtractive interface {
	Object
	Subtract(Object) Object
}

type Multiplicative interface {
	Object
	Multiply(Object) Object
}

type Division interface {
	Object
	Divide(Object) Object
}
