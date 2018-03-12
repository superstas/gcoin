package transaction

import (
	"testing"

	"github.com/gojuno/minimock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/superstas/gcoin/gcoin/amount"
	"github.com/superstas/gcoin/gcoin/keys"
)

func TestVerify(t *testing.T) {
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

	signedTX, err := New(NewUTXOSetFromTX(coinbaseTX), *alice.AddressPKH, *bob.AddressPKH, a)
	require.Nil(t, err)
	require.Nil(t, simpleSigner.Sign(&signedTX, alice.PrivateKey)) //sign TX

	simpleVerifier := NewSimpleVerifier(storageMock)
	assert.Nil(t, simpleVerifier.Verify(&signedTX)) // verify TX with Input.Signature and Input.PubKey
}

func TestVerify_InvalidSignature(t *testing.T) {
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
	storageMock.FindTransactionByIDMock.Return(nil, coinbaseTX, nil)

	simpleSigner := NewSimpleSigner(storageMock)

	require.Nil(t, simpleSigner.Sign(&coinbaseTX, alice.PrivateKey))
	assert.Empty(t, coinbaseTX.Inputs[0].Sign)

	a, err := amount.NewAmount(15.)
	require.Nil(t, err)

	signedTX, err := New(NewUTXOSetFromTX(coinbaseTX), *alice.AddressPKH, *bob.AddressPKH, a)
	require.Nil(t, err)
	require.Nil(t, simpleSigner.Sign(&signedTX, alice.PrivateKey)) //sign TX

	// break the signature
	signedTX.Inputs[0].Sign = []byte{1, 2, 3}

	simpleVerifier := NewSimpleVerifier(storageMock)
	err = simpleVerifier.Verify(&signedTX)
	assert.EqualError(t, err, "failed to parse input signature: malformed signature: too short")
}

func TestVerify_InvalidPubKey(t *testing.T) {
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
	storageMock.FindTransactionByIDMock.Return(nil, coinbaseTX, nil)

	simpleSigner := NewSimpleSigner(storageMock)

	require.Nil(t, simpleSigner.Sign(&coinbaseTX, alice.PrivateKey))
	assert.Empty(t, coinbaseTX.Inputs[0].Sign)

	a, err := amount.NewAmount(15.)
	require.Nil(t, err)

	signedTX, err := New(NewUTXOSetFromTX(coinbaseTX), *alice.AddressPKH, *bob.AddressPKH, a)
	require.Nil(t, err)
	require.Nil(t, simpleSigner.Sign(&signedTX, alice.PrivateKey)) //sign TX

	// break the pubkey
	signedTX.Inputs[0].PubKey = []byte{1, 2, 3}

	simpleVerifier := NewSimpleVerifier(storageMock)
	err = simpleVerifier.Verify(&signedTX)
	assert.EqualError(t, err, "failed to parse input pubkey: invalid pub key length 3")
}

func TestVerify_FakeSignature(t *testing.T) {
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
	storageMock.FindTransactionByIDMock.Return(nil, coinbaseTX, nil)

	simpleSigner := NewSimpleSigner(storageMock)

	require.Nil(t, simpleSigner.Sign(&coinbaseTX, alice.PrivateKey))
	assert.Empty(t, coinbaseTX.Inputs[0].Sign)

	amount, err := amount.NewAmount(15.)
	require.Nil(t, err)

	signedTX, err := New(NewUTXOSetFromTX(coinbaseTX), *alice.AddressPKH, *bob.AddressPKH, amount)
	require.Nil(t, err)
	require.Nil(t, simpleSigner.Sign(&signedTX, alice.PrivateKey)) //sign TX

	fakeTX, err := New(NewUTXOSetFromTX(coinbaseTX), *bob.AddressPKH, *alice.AddressPKH, amount)
	require.Nil(t, err)
	require.Nil(t, simpleSigner.Sign(&fakeTX, bob.PrivateKey))

	// fake the sign
	signedTX.Inputs[0].Sign = fakeTX.Inputs[0].Sign

	simpleVerifier := NewSimpleVerifier(storageMock)
	err = simpleVerifier.Verify(&signedTX)
	assert.EqualError(t, err, "signature is invalid")
}

func TestVerify_StorageReturnErr(t *testing.T) {
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
	storageMock.FindTransactionByIDMock.Return(nil, coinbaseTX, nil)

	simpleSigner := NewSimpleSigner(storageMock)

	require.Nil(t, simpleSigner.Sign(&coinbaseTX, alice.PrivateKey))
	assert.Empty(t, coinbaseTX.Inputs[0].Sign)

	a, err := amount.NewAmount(15.)
	require.Nil(t, err)

	signedTX, err := New(NewUTXOSetFromTX(coinbaseTX), *alice.AddressPKH, *bob.AddressPKH, a)
	require.Nil(t, err)
	require.Nil(t, simpleSigner.Sign(&signedTX, alice.PrivateKey)) //sign TX

	storageMock.FindTransactionByIDMock.Return(nil, coinbaseTX, errors.New("TX not found"))
	simpleVerifier := NewSimpleVerifier(storageMock)
	err = simpleVerifier.Verify(&signedTX)
	assert.EqualError(t, err, "failed to find TXByID: TX not found")
}
