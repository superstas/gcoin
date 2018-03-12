package mocks

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.8
The original interface "MessageService_MessageServer" can be found in github.com/superstas/gcoin/gcoin/network/message
*/
import (
	context "context"
	"sync/atomic"
	"time"

	"github.com/gojuno/minimock"
	message "github.com/superstas/gcoin/gcoin/network/message"
	metadata "google.golang.org/grpc/metadata"

	testify_assert "github.com/stretchr/testify/assert"
)

//MessageService_MessageServerMock implements github.com/superstas/gcoin/gcoin/network/message.MessageService_MessageServer
type MessageService_MessageServerMock struct {
	t minimock.Tester

	ContextFunc       func() (r context.Context)
	ContextCounter    uint64
	ContextPreCounter uint64
	ContextMock       mMessageService_MessageServerMockContext

	RecvFunc       func() (r *message.Msg, r1 error)
	RecvCounter    uint64
	RecvPreCounter uint64
	RecvMock       mMessageService_MessageServerMockRecv

	RecvMsgFunc       func(p interface{}) (r error)
	RecvMsgCounter    uint64
	RecvMsgPreCounter uint64
	RecvMsgMock       mMessageService_MessageServerMockRecvMsg

	SendFunc       func(p *message.Msg) (r error)
	SendCounter    uint64
	SendPreCounter uint64
	SendMock       mMessageService_MessageServerMockSend

	SendHeaderFunc       func(p metadata.MD) (r error)
	SendHeaderCounter    uint64
	SendHeaderPreCounter uint64
	SendHeaderMock       mMessageService_MessageServerMockSendHeader

	SendMsgFunc       func(p interface{}) (r error)
	SendMsgCounter    uint64
	SendMsgPreCounter uint64
	SendMsgMock       mMessageService_MessageServerMockSendMsg

	SetHeaderFunc       func(p metadata.MD) (r error)
	SetHeaderCounter    uint64
	SetHeaderPreCounter uint64
	SetHeaderMock       mMessageService_MessageServerMockSetHeader

	SetTrailerFunc       func(p metadata.MD)
	SetTrailerCounter    uint64
	SetTrailerPreCounter uint64
	SetTrailerMock       mMessageService_MessageServerMockSetTrailer
}

//NewMessageService_MessageServerMock returns a mock for github.com/superstas/gcoin/gcoin/network/message.MessageService_MessageServer
func NewMessageService_MessageServerMock(t minimock.Tester) *MessageService_MessageServerMock {
	m := &MessageService_MessageServerMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ContextMock = mMessageService_MessageServerMockContext{mock: m}
	m.RecvMock = mMessageService_MessageServerMockRecv{mock: m}
	m.RecvMsgMock = mMessageService_MessageServerMockRecvMsg{mock: m}
	m.SendMock = mMessageService_MessageServerMockSend{mock: m}
	m.SendHeaderMock = mMessageService_MessageServerMockSendHeader{mock: m}
	m.SendMsgMock = mMessageService_MessageServerMockSendMsg{mock: m}
	m.SetHeaderMock = mMessageService_MessageServerMockSetHeader{mock: m}
	m.SetTrailerMock = mMessageService_MessageServerMockSetTrailer{mock: m}

	return m
}

type mMessageService_MessageServerMockContext struct {
	mock *MessageService_MessageServerMock
}

//Return sets up a mock for MessageService_MessageServer.Context to return Return's arguments
func (m *mMessageService_MessageServerMockContext) Return(r context.Context) *MessageService_MessageServerMock {
	m.mock.ContextFunc = func() context.Context {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of MessageService_MessageServer.Context method
func (m *mMessageService_MessageServerMockContext) Set(f func() (r context.Context)) *MessageService_MessageServerMock {
	m.mock.ContextFunc = f
	return m.mock
}

//Context implements github.com/superstas/gcoin/gcoin/network/message.MessageService_MessageServer interface
func (m *MessageService_MessageServerMock) Context() (r context.Context) {
	atomic.AddUint64(&m.ContextPreCounter, 1)
	defer atomic.AddUint64(&m.ContextCounter, 1)

	if m.ContextFunc == nil {
		m.t.Fatal("Unexpected call to MessageService_MessageServerMock.Context")
		return
	}

	return m.ContextFunc()
}

//ContextMinimockCounter returns a count of MessageService_MessageServerMock.ContextFunc invocations
func (m *MessageService_MessageServerMock) ContextMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.ContextCounter)
}

