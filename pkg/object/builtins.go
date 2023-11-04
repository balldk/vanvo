package object

import (
	"fmt"
	"math"
	"math/big"
)

const (
	ErrorObj = "Lỗi"
)

func NewError(message string) *Error {
	return &Error{Message: message}
}

type Error struct {
	Message string
}

func (*Error) Type() ObjectType { return ErrorObj }
func (*Error) Display() string  { return ErrorObj }

func NewArgumentError(expected int, args []Object) *ArgumentError {
	return &ArgumentError{Expected: expected, Received: len(args)}
}

type ArgumentError struct {
	Expected int
	Received int
}

func (*ArgumentError) Type() ObjectType { return ErrorObj }
func (*ArgumentError) Display() string  { return ErrorObj }

var Builtins = map[string]Object{
	"Pi": &Real{Value: big.NewFloat(math.Pi)},
	"E":  &Real{Value: big.NewFloat(math.E)},
	"I":  &Complex{Real: NewInt(IntZero), Imagine: NewInt(IntOne)},
	"len": &Function{
		Builtin: func(args ...Object) Object {
			if len(args) != 1 {
				return NewArgumentError(1, args)
			}
			if arg, ok := args[0].(CountableSet); ok {
				val := big.NewInt(int64(arg.Length()))
				return &Int{Value: val}

			} else {
				errMsg := fmt.Sprintf("Không thể dùng '%s' làm tham số", args[0].Type())
				return NewError(errMsg)
			}
		},
	},
	"căn": &Function{
		Builtin: SquareRootBuiltin,
	},
	"sin": &Function{
		Builtin: sinBuiltin,
	},
	"cos": &Function{
		Builtin: cosBuiltin,
	},
	"tan": &Function{
		Builtin: tanBuiltin,
	},
	"ln": &Function{
		Builtin: naturalLogBuiltin,
	},
}

func SquareRootBuiltin(args ...Object) Object {
	if len(args) != 1 {
		return NewArgumentError(1, args)
	}
	if arg, ok := args[0].(Realness); ok {
		return arg.ToReal().Sqrt()

	} else {
		errMsg := fmt.Sprintf("Không thể dùng '%s' làm tham số", args[0].Type())
		return NewError(errMsg)
	}
}

func sinBuiltin(args ...Object) Object {
	if len(args) != 1 {
		return NewArgumentError(1, args)
	}
	if arg, ok := args[0].(Realness); ok {
		return arg.ToReal().Sin()

	} else {
		errMsg := fmt.Sprintf("Không thể dùng '%s' làm tham số", args[0].Type())
		return NewError(errMsg)
	}
}

func cosBuiltin(args ...Object) Object {
	if len(args) != 1 {
		return NewArgumentError(1, args)
	}
	if arg, ok := args[0].(Realness); ok {
		return arg.ToReal().Cos()

	} else {
		errMsg := fmt.Sprintf("Không thể dùng '%s' làm tham số", args[0].Type())
		return NewError(errMsg)
	}
}

func tanBuiltin(args ...Object) Object {
	if len(args) != 1 {
		return NewArgumentError(1, args)
	}
	if arg, ok := args[0].(Realness); ok {
		return arg.ToReal().Tan()

	} else {
		errMsg := fmt.Sprintf("Không thể dùng '%s' làm tham số", args[0].Type())
		return NewError(errMsg)
	}
}

func naturalLogBuiltin(args ...Object) Object {
	if len(args) != 1 {
		return NewArgumentError(1, args)
	}
	if arg, ok := args[0].(Realness); ok {
		return arg.ToReal().NaturalLog()

	} else {
		errMsg := fmt.Sprintf("Không thể dùng '%s' làm tham số", args[0].Type())
		return NewError(errMsg)
	}
}
