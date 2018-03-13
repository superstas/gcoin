package mocks

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.8
The original interface "MessageService_MessageClient" can be found in github.com/superstas/gcoin/gcoin/network/message
*/
import (
	"sync/atomic"
	"time"

	"github.com/gojuno/minimock"
	message "github.com/superstas/gcoin/gcoin/network/message"
	context "golang.org/x/net/context"
	metadata "google.golang.org/grpc/metadata"

	testify_assert "github.com/stretchr/testify/assert"
)

//MessageService_MessageClientMock implements github.com/superstas/gcoin/gcoin/network/message.MessageService_MessageClient
type MessageService_MessageClientMock struct {
	t minimock.Tester

	CloseSendFunc       func() (r error)
	CloseSendCounter    uint64
	CloseSendPreCounter uint64
	CloseSendMock       mMessageService_MessageClientMockCloseSend

	ContextFunc       func() (r context.Context)
	ContextCounter    uint64
	ContextPreCounter uint64
	ContextMock       mMessageService_MessageClientMockContext

	HeaderFunc       func() (r metadata.MD, r1 error)
	HeaderCounter    uint64
	HeaderPreCounter uint64
	HeaderMock       mMessageService_MessageClientMockHeader

	RecvFunc       func() (r *message.Msg, r1 error)
	RecvCounter    uint64
	RecvPreCounter uint64
	RecvMock       mMessageService_MessageClientMockRecv

	RecvMsgFunc       func(p interface{}) (r error)
	RecvMsgCounter    uint64
	RecvMsgPreCounter uint64
	RecvMsgMock       mMessageService_MessageClientMockRecvMsg

	SendFunc       func(p *message.Msg) (r error)
	SendCounter    uint64
	SendPreCounter uint64
	SendMock       mMessageService_MessageClientMockSend

	SendMsgFunc       func(p interface{}) (r error)
	SendMsgCounter    uint64
	SendMsgPreCounter uint64
	SendMsgMock       mMessageService_MessageClientMockSendMsg

	TrailerFunc       func() (r metadata.MD)
	TrailerCounter    uint64
	TrailerPreCounter uint64
	TrailerMock       mMessageService_MessageClientMockTrailer
}

//NewMessageService_MessageClientMock returns a mock for github.com/superstas/gcoin/gcoin/network/message.MessageService_MessageClient
func NewMessageService_MessageClientMock(t minimock.Tester) *MessageService_MessageClientMock {
	m := &MessageService_MessageClientMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CloseSendMock = mMessageService_MessageClientMockCloseSend{mock: m}
	m.ContextMock = mMessageService_MessageClientMockContext{mock: m}
	m.HeaderMock = mMessageService_MessageClientMockHeader{mock: m}
	m.RecvMock = mMessageService_MessageClientMockRecv{mock: m}
	m.RecvMsgMock = mMessageService_MessageClientMockRecvMsg{mock: m}
	m.SendMock = mMessageService_MessageClientMockSend{mock: m}
	m.SendMsgMock = mMessageService_MessageClientMockSendMsg{mock: m}
	m.TrailerMock = mMessageService_MessageClientMockTrailer{mock: m}

	return m
}

type mMessageService_MessageClientMockCloseSend struct {
	mock *MessageService_MessageClientMock
}

//Return sets up a mock for MessageService_MessageClient.CloseSend to return Return's arguments
func (m *mMessageService_MessageClientMockCloseSend) Return(r error) *MessageService_MessageClientMock {
	m.mock.CloseSendFunc = func() error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of MessageService_MessageClient.CloseSend method
func (m *mMessageService_MessageClientMockCloseSend) Set(f func() (r error)) *MessageService_MessageClientMock {
	m.mock.CloseSendFunc = f
	return m.mock
}

//CloseSend implements github.com/superstas/gcoin/gcoin/network/message.MessageService_MessageClient interface
func (m *MessageService_MessageClientMock) CloseSend() (r error) {
	atomic.AddUint64(&m.CloseSendPreCounter, 1)
	defer atomic.AddUint64(&m.CloseSendCounter, 1)

	if m.CloseSendFunc == nil {
		m.t.Fatal("Unexpected call to MessageService_MessageClientMock.CloseSend")
		return
	}

	return m.CloseSendFunc()
}

//CloseSendMinimockCounter returns a count of MessageService_MessageClientMock.CloseSendFunc invocations
func (m *MessageService_MessageClientMock) CloseSendMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.CloseSendCounter)
}

