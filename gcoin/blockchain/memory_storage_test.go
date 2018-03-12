package blockchain

import (
	"context"
	"testing"

	"encoding/hex"

	"math/big"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/superstas/gcoin/gcoin/amount"
	"github.com/superstas/gcoin/gcoin/block"
	"github.com/superstas/gcoin/gcoin/keys"
	"github.com/superstas/gcoin/gcoin/transaction"
)

func TestMemoryStorage_ReadGenesisBlock(t *testing.T) {
	s := &memoryStorage{}
	b, err := s.ReadGenesisBlock(context.Background())
	require.Nil(t, err)

	assert.Empty(t, b.PreviousBlockHash)
	require.Len(t, b.Transactions, 1)
	require.Len(t, b.Transactions[0].Inputs, 1)
	require.Len(t, b.Transactions[0].Outputs, 1)

	addrPKH, err := btcutil.DecodeAddress("16xDPt4RksjfSAaNP6tdB4Wj81CSf3dkwp", &chaincfg.Params{})
	require.Nil(t, err)

	differentAddrPKH, err := btcutil.DecodeAddress("12PLGoQb9usohESQAeB6rrYXPcua9M36Pc", &chaincfg.Params{})
	require.Nil(t, err)

	assert.Equal(t, b.Transactions[0].Outputs[0].Address, addrPKH.ScriptAddress())
	assert.NotEqual(t, b.Transactions[0].Outputs[0].Address, differentAddrPKH.ScriptAddress())

	a, err := amount.NewAmount(1000.)
	require.Nil(t, err)

	assert.EqualValues(t, b.Transactions[0].Outputs[0].Value, a)

	assert.Equal(t, "0000001f112450c8d1f24161c41fee37948b5d54fe055f16491bd17bd61ded9d",
		b.HexHash())
	assert.Equal(t, "ac81f1fd432e37f70841e117d4a73f3c0fbf15a722f538a015b08abc3b6947fe",
		hex.EncodeToString(b.Transactions[0].ID))
}

func TestMemoryStorage_ReadBlockByHash(t *testing.T) {
	s, err := NewMemoryStorage()
	require.Nil(t, err)
	defer assert.Nil(t, s.Close(context.Background()))

	gb, err := s.ReadGenesisBlock(context.Background())
	require.Nil(t, err)

	b, err := s.ReadBlockByHash(context.Background(), gb.Hash)
	require.Nil(t, err)
	assert.Equal(t, b, gb)

	b, err = s.ReadBlockByHash(context.Background(), []byte{1, 2, 3})
	require.NotNil(t, err)
	assert.Equal(t, block.Block{}, b)
}

func TestMemoryStorage_ReadLastBlockHash(t *testing.T) {
	s, err := NewMemoryStorage()
	require.Nil(t, err)

	gb, err := s.ReadGenesisBlock(context.Background())
	require.Nil(t, err)

	blockHash, err := s.ReadLastBlockHash(context.Background())
	require.Nil(t, err)
	assert.Equal(t, blockHash, gb.Hash)
}

func TestMemoryStorage_WriteBlock(t *testing.T) {
	s, err := NewMemoryStorage()
	require.Nil(t, err)

	genesisBlock, err := s.ReadGenesisBlock(context.Background())
	require.Nil(t, err)

	lastBlockHash, err := s.ReadLastBlockHash(context.Background())
	require.Nil(t, err)
	assert.Equal(t, lastBlockHash, genesisBlock.Hash)

	block2 := newMockBlock(lastBlockHash)
	s.WriteBlock(context.Background(), block2)

	lastBlockHash, err = s.ReadLastBlockHash(context.Background())
	require.Nil(t, err)
	assert.Equal(t, lastBlockHash, block2.Hash)

	readBlock, err := s.ReadBlockByHash(context.Background(), lastBlockHash)
	require.Nil(t, err)
	assert.Equal(t, block2, readBlock)

	block3 := newMockBlock(lastBlockHash)
	s.WriteBlock(context.Background(), block3)

	lastBlockHash, err = s.ReadLastBlockHash(context.Background())
	require.Nil(t, err)
	assert.Equal(t, lastBlockHash, block3.Hash)

	readBlock, err = s.ReadBlockByHash(context.Background(), lastBlockHash)
	require.Nil(t, err)
	assert.Equal(t, block3, readBlock)
}

