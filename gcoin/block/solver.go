package block

// Solver represents a simple interface that knows how to hash a given block
// You can read more about different hashing algorithms here: https://en.bitcoin.it/wiki/Proof_of_work
type Solver interface {
	Solve(*Block) error
	Verify(Block) error
}
