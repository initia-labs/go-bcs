package bcs

import (
	"math/big"
)

// decimals is the number of decimal digits
func NewFloat64FromUint64(i uint64, decimals int) (*Uint128, error) {
	return NewFloat64FromBigInt(NewBigIntFromUint64(i), decimals)
}

// decimals is the number of decimal digits
func NewFloat64FromBigInt(i *big.Int, decimals int) (*Uint128, error) {
	pow := big.NewInt(1)
	for p := 0; p < decimals; p++ {
		pow.Mul(pow, big.NewInt(10))
	}

	denominator, _ := new(big.Int).SetString("18446744073709551616", 10)
	i = i.Mul(i, denominator)
	i = i.Div(i, pow)

	return NewUint128FromBigInt(i)
}