//CloseSendMinimockPreCounter returns the value of MessageService_MessageClientMock.CloseSend invocations
func (m *MessageService_MessageClientMock) CloseSendMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.CloseSendPreCounter)
}

type mMessageService_MessageClientMockContext struct {
	mock *MessageService_MessageClientMock
}

//Return sets up a mock for MessageService_MessageClient.Context to return Return's arguments
func (m *mMessageService_MessageClientMockContext) Return(r context.Context) *MessageService_MessageClientMock {
	m.mock.ContextFunc = func() context.Context {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of MessageService_MessageClient.Context method
func (m *mMessageService_MessageClientMockContext) Set(f func() (r context.Context)) *MessageService_MessageClientMock {
	m.mock.ContextFunc = f
	return m.mock
}

//Context implements github.com/superstas/gcoin/gcoin/network/message.MessageService_MessageClient interface
func (m *MessageService_MessageClientMock) Context() (r context.Context) {
	atomic.AddUint64(&m.ContextPreCounter, 1)
	defer atomic.AddUint64(&m.ContextCounter, 1)

	if m.ContextFunc == nil {
		m.t.Fatal("Unexpected call to MessageService_MessageClientMock.Context")
		return
	}

	return m.ContextFunc()
}

//ContextMinimockCounter returns a count of MessageService_MessageClientMock.ContextFunc invocations
func (m *MessageService_MessageClientMock) ContextMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.ContextCounter)
}

//ContextMinimockPreCounter returns the value of MessageService_MessageClientMock.Context invocations
func (m *MessageService_MessageClientMock) ContextMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.ContextPreCounter)
}

type mMessageService_MessageClientMockHeader struct {
	mock *MessageService_MessageClientMock
}

//Return sets up a mock for MessageService_MessageClient.Header to return Return's arguments
func (m *mMessageService_MessageClientMockHeader) Return(r metadata.MD, r1 error) *MessageService_MessageClientMock {
	m.mock.HeaderFunc = func() (metadata.MD, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of MessageService_MessageClient.Header method
func (m *mMessageService_MessageClientMockHeader) Set(f func() (r metadata.MD, r1 error)) *MessageService_MessageClientMock {
	m.mock.HeaderFunc = f
	return m.mock
}

//Header implements github.com/superstas/gcoin/gcoin/network/message.MessageService_MessageClient interface
func (m *MessageService_MessageClientMock) Header() (r metadata.MD, r1 error) {
	atomic.AddUint64(&m.HeaderPreCounter, 1)
	defer atomic.AddUint64(&m.HeaderCounter, 1)

	if m.HeaderFunc == nil {
		m.t.Fatal("Unexpected call to MessageService_MessageClientMock.Header")
		return
	}

	return m.HeaderFunc()
}

//HeaderMinimockCounter returns a count of MessageService_MessageClientMock.HeaderFunc invocations
func (m *MessageService_MessageClientMock) HeaderMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.HeaderCounter)
}

//HeaderMinimockPreCounter returns the value of MessageService_MessageClientMock.Header invocations
func (m *MessageService_MessageClientMock) HeaderMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.HeaderPreCounter)
}

type mMessageService_MessageClientMockRecv struct {
	mock *MessageService_MessageClientMock
}

