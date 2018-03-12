package consensus

import (
	"context"

	"github.com/superstas/gcoin/gcoin/block"
	"github.com/superstas/gcoin/gcoin/transaction"
)

// Validator represents an interface that might has different realizations
// For better understanding what consensus does mean, read this: https://en.bitcoin.it/wiki/Consensus
type Validator interface {
	// ValidateBlock validates a given block for consensus rules
	// This is a list of checks in a real cryptocurrency https://en.bitcoin.it/wiki/Protocol_rules#.22block.22_messages
	ValidateBlock(context.Context, block.Block) error
	// ValidateTX validates a given TX for consensus rules
	// This is a list of checks in a real cryptocurrency https://en.bitcoin.it/wiki/Protocol_rules#.22tx.22_messages
	ValidateTX(context.Context, transaction.Transaction) error
}
