package blockchain

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/btcsuite/btcutil"
	"github.com/pkg/errors"
	"github.com/superstas/gcoin/gcoin/block"
	"github.com/superstas/gcoin/gcoin/transaction"
)

// memoryStorage is a very simple non-persistent storage implementation
// for demo purposes only
type memoryStorage struct {
	blocks        []block.Block
	lastBlockHash []byte
	sync.RWMutex
}

// NewMemoryStorage returns memory storage implementation
func NewMemoryStorage() (Storage, error) {
	s := &memoryStorage{
		blocks:        make([]block.Block, 0, 4096),
		lastBlockHash: make([]byte, 0, 64),
	}

	genesisBlock, err := s.ReadGenesisBlock(context.Background())
	if err != nil {
		return nil, err
	}

	s.blocks = append(s.blocks, genesisBlock)
	s.lastBlockHash = genesisBlock.Hash
	return s, nil
}

// ReadGenesisBlock returns the hard-coded genesis block
// This block was generated earlier, you can find a private key for the 1st coinbase TX output in demo/alice_node/wallet.dat
func (s *memoryStorage) ReadGenesisBlock(context.Context) (block.Block, error) {
	const genesisBlockHash = "7b2250726576696f7573426c6f636b48617368223a6e756c6c2c224d6" +
		"5726b6c65526f6f7448617368223a22724948782f554d754e2f63495165" +
		"4558314b632f50412b2f46616369395469674662434b76447470522f343" +
		"d222c2254696d657374616d70223a313532303531383431352c22546172" +
		"676574223a222f2f2f2f2f2f2f2f2f2f2f2f2f2f2f2f2f2f2f2f2f2f2f2" +
		"f2f2f2f2f2f2f2f2f2f2f2f2f2f2f383d222c224e6f6e6365223a393839" +
		"343437342c2248617368223a22414141414878456b554d6a52386b46687" +
		"8422f754e35534c5856542b425638575352765265395964375a303d222c" +
		"225472616e73616374696f6e73223a5b7b224944223a22724948782f554" +
		"d754e2f634951654558314b632f50412b2f46616369395469674662434b" +
		"76447470522f343d222c22496e70757473223a5b7b225472616e7361637" +
		"4696f6e4944223a6e756c6c2c224f7574496e646578223a2d312c225369" +
		"676e223a6e756c6c2c225075624b6579223a2272497858557a7235444e4" +
		"e4e68454744545438495445747372576d7a776174363041366941475075" +
		"4939453d227d5d2c224f757470757473223a5b7b2256616c7565223a313" +
		"0303030303030303030302c2241646472657373223a2251556e6162452b" +
		"79627957556b302b794b2b592f65424e3235736f3d227d5d7d5d7d"

	var b block.Block
	data, err := hex.DecodeString(genesisBlockHash)
	if err != nil {
		return b, err
	}

	return b, json.Unmarshal(data, &b)
}

// WriteBlock writes a given block to the storage
func (s *memoryStorage) WriteBlock(ctx context.Context, b block.Block) error {
	s.Lock()
	defer s.Unlock()

	s.blocks = append(s.blocks, b)
	s.lastBlockHash = b.Hash
	return nil
}

// ReadBlockByHash returns a block by given hash
func (s *memoryStorage) ReadBlockByHash(_ context.Context, hash []byte) (block.Block, error) {
	for _, b := range s.blocks {
		if bytes.Equal(b.Hash, hash) {
			return b, nil
		}
	}
	return block.Block{}, errors.New("block not found")
}

// ReadLastBlockHash returns a last known block hash
func (s *memoryStorage) ReadLastBlockHash(_ context.Context) ([]byte, error) {
	return s.lastBlockHash, nil
}

func (s *memoryStorage) Close(context.Context) error {
	return nil
}

func (s *memoryStorage) FindTransactionByID(TxID []byte) ([]byte, transaction.Transaction, error) {
	for _, b := range s.blocks {
		for _, tx := range b.Transactions {
			if bytes.Equal(tx.ID, TxID) {
				return b.Hash, tx, nil
			}
		}
	}
	return nil, transaction.Transaction{}, errors.New("transaction not found")
}

func (s *memoryStorage) FindUTXOByPKH(_ context.Context, addressPKH btcutil.AddressPubKeyHash) (transaction.UTXOSet, error) {
	// map[TxID]OutIndex
	spentOutputs := make(map[string]int64)
	UTXOs := make([]transaction.UTXO, 0, 64)

	for i := len(s.blocks) - 1; i >= 0; i-- {
		b := s.blocks[i]
		for _, tx := range b.Transactions {
			for _, in := range tx.Inputs {
				if in.OutIndex == -1 {
					continue
				}
				spentOutputs[fmt.Sprintf("%x_%d", in.TransactionID, in.OutIndex)] = in.OutIndex
			}

			for i, out := range tx.Outputs {
				_, ok := spentOutputs[fmt.Sprintf("%x_%d", tx.ID, i)]
				if !ok && bytes.Equal(out.Address, addressPKH.ScriptAddress()) {
					UTXOs = append(UTXOs, transaction.UTXO{
						TxID:     tx.ID,
						OutIndex: int64(i),
						Output:   out,
					})
				}
			}
		}
	}

	return UTXOs, nil
}

func (s *memoryStorage) TotalUTXOByPKH(ctx context.Context, addressPKH btcutil.AddressPubKeyHash) (int64, error) {
	UTXOs, err := s.FindUTXOByPKH(ctx, addressPKH)
	var t int64
	if err != nil {
		return t, err
	}

	for _, u := range UTXOs {
		t += u.Value
	}

	return t, nil
}