//Return sets up a mock for MessageService_MessageClient.Recv to return Return's arguments
func (m *mMessageService_MessageClientMockRecv) Return(r *message.Msg, r1 error) *MessageService_MessageClientMock {
	m.mock.RecvFunc = func() (*message.Msg, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of MessageService_MessageClient.Recv method
func (m *mMessageService_MessageClientMockRecv) Set(f func() (r *message.Msg, r1 error)) *MessageService_MessageClientMock {
	m.mock.RecvFunc = f
	return m.mock
}

//Recv implements github.com/superstas/gcoin/gcoin/network/message.MessageService_MessageClient interface
func (m *MessageService_MessageClientMock) Recv() (r *message.Msg, r1 error) {
	atomic.AddUint64(&m.RecvPreCounter, 1)
	defer atomic.AddUint64(&m.RecvCounter, 1)

	if m.RecvFunc == nil {
		m.t.Fatal("Unexpected call to MessageService_MessageClientMock.Recv")
		return
	}

	return m.RecvFunc()
}

//RecvMinimockCounter returns a count of MessageService_MessageClientMock.RecvFunc invocations
func (m *MessageService_MessageClientMock) RecvMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.RecvCounter)
}

//RecvMinimockPreCounter returns the value of MessageService_MessageClientMock.Recv invocations
func (m *MessageService_MessageClientMock) RecvMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.RecvPreCounter)
}

type mMessageService_MessageClientMockRecvMsg struct {
	mock             *MessageService_MessageClientMock
	mockExpectations *MessageService_MessageClientMockRecvMsgParams
}

//MessageService_MessageClientMockRecvMsgParams represents input parameters of the MessageService_MessageClient.RecvMsg
type MessageService_MessageClientMockRecvMsgParams struct {
	p interface{}
}

//Expect sets up expected params for the MessageService_MessageClient.RecvMsg
func (m *mMessageService_MessageClientMockRecvMsg) Expect(p interface{}) *mMessageService_MessageClientMockRecvMsg {
	m.mockExpectations = &MessageService_MessageClientMockRecvMsgParams{p}
	return m
}

//Return sets up a mock for MessageService_MessageClient.RecvMsg to return Return's arguments
func (m *mMessageService_MessageClientMockRecvMsg) Return(r error) *MessageService_MessageClientMock {
	m.mock.RecvMsgFunc = func(p interface{}) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of MessageService_MessageClient.RecvMsg method
func (m *mMessageService_MessageClientMockRecvMsg) Set(f func(p interface{}) (r error)) *MessageService_MessageClientMock {
	m.mock.RecvMsgFunc = f
	return m.mock
}

//RecvMsg implements github.com/superstas/gcoin/gcoin/network/message.MessageService_MessageClient interface
func (m *MessageService_MessageClientMock) RecvMsg(p interface{}) (r error) {
	atomic.AddUint64(&m.RecvMsgPreCounter, 1)
	defer atomic.AddUint64(&m.RecvMsgCounter, 1)

	if m.RecvMsgMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.RecvMsgMock.mockExpectations, MessageService_MessageClientMockRecvMsgParams{p},
			"MessageService_MessageClient.RecvMsg got unexpected parameters")

		if m.RecvMsgFunc == nil {

			m.t.Fatal("No results are set for the MessageService_MessageClientMock.RecvMsg")

			return
		}
	}

	if m.RecvMsgFunc == nil {
		m.t.Fatal("Unexpected call to MessageService_MessageClientMock.RecvMsg")
		return
	}

	return m.RecvMsgFunc(p)
}

//RecvMsgMinimockCounter returns a count of MessageService_MessageClientMock.RecvMsgFunc invocations
func (m *MessageService_MessageClientMock) RecvMsgMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.RecvMsgCounter)
}

//RecvMsgMinimockPreCounter returns the value of MessageService_MessageClientMock.RecvMsg invocations
func (m *MessageService_MessageClientMock) RecvMsgMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.RecvMsgPreCounter)
}

type mMessageService_MessageClientMockSend struct {
	mock             *MessageService_MessageClientMock
	mockExpectations *MessageService_MessageClientMockSendParams
}

//MessageService_MessageClientMockSendParams represents input parameters of the MessageService_MessageClient.Send
type MessageService_MessageClientMockSendParams struct {
	p *message.Msg
}

