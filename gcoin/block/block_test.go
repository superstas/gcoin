package block

import (
	"encoding/hex"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeaderSerialize(t *testing.T) {
	prevBlockHash, err := hex.DecodeString("000072bbe055cd287802d4e5e1f11090d449d4e82422b06f90ab84b8fa04e4cc")
	require.Nil(t, err)

	merkleRootHash, err := hex.DecodeString("d1735f3d6ed04ae6a00fdfeff31c4af475e37a5913179186e72433c39808e77d")
	require.Nil(t, err)

	h := Header{
		PreviousBlockHash: prevBlockHash,
		MerkleRootHash:    merkleRootHash,
		Nonce:             math.MaxUint64,
		Target:            []byte{1, 0}, // 256
		Timestamp:         1520194509,
	}

	wantHash := "000072bbe055cd287802d4e5e1f11090d449d4e82422b06f90ab84b8fa04e4ccd1735f3d6ed04ae6a00fdfeff" +
		"31c4af475e37a5913179186e72433c39808e77d01009acfe2a90b0000000000ffffffffffffffffff01"
	assert.Equal(t, wantHash, hex.EncodeToString(h.Serialize()))

	// change the header
	h.Target = []byte{1, 1}
	assert.NotEqual(t, wantHash, hex.EncodeToString(h.Serialize()))
}
