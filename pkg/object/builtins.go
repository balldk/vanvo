package object

import (
	"math"
)

var Builtins = map[string]Object{
	"Pi": &Real{Value: math.Pi},
	"Tổng": &Function{
		Builtin: func(args ...Object) Object {
			var s Additive
			s = &Int{Value: 0}

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
