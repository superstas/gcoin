package network

import (
	"context"
	"io"
	"net"
	"testing"

	"github.com/gojuno/minimock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/superstas/gcoin/gcoin/network/message"
	"github.com/superstas/gcoin/gcoin/network/message/mocks"
	"google.golang.org/grpc/peer"
)

func TestPeerManager_AddClient(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:12345")
	require.Nil(t, err)
	c := mocks.NewMessageService_MessageClientMock(mc)
	c.ContextMock.Return(peer.NewContext(context.Background(), &peer.Peer{Addr: tcpAddr}))

	sm := NewPeerManager()
	require.Nil(t, sm.AddClient(c))
	require.Nil(t, sm.AddClient(c))
	require.Len(t, sm.clientStreams, 1)
	require.Len(t, sm.serverStreams, 0)
	sm.RemoveClient("127.0.0.1:12345")
	require.Len(t, sm.clientStreams, 0)
}

func TestPeerManager_AddServer(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:12345")
	require.Nil(t, err)
	s := mocks.NewMessageService_MessageServerMock(mc)
	s.ContextMock.Return(peer.NewContext(context.Background(), &peer.Peer{Addr: tcpAddr}))

	sm := NewPeerManager()
	require.Nil(t, sm.AddServer(s))
	require.Nil(t, sm.AddServer(s))
	require.Len(t, sm.serverStreams, 1)
	require.Len(t, sm.clientStreams, 0)
	sm.RemoveServer("127.0.0.1:12345")
	require.Len(t, sm.serverStreams, 0)
}

func TestPeerManager_SendBroadcastMsg(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	serverAddr, _ := net.ResolveTCPAddr("tcp", "localhost:12345")
	clientAddr, _ := net.ResolveTCPAddr("tcp", "localhost:41232")

	wantMsg := NewAddTXMessage([]byte("test"))

	s := mocks.NewMessageService_MessageServerMock(mc)
	s.ContextMock.Return(peer.NewContext(context.Background(), &peer.Peer{Addr: serverAddr}))
	s.SendFunc = func(m *message.Msg) (r error) {
		assert.Equal(t, wantMsg, m)
		return nil
	}

	c := mocks.NewMessageService_MessageClientMock(mc)
	c.ContextMock.Return(peer.NewContext(context.Background(), &peer.Peer{Addr: clientAddr}))
	c.SendFunc = func(m *message.Msg) (r error) {
		assert.Equal(t, wantMsg, m)
		return nil
	}

	sm := NewPeerManager()
	require.Nil(t, sm.AddServer(s))
	require.Nil(t, sm.AddClient(c))
	sm.Send(context.Background(), wantMsg)
}

func TestPeerManager_SendBroadcastMsg_OneOfServerFailed(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	serverAddr, _ := net.ResolveTCPAddr("tcp", "localhost:12345")
	clientAddr, _ := net.ResolveTCPAddr("tcp", "localhost:41232")

	wantMsg := NewAddTXMessage([]byte("test"))

	s := mocks.NewMessageService_MessageServerMock(mc)
	s.ContextMock.Return(peer.NewContext(context.Background(), &peer.Peer{Addr: serverAddr}))
	s.SendMock.Return(io.EOF)

	c := mocks.NewMessageService_MessageClientMock(mc)
	c.ContextMock.Return(peer.NewContext(context.Background(), &peer.Peer{Addr: clientAddr}))
	c.SendFunc = func(m *message.Msg) (r error) {
		assert.Equal(t, wantMsg, m)
		return nil
	}

	sm := NewPeerManager()
	require.Nil(t, sm.AddServer(s))
	require.Nil(t, sm.AddClient(c))
	require.Len(t, sm.serverStreams, 1)
	require.Len(t, sm.clientStreams, 1)
	sm.Send(context.Background(), wantMsg)
	require.Len(t, sm.serverStreams, 0)
	require.Len(t, sm.clientStreams, 1)
}

func TestPeerManager_SendBroadcastMsg_OneOfClientFailed(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	serverAddr, _ := net.ResolveTCPAddr("tcp", "localhost:12345")
	clientAddr, _ := net.ResolveTCPAddr("tcp", "localhost:41232")

	wantMsg := NewAddTXMessage([]byte("test"))

	s := mocks.NewMessageService_MessageServerMock(mc)
	s.ContextMock.Return(peer.NewContext(context.Background(), &peer.Peer{Addr: serverAddr}))
	s.SendFunc = func(m *message.Msg) (r error) {
		assert.Equal(t, wantMsg, m)
		return nil
	}

	c := mocks.NewMessageService_MessageClientMock(mc)
	c.ContextMock.Return(peer.NewContext(context.Background(), &peer.Peer{Addr: clientAddr}))
	c.SendMock.Return(io.EOF)

	sm := NewPeerManager()
	require.Nil(t, sm.AddServer(s))
	require.Nil(t, sm.AddClient(c))
	require.Len(t, sm.serverStreams, 1)
	require.Len(t, sm.clientStreams, 1)
	sm.Send(context.Background(), wantMsg)
	require.Len(t, sm.serverStreams, 1)
	require.Len(t, sm.clientStreams, 0)
}
