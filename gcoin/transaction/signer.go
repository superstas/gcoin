package transaction

import (
	"github.com/btcsuite/btcd/btcec"
	"github.com/pkg/errors"
)

// Signer represents a signer interface
type Signer interface {
	Sign(*Transaction, *btcec.PrivateKey) error
}

type simpleSigner struct {
	storage Storage
}

// NewSimpleSigner returns a new Signer
func NewSimpleSigner(s Storage) Signer {
	return &simpleSigner{s}
}

// Sign computes a signature for given TX.
// You can read more about that algorithm here:
// - https://en.bitcoin.it/wiki/OP_CHECKSIG
// - https://bitcoin.stackexchange.com/a/5241
// This implementation is inspired by Jeiwan signing implementation https://github.com/Jeiwan/blockchain_go/blob/master/transaction.go#L58
func (s *simpleSigner) Sign(tx *Transaction, privKey *btcec.PrivateKey) error {
	if tx == nil {
		return errors.New("input TX is nil")
	}

	if isCoinBase(tx) {
		return nil
	}

	txClone := cloneTX(tx)
	for i, input := range txClone.Inputs {
		_, inputTX, err := s.storage.FindTransactionByID(input.TransactionID)
		if err != nil {
			return errors.Wrap(err, "failed to find TXByID")
		}

		// Since we don't have an internal script language
		// this step is simplified step from this manual https://en.bitcoin.it/wiki/OP_CHECKSIG
		// See "How it works", steps 7 and 8
		txClone.Inputs[i].PubKey = inputTX.Outputs[input.OutIndex].Address

		h, err := txClone.Hash()
		if err != nil {
			return errors.Wrap(err, "failed to calculate TX hash")
		}

		sign, err := privKey.Sign(h)
		if err != nil {
			return errors.Wrap(err, "failed to sign TX")
		}

		tx.Inputs[i].Sign = sign.Serialize()
		tx.Inputs[i].PubKey = privKey.PubKey().SerializeCompressed()
	}
	return nil
}

func isCoinBase(tx *Transaction) bool {
	return len(tx.Inputs) == 1 && tx.Inputs[0].OutIndex == -1 && len(tx.Outputs) == 1
}

func cloneTX(tx *Transaction) Transaction {
	txClone := Transaction{
		Inputs:  make([]Input, len(tx.Inputs)),
		Outputs: make([]Output, len(tx.Outputs)),
	}

	for i, input := range tx.Inputs {
		txClone.Inputs[i] = Input{
			TransactionID: input.TransactionID,
			OutIndex:      input.OutIndex,
		}
	}

	for i, output := range tx.Outputs {
		txClone.Outputs[i] = Output{
			Value:   output.Value,
			Address: output.Address,
		}
	}
	return txClone
}