//ContextMinimockPreCounter returns the value of MessageService_MessageServerMock.Context invocations
func (m *MessageService_MessageServerMock) ContextMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.ContextPreCounter)
}

type mMessageService_MessageServerMockRecv struct {
	mock *MessageService_MessageServerMock
}

//Return sets up a mock for MessageService_MessageServer.Recv to return Return's arguments
func (m *mMessageService_MessageServerMockRecv) Return(r *message.Msg, r1 error) *MessageService_MessageServerMock {
	m.mock.RecvFunc = func() (*message.Msg, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of MessageService_MessageServer.Recv method
func (m *mMessageService_MessageServerMockRecv) Set(f func() (r *message.Msg, r1 error)) *MessageService_MessageServerMock {
	m.mock.RecvFunc = f
	return m.mock
}

//Recv implements github.com/superstas/gcoin/gcoin/network/message.MessageService_MessageServer interface
func (m *MessageService_MessageServerMock) Recv() (r *message.Msg, r1 error) {
	atomic.AddUint64(&m.RecvPreCounter, 1)
	defer atomic.AddUint64(&m.RecvCounter, 1)

	if m.RecvFunc == nil {
		m.t.Fatal("Unexpected call to MessageService_MessageServerMock.Recv")
		return
	}

	return m.RecvFunc()
}

//RecvMinimockCounter returns a count of MessageService_MessageServerMock.RecvFunc invocations
func (m *MessageService_MessageServerMock) RecvMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.RecvCounter)
}

//RecvMinimockPreCounter returns the value of MessageService_MessageServerMock.Recv invocations
func (m *MessageService_MessageServerMock) RecvMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.RecvPreCounter)
}

type mMessageService_MessageServerMockRecvMsg struct {
	mock             *MessageService_MessageServerMock
	mockExpectations *MessageService_MessageServerMockRecvMsgParams
}

//MessageService_MessageServerMockRecvMsgParams represents input parameters of the MessageService_MessageServer.RecvMsg
type MessageService_MessageServerMockRecvMsgParams struct {
	p interface{}
}

//Expect sets up expected params for the MessageService_MessageServer.RecvMsg
func (m *mMessageService_MessageServerMockRecvMsg) Expect(p interface{}) *mMessageService_MessageServerMockRecvMsg {
	m.mockExpectations = &MessageService_MessageServerMockRecvMsgParams{p}
	return m
}

//Return sets up a mock for MessageService_MessageServer.RecvMsg to return Return's arguments
func (m *mMessageService_MessageServerMockRecvMsg) Return(r error) *MessageService_MessageServerMock {
	m.mock.RecvMsgFunc = func(p interface{}) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of MessageService_MessageServer.RecvMsg method
func (m *mMessageService_MessageServerMockRecvMsg) Set(f func(p interface{}) (r error)) *MessageService_MessageServerMock {
	m.mock.RecvMsgFunc = f
	return m.mock
}

//RecvMsg implements github.com/superstas/gcoin/gcoin/network/message.MessageService_MessageServer interface
func (m *MessageService_MessageServerMock) RecvMsg(p interface{}) (r error) {
	atomic.AddUint64(&m.RecvMsgPreCounter, 1)
	defer atomic.AddUint64(&m.RecvMsgCounter, 1)

	if m.RecvMsgMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.RecvMsgMock.mockExpectations, MessageService_MessageServerMockRecvMsgParams{p},
			"MessageService_MessageServer.RecvMsg got unexpected parameters")

		if m.RecvMsgFunc == nil {

			m.t.Fatal("No results are set for the MessageService_MessageServerMock.RecvMsg")

			return
		}
	}

	if m.RecvMsgFunc == nil {
		m.t.Fatal("Unexpected call to MessageService_MessageServerMock.RecvMsg")
		return
	}

	return m.RecvMsgFunc(p)
}

//RecvMsgMinimockCounter returns a count of MessageService_MessageServerMock.RecvMsgFunc invocations
func (m *MessageService_MessageServerMock) RecvMsgMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.RecvMsgCounter)
}

//RecvMsgMinimockPreCounter returns the value of MessageService_MessageServerMock.RecvMsg invocations
func (m *MessageService_MessageServerMock) RecvMsgMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.RecvMsgPreCounter)
}

type mMessageService_MessageServerMockSend struct {
	mock             *MessageService_MessageServerMock
	mockExpectations *MessageService_MessageServerMockSendParams
}