//Expect sets up expected params for the MessageService_MessageClient.Send
func (m *mMessageService_MessageClientMockSend) Expect(p *message.Msg) *mMessageService_MessageClientMockSend {
	m.mockExpectations = &MessageService_MessageClientMockSendParams{p}
	return m
}

//Return sets up a mock for MessageService_MessageClient.Send to return Return's arguments
func (m *mMessageService_MessageClientMockSend) Return(r error) *MessageService_MessageClientMock {
	m.mock.SendFunc = func(p *message.Msg) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of MessageService_MessageClient.Send method
func (m *mMessageService_MessageClientMockSend) Set(f func(p *message.Msg) (r error)) *MessageService_MessageClientMock {
	m.mock.SendFunc = f
	return m.mock
}

//Send implements github.com/superstas/gcoin/gcoin/network/message.MessageService_MessageClient interface
func (m *MessageService_MessageClientMock) Send(p *message.Msg) (r error) {
	atomic.AddUint64(&m.SendPreCounter, 1)
	defer atomic.AddUint64(&m.SendCounter, 1)

	if m.SendMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.SendMock.mockExpectations, MessageService_MessageClientMockSendParams{p},
			"MessageService_MessageClient.Send got unexpected parameters")

		if m.SendFunc == nil {

			m.t.Fatal("No results are set for the MessageService_MessageClientMock.Send")

			return
		}
	}

	if m.SendFunc == nil {
		m.t.Fatal("Unexpected call to MessageService_MessageClientMock.Send")
		return
	}

	return m.SendFunc(p)
}

//SendMinimockCounter returns a count of MessageService_MessageClientMock.SendFunc invocations
func (m *MessageService_MessageClientMock) SendMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.SendCounter)
}

//SendMinimockPreCounter returns the value of MessageService_MessageClientMock.Send invocations
func (m *MessageService_MessageClientMock) SendMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.SendPreCounter)
}

type mMessageService_MessageClientMockSendMsg struct {
	mock             *MessageService_MessageClientMock
	mockExpectations *MessageService_MessageClientMockSendMsgParams
}

//MessageService_MessageClientMockSendMsgParams represents input parameters of the MessageService_MessageClient.SendMsg
type MessageService_MessageClientMockSendMsgParams struct {
	p interface{}
}

//Expect sets up expected params for the MessageService_MessageClient.SendMsg
func (m *mMessageService_MessageClientMockSendMsg) Expect(p interface{}) *mMessageService_MessageClientMockSendMsg {
	m.mockExpectations = &MessageService_MessageClientMockSendMsgParams{p}
	return m
}

//Return sets up a mock for MessageService_MessageClient.SendMsg to return Return's arguments
func (m *mMessageService_MessageClientMockSendMsg) Return(r error) *MessageService_MessageClientMock {
	m.mock.SendMsgFunc = func(p interface{}) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of MessageService_MessageClient.SendMsg method
func (m *mMessageService_MessageClientMockSendMsg) Set(f func(p interface{}) (r error)) *MessageService_MessageClientMock {
	m.mock.SendMsgFunc = f
	return m.mock
}

//SendMsg implements github.com/superstas/gcoin/gcoin/network/message.MessageService_MessageClient interface
func (m *MessageService_MessageClientMock) SendMsg(p interface{}) (r error) {
	atomic.AddUint64(&m.SendMsgPreCounter, 1)
	defer atomic.AddUint64(&m.SendMsgCounter, 1)

	if m.SendMsgMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.SendMsgMock.mockExpectations, MessageService_MessageClientMockSendMsgParams{p},
			"MessageService_MessageClient.SendMsg got unexpected parameters")

		if m.SendMsgFunc == nil {

			m.t.Fatal("No results are set for the MessageService_MessageClientMock.SendMsg")

			return
		}
	}

	if m.SendMsgFunc == nil {
		m.t.Fatal("Unexpected call to MessageService_MessageClientMock.SendMsg")
		return
	}

	return m.SendMsgFunc(p)
}

