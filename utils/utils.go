package utils

import (
	"fmt"
	"math/big"

	"github.com/ALTree/bigfloat"
)

var ten = big.NewFloat(10)

func ScaleInt(amount *big.Int, decimals int64) *big.Float {
	return ScaleFloat(new(big.Float).SetInt(amount), decimals)
}

func ScaleFloat(amount *big.Float, decimals int64) *big.Float {
	dec := big.NewFloat(float64(decimals))
	tenPowDec := bigfloat.Pow(ten, dec)

	return new(big.Float).Quo(amount, tenPowDec)
}

func FormatBigFloat(x *big.Float, opts ...int64) string {
	var decimals int64 = 2
	if len(opts) > 0 {
		decimals = opts[0]
	}

	format := fmt.Sprintf("%%0.%df", decimals)

	return fmt.Sprintf(format, x)
}

func BigIntToBigFloat(x *big.Int) *big.Float {
	return new(big.Float).SetInt(x)
}