//MessageService_MessageServerMockSendParams represents input parameters of the MessageService_MessageServer.Send
type MessageService_MessageServerMockSendParams struct {
	p *message.Msg
}

//Expect sets up expected params for the MessageService_MessageServer.Send
func (m *mMessageService_MessageServerMockSend) Expect(p *message.Msg) *mMessageService_MessageServerMockSend {
	m.mockExpectations = &MessageService_MessageServerMockSendParams{p}
	return m
}

//Return sets up a mock for MessageService_MessageServer.Send to return Return's arguments
func (m *mMessageService_MessageServerMockSend) Return(r error) *MessageService_MessageServerMock {
	m.mock.SendFunc = func(p *message.Msg) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of MessageService_MessageServer.Send method
func (m *mMessageService_MessageServerMockSend) Set(f func(p *message.Msg) (r error)) *MessageService_MessageServerMock {
	m.mock.SendFunc = f
	return m.mock
}

//Send implements github.com/superstas/gcoin/gcoin/network/message.MessageService_MessageServer interface
func (m *MessageService_MessageServerMock) Send(p *message.Msg) (r error) {
	atomic.AddUint64(&m.SendPreCounter, 1)
	defer atomic.AddUint64(&m.SendCounter, 1)

	if m.SendMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.SendMock.mockExpectations, MessageService_MessageServerMockSendParams{p},
			"MessageService_MessageServer.Send got unexpected parameters")

		if m.SendFunc == nil {

			m.t.Fatal("No results are set for the MessageService_MessageServerMock.Send")

			return
		}
	}

	if m.SendFunc == nil {
		m.t.Fatal("Unexpected call to MessageService_MessageServerMock.Send")
		return
	}

	return m.SendFunc(p)
}

//SendMinimockCounter returns a count of MessageService_MessageServerMock.SendFunc invocations
func (m *MessageService_MessageServerMock) SendMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.SendCounter)
}

//SendMinimockPreCounter returns the value of MessageService_MessageServerMock.Send invocations
func (m *MessageService_MessageServerMock) SendMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.SendPreCounter)
}

type mMessageService_MessageServerMockSendHeader struct {
	mock             *MessageService_MessageServerMock
	mockExpectations *MessageService_MessageServerMockSendHeaderParams
}

//MessageService_MessageServerMockSendHeaderParams represents input parameters of the MessageService_MessageServer.SendHeader
type MessageService_MessageServerMockSendHeaderParams struct {
	p metadata.MD
}

//Expect sets up expected params for the MessageService_MessageServer.SendHeader
func (m *mMessageService_MessageServerMockSendHeader) Expect(p metadata.MD) *mMessageService_MessageServerMockSendHeader {
	m.mockExpectations = &MessageService_MessageServerMockSendHeaderParams{p}
	return m
}

//Return sets up a mock for MessageService_MessageServer.SendHeader to return Return's arguments
func (m *mMessageService_MessageServerMockSendHeader) Return(r error) *MessageService_MessageServerMock {
	m.mock.SendHeaderFunc = func(p metadata.MD) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of MessageService_MessageServer.SendHeader method
func (m *mMessageService_MessageServerMockSendHeader) Set(f func(p metadata.MD) (r error)) *MessageService_MessageServerMock {
	m.mock.SendHeaderFunc = f
	return m.mock
}

//SendHeader implements github.com/superstas/gcoin/gcoin/network/message.MessageService_MessageServer interface
func (m *MessageService_MessageServerMock) SendHeader(p metadata.MD) (r error) {
	atomic.AddUint64(&m.SendHeaderPreCounter, 1)
	defer atomic.AddUint64(&m.SendHeaderCounter, 1)

	if m.SendHeaderMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.SendHeaderMock.mockExpectations, MessageService_MessageServerMockSendHeaderParams{p},
			"MessageService_MessageServer.SendHeader got unexpected parameters")

		if m.SendHeaderFunc == nil {

			m.t.Fatal("No results are set for the MessageService_MessageServerMock.SendHeader")

			return
		}
	}

	if m.SendHeaderFunc == nil {
		m.t.Fatal("Unexpected call to MessageService_MessageServerMock.SendHeader")
		return
	}

	return m.SendHeaderFunc(p)
}

//SendHeaderMinimockCounter returns a count of MessageService_MessageServerMock.SendHeaderFunc invocations
func (m *MessageService_MessageServerMock) SendHeaderMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.SendHeaderCounter)
}

//SendHeaderMinimockPreCounter returns the value of MessageService_MessageServerMock.SendHeader invocations
func (m *MessageService_MessageServerMock) SendHeaderMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.SendHeaderPreCounter)
}

