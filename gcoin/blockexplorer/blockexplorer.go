package blockexplorer

import (
	"context"
	"encoding/hex"
	"encoding/json"
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

// SimpleBlockExplorer represents a very simple HTTP block explorer
type SimpleBlockExplorer struct {
	storage  blockchain.Storage
	mp       *mempool.MemPool
	indexTpl *template.Template
	txTpl    *template.Template
	blockTpl *template.Template
}

// New creates new block explorer
func New(s blockchain.Storage, mp *mempool.MemPool) (*SimpleBlockExplorer, error) {
	layer := templates.HeaderTPL + templates.FooterTPL
	indexTpl, err := template.New("i").Parse(templates.IndexTPL + layer)
	if err != nil {
		return nil, err
	}

	txTpl, err := template.New("tx").Parse(templates.TXTPL + layer)
	if err != nil {
		return nil, err
	}

	blockTpl, err := template.New("block").Parse(templates.BlockTPL + layer)
	if err != nil {
		return nil, err
	}

	return &SimpleBlockExplorer{
		storage:  s,
		mp:       mp,
		indexTpl: indexTpl,
		txTpl:    txTpl,
		blockTpl: blockTpl,
	}, nil
}

// ViewIndex executes index page template with general data
func (e *SimpleBlockExplorer) ViewIndex(w http.ResponseWriter, r *http.Request) {
	blocks, err := e.storage.ReadLastNBlocks(context.Background(), 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := e.indexTpl.Execute(w, struct {
		Blocks       []block.Block
		Transactions []transaction.Transaction
	}{blocks, e.mp.Read(unconfirmedTXLimit)}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ViewTX executes transaction page template with TX data
func (e *SimpleBlockExplorer) ViewTX(w http.ResponseWriter, r *http.Request) {
	txID := r.URL.Path[len("/tx/"):]
	txIDBytes, err := hex.DecodeString(txID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tplData := struct {
		Confirmed bool
		TX        tx
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
		tplData.TX = newInternalTransaction(tx)
		tplData.BlockHash = hex.EncodeToString(blockHash)
	} else {
		tplData.TX = newInternalTransaction(tx)
	}

	jsonRepresentation, err := json.MarshalIndent(tplData.TX, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tplData.RawJSON = string(jsonRepresentation)
	if err := e.txTpl.Execute(w, tplData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ViewBlock executes block page template with Block data
func (e *SimpleBlockExplorer) ViewBlock(w http.ResponseWriter, r *http.Request) {
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

	ib := newInternalBlock(b)
	jsonRepresentation, err := json.MarshalIndent(ib, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := e.blockTpl.Execute(w, struct {
		Block   internalBlock
		RawJSON string
	}{ib, string(jsonRepresentation)}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
