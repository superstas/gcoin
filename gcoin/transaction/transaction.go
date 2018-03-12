package transaction

import (
	"crypto/rand"
	"encoding/json"

	"encoding/hex"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcutil"
	"github.com/pkg/errors"
	"github.com/superstas/gcoin/gcoin/amount"
)

// Transaction represents a TX
type Transaction struct {
	// ID calculated with Hash method
	ID []byte
	// Inputs has a set of Input
	Inputs []Input
	// Outputs has a set of Output
	Outputs []Output
}

// Input represents an input
type Input struct {
	// TransactionID points to a particular TX
	TransactionID []byte
	// OutIndex points to a particular output in the TX
	OutIndex int64
	// Sign is calculated with Signer
	Sign []byte
	// PubKey represents a public key of a sender
	PubKey []byte
}

// Output represents an output
type Output struct {
	// Value has amount of coins
	Value int64
	// Address represents a PKH (public key hash) of a destination
	Address []byte
}

// New creates a TX
func New(us UTXOSet, from, to btcutil.AddressPubKeyHash, a amount.Amount) (Transaction, error) {
	tx := Transaction{}

	var (
		total amount.Amount
		err   error
	)

	// find needed UTXO's
	for i, u := range us {
		if i == 0 || total < a {
			tx.Inputs = append(tx.Inputs, Input{
				TransactionID: u.TxID,
				OutIndex:      u.OutIndex,
			})
		} else {
			break
		}
		total += amount.Amount(u.Value)
	}

	if total < a {
		return tx, errors.New("there are no enough coins")
	}

	//split outputs
	tx.Outputs = []Output{
		{
			Value:   int64(a),
			Address: to.ScriptAddress(),
		},
	}

	// change
	if total > a {
		tx.Outputs = append(tx.Outputs, Output{
			Value:   int64(total - a),
			Address: from.ScriptAddress(),
		})
	}

	tx.ID, err = tx.Hash()
	return tx, err
}

// NewCoinBase creates new coinbase TX with given reward
func NewCoinBase(to btcutil.AddressPubKeyHash, reward amount.Amount) (Transaction, error) {
	tx := Transaction{}

	// a bit of entropy for prevent same txID
	randData := make([]byte, 32)
	_, err := rand.Read(randData)
	if err != nil {
		return tx, err
	}

	tx.Inputs = []Input{{OutIndex: -1, PubKey: randData}}
	tx.Outputs = []Output{
		{
			Value:   int64(reward),
			Address: to.ScriptAddress(),
		}}

	tx.ID, err = tx.Hash()
	return tx, err
}

// Serialize serializes TX
// This is not the best way to serialize TX.
func (tx Transaction) Serialize() ([]byte, error) {
	tx.ID = nil
	return json.Marshal(tx)
}

// Hash returns SHA256(SHA256(TX))
func (tx Transaction) Hash() ([]byte, error) {
	d, err := tx.Serialize()
	if err != nil {
		return nil, err
	}

	return chainhash.DoubleHashB(d), nil
}

// HexID returns hex encoded tx ID
func (tx Transaction) HexID() string {
	return hex.EncodeToString(tx.ID)
}