type mMessageService_MessageServerMockSendMsg struct {
	mock             *MessageService_MessageServerMock
	mockExpectations *MessageService_MessageServerMockSendMsgParams
}

//MessageService_MessageServerMockSendMsgParams represents input parameters of the MessageService_MessageServer.SendMsg
type MessageService_MessageServerMockSendMsgParams struct {
	p interface{}
}

//Expect sets up expected params for the MessageService_MessageServer.SendMsg
func (m *mMessageService_MessageServerMockSendMsg) Expect(p interface{}) *mMessageService_MessageServerMockSendMsg {
	m.mockExpectations = &MessageService_MessageServerMockSendMsgParams{p}
	return m
}

//Return sets up a mock for MessageService_MessageServer.SendMsg to return Return's arguments
func (m *mMessageService_MessageServerMockSendMsg) Return(r error) *MessageService_MessageServerMock {
	m.mock.SendMsgFunc = func(p interface{}) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of MessageService_MessageServer.SendMsg method
func (m *mMessageService_MessageServerMockSendMsg) Set(f func(p interface{}) (r error)) *MessageService_MessageServerMock {
	m.mock.SendMsgFunc = f
	return m.mock
}

//SendMsg implements github.com/superstas/gcoin/gcoin/network/message.MessageService_MessageServer interface
func (m *MessageService_MessageServerMock) SendMsg(p interface{}) (r error) {
	atomic.AddUint64(&m.SendMsgPreCounter, 1)
	defer atomic.AddUint64(&m.SendMsgCounter, 1)

	if m.SendMsgMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.SendMsgMock.mockExpectations, MessageService_MessageServerMockSendMsgParams{p},
			"MessageService_MessageServer.SendMsg got unexpected parameters")

		if m.SendMsgFunc == nil {

			m.t.Fatal("No results are set for the MessageService_MessageServerMock.SendMsg")

			return
		}
	}

	if m.SendMsgFunc == nil {
		m.t.Fatal("Unexpected call to MessageService_MessageServerMock.SendMsg")
		return
	}

	return m.SendMsgFunc(p)
}

//SendMsgMinimockCounter returns a count of MessageService_MessageServerMock.SendMsgFunc invocations
func (m *MessageService_MessageServerMock) SendMsgMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.SendMsgCounter)
}

//SendMsgMinimockPreCounter returns the value of MessageService_MessageServerMock.SendMsg invocations
func (m *MessageService_MessageServerMock) SendMsgMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.SendMsgPreCounter)
}

type mMessageService_MessageServerMockSetHeader struct {
	mock             *MessageService_MessageServerMock
	mockExpectations *MessageService_MessageServerMockSetHeaderParams
}

//MessageService_MessageServerMockSetHeaderParams represents input parameters of the MessageService_MessageServer.SetHeader
type MessageService_MessageServerMockSetHeaderParams struct {
	p metadata.MD
}

//Expect sets up expected params for the MessageService_MessageServer.SetHeader
func (m *mMessageService_MessageServerMockSetHeader) Expect(p metadata.MD) *mMessageService_MessageServerMockSetHeader {
	m.mockExpectations = &MessageService_MessageServerMockSetHeaderParams{p}
	return m
}

//Return sets up a mock for MessageService_MessageServer.SetHeader to return Return's arguments
func (m *mMessageService_MessageServerMockSetHeader) Return(r error) *MessageService_MessageServerMock {
	m.mock.SetHeaderFunc = func(p metadata.MD) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of MessageService_MessageServer.SetHeader method
func (m *mMessageService_MessageServerMockSetHeader) Set(f func(p metadata.MD) (r error)) *MessageService_MessageServerMock {
	m.mock.SetHeaderFunc = f
	return m.mock
}

//SetHeader implements github.com/superstas/gcoin/gcoin/network/message.MessageService_MessageServer interface
func (m *MessageService_MessageServerMock) SetHeader(p metadata.MD) (r error) {
	atomic.AddUint64(&m.SetHeaderPreCounter, 1)
	defer atomic.AddUint64(&m.SetHeaderCounter, 1)

	if m.SetHeaderMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.SetHeaderMock.mockExpectations, MessageService_MessageServerMockSetHeaderParams{p},
			"MessageService_MessageServer.SetHeader got unexpected parameters")

		if m.SetHeaderFunc == nil {

			m.t.Fatal("No results are set for the MessageService_MessageServerMock.SetHeader")

			return
		}
	}

	if m.SetHeaderFunc == nil {
		m.t.Fatal("Unexpected call to MessageService_MessageServerMock.SetHeader")
		return
	}

	return m.SetHeaderFunc(p)
}

