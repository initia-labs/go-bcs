package bcs

import (
	"errors"
	"math/big"
)

// Floats is the number of Float digits
func NewFloat32FromUint64(i uint64, decimals int) (*uint64, error) {
	return NewFloat32FromBigInt(NewBigIntFromUint64(i), decimals)
}

// Floats is the number of Float digits
func NewFloat32FromBigInt(i *big.Int, decimals int) (*uint64, error) {
	pow := big.NewInt(1)
	for p := 0; p < decimals; p++ {
		pow.Mul(pow, big.NewInt(10))
	}

	i = i.Mul(i, big.NewInt(4_294_967_296))
	i = i.Div(i, pow)

	if !i.IsUint64() {
		return nil, errors.New("failed to create float32 from big.Int")
	}

	num := i.Uint64()
	return &num, nil
}
