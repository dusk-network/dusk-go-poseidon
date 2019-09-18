package scalar

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewScalar(t *testing.T) {
	bytes := make([]byte, 32)
	for i := 0; i < 32; i++ {
		bytes[i] = byte(31 - i)
	}
	s, err := New(bytes)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 32; i++ {
		assert.Equal(t, byte(31-i), s.Bytes[i])
	}
}

func TestNewScalarFromSmallSlice(t *testing.T) {
	bytes := make([]byte, 16)
	for i := 0; i < 16; i++ {
		bytes[i] = byte(16 - i)
	}
	s, err := New(bytes)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 16; i++ {
		assert.Equal(t, byte(16-i), s.Bytes[i])
	}
}

func TestNewScalarFromSingleByte(t *testing.T) {
	s, err := New([]byte{0x01})
	if err != nil {
		log.Fatal(err)
	}

	for i, b := range s.Bytes {
		if i == 0 {
			assert.Equal(t, byte(0x01), b)
		} else {
			assert.Equal(t, byte(0x00), b)
		}
	}
}

func TestScalarOverflow(t *testing.T) {
	bytes := make([]byte, 33)
	_, err := New(bytes)
	if err == nil {
		log.Fatal("Expected overflow")
	}
}

func TestScalarMaximumValue(t *testing.T) {
	bytes := make([]byte, 32)
	bytes[31] = 0xff
	_, err := New(bytes)
	if err == nil {
		log.Fatal("Expected overflow")
	}
}
