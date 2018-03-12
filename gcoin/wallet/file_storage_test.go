package wallet

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/superstas/gcoin/gcoin/keys"
)

func TestNewFileStorage(t *testing.T) {
	s := NewFileStorage("")
	require.NotNil(t, s)
	filename := "wallet.dat"
	defer tearDown(t, filename)

	// check that file exists
	_, err := os.Stat(filename)
	require.Nil(t, err)

	// check that file read again
	s2 := NewFileStorage("wallet.dat")
	require.NotNil(t, s2)

	require.Nil(t, s.Close())
	require.Nil(t, s2.Close())
}

func TestFileStorage_Init(t *testing.T) {
	filename := "wallet.dat"
	s := NewFileStorage(filename)
	defer tearDown(t, filename)

	w, err := s.Init(context.Background())
	require.Nil(t, err)
	require.Len(t, w.Keys(), defaultKeysAmount)
	require.Nil(t, s.Close())
}

func TestFileStorage_Write(t *testing.T) {
	filename := "wallet.dat"
	s := NewFileStorage(filename)
	defer tearDown(t, filename)

	w, err := s.Init(context.Background())
	require.Nil(t, err)
	require.Len(t, w.Keys(), defaultKeysAmount)

	k, err := keys.New()
	require.Nil(t, err)
	require.NotNil(t, k)

	w.AddKeys(k)
	require.Len(t, w.Keys(), defaultKeysAmount+1)
	err = s.Write(context.Background(), w)
	require.Nil(t, err)
	require.Nil(t, s.Close())

	s = NewFileStorage(filename)
	w, err = s.Init(context.Background())
	require.Nil(t, err)
	require.Len(t, w.Keys(), defaultKeysAmount+1)

	k, err = keys.New()
	require.Nil(t, err)
	require.NotNil(t, k)
	w.AddKeys(k)
	require.Len(t, w.Keys(), defaultKeysAmount+2)
	err = s.Write(context.Background(), w)
	require.Nil(t, err)
	require.Nil(t, s.Close())

	s = NewFileStorage(filename)
	w, err = s.Init(context.Background())
	require.Nil(t, err)
	require.Len(t, w.Keys(), defaultKeysAmount+2)

}

func tearDown(t *testing.T, filename string) {
	err := os.Remove(filename)
	require.Nil(t, err)

	_, err = os.Stat(filename)
	require.NotNil(t, err)
}

//
//func TestGenerateGenesisBlock(t *testing.T) {
//	s := NewFileStorage("alice_wallet.dat")
//	//defer tearDown(t, "alice_wallet.dat")
//	w, err := s.Init(context.Background())
//	require.Nil(t, err)
//	require.Len(t, w.Keys(), defaultKeysAmount)
//	s.Write(context.Background(), w)
//
//	addr := w.Addresses()[0]
//	rew, err := amount.NewAmount(1000.)
//	require.Nil(t, err)
//
//	coinbaseTX, err := transaction.NewCoinBase(*addr, rew)
//	require.Nil(t, err)
//
//	genesisBlock := block.NewGenesis([]transaction.Transaction{coinbaseTX})
//
//	targetBytes, err := hex.DecodeString("000000FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")
//	require.Nil(t, err)
//
//	solver := block.NewDSHA256Solver(new(big.Int).SetBytes(targetBytes))
//	err = solver.Solve(&genesisBlock)
//	require.Nil(t, err)
//
//	d, _ := json.Marshal(genesisBlock)
//	fmt.Println(hex.EncodeToString(d))
//}
