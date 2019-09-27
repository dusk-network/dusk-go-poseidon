# Dusk Poseidon - Go

[![Build Status](https://travis-ci.com/dusk-network/dusk-go-poseidon.svg?token=czzGwcZEd8hUsCLG3xJC&branch=master)](https://travis-ci.com/dusk-network/dusk-go-poseidon)

GoLang implementation of Poseidon hash-only function. This follows-up [Dusk Hades252](https://github.com/dusk-network/Hades252)

## Guidelines

### Custom parameters

The parameters are defined by the number of beginning, end and partial rounds. It is recommended, though not mandatory, that the number of beginning rounds equals the number of end rounds.

Also, there will be the arity of the Merkle tree; the maximum number of allowed scalars will equal the arity. Therefore, the `Width` attribute of the parametrization will be `arity + 1`, since the first element of the input set will always be used internally for the hash computation.

The number of constants must be, optimally, `(arity + 1) * (beginning_rounds + end_rounds + partial_rounds)`. Even if more constants are provided, they will not be used. If a lower number of constants is provided, the hash cannot be computed.

The matrix must be MDS, with width and height defined by `arity + 1`

If you are unsure, just call the `poseidon.New()` constructor, and an instance with default parameters will be returned.

## Example

```go
package main

import (
	"reflect"

	ristretto "github.com/bwesterb/go-ristretto"
	"github.com/dusk-network/dusk-go-poseidon/pkg/core/poseidon"
)

func main() {
	// Instantiate a new poseidon hasher
	p := poseidon.New()

	// Prepare some arbitrary scalars to be inserted
	a := ristretto.Scalar{0x2e135665, 0x7d5a70f0, 0xe16c44d8, 0x34937f95, 0x918f4146, 0xcf4c92eb, 0xed739c42, 0x054a144f}
	b := ristretto.Scalar{0xcef5ca00, 0x645d9237, 0x570e0c2f, 0xf235c67d, 0x92d59395, 0x461a4091, 0x7ce77ad9, 0x00e01594}

	// Input the provided scalars to the set
	p.WriteScalar(a)
	p.WriteScalar(b)

	// It's also possible to treat the hasher as an implementation of io.Write,
	// by inserting raw bytes. These bytes can be of any size, for internally
	// ristretto implementation will reduce them to 64 bytes as instance of SHA512,
	// and then to 32 as `t mod l`
	p.Write([]byte("Time is an illusion. Lunchtime doubly so."))

	// Expected result
	digest := []byte{0x83, 0xc6, 0x9d, 0xf7, 0x92, 0x08, 0x59, 0xbf, 0x6a, 0x7d, 0x13, 0x80, 0x45, 0xe3, 0x21, 0x7b, 0x57, 0xd7, 0x36, 0x61, 0xf2, 0xba, 0x3a, 0x52, 0xfd, 0xd0, 0x1d, 0x81, 0xfb, 0xa9, 0x08, 0x0b}

	// Perform the computation of the hash
	result := p.Sum(nil)

	// The calculated hash must equal the expected value. It is always
	// ristretto.Scalar form ([32]byte)
	if !reflect.DeepEqual(digest, result) {
		panic("This is unexpected")
	}
}
```

## Reference

[Starkad and Poseidon: New Hash Functions for Zero Knowledge Proof Systems](https://eprint.iacr.org/2019/458.pdf)
