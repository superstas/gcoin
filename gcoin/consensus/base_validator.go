package consensus

import (
	"context"

	"github.com/pkg/errors"
	"github.com/superstas/gcoin/gcoin/block"
	"github.com/superstas/gcoin/gcoin/blockchain"
	"github.com/superstas/gcoin/gcoin/mempool"
	"github.com/superstas/gcoin/gcoin/transaction"
)

// baseValidator is a very simple validator
type baseValidator struct {
	storage    blockchain.Storage
	mempool    *mempool.MemPool
	solver     block.Solver
	txVerifier transaction.Verifier
}

// NewBaseValidator returns baseValidator
func NewBaseValidator(s blockchain.Storage, mp *mempool.MemPool, sv block.Solver, v transaction.Verifier) Validator {
	return &baseValidator{
		storage:    s,
		mempool:    mp,
		solver:     sv,
		txVerifier: v,
	}
}

// ValidateBlock validates a given block for consensus rules
// This is a list of checks in a real cryptocurrency https://en.bitcoin.it/wiki/Protocol_rules#.22block.22_messages
func (v *baseValidator) ValidateBlock(ctx context.Context, b block.Block) error {
	// Actually, this check shouldn't be here.
	// Because it's ok when a node receives a block that points to a non existed block in the local storage ( it's called "orphan blocks" ).
	// Since we haven't sync algorithm, it's ok for us.
	if _, err := v.storage.ReadBlockByHash(ctx, b.PreviousBlockHash); err != nil {
		return errors.Wrap(err, "previous block not found")
	}

	if _, err := v.storage.ReadBlockByHash(ctx, b.Hash); err == nil {
		return errors.Wrap(err, "block already exists")
	}

	//TODO: check that block target equals to the current target
	return errors.Wrap(v.solver.Verify(b), "failed to validate block")
}

// ValidateTX validates a given TX for consensus rules
// This is a list of checks in a real cryptocurrency https://en.bitcoin.it/wiki/Protocol_rules#.22tx.22_messages
func (v *baseValidator) ValidateTX(ctx context.Context, tx transaction.Transaction) error {
	if _, err := v.mempool.GetByID(tx.HexID()); err == nil {
		return errors.Wrap(err, "transaction already exists")
	}

	//TODO: check used UTXO
	return errors.Wrap(v.txVerifier.Verify(&tx), "failed to validate TX")
}
