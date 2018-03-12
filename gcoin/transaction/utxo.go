package transaction

// UTXO represents an unspent transaction output
// https://bitcoin.org/en/glossary/unspent-transaction-output
type UTXO struct {
	TxID     []byte
	OutIndex int64
	Output
}

// UTXOSet represents a set of UTXOs
type UTXOSet []UTXO

// NewUTXOSetFromTX create UTXOSet from a TX
// It used in tests only
func NewUTXOSetFromTX(tx Transaction) UTXOSet {
	utxoSet := make([]UTXO, len(tx.Outputs))
	for i, output := range tx.Outputs {
		utxoSet[i] = UTXO{
			TxID:     tx.ID,
			OutIndex: 0,
			Output:   output,
		}
	}
	return utxoSet
}
