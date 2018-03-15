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

type tx struct {
	ID      string   `json:"id"`
	Inputs  []input  `json:"inputs"`
	Outputs []output `json:"outputs"`
}

type output struct {
	Value   amount.Amount `json:"amount"`
	Address string        `json:"address"`
}

type input struct {
	TransactionID string `json:"tx_id,omitempty"`
	OutIndex      int64  `json:"out_index"`
	Sign          string `json:"sign,omitempty"`
	PubKey        string `json:"pub_key,omitempty"`
}

type header struct {
	PreviousBlockHash string `json:"prev_block_hash,omitempty"`
	MerkleRootHash    string `json:"merkle_root_hash"`
	Timestamp         int64  `json:"ts"`
	Target            string `json:"target"`
	Nonce             uint64 `json:"nonce"`
}

type internalBlock struct {
	Hash         string `json:"hash"`
	Header       header `json:"header"`
	Transactions []tx   `json:"transactions"`
}

func newInternalBlock(b block.Block) internalBlock {
	newBlock := internalBlock{
		Hash: b.HexHash(),
		Header: header{
			PreviousBlockHash: hex.EncodeToString(b.PreviousBlockHash),
			MerkleRootHash:    hex.EncodeToString(b.MerkleRootHash),
			Timestamp:         b.Timestamp,
			Nonce:             b.Nonce,
			Target:            fmt.Sprintf("%064x", b.Target),
		},
		Transactions: make([]tx, 0, len(b.Transactions)),
	}

	for _, tx := range b.Transactions {
		newBlock.Transactions = append(newBlock.Transactions, newInternalTransaction(tx))
	}

	return newBlock
}

func newInternalTransaction(t transaction.Transaction) tx {
	newTX := tx{
		ID:      t.HexID(),
		Outputs: make([]output, 0, len(t.Outputs)),
		Inputs:  make([]input, 0, len(t.Inputs)),
	}

	for _, in := range t.Inputs {
		newTX.Inputs = append(newTX.Inputs, input{
			TransactionID: hex.EncodeToString(in.TransactionID),
			OutIndex:      in.OutIndex,
			Sign:          hex.EncodeToString(in.Sign),
			PubKey:        hex.EncodeToString(in.PubKey),
		})
	}

	for _, out := range t.Outputs {
		addr, _ := btcutil.NewAddressPubKeyHash(out.Address, &chaincfg.Params{})
		// todo: handle err
		newTX.Outputs = append(newTX.Outputs, output{
			Value:   amount.Amount(out.Value),
			Address: addr.EncodeAddress(),
		})
	}
	return newTX
}
