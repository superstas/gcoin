package blockexplorer

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/superstas/gcoin/gcoin/amount"
	"github.com/superstas/gcoin/gcoin/block"
	"github.com/superstas/gcoin/gcoin/transaction"
)

type Transaction struct {
	ID      string
	Inputs  []Input
	Outputs []Output
}

type Output struct {
	Value   amount.Amount
	Address string
}

type Input struct {
	TransactionID string
	OutIndex      int64
	Sign          string
	PubKey        string
}

type BlockHeader struct {
	PreviousBlockHash string
	MerkleRootHash    string
	Timestamp         int64
	Target            string
	Nonce             uint64
}

type Block struct {
	BlockHeader
	Hash         string
	Transactions []Transaction
}

func newBlockFromLegacy(b block.Block) Block {
	newBlock := Block{
		Hash: b.HexHash(),
		BlockHeader: BlockHeader{
			PreviousBlockHash: hex.EncodeToString(b.PreviousBlockHash),
			MerkleRootHash:    hex.EncodeToString(b.MerkleRootHash),
			Timestamp:         b.Timestamp,
			Nonce:             b.Nonce,
			Target:            fmt.Sprintf("%064x", b.Target),
		},
		Transactions: make([]Transaction, 0, len(b.Transactions)),
	}

	for _, tx := range b.Transactions {
		newBlock.Transactions = append(newBlock.Transactions, newTransactionFromLegacy(tx))
	}

	return newBlock
}

func newTransactionFromLegacy(tx transaction.Transaction) Transaction {
	newTX := Transaction{
		ID:      tx.HexID(),
		Outputs: make([]Output, 0, len(tx.Outputs)),
		Inputs:  make([]Input, 0, len(tx.Inputs)),
	}

	for _, in := range tx.Inputs {
		newTX.Inputs = append(newTX.Inputs, Input{
			TransactionID: hex.EncodeToString(in.TransactionID),
			OutIndex:      in.OutIndex,
			Sign:          hex.EncodeToString(in.Sign),
			PubKey:        hex.EncodeToString(in.PubKey),
		})
	}

	for _, out := range tx.Outputs {
		addr, _ := btcutil.NewAddressPubKeyHash(out.Address, &chaincfg.Params{})
		// todo: handle err
		newTX.Outputs = append(newTX.Outputs, Output{
			Value:   amount.Amount(out.Value),
			Address: addr.EncodeAddress(),
		})
	}
	return newTX
}
