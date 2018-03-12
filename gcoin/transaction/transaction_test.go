package transaction

import (
	"testing"

	"encoding/hex"

	"github.com/btcsuite/btcutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/superstas/gcoin/gcoin/amount"
	"github.com/superstas/gcoin/gcoin/keys"
)

func TestNewCoinbase(t *testing.T) {
	p, err := keys.New()
	require.Nil(t, err)

	reward, err := amount.NewAmount(1.)
	require.Nil(t, err)

	tx, err := NewCoinBase(*p.AddressPKH, reward)
	require.Nil(t, err)

	require.Len(t, tx.Inputs, 1)
	require.Len(t, tx.Outputs, 1)
	assert.NotEmpty(t, tx.ID)
	assert.Equal(t, hex.EncodeToString(tx.ID), tx.HexID())

	assert.EqualValues(t, -1, tx.Inputs[0].OutIndex)
	assert.NotEmpty(t, tx.Inputs[0].PubKey)
	assert.Empty(t, tx.Inputs[0].TransactionID)
	assert.Empty(t, tx.Inputs[0].Sign)

	assert.EqualValues(t, reward, tx.Outputs[0].Value)
	assert.Equal(t, p.AddressPKH.ScriptAddress(), tx.Outputs[0].Address)
}

func TestNew_OneInput(t *testing.T) {
	alice, err := keys.New()
	require.Nil(t, err)

	bob, err := keys.New()
	require.Nil(t, err)

	reward, err := amount.NewAmount(50.)
	require.Nil(t, err)

	coinbaseTX, err := NewCoinBase(*alice.AddressPKH, reward)
	require.Nil(t, err)

	amount, err := amount.NewAmount(15.)
	require.Nil(t, err)

	newTX, err := New(NewUTXOSetFromTX(coinbaseTX), *alice.AddressPKH, *bob.AddressPKH, amount)
	require.Nil(t, err)

	assert.Len(t, newTX.Inputs, 1)
	assert.Len(t, newTX.Outputs, 2)
	assert.NotEmpty(t, newTX.ID)

	assert.Equal(t, coinbaseTX.ID, newTX.Inputs[0].TransactionID)
	assert.EqualValues(t, 0, newTX.Inputs[0].OutIndex)
	assert.Empty(t, newTX.Inputs[0].Sign)
	assert.Empty(t, newTX.Inputs[0].PubKey)

	assert.EqualValues(t, 15*btcutil.SatoshiPerBitcoin, newTX.Outputs[0].Value)
	assert.EqualValues(t, bob.AddressPKH.ScriptAddress(), newTX.Outputs[0].Address)

	assert.EqualValues(t, 35*btcutil.SatoshiPerBitcoin, newTX.Outputs[1].Value)
	assert.EqualValues(t, alice.AddressPKH.ScriptAddress(), newTX.Outputs[1].Address)
}

func TestNew_TwoInputs(t *testing.T) {
	alice, err := keys.New()
	require.Nil(t, err)

	bob, err := keys.New()
	require.Nil(t, err)

	reward, err := amount.NewAmount(50.)
	require.Nil(t, err)

	// first TX from Alice to Bob
	coinbaseTX, err := NewCoinBase(*alice.AddressPKH, reward)
	require.Nil(t, err)

	aliceToBobAmount, err := amount.NewAmount(15.)
	require.Nil(t, err)

	fromAliceToBobFirstTX, err := New(NewUTXOSetFromTX(coinbaseTX), *alice.AddressPKH, *bob.AddressPKH, aliceToBobAmount)
	require.Nil(t, err)

	// second TX from Alice to Bob
	coinbaseTX, err = NewCoinBase(*alice.AddressPKH, reward)
	require.Nil(t, err)

	aliceToBobAmount, err = amount.NewAmount(10.)
	require.Nil(t, err)

	fromAliceToBobSecondTX, err := New(NewUTXOSetFromTX(coinbaseTX), *alice.AddressPKH, *bob.AddressPKH, aliceToBobAmount)
	require.Nil(t, err)

	// summary UTXOSet that includes outputs from both TXs
	UTXOSet := UTXOSet{
		NewUTXOSetFromTX(fromAliceToBobFirstTX)[0],
		NewUTXOSetFromTX(fromAliceToBobSecondTX)[0],
	}

	bobToAliceAmount, err := amount.NewAmount(20.)
	require.Nil(t, err)

	newTX, err := New(UTXOSet, *bob.AddressPKH, *alice.AddressPKH, bobToAliceAmount)
	require.Nil(t, err)

	require.Len(t, newTX.Inputs, 2)
	require.Len(t, newTX.Outputs, 2)

	// check first input
	assert.Equal(t, fromAliceToBobFirstTX.ID, newTX.Inputs[0].TransactionID)
	assert.EqualValues(t, 0, newTX.Inputs[0].OutIndex)
	// check second input
	assert.Equal(t, fromAliceToBobSecondTX.ID, newTX.Inputs[1].TransactionID)
	assert.EqualValues(t, 0, newTX.Inputs[1].OutIndex)

	assert.EqualValues(t, 20*btcutil.SatoshiPerBitcoin, newTX.Outputs[0].Value)
	assert.EqualValues(t, alice.AddressPKH.ScriptAddress(), newTX.Outputs[0].Address)

	// check the change to the sender
	assert.EqualValues(t, 5*btcutil.SatoshiPerBitcoin, newTX.Outputs[1].Value)
	assert.EqualValues(t, bob.AddressPKH.ScriptAddress(), newTX.Outputs[1].Address)
}

func TestNew_ExpectedNonEnoughErr(t *testing.T) {
	alice, err := keys.New()
	require.Nil(t, err)

	bob, err := keys.New()
	require.Nil(t, err)

	a, err := amount.NewAmount(50.)
	require.Nil(t, err)

	_, err = New(UTXOSet{}, *bob.AddressPKH, *alice.AddressPKH, a)
	require.EqualError(t, err, "there are no enough coins")
}

func TestHash(t *testing.T) {
	alice, err := keys.New()
	require.Nil(t, err)

	reward, err := amount.NewAmount(50.)
	require.Nil(t, err)

	// first TX from Alice to Bob
	coinbaseTX, err := NewCoinBase(*alice.AddressPKH, reward)
	require.Nil(t, err)

	h, err := coinbaseTX.Hash()
	require.Nil(t, err)
	assert.Equal(t, coinbaseTX.ID, h)
}
