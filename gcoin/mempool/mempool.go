package mempool

import (
	"sync"

	"github.com/pkg/errors"
	"github.com/superstas/gcoin/gcoin/transaction"
)

// MemPool represents a very simple memory pool for unconfirmed transactions
// There is an expiration period for each TX in real implementations.
//
// You can read more about mempool:
// https://bitcoin.stackexchange.com/questions/46152/how-do-transactions-leave-the-memory-pool
// https://bitcoin.stackexchange.com/questions/41536/why-does-bitcoin-keep-transactions-in-a-memory-pool
type MemPool struct {
	pool map[string]transaction.Transaction
	l    sync.RWMutex
}

// New returns new MemPool
func New() *MemPool {
	return &MemPool{
		pool: make(map[string]transaction.Transaction, 1024),
	}
}

// Add adds new transaction to the pool
func (m *MemPool) Add(tx transaction.Transaction) {
	m.l.Lock()
	m.pool[tx.HexID()] = tx
	m.l.Unlock()
}

// GetByID returns a transaction with a given transaction hex-encoded ID
func (m *MemPool) GetByID(txHexID string) (transaction.Transaction, error) {
	m.l.RLock()
	tx, ok := m.pool[txHexID]
	m.l.RUnlock()
	if !ok {
		return transaction.Transaction{}, errors.New("tx not found")
	}
	return tx, nil
}

// Get returns N transactions and cleans the pool
func (m *MemPool) Get(n int) []transaction.Transaction {
	foundTXs := make([]transaction.Transaction, 0, n)
	m.l.RLock()
	defer func() {
		for _, tx := range foundTXs {
			delete(m.pool, tx.HexID())
		}
		m.l.RUnlock()
	}()

	for _, t := range m.pool {
		if len(foundTXs) >= n {
			return foundTXs
		}
		foundTXs = append(foundTXs, t)
	}
	return foundTXs
}

// DeleteByID removes a transaction with a given transaction hex-encoded ID
func (m *MemPool) DeleteByID(txHexID string) error {
	m.l.Lock()
	delete(m.pool, txHexID)
	m.l.Unlock()
	return nil
}

// Size return size of the MemPool
func (m *MemPool) Size() int {
	return len(m.pool)
}
