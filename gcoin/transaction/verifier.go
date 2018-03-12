package transaction

import (
	"github.com/btcsuite/btcd/btcec"
	"github.com/pkg/errors"
)

// Verifier represents a simpleSigner TX verifier
type Verifier interface {
	Verify(*Transaction) error
}

type simpleVerifier struct {
	storage Storage
}

// NewSimpleVerifier return a new Verifier
func NewSimpleVerifier(s Storage) Verifier {
	return &simpleVerifier{s}
}

// Verify checks that each input signature of a given transaction is valid.
// Read more about signing at Signer.Sign
func (v *simpleVerifier) Verify(tx *Transaction) error {
	if isCoinBase(tx) {
		return nil
	}

	txClone := cloneTX(tx)
	for i, input := range tx.Inputs {
		sign, err := btcec.ParseSignature(input.Sign, btcec.S256())
		if err != nil {
			return errors.Wrap(err, "failed to parse input signature")
		}

		pubkey, err := btcec.ParsePubKey(input.PubKey, btcec.S256())
		if err != nil {
			return errors.Wrap(err, "failed to parse input pubkey")
		}

		_, inputTX, err := v.storage.FindTransactionByID(input.TransactionID)
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

		if ok := sign.Verify(h, pubkey); !ok {
			return errors.New("signature is invalid")
		}
	}
	return nil
}
