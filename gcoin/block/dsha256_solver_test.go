package block

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/superstas/gcoin/gcoin/amount"
	"github.com/superstas/gcoin/gcoin/keys"
	"github.com/superstas/gcoin/gcoin/transaction"
)

func TestDSHA256Solver_Solve(t *testing.T) {
	alice, err := keys.New()
	require.Nil(t, err)

	reward, err := amount.NewAmount(50.)
	require.Nil(t, err)

	coinbaseTX, err := transaction.NewCoinBase(*alice.AddressPKH, reward)
	require.Nil(t, err)

	targetBytes, err := hex.DecodeString("00FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")
	require.Nil(t, err)
	targetInt := new(big.Int).SetBytes(targetBytes)

	s := NewDSHA256Solver(targetInt)
	gb := NewGenesis([]transaction.Transaction{coinbaseTX})
	// mine the block
	err = s.Solve(&gb)
	require.Nil(t, err)
	assert.Equal(t, targetInt.Bytes(), gb.Target)
	require.Nil(t, s.Verify(gb))

	// break the target
	gb.Target = []byte{1, 0} // 256
	require.NotNil(t, s.Verify(gb))
}
