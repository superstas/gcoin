package transaction

import (
	"context"

	"github.com/btcsuite/btcutil"
)

// Storage represents a storage for transactions
// A storage for transaction might differ from blockchain storage
type Storage interface {
	// FindTransactionByID finds a tx by given ID
	FindTransactionByID(TxID []byte) ([]byte, Transaction, error)
	// FindUTXOByPKH finds a set of UTXO by a given public key hash ( address )
	FindUTXOByPKH(context.Context, btcutil.AddressPubKeyHash) (UTXOSet, error)
	// TotalUTXOByPKH returns a sum of UTXO by a given public key hash ( address )
	TotalUTXOByPKH(context.Context, btcutil.AddressPubKeyHash) (int64, error)
}