//SendMsgMinimockCounter returns a count of MessageService_MessageClientMock.SendMsgFunc invocations
func (m *MessageService_MessageClientMock) SendMsgMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.SendMsgCounter)
}

//SendMsgMinimockPreCounter returns the value of MessageService_MessageClientMock.SendMsg invocations
func (m *MessageService_MessageClientMock) SendMsgMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.SendMsgPreCounter)
}

type mMessageService_MessageClientMockTrailer struct {
	mock *MessageService_MessageClientMock
}

//Return sets up a mock for MessageService_MessageClient.Trailer to return Return's arguments
func (m *mMessageService_MessageClientMockTrailer) Return(r metadata.MD) *MessageService_MessageClientMock {
	m.mock.TrailerFunc = func() metadata.MD {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of MessageService_MessageClient.Trailer method
func (m *mMessageService_MessageClientMockTrailer) Set(f func() (r metadata.MD)) *MessageService_MessageClientMock {
	m.mock.TrailerFunc = f
	return m.mock
}

//Trailer implements github.com/superstas/gcoin/gcoin/network/message.MessageService_MessageClient interface
func (m *MessageService_MessageClientMock) Trailer() (r metadata.MD) {
	atomic.AddUint64(&m.TrailerPreCounter, 1)
	defer atomic.AddUint64(&m.TrailerCounter, 1)

	if m.TrailerFunc == nil {
		m.t.Fatal("Unexpected call to MessageService_MessageClientMock.Trailer")
		return
	}

	return m.TrailerFunc()
}

//TrailerMinimockCounter returns a count of MessageService_MessageClientMock.TrailerFunc invocations
func (m *MessageService_MessageClientMock) TrailerMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.TrailerCounter)
}

