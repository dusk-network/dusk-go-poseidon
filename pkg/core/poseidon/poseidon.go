package poseidon

// Poseidon represents the structure for the main hash unit
type Poseidon struct {
	params *Params
}

// New returns an instance of Poseidon
func New(params *Params) Poseidon {
	return Poseidon{params: params}
}

// DefaultPoseidon will generate a Poseidon instance with DefaultParams
func DefaultPoseidon() Poseidon {
	params := DefaultParams()
	return New(&params)
}

// Size ...
func (p *Poseidon) Size() int {
	// TODO not implemented
	return 0
}

// BlockSize ..
func (p *Poseidon) BlockSize() int {
	// TODO not implemented
	return 0
}

// Write ..
func (p *Poseidon) Write(b []byte) (int, error) {
	// TODO not implemented
	return 0, nil
}

// Sum ..
func (p *Poseidon) Sum(in []byte) []byte {
	// TODO not implemented
	return []byte{}
}