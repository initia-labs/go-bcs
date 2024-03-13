package bcs_test

import (
	"math/big"
	"testing"

	"github.com/initia-labs/go-bcs/bcs"
)

func TestNewUint256FromUint64(t *testing.T) {
	expected := &big.Int{}
	expected.Add(
		big.NewInt(0).Add(
			big.NewInt(0).Lsh(big.NewInt(96), 64),
			big.NewInt(50),
		),
		big.NewInt(0).Add(
			big.NewInt(0).Lsh(big.NewInt(24), 128),
			big.NewInt(0).Lsh(big.NewInt(19), 192),
		),
	)
	result := bcs.NewUint256FromUint64(50, 96, 24, 19).Big()
	if result.Cmp(expected) != 0 {
		t.Fatalf("want: %s, got: %s", expected.String(), result.String())
	}
}

func TestNewUint256(t *testing.T) {
	s := "119264932972346934521046775847048185030836631454255209775154"
	expected := &big.Int{}
	expected.Add(
		big.NewInt(0).Add(
			big.NewInt(0).Lsh(big.NewInt(96), 64),
			big.NewInt(50),
		),
		big.NewInt(0).Add(
			big.NewInt(0).Lsh(big.NewInt(24), 128),
			big.NewInt(0).Lsh(big.NewInt(19), 192),
		),
	)

	result, err := bcs.NewUint256(s)
	if err != nil {
		t.Fatal(err)
	}

	if result.Big().Cmp(expected) != 0 {
		t.Fatalf("want: %s, got: %s", expected.String(), result.String())
	}
}
