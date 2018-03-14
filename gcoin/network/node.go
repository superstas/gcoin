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
func (n *Node) Message(msgStream message.MessageService_MessageServer) error {
	n.PeerManager.AddServer(msgStream)
	for {
		in, err := msgStream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}
		n.handleMsg(msgStream.Context(), in)
	}
}

// ServeClient starts listening a given client stream
func (n *Node) ServeClient(stream message.MessageService_MessageClient) {
	n.PeerManager.AddClient(stream)
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Printf("[network]: failed to receive message: %v\n", err)
			return
		}

		n.handleMsg(stream.Context(), in)
	}
}

func (n *Node) handleMsg(ctx context.Context, msg *message.Msg) {
	p, _ := peer.FromContext(ctx)
	log.Printf("[network]: new incoming %q msg from peer %s\n", msg.Type, p.Addr)
	switch msg.Type {
	case messageAddBlock:
		if err := n.handleBlockMsg(ctx, msg); err != nil {
			log.Printf("[network]: failed to handle %q msg: %v\n", messageAddBlock, err)
		}
		n.PeerManager.Send(ctx, msg)
	case messageAddTx:
		if err := n.handleTXMsg(ctx, msg); err != nil {
			log.Printf("[network]: failed to handle %q msg: %v\n", messageAddTx, err)
		}
		n.PeerManager.Send(ctx, msg)
	}
}

func (n *Node) handleBlockMsg(ctx context.Context, msg *message.Msg) error {
	var b block.Block
	err := json.Unmarshal(msg.Data, &b)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal block data")
	}

	if err := n.Validator.ValidateBlock(ctx, b); err != nil {
		return errors.Wrap(err, "failed to validate block")
	}

	for _, tx := range b.Transactions {
		n.MemPool.DeleteByID(tx.HexID())
	}

	return errors.Wrap(n.Storage.WriteBlock(ctx, b), "failed to write block")
}

func (n *Node) handleTXMsg(ctx context.Context, msg *message.Msg) error {
	var tx transaction.Transaction
	err := json.Unmarshal(msg.Data, &tx)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal TX data")
	}

	if err := n.Validator.ValidateTX(ctx, tx); err != nil {
		return err
	}

	n.MemPool.Add(tx)
	return nil
}

// Send sends a TX message to the network
func (n *Node) Send(ctx context.Context, r *cli.SendRequest) (*cli.SendResponse, error) {
	// TODO: UTXO should be found by all existed addresses
	from := n.Wallet.Keys()[0]
	utxos, err := n.Storage.FindUTXOByPKH(ctx, *from.AddressPKH)
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

	if err := n.Signer.Sign(&tx, from.PrivateKey); err != nil {
		return nil, errors.Wrap(err, "failed to sign")
	}

	if err := n.Validator.ValidateTX(ctx, tx); err != nil {
		return nil, errors.Wrap(err, "failed to validate tx")
	}

	n.MemPool.Add(tx)
	txData, err := json.Marshal(tx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal tx")
	}

	n.PeerManager.Send(ctx, NewAddTXMessage(txData))

	return &cli.SendResponse{
		TxId: tx.HexID(),
	}, nil
}

// GetBalance returns sum of all UTXO of all addresses in the wallet
func (n *Node) GetBalance(ctx context.Context, in *cli.GetBalanceRequest) (*cli.GetBalanceResponse, error) {
	keys := n.Wallet.Keys()
	if len(keys) == 0 {
		return nil, errors.New("failed to find keys in the wallet")
	}
	r := &cli.GetBalanceResponse{
		Addresses: make([]*cli.Balance, 0, len(keys)),
	}

	for _, k := range n.Wallet.Keys() {
		total, err := n.Storage.TotalUTXOByPKH(ctx, *k.AddressPKH)
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
