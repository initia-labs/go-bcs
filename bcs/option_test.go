package bcs_test

import (
	"testing"

	"github.com/initia-labs/go-bcs/bcs"
	"github.com/stretchr/testify/require"
)

func Test_Option(t *testing.T) {
	// some
	bz, err := bcs.Marshal(bcs.Some(true))
	require.NoError(t, err)
	require.Equal(t, []byte{1, 1}, bz)

	var opt bcs.Option[bool]
	_, err = bcs.Unmarshal(bz, &opt)
	require.NoError(t, err)
	require.True(t, opt.IsSome())
	require.False(t, opt.IsNone())

	// none
	bz, err = bcs.Marshal(bcs.None[any]())
	require.NoError(t, err)
	require.Equal(t, []byte{0}, bz)

	_, err = bcs.Unmarshal(bz, &opt)
	require.NoError(t, err)
	require.False(t, opt.IsSome())
	require.True(t, opt.IsNone())
}
