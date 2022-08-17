package object

import (
	"math"
	"math/big"
)

func toSmallAngle(ang *big.Float) {
	if ang.Cmp(big.NewFloat(math.MaxFloat64)) == -1 {
		return
	}

	kFloat := new(big.Float)
	kFloat.Quo(kFloat, big.NewFloat(2*math.Pi))

	k := new(big.Int)
	kFloat.Int(k)
	k.Add(k, IntOne)
	kFloat.SetInt(k)

	kFloat.Mul(kFloat, big.NewFloat(2*math.Pi))
	ang.Sub(ang, big.NewFloat(2*math.Pi))
}

func setHasLength(set CountableSet, length int) bool {
	for i := 0; i < length; i++ {
		if set.At(i) == IndexError {
			return false
		}
	}
	return set.At(length) != IndexError
}