//SetHeaderMinimockCounter returns a count of MessageService_MessageServerMock.SetHeaderFunc invocations
func (m *MessageService_MessageServerMock) SetHeaderMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.SetHeaderCounter)
}

//SetHeaderMinimockPreCounter returns the value of MessageService_MessageServerMock.SetHeader invocations
func (m *MessageService_MessageServerMock) SetHeaderMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.SetHeaderPreCounter)
}

type mMessageService_MessageServerMockSetTrailer struct {
	mock             *MessageService_MessageServerMock
	mockExpectations *MessageService_MessageServerMockSetTrailerParams
}

//MessageService_MessageServerMockSetTrailerParams represents input parameters of the MessageService_MessageServer.SetTrailer
type MessageService_MessageServerMockSetTrailerParams struct {
	p metadata.MD
}

//Expect sets up expected params for the MessageService_MessageServer.SetTrailer
func (m *mMessageService_MessageServerMockSetTrailer) Expect(p metadata.MD) *mMessageService_MessageServerMockSetTrailer {
	m.mockExpectations = &MessageService_MessageServerMockSetTrailerParams{p}
	return m
}

//Return sets up a mock for MessageService_MessageServer.SetTrailer to return Return's arguments
func (m *mMessageService_MessageServerMockSetTrailer) Return() *MessageService_MessageServerMock {
	m.mock.SetTrailerFunc = func(p metadata.MD) {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of MessageService_MessageServer.SetTrailer method
func (m *mMessageService_MessageServerMockSetTrailer) Set(f func(p metadata.MD)) *MessageService_MessageServerMock {
	m.mock.SetTrailerFunc = f
	return m.mock
}

//SetTrailer implements github.com/superstas/gcoin/gcoin/network/message.MessageService_MessageServer interface
func (m *MessageService_MessageServerMock) SetTrailer(p metadata.MD) {
	atomic.AddUint64(&m.SetTrailerPreCounter, 1)
	defer atomic.AddUint64(&m.SetTrailerCounter, 1)

	if m.SetTrailerMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.SetTrailerMock.mockExpectations, MessageService_MessageServerMockSetTrailerParams{p},
			"MessageService_MessageServer.SetTrailer got unexpected parameters")

		if m.SetTrailerFunc == nil {

			m.t.Fatal("No results are set for the MessageService_MessageServerMock.SetTrailer")

			return
		}
	}

	if m.SetTrailerFunc == nil {
		m.t.Fatal("Unexpected call to MessageService_MessageServerMock.SetTrailer")
		return
	}

	m.SetTrailerFunc(p)
}

//SetTrailerMinimockCounter returns a count of MessageService_MessageServerMock.SetTrailerFunc invocations
func (m *MessageService_MessageServerMock) SetTrailerMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.SetTrailerCounter)
}

