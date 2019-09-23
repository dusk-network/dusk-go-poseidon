package poseidon

import (
	"errors"

	ristretto "github.com/bwesterb/go-ristretto"
)

// Poseidon represents the structure for the main hash unit
type Poseidon struct {
	params *Params
	input  []ristretto.Scalar
}

// New returns an instance of Poseidon
func New(params *Params) Poseidon {
	return Poseidon{params: params, input: []ristretto.Scalar{}}
}

// DefaultPoseidon will generate a Poseidon instance with DefaultParams
func DefaultPoseidon() Poseidon {
	params := DefaultParams()
	return New(&params)
}

// Size will return the maximum allowed length for input elements
func (p *Poseidon) Size() int {
	return p.params.Width
}

// BlockSize ..
func (p *Poseidon) BlockSize() int {
	// TODO not implemented
	return 0
}

// Write will try to append the provided bytes to the input, converted as Scalar
func (p *Poseidon) Write(b []byte) (int, error) {
	if p.params.Width == len(p.input) {
		return 0, errors.New("Maximum width reached")
	}

	s := ristretto.Scalar{}
	s.Derive(b)
	p.input = append(p.input, s)

	return len(b), nil
}

// Sum will compute the Poseidon digest value. The usage of the bytes parameter is currently not implemented
func (p *Poseidon) Sum(in []byte) []byte {
	p.Pad()

	keysOffset := 0
	result := append([]ristretto.Scalar{}, p.input...)

	for i := 0; i < p.params.FullRoundsBeginning; i++ {
		p.applyFullRound(&p.params.RoundKeys[keysOffset])
		keysOffset++
	}

	for i := 0; i < p.params.PartialRounds; i++ {
		p.applyPartialRound(&p.params.RoundKeys[keysOffset])
		keysOffset++
	}

	for i := 0; i < p.params.FullRoundsEnd; i++ {
		p.applyFullRound(&p.params.RoundKeys[keysOffset])
		keysOffset++
	}

	return result[1].Bytes()
}

func (p *Poseidon) applyFullRound(roundKey *ristretto.Scalar) {
	// Apply quintic SBox
	for i := range p.input {
		for k := 0; k < 5; k++ {
			p.input[i] = *p.input[i].Mul(&p.input[i], &p.input[i])
		}
	}

	p.input = *mulVec(&p.params.MDSMatrix, &p.input)
}

func (p *Poseidon) applyPartialRound(roundKey *ristretto.Scalar) {
	// Apply quintic SBox
	for k := 0; k < 5; k++ {
		p.input[0] = *p.input[0].Mul(&p.input[0], &p.input[0])
	}

	p.input = *mulVec(&p.params.MDSMatrix, &p.input)
}

func mulVec(a *[][]ristretto.Scalar, b *[]ristretto.Scalar) *[]ristretto.Scalar {
	result := make([]ristretto.Scalar, len(*b))

	for j, row := range *a {
		line := make([]ristretto.Scalar, len(*b))
		for k, cell := range row {
			line[k].Mul(&cell, &(*b)[k])
		}

		for _, cell := range line {
			result[j].Add(&result[j], &cell)
		}
	}

	return &result
}

// Pad will fill the input with zeroed scalars until its length equal the parametrization width
func (p *Poseidon) Pad() {
	dif := p.params.Width - len(p.input)
	if dif > 0 {
		pad := make([]ristretto.Scalar, dif)
		p.input = append(p.input, pad...)
	}
}
