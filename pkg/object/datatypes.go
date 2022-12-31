package object

import "math/big"

type checkTypeFunc func(Object) bool

var DataTypes = map[string]checkTypeFunc{
	"Số nguyên": func(obj Object) bool {
		_, ok := obj.(*Int)
		return ok
	},
	"Hữu tỉ": func(obj Object) bool {
		_, ok := obj.(*Quotient)
		return ok
	},
	"Số thực": func(obj Object) bool {
		_, ok := obj.(*Real)
		return ok
	},
	"Chuỗi": func(obj Object) bool {
		_, ok := obj.(*String)
		return ok
	},
	"Logic": func(obj Object) bool {
		_, ok := obj.(*Boolean)
		return ok
	},
	"Nguyên tố": func(obj Object) bool {
		if obj, ok := obj.(*Int); ok {
			value := obj.Value

			comp := value.Cmp(IntOne)
			if comp == -1 || comp == 0 {
				return false
			}

			two := big.NewInt(2)
			three := big.NewInt(3)
			if value.Cmp(two) == 0 || value.Cmp(three) == 0 {
				return true
			}

			m1 := new(big.Int).Mod(value, two)
			m2 := new(big.Int).Mod(value, three)
			if m1.Cmp(IntZero) == 0 || m2.Cmp(IntZero) == 0 {
				return false
			}

			six := big.NewInt(6)
			for i := big.NewInt(5); ; i.Add(i, six) {
				comp = new(big.Int).Mul(i, i).Cmp(value)
				if comp == 1 {
					break
				}

				comp = new(big.Int).Mod(value, i).Cmp(IntZero)
				comp2 := new(big.Int).Mod(value, new(big.Int).Add(i, two)).Cmp(IntZero)
				if comp == 0 || comp2 == 0 {
					return false
				}
			}
			return true
		}
		return false
	},
	"Chẵn": func(obj Object) bool {
		if obj, ok := obj.(*Int); ok {
			return obj.IsEven()
		}
		return false
	},
	"Lẻ": func(obj Object) bool {
		if obj, ok := obj.(*Int); ok {
			return !obj.IsEven()
		}
		return false
	},
}