func TestMemoryStorage_FindTransactionByID(t *testing.T) {
	s, err := NewMemoryStorage()
	require.Nil(t, err)

	genesisBlock, err := s.ReadGenesisBlock(context.Background())
	require.Nil(t, err)

	lastBlockHash, err := s.ReadLastBlockHash(context.Background())
	require.Nil(t, err)
	assert.Equal(t, lastBlockHash, genesisBlock.Hash)

	block2 := newMockBlock(lastBlockHash)
	s.WriteBlock(context.Background(), block2)
	wantTX := block2.Transactions[0]

	lastBlockHash, err = s.ReadLastBlockHash(context.Background())
	require.Nil(t, err)
	assert.Equal(t, lastBlockHash, block2.Hash)

	readBlock, err := s.ReadBlockByHash(context.Background(), lastBlockHash)
	require.Nil(t, err)
	assert.Equal(t, block2, readBlock)

	block3 := newMockBlock(lastBlockHash)
	s.WriteBlock(context.Background(), block3)

	lastBlockHash, err = s.ReadLastBlockHash(context.Background())
	require.Nil(t, err)
	assert.Equal(t, lastBlockHash, block3.Hash)

	readBlock, err = s.ReadBlockByHash(context.Background(), lastBlockHash)
	require.Nil(t, err)
	assert.Equal(t, block3, readBlock)

	// read foundTX from block2
	_, foundTX, err := s.FindTransactionByID(block2.Transactions[0].ID)
	require.Nil(t, err)
	assert.Equal(t, wantTX, foundTX)
}

func newMockBlock(prevHash []byte) block.Block {
	alice, _ := keys.New()
	bob, _ := keys.New()
	reward, _ := amount.NewAmount(50.)
	coinbaseTX, _ := transaction.NewCoinBase(*alice.AddressPKH, reward)
	amount, _ := amount.NewAmount(15.)
	newTX, _ := transaction.New(transaction.NewUTXOSetFromTX(coinbaseTX), *alice.AddressPKH, *bob.AddressPKH, amount)

	b := block.New(prevHash, []transaction.Transaction{newTX})
	targetBytes, _ := hex.DecodeString("0FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")
	targetInt := new(big.Int).SetBytes(targetBytes)

	s := block.NewDSHA256Solver(targetInt)
	s.Solve(&b)
	return b
}

