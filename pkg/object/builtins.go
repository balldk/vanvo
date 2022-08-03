package object

import (
	"math"
	"math/big"
)

var Builtins = map[string]Object{
	"Pi": &Real{Value: big.NewFloat(math.Pi)},
	"Tổng": &Function{
		Builtin: func(args ...Object) Object {
			var s Additive
			s = &Int{Value: IntZero}

			for _, arg := range args {
				if arg, ok := arg.(Additive); ok {
					s = arg.Add(s).(Additive)
				} else {
					return NULL
				}
			}
			return s
		},
	},
}
