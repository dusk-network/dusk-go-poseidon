package poseidon

import (
	"github.com/dusk-network/dusk-go-poseidon/pkg/core/scalar"
)

// Params represents the parameters for the hash calculation
type Params struct {
	Width               uint
	FullRoundsBeginning uint
	PartialRounds       uint
	FullRoundsEnd       uint
	RoundKeys           []scalar.Scalar
	MDSMatrix           [][]scalar.Scalar
}

// DefaultParams will generate a default parametrization for Poseidon
func DefaultParams() Params {
	return Params{
		Width:               9,
		FullRoundsBeginning: 4,
		PartialRounds:       59,
		FullRoundsEnd:       4,
		RoundKeys:           []scalar.Scalar{},
		MDSMatrix:           [][]scalar.Scalar{},
	}
}
