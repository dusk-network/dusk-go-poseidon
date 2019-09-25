package poseidon

import (
	"testing"

	ristretto "github.com/bwesterb/go-ristretto"
	"github.com/stretchr/testify/assert"
)

func TestPoseidonHash(t *testing.T) {
	p := DefaultPoseidon()

	p.Write([]byte("hello"))
	p.Write([]byte("world"))

	digest := []byte{0xd9, 0x2a, 0x01, 0x93, 0x79, 0xb8, 0xa2, 0xdf, 0xf3, 0xb3, 0x7d, 0x4b, 0x3b, 0x59, 0xe6, 0x88, 0x38, 0x89, 0x12, 0xc0, 0x6f, 0xfd, 0x31, 0x69, 0x3e, 0x0d, 0xad, 0xcb, 0xf3, 0x59, 0x55, 0x06}

	assert.Equal(t, digest, p.Sum(nil))
}

func TestQuinticSbox(t *testing.T) {
	a := ristretto.Scalar{0xe2d76bf9, 0xbb6e333c, 0x2ec4e479, 0xba272f09, 0x046d4aca, 0x6aadbd72, 0x95c9842a, 0x0f0cdba9}
	b := ristretto.Scalar{0xe41683eb, 0xc0f550ab, 0x7f547e18, 0x935175b2, 0xf72488bf, 0x03384905, 0x30658415, 0x0cf11c8e}

	QuinticSbox(&a)
	assert.Equal(t, b, a)
}
