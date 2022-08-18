package object

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ALTree/bigfloat"
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
		Builtin: squareRootBuiltin,
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

func squareRootBuiltin(args ...Object) Object {
	if len(args) != 1 {
		return NewArgumentError(1, args)
	}
	if arg, ok := args[0].(Realness); ok {
		real := arg.ToReal().Value
		if real.Cmp(RealZero) == -1 {
			real = new(big.Float).Abs(real)
			return NewComplex(NewInt(IntZero), NewReal(real))
		}
		return NewReal(new(big.Float).Sqrt(real))

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
		real := arg.ToReal().Value
		toSmallAngle(real)

		val, _ := real.Float64()
		val = math.Sin(val)
		if math.Abs(val) < 1e-15 {
			return NewReal(big.NewFloat(0))
		}

		bigVal := big.NewFloat(val)
		return NewReal(bigVal)

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
		real := arg.ToReal().Value
		toSmallAngle(real)

		val, _ := real.Float64()
		val = math.Cos(val)
		if math.Abs(val) < 1e-15 {
			return NewReal(big.NewFloat(0))
		}

		bigVal := big.NewFloat(val)
		return NewReal(bigVal)

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
		real := arg.ToReal().Value
		toSmallAngle(real)

		val, _ := real.Float64()
		val = math.Tan(val)
		if math.Abs(val) < 1e-15 {
			return NewReal(big.NewFloat(0))
		}

		bigVal := big.NewFloat(val)
		return NewReal(bigVal)

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
		real := arg.ToReal().Value
		return NewReal(bigfloat.Log(real))

	} else {
		errMsg := fmt.Sprintf("Không thể dùng '%s' làm tham số", args[0].Type())
		return NewError(errMsg)
	}
}
