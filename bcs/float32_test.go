package bcs_test

import (
	"math/big"
	"testing"

	"github.com/initia-labs/go-bcs/bcs"
	"github.com/stretchr/testify/require"
)

func Test_Float32(t *testing.T) {
	float, err := bcs.NewFloat32FromUint64(105, 2)
	require.NoError(t, err)

	num := big.NewInt(0)
	num.Div(num.Mul(big.NewInt(105), big.NewInt(1<<32)), big.NewInt(100))
	require.Equal(t, num.Uint64(), *float)
}
