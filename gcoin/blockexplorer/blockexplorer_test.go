package blockexplorer

import (
	"context"
	"encoding/hex"
	"html/template"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/superstas/gcoin/gcoin/block"
	"github.com/superstas/gcoin/gcoin/blockchain"
	"github.com/superstas/gcoin/gcoin/mempool"
	"github.com/superstas/gcoin/gcoin/transaction"
)

func TestSimpleBlockExplorer_ViewIndex(t *testing.T) {
	t.Skip()
	mp := mempool.New()

	txID, err := hex.DecodeString("ac81f1fd432e37f70841e117d4a73f3c0fbf15a722f538a015b08abc3b6947fe")
	require.Nil(t, err)
	mp.Add(transaction.Transaction{ID: txID})

	s, err := blockchain.NewMemoryStorage()
	require.Nil(t, err)

	//be := New(s, mp)
	//be.ViewIndex()

	blocks, _ := s.ReadLastNBlocks(context.Background(), 10)

	type data struct {
		Blocks       []block.Block
		Transactions []transaction.Transaction
	}

	tpl := template.Must(template.ParseFiles("./templates/index.html.tmpl",
		"./templates/header.html.tmpl",
		"./templates/footer.html.tmpl"))

	err = tpl.Execute(os.Stdout, data{
		Blocks:       blocks,
		Transactions: mp.Get(10),
	})
	assert.Nil(t, err)
}
