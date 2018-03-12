package block

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"time"

	"github.com/onrik/gomerkle"
	"github.com/superstas/gcoin/gcoin/transaction"
)

// Block represents a simple block
type Block struct {
	Header
	Hash         []byte
	Transactions []transaction.Transaction
}

// HexHash returns hex encoded block hash
func (b Block) HexHash() string {
	return hex.EncodeToString(b.Hash)
}

// Header represents a block header
type Header struct {
	PreviousBlockHash []byte
	MerkleRootHash    []byte
	Timestamp         int64
	Target            []byte
	Nonce             uint64
}

// Serialize serializes header. It used in the solver.
func (h Header) Serialize() []byte {
	buf := append(h.PreviousBlockHash, h.MerkleRootHash...)
	buf = append(buf, h.Target...)
	bLen := len(buf)

	buf = append(buf, make([]byte, binary.MaxVarintLen64*2)...)
	binary.PutVarint(buf[bLen:], h.Timestamp)
	binary.PutUvarint(buf[bLen+binary.MaxVarintLen64:], h.Nonce)
	return buf
}

// NewGenesis creates genesis block with given transactions
// It's used for test only.
// You can get generated genesis block with blockchain.ReadGenesisBlock
func NewGenesis(txs []transaction.Transaction) Block {
	return Block{
		Header: Header{
			MerkleRootHash: merkleRootHash(txs),
			Timestamp:      time.Now().Unix(),
		},
		Transactions: txs,
	}
}

// New creates new block with given transactions. New block points to a previous block.
func New(previousBlockHash []byte, txs []transaction.Transaction) Block {
	return Block{
		Header: Header{
			PreviousBlockHash: previousBlockHash,
			MerkleRootHash:    merkleRootHash(txs),
			Timestamp:         time.Now().Unix(),
		},
		Transactions: txs,
	}
}

// For better understanding usage of the Merkle Tree
// see: https://bitcoin.stackexchange.com/questions/10479/what-is-the-merkle-root
func merkleRootHash(txs []transaction.Transaction) []byte {
	tree := gomerkle.NewTree(sha256.New())
	for _, tx := range txs {
		tree.AddHash(tx.ID)
	}

	tree.Generate()
	return tree.Root()
}
