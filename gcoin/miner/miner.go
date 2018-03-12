package miner

import (
	"context"
	"log"
	"time"

	"github.com/btcsuite/btcutil"
	"github.com/pkg/errors"
	"github.com/superstas/gcoin/gcoin/block"
	"github.com/superstas/gcoin/gcoin/blockchain"
	"github.com/superstas/gcoin/gcoin/mempool"
	"github.com/superstas/gcoin/gcoin/transaction"
)

const (
	// this limit for demo purposes only
	// https://bitcoin.stackexchange.com/questions/10457/what-is-the-number-of-transactions-in-a-block
	// https://bitcoin.stackexchange.com/questions/7311/how-do-miners-select-which-transactions-to-include-in-a-block
	maxTransactionsInBlock = 10
	// this value shouldn't be static, but for the prototype it's ok
	rewardForBlock = 5000000000
)

// SimpleMiner represents a very simple miner implementation
// A few real implementation:
// - https://github.com/amir20/sha-miner
// - https://github.com/btcsuite/btcd/tree/master/mining/cpuminer
type SimpleMiner struct {
	storage      blockchain.Storage
	mempool      *mempool.MemPool
	solver       block.Solver
	minerAddress btcutil.AddressPubKeyHash
}

// New creates a simple miner
func New(s blockchain.Storage, m *mempool.MemPool, sl block.Solver, a btcutil.AddressPubKeyHash) *SimpleMiner {
	return &SimpleMiner{
		storage:      s,
		mempool:      m,
		solver:       sl,
		minerAddress: a,
	}
}

// Run starts listening mempool and mining
func (m *SimpleMiner) Run(ctx context.Context, foundBlock chan block.Block) {
	log.Println("[miner]: waiting for transactions")
	for range time.Tick(time.Second) {
		// TODO: Size+Get should be atomic
		if m.mempool.Size() == 0 {
			continue
		}

		log.Println("[miner]: started hashing...")
		b, err := m.mine(ctx)
		if err != nil {
			log.Printf("[miner]: failed to mine: %v\n", err)
		} else {
			log.Printf("[miner]: block %q found with nonce %d\n", b.HexHash(), b.Nonce)

			err := m.storage.WriteBlock(ctx, b)
			if err != nil {
				log.Println("[miner]: failed to save new block")
			} else {
				foundBlock <- b
			}
		}
	}
}

func (m *SimpleMiner) mine(ctx context.Context) (block.Block, error) {
	bc, err := m.blockCandidate(ctx)
	if err != nil {
		return bc, err
	}

	if err := m.solver.Solve(&bc); err != nil {
		return bc, err
	}

	return bc, nil
}

func (m *SimpleMiner) blockCandidate(ctx context.Context) (block.Block, error) {
	var b block.Block
	if m.mempool.Size() == 0 {
		return b, errors.New("mempool is empty")
	}

	lastBlockHash, err := m.storage.ReadLastBlockHash(ctx)
	if err != nil {
		return b, err
	}

	txs := m.mempool.Get(maxTransactionsInBlock)
	coinbaseTX, err := transaction.NewCoinBase(m.minerAddress, rewardForBlock)
	if err != nil {
		return b, err
	}

	return block.New(
		lastBlockHash,
		append([]transaction.Transaction{coinbaseTX}, txs...),
	), nil
}
