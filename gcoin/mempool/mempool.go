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
	defer m.l.RUnlock()

	tx, ok := m.pool[txHexID]
	if !ok {
		return transaction.Transaction{}, errors.New("tx not found")
	}
	return tx, nil
}

// Get returns N transactions and cleans the pool
func (m *MemPool) Get(n int) []transaction.Transaction {
	foundTXs := make([]transaction.Transaction, 0, n)
	m.l.Lock()
	defer func() {
		for _, tx := range foundTXs {
			delete(m.pool, tx.HexID())
		}
		m.l.Unlock()
	}()

	for _, t := range m.pool {
		if len(foundTXs) >= n {
			return foundTXs
		}
		foundTXs = append(foundTXs, t)
	}
	return foundTXs
}

// Read returns N transactions without cleaning
func (m *MemPool) Read(n int) []transaction.Transaction {
	m.l.RLock()
	defer m.l.RUnlock()

	txs := make([]transaction.Transaction, 0, n)
	for _, t := range m.pool {
		txs = append(txs, t)
	}
	return txs
}

// DeleteByID removes a transaction with a given transaction hex-encoded ID
func (m *MemPool) DeleteByID(txHexID string) {
	m.l.Lock()
	defer m.l.Unlock()
	delete(m.pool, txHexID)
}

// Size return size of the MemPool
func (m *MemPool) Size() int {
	return len(m.pool)
}
