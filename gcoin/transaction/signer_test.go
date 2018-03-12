package transaction

import (
	"testing"

	"github.com/gojuno/minimock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/superstas/gcoin/gcoin/amount"
	"github.com/superstas/gcoin/gcoin/keys"
)

func TestSign(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	alice, err := keys.New()
	require.Nil(t, err)

	bob, err := keys.New()
	require.Nil(t, err)

	reward, err := amount.NewAmount(50.)
	require.Nil(t, err)

	coinbaseTX, err := NewCoinBase(*alice.AddressPKH, reward)
	require.Nil(t, err)

	storageMock := NewStorageMock(t)
	storageMock.FindTransactionByIDFunc = func(txID []byte) ([]byte, Transaction, error) {
		require.Equal(t, txID, coinbaseTX.ID)
		return nil, coinbaseTX, nil
	}

	simpleSigner := NewSimpleSigner(storageMock)

	require.Nil(t, simpleSigner.Sign(&coinbaseTX, alice.PrivateKey))
	assert.Empty(t, coinbaseTX.Inputs[0].Sign)

	a, err := amount.NewAmount(15.)
	require.Nil(t, err)

	newTX, err := New(NewUTXOSetFromTX(coinbaseTX), *alice.AddressPKH, *bob.AddressPKH, a)
	require.Nil(t, err)
	require.Nil(t, simpleSigner.Sign(&newTX, alice.PrivateKey))

	require.NotEmpty(t, newTX.Inputs[0].Sign)
	require.NotEmpty(t, newTX.Inputs[0].PubKey)
}
