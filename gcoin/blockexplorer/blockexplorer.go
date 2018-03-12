package blockexplorer

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/superstas/gcoin/gcoin/blockchain"
	"github.com/superstas/gcoin/gcoin/mempool"
)

type SimpleBlockExplorer struct {
	storage blockchain.Storage
	mp      *mempool.MemPool
}

func New(s blockchain.Storage, mp *mempool.MemPool) *SimpleBlockExplorer {
	return &SimpleBlockExplorer{storage: s, mp: mp}
}

func (e *SimpleBlockExplorer) ViewTXHandler(w http.ResponseWriter, r *http.Request) {
	txID := r.URL.Path[len("/tx/"):]
	txIDBytes, err := hex.DecodeString(txID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// check tx in MemPool
	memPoolTX, err := e.mp.GetByID(txID)
	if err != nil {
		// check in Storage
		blockHash, tx, err := e.storage.FindTransactionByID(txIDBytes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonRepresentation, err := json.MarshalIndent(newTransactionFromLegacy(tx), "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "<h1>TX ID: %[1]s</h1><h2 style=\"color:green\">confirmed</h2>Included in block <a href=\"/block/%[2]s\">%[2]s</a><pre>%[3]s</pre>",
			txID, hex.EncodeToString(blockHash), jsonRepresentation)
		return
	}

	jsonRepresentation, err := json.MarshalIndent(newTransactionFromLegacy(memPoolTX), "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "<h1>TX ID: %[1]s</h1><h2 style=\"color:red\">unconfirmed</h2><pre>%[2]s</pre>", txID, jsonRepresentation)
}

func (e *SimpleBlockExplorer) ViewBlockHandler(w http.ResponseWriter, r *http.Request) {
	blockHash := r.URL.Path[len("/block/"):]
	blockHashBytes, err := hex.DecodeString(blockHash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := e.storage.ReadBlockByHash(context.Background(), blockHashBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonRepresentation, err := json.MarshalIndent(newBlockFromLegacy(b), "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "<h1>Block hash: %s</h1><pre>%s</pre>", blockHash, jsonRepresentation)
}