//TrailerMinimockPreCounter returns the value of MessageService_MessageClientMock.Trailer invocations
func (m *MessageService_MessageClientMock) TrailerMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.TrailerPreCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *MessageService_MessageClientMock) ValidateCallCounters() {

	if m.CloseSendFunc != nil && atomic.LoadUint64(&m.CloseSendCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageClientMock.CloseSend")
	}

	if m.ContextFunc != nil && atomic.LoadUint64(&m.ContextCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageClientMock.Context")
	}

	if m.HeaderFunc != nil && atomic.LoadUint64(&m.HeaderCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageClientMock.Header")
	}

	if m.RecvFunc != nil && atomic.LoadUint64(&m.RecvCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageClientMock.Recv")
	}

	if m.RecvMsgFunc != nil && atomic.LoadUint64(&m.RecvMsgCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageClientMock.RecvMsg")
	}

	if m.SendFunc != nil && atomic.LoadUint64(&m.SendCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageClientMock.Send")
	}

	if m.SendMsgFunc != nil && atomic.LoadUint64(&m.SendMsgCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageClientMock.SendMsg")
	}

	if m.TrailerFunc != nil && atomic.LoadUint64(&m.TrailerCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageClientMock.Trailer")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *MessageService_MessageClientMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *MessageService_MessageClientMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *MessageService_MessageClientMock) MinimockFinish() {

	if m.CloseSendFunc != nil && atomic.LoadUint64(&m.CloseSendCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageClientMock.CloseSend")
	}

	if m.ContextFunc != nil && atomic.LoadUint64(&m.ContextCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageClientMock.Context")
	}

	if m.HeaderFunc != nil && atomic.LoadUint64(&m.HeaderCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageClientMock.Header")
	}

	if m.RecvFunc != nil && atomic.LoadUint64(&m.RecvCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageClientMock.Recv")
	}

	if m.RecvMsgFunc != nil && atomic.LoadUint64(&m.RecvMsgCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageClientMock.RecvMsg")
	}

	if m.SendFunc != nil && atomic.LoadUint64(&m.SendCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageClientMock.Send")
	}

	if m.SendMsgFunc != nil && atomic.LoadUint64(&m.SendMsgCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageClientMock.SendMsg")
	}

	if m.TrailerFunc != nil && atomic.LoadUint64(&m.TrailerCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageClientMock.Trailer")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *MessageService_MessageClientMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *MessageService_MessageClientMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && (m.CloseSendFunc == nil || atomic.LoadUint64(&m.CloseSendCounter) > 0)
		ok = ok && (m.ContextFunc == nil || atomic.LoadUint64(&m.ContextCounter) > 0)
		ok = ok && (m.HeaderFunc == nil || atomic.LoadUint64(&m.HeaderCounter) > 0)
		ok = ok && (m.RecvFunc == nil || atomic.LoadUint64(&m.RecvCounter) > 0)
		ok = ok && (m.RecvMsgFunc == nil || atomic.LoadUint64(&m.RecvMsgCounter) > 0)
		ok = ok && (m.SendFunc == nil || atomic.LoadUint64(&m.SendCounter) > 0)
		ok = ok && (m.SendMsgFunc == nil || atomic.LoadUint64(&m.SendMsgCounter) > 0)
		ok = ok && (m.TrailerFunc == nil || atomic.LoadUint64(&m.TrailerCounter) > 0)

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if m.CloseSendFunc != nil && atomic.LoadUint64(&m.CloseSendCounter) == 0 {
				m.t.Error("Expected call to MessageService_MessageClientMock.CloseSend")
			}

			if m.ContextFunc != nil && atomic.LoadUint64(&m.ContextCounter) == 0 {
				m.t.Error("Expected call to MessageService_MessageClientMock.Context")
			}

			if m.HeaderFunc != nil && atomic.LoadUint64(&m.HeaderCounter) == 0 {
				m.t.Error("Expected call to MessageService_MessageClientMock.Header")
			}

			if m.RecvFunc != nil && atomic.LoadUint64(&m.RecvCounter) == 0 {
				m.t.Error("Expected call to MessageService_MessageClientMock.Recv")
			}

			if m.RecvMsgFunc != nil && atomic.LoadUint64(&m.RecvMsgCounter) == 0 {
				m.t.Error("Expected call to MessageService_MessageClientMock.RecvMsg")
			}

			if m.SendFunc != nil && atomic.LoadUint64(&m.SendCounter) == 0 {
				m.t.Error("Expected call to MessageService_MessageClientMock.Send")
			}

			if m.SendMsgFunc != nil && atomic.LoadUint64(&m.SendMsgCounter) == 0 {
				m.t.Error("Expected call to MessageService_MessageClientMock.SendMsg")
			}

			if m.TrailerFunc != nil && atomic.LoadUint64(&m.TrailerCounter) == 0 {
				m.t.Error("Expected call to MessageService_MessageClientMock.Trailer")
			}

			m.t.Fatalf("Some mocks were not called on time: %s", timeout)
			return
		default:
			time.Sleep(time.Millisecond)
		}
	}
}

//AllMocksCalled returns true if all mocked methods were called before the execution of AllMocksCalled,
//it can be used with assert/require, i.e. assert.True(mock.AllMocksCalled())
func (m *MessageService_MessageClientMock) AllMocksCalled() bool {

	if m.CloseSendFunc != nil && atomic.LoadUint64(&m.CloseSendCounter) == 0 {
		return false
	}

	if m.ContextFunc != nil && atomic.LoadUint64(&m.ContextCounter) == 0 {
		return false
	}

	if m.HeaderFunc != nil && atomic.LoadUint64(&m.HeaderCounter) == 0 {
		return false
	}

	if m.RecvFunc != nil && atomic.LoadUint64(&m.RecvCounter) == 0 {
		return false
	}

	if m.RecvMsgFunc != nil && atomic.LoadUint64(&m.RecvMsgCounter) == 0 {
		return false
	}

	if m.SendFunc != nil && atomic.LoadUint64(&m.SendCounter) == 0 {
		return false
	}

	if m.SendMsgFunc != nil && atomic.LoadUint64(&m.SendMsgCounter) == 0 {
		return false
	}

	if m.TrailerFunc != nil && atomic.LoadUint64(&m.TrailerCounter) == 0 {
		return false
	}

	return true
}
