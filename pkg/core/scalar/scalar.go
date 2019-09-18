package scalar

import (
	"errors"
)

// Scalar is an interger s < 2^255, and belongs to Z / l
type Scalar struct {
	// Little endian order
	Bytes [32]byte
}

// New returns an instance of Scalar. The high bit of the last byte must be 0, so s < 2^255 is true
func New(b []byte) (*Scalar, error) {
	if len(b) > 32 {
		return nil, errors.New("Scalar size overflow")
	}

	bytes := [32]byte{}
	copy(bytes[:], b)

	if (bytes[31] >> 7) != 0 {
		return nil, errors.New("Scalar size overflow")
	}

	scalar := Scalar{Bytes: bytes}

	return &scalar, nil
}
