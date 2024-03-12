package bcs_test

import (
	"math/big"
	"testing"

	"github.com/fardream/go-bcs/bcs"
	"github.com/stretchr/testify/require"
)

func Test_Decimal128(t *testing.T) {
	// small decimals
	dec, err := bcs.NewDecimal128FromUint64(105, 2)
	require.NoError(t, err)

	expectedDec := big.NewInt(0).Mul(
		big.NewInt(105),
		big.NewInt(10_000_000_000_000_000), // 18 - 2
	)
	require.Equal(t, expectedDec, dec.Big())

	// bigger decimals
	dec, err = bcs.NewDecimal128FromUint64(105, 20)
	require.NoError(t, err)

	expectedDec = big.NewInt(0).Div(
		big.NewInt(105),
		big.NewInt(100), // 20 - 18
	)
	require.Equal(t, expectedDec, dec.Big())
}
