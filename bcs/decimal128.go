package bcs

import (
	"math/big"
)

const DEFAULT_DECIMALS int = 18

// decimals is the number of decimal digits
func NewDecimal128FromUint64(i uint64, decimals int) (*Uint128, error) {
	return NewDecimal128FromBigInt(NewBigIntFromUint64(i), decimals)
}

// decimals is the number of decimal digits
func NewDecimal128FromBigInt(i *big.Int, decimals int) (*Uint128, error) {
	diff := abs(DEFAULT_DECIMALS - decimals)

	pow := big.NewInt(1)
	for p := 0; p < diff; p++ {
		pow.Mul(pow, big.NewInt(10))
	}

	num := big.NewInt(0)
	if decimals > DEFAULT_DECIMALS {
		num.Div(i, pow)
	} else {
		num.Mul(i, pow)
	}

	return NewUint128FromBigInt(num)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
