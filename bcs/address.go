package bcs

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// Address is like `address` or `object` in move.
type Address [32]byte

var (
	_ json.Marshaler   = (*Address)(nil)
	_ json.Unmarshaler = (*Address)(nil)
	_ Marshaler        = (*Address)(nil)
	_ Unmarshaler      = (*Address)(nil)
)

func NewAddressFromBytes(bz []byte) Address {
	r := [32]byte{}
	for i := 0; i+len(bz) < 32; i++ {
		r[i] = 0
	}
	copy(r[32-len(bz):], bz)
	return r
}

// UnmarshalBCS implements Unmarshaler.
func (addr *Address) UnmarshalBCS(r io.Reader) (int, error) {
	n, err := r.Read(addr[:])
	if err != nil {
		return n, err
	}
	if n != 32 {
		return n, fmt.Errorf("failed to read 32 bytes for Address (read %d bytes)", n)
	}

	return n, nil
}

// MarshalBCS implements Marshaler.
func (addr Address) MarshalBCS() ([]byte, error) {
	return addr[:], nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (addr *Address) UnmarshalJSON(bz []byte) error {
	var dataStr string
	if err := json.Unmarshal(bz, &dataStr); err != nil {
		return err
	}

	bz, err := hex.DecodeString(strings.TrimPrefix("0x", dataStr))
	if err != nil {
		return err
	}

	if len(bz) != 32 {
		return fmt.Errorf("failed to read 32 bytes for Address (read %d bytes)", len(bz))
	}

	*addr = [32]byte(bz)
	return nil
}

// MarshalJSON implements json.Marshaler.
func (addr Address) MarshalJSON() ([]byte, error) {
	return json.Marshal("0x" + hex.EncodeToString(addr[:]))
}
