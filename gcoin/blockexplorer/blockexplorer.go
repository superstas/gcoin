package blockexplorer

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/superstas/gcoin/gcoin/block"
	"github.com/superstas/gcoin/gcoin/blockchain"
	"github.com/superstas/gcoin/gcoin/blockexplorer/templates"
	"github.com/superstas/gcoin/gcoin/mempool"
	"github.com/superstas/gcoin/gcoin/transaction"
)

const (
	unconfirmedTXLimit = 10
)

type SimpleBlockExplorer struct {
	storage  blockchain.Storage
	mp       *mempool.MemPool
	indexTpl *template.Template
	txTpl    *template.Template
	blockTpl *template.Template
}

func New(s blockchain.Storage, mp *mempool.MemPool) (*SimpleBlockExplorer, error) {
	indexTpl, err := template.New("i").Parse(templates.IndexTpl + templates.HeaderTpl + templates.FooterTpl)
	if err != nil {
		return nil, err
	}

	txTpl, err := template.New("tx").Parse(templates.TXTpl + templates.HeaderTpl + templates.FooterTpl)
	if err != nil {
		return nil, err
	}

	return &SimpleBlockExplorer{
		storage:  s,
		mp:       mp,
		indexTpl: indexTpl,
		txTpl:    txTpl,
	}, nil
}

func (e *SimpleBlockExplorer) ViewIndex(w http.ResponseWriter, r *http.Request) {
	blocks, err := e.storage.ReadLastNBlocks(context.Background(), 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	e.indexTpl.Execute(w, struct {
		Blocks       []block.Block
		Transactions []transaction.Transaction
	}{blocks, e.mp.Read(unconfirmedTXLimit)})
}

func (e *SimpleBlockExplorer) ViewTXHandler(w http.ResponseWriter, r *http.Request) {
	txID := r.URL.Path[len("/tx/"):]
	txIDBytes, err := hex.DecodeString(txID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tplData := struct {
		Confirmed bool
		TX        Transaction
		RawJSON   string
		BlockHash string
	}{
		Confirmed: false,
	}

	tx, err := e.mp.GetByID(txID)
	if err != nil {
		// check in Storage
		blockHash, tx, err := e.storage.FindTransactionByID(txIDBytes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tplData.Confirmed = true
		tplData.TX = newTransactionFromLegacy(tx)
		tplData.BlockHash = hex.EncodeToString(blockHash)
	} else {
		tplData.TX = newTransactionFromLegacy(tx)
	}

	jsonRepresentation, err := json.MarshalIndent(tplData.TX, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tplData.RawJSON = string(jsonRepresentation)
	e.txTpl.Execute(w, tplData)
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
