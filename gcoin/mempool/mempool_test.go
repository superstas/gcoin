package mempool

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/superstas/gcoin/gcoin/transaction"
)

func TestMempool_Get(t *testing.T) {
	tx1 := newMockTX(t, "01")
	tx2 := newMockTX(t, "02")
	tx3 := newMockTX(t, "03")
	tx3Dupl := newMockTX(t, "03")

	p := New()
	p.Add(tx1)
	p.Add(tx2)
	p.Add(tx3)
	p.Add(tx3Dupl)
	assert.Equal(t, 3, p.Size())

	txs := p.Get(1)
	assert.Equal(t, 1, len(txs))
	assert.Equal(t, 2, p.Size())
	txs = p.Get(2)
	assert.Equal(t, 2, len(txs))
	assert.Equal(t, 0, p.Size())
}

func TestMempool_GetByID(t *testing.T) {
	tx1 := newMockTX(t, "01")
	tx2 := newMockTX(t, "02")
	tx3 := newMockTX(t, "03")

	p := New()
	p.Add(tx1)
	p.Add(tx2)
	p.Add(tx3)

	_, err := p.GetByID("111")
	assert.NotNil(t, err)
	assert.Equal(t, 3, p.Size())

	tx, err := p.GetByID("01")
	assert.Nil(t, err)
	assert.Equal(t, tx1, tx)
	assert.Equal(t, 3, p.Size())
}

func TestMempool_DeleteByID(t *testing.T) {
	tx1 := newMockTX(t, "01")
	tx2 := newMockTX(t, "02")
	tx3 := newMockTX(t, "03")

	p := New()
	p.Add(tx1)
	p.Add(tx2)
	p.Add(tx3)
	assert.Equal(t, 3, p.Size())

	err := p.DeleteByID("01")
	assert.Nil(t, err)
	assert.Equal(t, 2, p.Size())
	_, err = p.GetByID("01")
	assert.NotNil(t, err)

	err = p.DeleteByID("01")
	assert.Nil(t, err)
	assert.Equal(t, 2, p.Size())

	err = p.DeleteByID("02")
	assert.Nil(t, err)
	assert.Equal(t, 1, p.Size())

}

func newMockTX(t *testing.T, txHexID string) transaction.Transaction {
	txID, err := hex.DecodeString(txHexID)
	assert.Nil(t, err)
	return transaction.Transaction{ID: txID}
}
