package network

import (
	"context"
	"log"
	"sync"

	"github.com/pkg/errors"
	"github.com/superstas/gcoin/gcoin/network/message"
	"google.golang.org/grpc/peer"
)

// PeerManager represents a manager that knows all about the node connections
type PeerManager struct {
	serverStreams map[string]message.MessageService_MessageServer
	clientStreams map[string]message.MessageService_MessageClient
	sync.Mutex
}

// NewPeerManager return a new peer manager
func NewPeerManager() *PeerManager {
	return &PeerManager{
		serverStreams: make(map[string]message.MessageService_MessageServer, 8),
		clientStreams: make(map[string]message.MessageService_MessageClient, 8),
	}
}

// AddClient adds a client stream
func (m *PeerManager) AddClient(c message.MessageService_MessageClient) error {
	p, ok := peer.FromContext(c.Context())
	if !ok {
		return errors.New("failed to get client peer info")
	}

	m.Lock()
	m.clientStreams[p.Addr.String()] = c
	m.Unlock()
	return nil
}

// RemoveClient removes a client stream
func (m *PeerManager) RemoveClient(addr string) {
	m.Lock()
	delete(m.clientStreams, addr)
	m.Unlock()
}

// AddServer adds a server stream
func (m *PeerManager) AddServer(s message.MessageService_MessageServer) error {
	p, ok := peer.FromContext(s.Context())
	if !ok {
		return errors.New("failed to get server peer info")
	}
	log.Printf("[network]: a new client %q connected\n", p.Addr)

	m.Lock()
	m.serverStreams[p.Addr.String()] = s
	m.Unlock()
	return nil
}

// RemoveServer removes a server stream
func (m *PeerManager) RemoveServer(addr string) {
	m.Lock()
	delete(m.serverStreams, addr)
	m.Unlock()
}

// Send sends a given message to all known peers except a peer that sent this message to the node.
func (m *PeerManager) Send(ctx context.Context, msg *message.Msg) {
	peerFrom, _ := peer.FromContext(ctx)

	for addr, s := range m.serverStreams {
		if peerFrom != nil && addr == peerFrom.Addr.String() {
			continue
		}

		if err := s.Send(msg); err != nil {
			m.RemoveServer(addr)
		}
	}

	for addr, c := range m.clientStreams {
		if peerFrom != nil && addr == peerFrom.Addr.String() {
			continue
		}

		if err := c.Send(msg); err != nil {
			m.RemoveClient(addr)
		}
	}
}