//SetTrailerMinimockPreCounter returns the value of MessageService_MessageServerMock.SetTrailer invocations
func (m *MessageService_MessageServerMock) SetTrailerMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.SetTrailerPreCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *MessageService_MessageServerMock) ValidateCallCounters() {

	if m.ContextFunc != nil && atomic.LoadUint64(&m.ContextCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageServerMock.Context")
	}

	if m.RecvFunc != nil && atomic.LoadUint64(&m.RecvCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageServerMock.Recv")
	}

	if m.RecvMsgFunc != nil && atomic.LoadUint64(&m.RecvMsgCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageServerMock.RecvMsg")
	}

	if m.SendFunc != nil && atomic.LoadUint64(&m.SendCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageServerMock.Send")
	}

	if m.SendHeaderFunc != nil && atomic.LoadUint64(&m.SendHeaderCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageServerMock.SendHeader")
	}

	if m.SendMsgFunc != nil && atomic.LoadUint64(&m.SendMsgCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageServerMock.SendMsg")
	}

	if m.SetHeaderFunc != nil && atomic.LoadUint64(&m.SetHeaderCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageServerMock.SetHeader")
	}

	if m.SetTrailerFunc != nil && atomic.LoadUint64(&m.SetTrailerCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageServerMock.SetTrailer")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *MessageService_MessageServerMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *MessageService_MessageServerMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *MessageService_MessageServerMock) MinimockFinish() {

	if m.ContextFunc != nil && atomic.LoadUint64(&m.ContextCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageServerMock.Context")
	}

	if m.RecvFunc != nil && atomic.LoadUint64(&m.RecvCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageServerMock.Recv")
	}

	if m.RecvMsgFunc != nil && atomic.LoadUint64(&m.RecvMsgCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageServerMock.RecvMsg")
	}

	if m.SendFunc != nil && atomic.LoadUint64(&m.SendCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageServerMock.Send")
	}

	if m.SendHeaderFunc != nil && atomic.LoadUint64(&m.SendHeaderCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageServerMock.SendHeader")
	}

	if m.SendMsgFunc != nil && atomic.LoadUint64(&m.SendMsgCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageServerMock.SendMsg")
	}

	if m.SetHeaderFunc != nil && atomic.LoadUint64(&m.SetHeaderCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageServerMock.SetHeader")
	}

	if m.SetTrailerFunc != nil && atomic.LoadUint64(&m.SetTrailerCounter) == 0 {
		m.t.Fatal("Expected call to MessageService_MessageServerMock.SetTrailer")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *MessageService_MessageServerMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *MessageService_MessageServerMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && (m.ContextFunc == nil || atomic.LoadUint64(&m.ContextCounter) > 0)
		ok = ok && (m.RecvFunc == nil || atomic.LoadUint64(&m.RecvCounter) > 0)
		ok = ok && (m.RecvMsgFunc == nil || atomic.LoadUint64(&m.RecvMsgCounter) > 0)
		ok = ok && (m.SendFunc == nil || atomic.LoadUint64(&m.SendCounter) > 0)
		ok = ok && (m.SendHeaderFunc == nil || atomic.LoadUint64(&m.SendHeaderCounter) > 0)
		ok = ok && (m.SendMsgFunc == nil || atomic.LoadUint64(&m.SendMsgCounter) > 0)
		ok = ok && (m.SetHeaderFunc == nil || atomic.LoadUint64(&m.SetHeaderCounter) > 0)
		ok = ok && (m.SetTrailerFunc == nil || atomic.LoadUint64(&m.SetTrailerCounter) > 0)

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if m.ContextFunc != nil && atomic.LoadUint64(&m.ContextCounter) == 0 {
				m.t.Error("Expected call to MessageService_MessageServerMock.Context")
			}

			if m.RecvFunc != nil && atomic.LoadUint64(&m.RecvCounter) == 0 {
				m.t.Error("Expected call to MessageService_MessageServerMock.Recv")
			}

			if m.RecvMsgFunc != nil && atomic.LoadUint64(&m.RecvMsgCounter) == 0 {
				m.t.Error("Expected call to MessageService_MessageServerMock.RecvMsg")
			}

			if m.SendFunc != nil && atomic.LoadUint64(&m.SendCounter) == 0 {
				m.t.Error("Expected call to MessageService_MessageServerMock.Send")
			}

			if m.SendHeaderFunc != nil && atomic.LoadUint64(&m.SendHeaderCounter) == 0 {
				m.t.Error("Expected call to MessageService_MessageServerMock.SendHeader")
			}

			if m.SendMsgFunc != nil && atomic.LoadUint64(&m.SendMsgCounter) == 0 {
				m.t.Error("Expected call to MessageService_MessageServerMock.SendMsg")
			}

			if m.SetHeaderFunc != nil && atomic.LoadUint64(&m.SetHeaderCounter) == 0 {
				m.t.Error("Expected call to MessageService_MessageServerMock.SetHeader")
			}

			if m.SetTrailerFunc != nil && atomic.LoadUint64(&m.SetTrailerCounter) == 0 {
				m.t.Error("Expected call to MessageService_MessageServerMock.SetTrailer")
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
func (m *MessageService_MessageServerMock) AllMocksCalled() bool {

	if m.ContextFunc != nil && atomic.LoadUint64(&m.ContextCounter) == 0 {
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

	if m.SendHeaderFunc != nil && atomic.LoadUint64(&m.SendHeaderCounter) == 0 {
		return false
	}

	if m.SendMsgFunc != nil && atomic.LoadUint64(&m.SendMsgCounter) == 0 {
		return false
	}

	if m.SetHeaderFunc != nil && atomic.LoadUint64(&m.SetHeaderCounter) == 0 {
		return false
	}

	if m.SetTrailerFunc != nil && atomic.LoadUint64(&m.SetTrailerCounter) == 0 {
		return false
	}

	return true
}
