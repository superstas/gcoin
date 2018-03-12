package main

import (
	"crypto/sha256"
	"hash"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://play.golang.org/p/vePTwCkQBFS

//BenchmarkIneffectiveDoubleHash-4   	 1000000	      1500 ns/op	      32 B/op	       1 allocs/op
func IneffectiveDoubleHash(b []byte) []byte {
	f := sha256.Sum256(b)
	s := sha256.Sum256(f[:])
	return s[:]
}

//BenchmarkEffectiveDoubleHash-4     	 1000000	      1427 ns/op	       0 B/op	       0 allocs/op
func EffectiveDoubleHash(h hash.Hash, d, b []byte) []byte {
	h.Reset()
	h.Write(d)
	sum := h.Sum(b[:0])
	h.Reset()
	h.Write(sum)
	return h.Sum(b[:0])
}

func TestIneffectiveDoubleHash(t *testing.T) {
	data := []byte{1}
	h := sha256.Sum256(data)
	wantHash := sha256.Sum256(h[:])

	assert.Equal(t, wantHash[:], IneffectiveDoubleHash(data))
	assert.NotEqual(t, wantHash[:], IneffectiveDoubleHash(append(data, 2)))
}

func TestEffectiveDoubleHash(t *testing.T) {
	data := []byte{1}
	buf := make([]byte, 0, sha256.Size)

	h := sha256.Sum256(data)
	wantHash := sha256.Sum256(h[:])

	assert.Equal(t, wantHash[:], EffectiveDoubleHash(sha256.New(), data, buf))
	assert.NotEqual(t, wantHash[:], EffectiveDoubleHash(sha256.New(), append(data, 2), buf))
}

//BenchmarkIneffectiveDoubleHash-4   	 1000000	      1500 ns/op	      32 B/op	       1 allocs/op
func BenchmarkIneffectiveDoubleHash(b *testing.B) {
	data := []byte{1}
	for i := 0; i < b.N; i++ {
		IneffectiveDoubleHash(data)
	}
}

//BenchmarkEffectiveDoubleHash-4     	 1000000	      1427 ns/op	       0 B/op	       0 allocs/op
func BenchmarkEffectiveDoubleHash(b *testing.B) {
	data := []byte{1}
	buf := make([]byte, 0, sha256.Size)
	hasher := sha256.New()
	for i := 0; i < b.N; i++ {
		EffectiveDoubleHash(hasher, data, buf)
	}
}
