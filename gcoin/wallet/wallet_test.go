package wallet

import (
	"testing"

	"encoding/json"

	"github.com/btcsuite/btcutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/superstas/gcoin/gcoin/keys"
)

func TestInit(t *testing.T) {
	n := 5
	w, err := Init(n)
	require.Nil(t, err)
	assert.Len(t, w.Keys(), n)
}

func TestWallet_AddKeys(t *testing.T) {
	n := 5
	w, err := Init(n)
	require.Nil(t, err)
	assert.Len(t, w.Keys(), n)

	k, err := keys.New()
	require.Nil(t, err)

	w.AddKeys(k)
	assert.Len(t, w.Keys(), n+1)
}

func TestWallet_Addresses(t *testing.T) {
	k, err := keys.New()
	require.Nil(t, err)

	k2, err := keys.New()
	require.Nil(t, err)

	w := &Wallet{}
	w.AddKeys(k)
	w.AddKeys(k2)

	assert.Equal(t, w.Addresses(), []*btcutil.AddressPubKeyHash{k.AddressPKH, k2.AddressPKH})
}

func TestWalletJSON(t *testing.T) {
	w, err := Init(5)
	require.Nil(t, err)
	data, err := json.Marshal(w)
	require.Nil(t, err)
	assert.NotNil(t, data)

	var w2 Wallet
	err = json.Unmarshal(data, &w2)
	require.Nil(t, err)
	assert.Equal(t, w, &w2)
}
