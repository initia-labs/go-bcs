package bcs_test

import (
	"bytes"
	"testing"

	"github.com/initia-labs/go-bcs/bcs"
	"github.com/stretchr/testify/require"
)

func Test_Address(t *testing.T) {
	addrBz := []byte{1}
	addr := bcs.NewAddressFromBytes(addrBz)
	bz, err := bcs.Marshal(addr)
	require.NoError(t, err)
	require.Equal(t, append(bytes.Repeat([]byte{0}, 31), 1), bz)

	var _addr bcs.Address
	_, err = bcs.Unmarshal(bz, &_addr)
	require.NoError(t, err)
	require.Equal(t, addr, _addr)
}
