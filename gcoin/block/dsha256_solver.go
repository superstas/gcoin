package block

import (
	"math"
	"math/big"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/pkg/errors"
)

// dSHA256Solver represents DoubleSHA256 solver
// This is the most widely used hashing algorithm
type dSHA256Solver struct {
	targetDifficulty *big.Int
}

// NewDSHA256Solver return a double sha256 solver implementation
func NewDSHA256Solver(td *big.Int) Solver {
	return &dSHA256Solver{
		targetDifficulty: td,
	}
}

// Solve tries to calculate a given block hash.
// This is a quite naive implementation.
// In real life, Nonce overflows quite frequently, that's why extra_nonce is used.
// For more details read this doc: https://en.bitcoin.it/wiki/Block_hashing_algorithm
// Here you can find a good example of using extraNonce with Nonce:
// https://github.com/btcsuite/btcd/blob/master/mining/cpuminer/cpuminer.go#L231
func (s *dSHA256Solver) Solve(b *Block) error {
	var i uint64
	for i < math.MaxUint64 {
		b.Header.Nonce = i
		h := chainhash.DoubleHashB(b.Header.Serialize())
		hashInt := new(big.Int).SetBytes(h)

		// As for comparing there is some inconsistency in docs.
		// Some of docs say that hash must be lower to the current target.
		// Some of docs say that lower OR EQUAL ( e.g. https://en.bitcoin.it/wiki/Target "...The SHA-256 hash of a block's header must be lower than or equal to the current target for the block to be accepted by the network.")
		// This comparing implemented according to this implementation:
		// https://github.com/btcsuite/btcd/blob/master/mining/cpuminer/cpuminer.go#L282
		if hashInt.Cmp(s.targetDifficulty) <= 0 {
			b.Hash = h
			b.Target = s.targetDifficulty.Bytes()
			return nil
		}
		i++
	}

	return errors.New("failed to solve DSHA-256")
}

// Verify verifies that the given block solved correctly with header.Target
func (s *dSHA256Solver) Verify(b Block) error {
	target := new(big.Int).SetBytes(b.Header.Target)
	hash := new(big.Int).SetBytes(b.Hash)

	if hash.Cmp(target) > 0 {
		return errors.New("block hash is invalid")
	}

	return nil
}
