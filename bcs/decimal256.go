package bcs

import (
	"math/big"
)

// decimals is the number of decimal digits
func NewDecimal256FromUint64(i uint64, decimals int) (*Uint256, error) {
	return NewDecimal256FromBigInt(NewBigIntFromUint64(i), decimals)
}

// decimals is the number of decimal digits
func NewDecimal256FromBigInt(i *big.Int, decimals int) (*Uint256, error) {
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

	return NewUint256FromBigInt(num)
}
