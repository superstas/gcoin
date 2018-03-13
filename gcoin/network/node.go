package network

import (
	"context"
	"encoding/json"
	"io"
	"log"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/pkg/errors"
	"github.com/superstas/gcoin/gcoin/amount"
	"github.com/superstas/gcoin/gcoin/block"
	"github.com/superstas/gcoin/gcoin/blockchain"
	"github.com/superstas/gcoin/gcoin/consensus"
	"github.com/superstas/gcoin/gcoin/mempool"
	"github.com/superstas/gcoin/gcoin/network/cli"
	"github.com/superstas/gcoin/gcoin/network/message"
	"github.com/superstas/gcoin/gcoin/transaction"
	"github.com/superstas/gcoin/gcoin/wallet"
	"google.golang.org/grpc/peer"
)

// Node represents a single node that can send/relay/validate messages.
// It has access to the wallet for GetBalance command.
type Node struct {
	PeerManager *PeerManager
	MemPool     *mempool.MemPool
	Storage     blockchain.Storage
	Signer      transaction.Signer
	Validator   consensus.Validator
	Wallet      *wallet.Wallet
}

// Message starts listening a given server stream
func (s *Node) Message(msgStream message.MessageService_MessageServer) error {
	s.PeerManager.AddServer(msgStream)
	for {
		in, err := msgStream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}
		s.handleMsg(msgStream.Context(), in)
	}
}

// ServeClient starts listening a given client stream
func (s *Node) ServeClient(stream message.MessageService_MessageClient) {
	s.PeerManager.AddClient(stream)
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Printf("[network]: failed to receive message: %v\n", err)
			return
		}

		s.handleMsg(stream.Context(), in)
	}
}

func (s *Node) handleMsg(ctx context.Context, msg *message.Msg) {
	p, _ := peer.FromContext(ctx)
	log.Printf("[network]: new incoming %q msg from peer %s\n", msg.Type, p.Addr)
	switch msg.Type {
	case messageAddBlock:
		if err := s.handleBlockMsg(ctx, msg); err != nil {
			log.Printf("[network]: failed to handle %q msg: %v\n", messageAddBlock, err)
		}
		s.PeerManager.Send(ctx, msg)
	case messageAddTx:
		if err := s.handleTXMsg(ctx, msg); err != nil {
			log.Printf("[network]: failed to handle %q msg: %v\n", messageAddTx, err)
		}
		s.PeerManager.Send(ctx, msg)
	}
}

func (s *Node) handleBlockMsg(ctx context.Context, msg *message.Msg) error {
	var b block.Block
	err := json.Unmarshal(msg.Data, &b)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal block data")
	}

	if err := s.Validator.ValidateBlock(ctx, b); err != nil {
		return errors.Wrap(err, "failed to validate block")
	}

	for _, tx := range b.Transactions {
		if err := s.MemPool.DeleteByID(tx.HexID()); err != nil {
			log.Printf("[mempool]: failed to remove TX %q\n", tx.HexID())
		}
	}

	return errors.Wrap(s.Storage.WriteBlock(ctx, b), "failed to write block")
}

func (s *Node) handleTXMsg(ctx context.Context, msg *message.Msg) error {
	var tx transaction.Transaction
	err := json.Unmarshal(msg.Data, &tx)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal TX data")
	}

	if err := s.Validator.ValidateTX(ctx, tx); err != nil {
		return errors.Wrap(err, "failed to validate TX")
	}

	s.MemPool.Add(tx)
	return nil
}

// Send sends a TX message to the network
func (s *Node) Send(ctx context.Context, r *cli.SendRequest) (*cli.SendResponse, error) {
	// TODO: UTXO should be found by all existed addresses
	from := s.Wallet.Keys()[0]
	utxos, err := s.Storage.FindUTXOByPKH(ctx, *from.AddressPKH)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find UTXO")
	}

	toAddr, err := btcutil.DecodeAddress(r.ToAddress, &chaincfg.Params{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse address")
	}

	toAddrPKH, ok := toAddr.(*btcutil.AddressPubKeyHash)
	if !ok {
		return nil, errors.Wrap(err, "this is not PKH address")
	}

	tx, err := transaction.New(utxos, *from.AddressPKH, *toAddrPKH, amount.Amount(r.Amount))
	if err != nil {
		return nil, err
	}

	if err := s.Signer.Sign(&tx, from.PrivateKey); err != nil {
		return nil, errors.Wrap(err, "failed to sign")
	}

	if err := s.Validator.ValidateTX(ctx, tx); err != nil {
		return nil, errors.Wrap(err, "failed to validate tx")
	}

	s.MemPool.Add(tx)
	txData, err := json.Marshal(tx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal tx")
	}

	s.PeerManager.Send(ctx, NewAddTXMessage(txData))

	return &cli.SendResponse{
		TxId: tx.HexID(),
	}, nil
}

// GetBalance returns sum of all UTXO of all addresses in the wallet
func (s *Node) GetBalance(ctx context.Context, in *cli.GetBalanceRequest) (*cli.GetBalanceResponse, error) {
	keys := s.Wallet.Keys()
	if len(keys) == 0 {
		return nil, errors.New("failed to find keys in the wallet")
	}
	r := &cli.GetBalanceResponse{
		Addresses: make([]*cli.Balance, 0, len(keys)),
	}

	for _, k := range s.Wallet.Keys() {
		total, err := s.Storage.TotalUTXOByPKH(ctx, *k.AddressPKH)
		if err != nil {
			errors.Wrapf(err, "failed to read total UTXO by %s", k.AddressPKH.EncodeAddress())
		}

		r.Addresses = append(r.Addresses, &cli.Balance{
			Address: k.AddressPKH.EncodeAddress(),
			Balance: total,
		})
	}
	return r, nil
}