// Test case
// 1) Alice receives 50 coins as a miner reward ( Alice's UTXO(50); Bob's has no UTXO )
// 2) Alice sends 45 coins to Bob ( Alice's UTXO(5); Bob's UTXO(45) )
// 3) Bob sends 20 coins to Alice ( Alice's UTXO(20,5); Bob's sum UTXO(25) )
// 4) Alice sends 7 coins to Bob ( Alice's UTXO(13,5); Bob's UTXO(25,7) )
// 5) Bob sends 5.5 coins to Alice ( Alice's UTXO(13,5,5.5); Bob's UTXO(25,1.5) )
// 6) Alice sends to Chris 1 coins; Bob sends to Chris 5 coins;
// ( Alice's UTXO(13,5,4.5); Bob's UTXO(21.5); Chris's UTXO(1,5) )
// Todo: refactor it!
func TestMemoryStorage_FindUTXOByPKH(t *testing.T) {
	alice, _ := keys.New()
	bob, _ := keys.New()
	chris, _ := keys.New()
	reward, _ := amount.NewAmount(50.)
	coinbaseTX, _ := transaction.NewCoinBase(*alice.AddressPKH, reward)

	s, _ := NewMemoryStorage()
	lastBlockHash, _ := s.ReadLastBlockHash(context.Background())

	// 1) Alice receives 50 coins as a miner reward ( Alice's UTXO(50); Bob's has no UTXO )
	b := block.New(lastBlockHash, []transaction.Transaction{coinbaseTX})
	targetBytes, _ := hex.DecodeString("0FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")
	targetInt := new(big.Int).SetBytes(targetBytes)
	solver := block.NewDSHA256Solver(targetInt)

	solver.Solve(&b)
	s.WriteBlock(context.Background(), b)

	aliceUTXOSet, err := s.FindUTXOByPKH(context.Background(), *alice.AddressPKH)
	require.Nil(t, err)
	require.Len(t, aliceUTXOSet, 1)
	assert.EqualValues(t, reward, aliceUTXOSet[0].Value)
	assert.Equal(t, alice.AddressPKH.ScriptAddress(), aliceUTXOSet[0].Address)

	bobUTXOSet, err := s.FindUTXOByPKH(context.Background(), *bob.AddressPKH)
	require.Nil(t, err)
	require.Len(t, bobUTXOSet, 0)

	// 2) Alice sends 45 coins to Bob ( Alice's UTXO(5); Bob's UTXO(45) )
	a, _ := amount.NewAmount(45.)
	tx, _ := transaction.New(aliceUTXOSet, *alice.AddressPKH, *bob.AddressPKH, a)

	lastBlockHash, _ = s.ReadLastBlockHash(context.Background())
	b = block.New(lastBlockHash, []transaction.Transaction{tx})
	solver.Solve(&b)
	s.WriteBlock(context.Background(), b)

	aliceUTXOSet, err = s.FindUTXOByPKH(context.Background(), *alice.AddressPKH)
	require.Nil(t, err)
	require.Len(t, aliceUTXOSet, 1)
	assert.EqualValues(t, 500000000, aliceUTXOSet[0].Value)
	assert.Equal(t, alice.AddressPKH.ScriptAddress(), aliceUTXOSet[0].Address)

	bobUTXOSet, err = s.FindUTXOByPKH(context.Background(), *bob.AddressPKH)
	require.Nil(t, err)
	require.Len(t, bobUTXOSet, 1)
	assert.EqualValues(t, a, bobUTXOSet[0].Value)
	assert.Equal(t, bob.AddressPKH.ScriptAddress(), bobUTXOSet[0].Address)

	// 3) Bob sends 20 coins to Alice ( Alice's UTXO(20,5); Bob's sum UTXO(25) )
	a, _ = amount.NewAmount(20.)
	tx, _ = transaction.New(bobUTXOSet, *bob.AddressPKH, *alice.AddressPKH, a)

	lastBlockHash, _ = s.ReadLastBlockHash(context.Background())
	b = block.New(lastBlockHash, []transaction.Transaction{tx})
	solver.Solve(&b)
	s.WriteBlock(context.Background(), b)

	aliceUTXOSet, err = s.FindUTXOByPKH(context.Background(), *alice.AddressPKH)
	require.Nil(t, err)
	require.Len(t, aliceUTXOSet, 2)
	// 20 coins
	assert.EqualValues(t, 2000000000, aliceUTXOSet[0].Value)
	assert.Equal(t, alice.AddressPKH.ScriptAddress(), aliceUTXOSet[0].Address)
	// 5 coins
	assert.EqualValues(t, 500000000, aliceUTXOSet[1].Value)
	assert.Equal(t, alice.AddressPKH.ScriptAddress(), aliceUTXOSet[1].Address)

	total, err := s.TotalUTXOByPKH(context.Background(), *alice.AddressPKH)
	require.Nil(t, err)
	assert.EqualValues(t, 2500000000, total)

	bobUTXOSet, err = s.FindUTXOByPKH(context.Background(), *bob.AddressPKH)
	require.Nil(t, err)
	require.Len(t, bobUTXOSet, 1)
	// 45 coins
	assert.EqualValues(t, 2500000000, bobUTXOSet[0].Value)
	assert.Equal(t, bob.AddressPKH.ScriptAddress(), bobUTXOSet[0].Address)

	// 4) Alice sends 7 coins to Bob ( Alice's UTXO(13,5); Bob's UTXO(25,7) )
	a, _ = amount.NewAmount(7.)
	tx, _ = transaction.New(aliceUTXOSet, *alice.AddressPKH, *bob.AddressPKH, a)

	lastBlockHash, _ = s.ReadLastBlockHash(context.Background())
	b = block.New(lastBlockHash, []transaction.Transaction{tx})
	solver.Solve(&b)
	s.WriteBlock(context.Background(), b)

	aliceUTXOSet, err = s.FindUTXOByPKH(context.Background(), *alice.AddressPKH)
	require.Nil(t, err)
	require.Len(t, aliceUTXOSet, 2)
	// 13 coins
	assert.EqualValues(t, 1300000000, aliceUTXOSet[0].Value)
	assert.Equal(t, alice.AddressPKH.ScriptAddress(), aliceUTXOSet[0].Address)
	// 5 coins
	assert.EqualValues(t, 500000000, aliceUTXOSet[1].Value)
	assert.Equal(t, alice.AddressPKH.ScriptAddress(), aliceUTXOSet[1].Address)

	bobUTXOSet, err = s.FindUTXOByPKH(context.Background(), *bob.AddressPKH)
	require.Nil(t, err)
	require.Len(t, bobUTXOSet, 2)
	// 7 coins
	assert.EqualValues(t, 700000000, bobUTXOSet[0].Value)
	assert.Equal(t, bob.AddressPKH.ScriptAddress(), bobUTXOSet[0].Address)
	// 25 coins
	assert.EqualValues(t, 2500000000, bobUTXOSet[1].Value)
	assert.Equal(t, bob.AddressPKH.ScriptAddress(), bobUTXOSet[1].Address)

	// 5) Bob sends 5.5 coins to Alice ( Alice's UTXO(13,5,5.5); Bob's sum UTXO(25,1.5) )
	a, _ = amount.NewAmount(5.5)
	tx, _ = transaction.New(bobUTXOSet, *bob.AddressPKH, *alice.AddressPKH, a)

	lastBlockHash, _ = s.ReadLastBlockHash(context.Background())
	b = block.New(lastBlockHash, []transaction.Transaction{tx})
	solver.Solve(&b)
	s.WriteBlock(context.Background(), b)

	aliceUTXOSet, err = s.FindUTXOByPKH(context.Background(), *alice.AddressPKH)
	require.Nil(t, err)
	require.Len(t, aliceUTXOSet, 3)
	// 5.5 coins
	assert.EqualValues(t, 550000000, aliceUTXOSet[0].Value)
	assert.Equal(t, alice.AddressPKH.ScriptAddress(), aliceUTXOSet[0].Address)
	// 13 coins
	assert.EqualValues(t, 1300000000, aliceUTXOSet[1].Value)
	assert.Equal(t, alice.AddressPKH.ScriptAddress(), aliceUTXOSet[1].Address)
	// 5 coins
	assert.EqualValues(t, 500000000, aliceUTXOSet[2].Value)
	assert.Equal(t, alice.AddressPKH.ScriptAddress(), aliceUTXOSet[2].Address)

	bobUTXOSet, err = s.FindUTXOByPKH(context.Background(), *bob.AddressPKH)
	require.Nil(t, err)
	require.Len(t, bobUTXOSet, 2)
	// 1.5 coins
	assert.EqualValues(t, 150000000, bobUTXOSet[0].Value)
	assert.Equal(t, bob.AddressPKH.ScriptAddress(), bobUTXOSet[0].Address)
	// 25 coins
	assert.EqualValues(t, 2500000000, bobUTXOSet[1].Value)
	assert.Equal(t, bob.AddressPKH.ScriptAddress(), bobUTXOSet[1].Address)

	// 6) Alice sends to Chris 1 coins; Bob sends to Chris 5 coins;
	// ( Alice's UTXO(13,5,4.5); Bob's UTXO(21.5); Chris's UTXO(1,5) )
	a, _ = amount.NewAmount(1.)
	tx, _ = transaction.New(aliceUTXOSet, *alice.AddressPKH, *chris.AddressPKH, a)

	a, _ = amount.NewAmount(5.)
	tx2, _ := transaction.New(bobUTXOSet, *bob.AddressPKH, *chris.AddressPKH, a)

	lastBlockHash, _ = s.ReadLastBlockHash(context.Background())
	b = block.New(lastBlockHash, []transaction.Transaction{tx, tx2})
	solver.Solve(&b)
	s.WriteBlock(context.Background(), b)

	aliceUTXOSet, err = s.FindUTXOByPKH(context.Background(), *alice.AddressPKH)
	require.Nil(t, err)
	require.Len(t, aliceUTXOSet, 3)
	// 4.5 coins
	assert.EqualValues(t, 450000000, aliceUTXOSet[0].Value)
	assert.Equal(t, alice.AddressPKH.ScriptAddress(), aliceUTXOSet[0].Address)
	// 13 coins
	assert.EqualValues(t, 1300000000, aliceUTXOSet[1].Value)
	assert.Equal(t, alice.AddressPKH.ScriptAddress(), aliceUTXOSet[1].Address)
	// 5 coins
	assert.EqualValues(t, 500000000, aliceUTXOSet[2].Value)
	assert.Equal(t, alice.AddressPKH.ScriptAddress(), aliceUTXOSet[2].Address)

	bobUTXOSet, err = s.FindUTXOByPKH(context.Background(), *bob.AddressPKH)
	require.Nil(t, err)
	require.Len(t, bobUTXOSet, 1)
	// 21.5 coins
	assert.EqualValues(t, 2150000000, bobUTXOSet[0].Value)
	assert.Equal(t, bob.AddressPKH.ScriptAddress(), bobUTXOSet[0].Address)

	chrisUTXOSet, err := s.FindUTXOByPKH(context.Background(), *chris.AddressPKH)
	require.Nil(t, err)
	require.Len(t, chrisUTXOSet, 2)
	// 1 coins
	assert.EqualValues(t, 100000000, chrisUTXOSet[0].Value)
	assert.Equal(t, chris.AddressPKH.ScriptAddress(), chrisUTXOSet[0].Address)
	// 5 coins
	assert.EqualValues(t, 500000000, chrisUTXOSet[1].Value)
	assert.Equal(t, chris.AddressPKH.ScriptAddress(), chrisUTXOSet[1].Address)

}
