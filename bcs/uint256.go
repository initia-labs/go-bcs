package bcs

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
)

// Uint256 is like `u256` in move.
type Uint256 struct {
	lo   uint64
	hi   uint64
	hilo uint64
	hihi uint64
}

var (
	_ json.Marshaler   = (*Uint256)(nil)
	_ json.Unmarshaler = (*Uint256)(nil)
	_ Marshaler        = (*Uint256)(nil)
	_ Unmarshaler      = (*Uint256)(nil)
)

func (i Uint256) Big() *big.Int {
	loBig := NewBigIntFromUint64(i.lo)
	hiBig := NewBigIntFromUint64(i.hi)
	hiloBig := NewBigIntFromUint64(i.hilo)
	hihiBig := NewBigIntFromUint64(i.hihi)
	hiBig = hiBig.Lsh(hiBig, 64)
	hiloBig = hiloBig.Lsh(hiloBig, 128)
	hihiBig = hihiBig.Lsh(hihiBig, 192)

	return hihiBig.Add(hiBig.Add(hiBig, loBig), hihiBig.Add(hihiBig, hiloBig))
}

func (i Uint256) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Big().String())
}

var maxu256 = (&big.Int{}).Lsh(big.NewInt(1), 256)

func checkUint256(bigI *big.Int) error {
	if bigI.Sign() < 0 {
		return fmt.Errorf("%s is negative", bigI.String())
	}

	if bigI.Cmp(maxu256) >= 0 {
		return fmt.Errorf("%s is greater than Max Uint 256", bigI.String())
	}

	return nil
}

func (i *Uint256) SetBigInt(bigI *big.Int) error {
	if err := checkUint256(bigI); err != nil {
		return err
	}

	r := make([]byte, 0, 32)
	bs := bigI.Bytes()
	for i := 0; i+len(bs) < 32; i++ {
		r = append(r, 0)
	}
	r = append(r, bs...)

	hihi := binary.BigEndian.Uint64(r[0:8])
	hilo := binary.BigEndian.Uint64(r[8:16])
	hi := binary.BigEndian.Uint64(r[16:24])
	lo := binary.BigEndian.Uint64(r[24:32])

	i.hi = hi
	i.lo = lo
	i.hilo = hilo
	i.hihi = hihi

	return nil
}

func (i *Uint256) UnmarshalText(data []byte) error {
	bigI := &big.Int{}
	_, ok := bigI.SetString(string(data), 10)
	if !ok {
		return fmt.Errorf("failed to parse %s as an integer", string(data))
	}

	return i.SetBigInt(bigI)
}

func (i *Uint256) UnmarshalJSON(data []byte) error {
	var dataStr string
	if err := json.Unmarshal(data, &dataStr); err != nil {
		return err
	}

	bigI := &big.Int{}
	_, ok := bigI.SetString(dataStr, 10)
	if !ok {
		return fmt.Errorf("failed to parse %s as an integer", dataStr)
	}

	return i.SetBigInt(bigI)
}

func NewUint256FromBigInt(bigI *big.Int) (*Uint256, error) {
	i := &Uint256{}

	if err := i.SetBigInt(bigI); err != nil {
		return nil, err
	}

	return i, nil
}

func NewUint256(s string) (*Uint256, error) {
	r := &big.Int{}
	r, ok := r.SetString(s, 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse %s as an integer", s)
	}

	return NewUint256FromBigInt(r)
}

func (i Uint256) MarshalBCS() ([]byte, error) {
	r := make([]byte, 32)

	binary.LittleEndian.PutUint64(r, i.lo)
	binary.LittleEndian.PutUint64(r[8:], i.hi)
	binary.LittleEndian.PutUint64(r[16:], i.hilo)
	binary.LittleEndian.PutUint64(r[24:], i.hihi)

	return r, nil
}

func (i *Uint256) UnmarshalBCS(r io.Reader) (int, error) {
	buf := make([]byte, 32)
	n, err := r.Read(buf)
	if err != nil {
		return n, err
	}
	if n != 32 {
		return n, fmt.Errorf("failed to read 32 bytes for Uint256 (read %d bytes)", n)
	}

	i.lo = binary.LittleEndian.Uint64(buf[0:8])
	i.hi = binary.LittleEndian.Uint64(buf[8:16])
	i.hilo = binary.LittleEndian.Uint64(buf[16:24])
	i.hihi = binary.LittleEndian.Uint64(buf[24:32])

	return n, nil
}

func (i *Uint256) Cmp(j *Uint256) int {
	switch {
	case i.hihi > j.hihi ||
		(i.hihi == j.hihi && i.hilo > j.hilo) ||
		(i.hihi == j.hihi && i.hilo == j.hilo && i.hi > j.hi) ||
		(i.hihi == j.hihi && i.hilo == j.hilo && i.hi == j.hi && i.lo > j.lo):
		return 1
	case i.hihi == j.hihi && i.hilo == j.hilo && i.hi == j.hi && i.lo == j.lo:
		return 0
	default:
		return -1
	}
}

func NewUint256FromUint64(lo, hi, hilo, hihi uint64) *Uint256 {
	return &Uint256{
		lo:   lo,
		hi:   hi,
		hilo: hilo,
		hihi: hihi,
	}
}

func (u Uint256) String() string {
	return u.Big().String()
}
