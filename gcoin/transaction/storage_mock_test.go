package transaction

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "Storage" can be found in github.com/superstas/gcoin/gcoin/transaction
*/
import (
	context "context"
	"sync/atomic"
	"time"

	btcutil "github.com/btcsuite/btcutil"
	"github.com/gojuno/minimock"

	testify_assert "github.com/stretchr/testify/assert"
)

//StorageMock implements github.com/superstas/gcoin/gcoin/transaction.Storage
type StorageMock struct {
	t minimock.Tester

	FindTransactionByIDFunc       func(p []byte) (r []byte, r1 Transaction, r2 error)
	FindTransactionByIDCounter    uint64
	FindTransactionByIDPreCounter uint64
	FindTransactionByIDMock       mStorageMockFindTransactionByID

	FindUTXOByPKHFunc       func(p context.Context, p1 btcutil.AddressPubKeyHash) (r UTXOSet, r1 error)
	FindUTXOByPKHCounter    uint64
	FindUTXOByPKHPreCounter uint64
	FindUTXOByPKHMock       mStorageMockFindUTXOByPKH

	TotalUTXOByPKHFunc       func(p context.Context, p1 btcutil.AddressPubKeyHash) (r int64, r1 error)
	TotalUTXOByPKHCounter    uint64
	TotalUTXOByPKHPreCounter uint64
	TotalUTXOByPKHMock       mStorageMockTotalUTXOByPKH
}

//NewStorageMock returns a mock for github.com/superstas/gcoin/gcoin/transaction.Storage
func NewStorageMock(t minimock.Tester) *StorageMock {
	m := &StorageMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.FindTransactionByIDMock = mStorageMockFindTransactionByID{mock: m}
	m.FindUTXOByPKHMock = mStorageMockFindUTXOByPKH{mock: m}
	m.TotalUTXOByPKHMock = mStorageMockTotalUTXOByPKH{mock: m}

	return m
}

type mStorageMockFindTransactionByID struct {
	mock             *StorageMock
	mockExpectations *StorageMockFindTransactionByIDParams
}

//StorageMockFindTransactionByIDParams represents input parameters of the Storage.FindTransactionByID
type StorageMockFindTransactionByIDParams struct {
	p []byte
}

//Expect sets up expected params for the Storage.FindTransactionByID
func (m *mStorageMockFindTransactionByID) Expect(p []byte) *mStorageMockFindTransactionByID {
	m.mockExpectations = &StorageMockFindTransactionByIDParams{p}
	return m
}

//Return sets up a mock for Storage.FindTransactionByID to return Return's arguments
func (m *mStorageMockFindTransactionByID) Return(r []byte, r1 Transaction, r2 error) *StorageMock {
	m.mock.FindTransactionByIDFunc = func(p []byte) ([]byte, Transaction, error) {
		return r, r1, r2
	}
	return m.mock
}

//Set uses given function f as a mock of Storage.FindTransactionByID method
func (m *mStorageMockFindTransactionByID) Set(f func(p []byte) (r []byte, r1 Transaction, r2 error)) *StorageMock {
	m.mock.FindTransactionByIDFunc = f
	return m.mock
}

//FindTransactionByID implements github.com/superstas/gcoin/gcoin/transaction.Storage interface
func (m *StorageMock) FindTransactionByID(p []byte) (r []byte, r1 Transaction, r2 error) {
	atomic.AddUint64(&m.FindTransactionByIDPreCounter, 1)
	defer atomic.AddUint64(&m.FindTransactionByIDCounter, 1)

	if m.FindTransactionByIDMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.FindTransactionByIDMock.mockExpectations, StorageMockFindTransactionByIDParams{p},
			"Storage.FindTransactionByID got unexpected parameters")

		if m.FindTransactionByIDFunc == nil {

			m.t.Fatal("No results are set for the StorageMock.FindTransactionByID")

			return
		}
	}

	if m.FindTransactionByIDFunc == nil {
		m.t.Fatal("Unexpected call to StorageMock.FindTransactionByID")
		return
	}

	return m.FindTransactionByIDFunc(p)
}

//FindTransactionByIDMinimockCounter returns a count of StorageMock.FindTransactionByIDFunc invocations
func (m *StorageMock) FindTransactionByIDMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.FindTransactionByIDCounter)
}

//FindTransactionByIDMinimockPreCounter returns the value of StorageMock.FindTransactionByID invocations
func (m *StorageMock) FindTransactionByIDMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.FindTransactionByIDPreCounter)
}

type mStorageMockFindUTXOByPKH struct {
	mock             *StorageMock
	mockExpectations *StorageMockFindUTXOByPKHParams
}

//StorageMockFindUTXOByPKHParams represents input parameters of the Storage.FindUTXOByPKH
type StorageMockFindUTXOByPKHParams struct {
	p  context.Context
	p1 btcutil.AddressPubKeyHash
}

//Expect sets up expected params for the Storage.FindUTXOByPKH
func (m *mStorageMockFindUTXOByPKH) Expect(p context.Context, p1 btcutil.AddressPubKeyHash) *mStorageMockFindUTXOByPKH {
	m.mockExpectations = &StorageMockFindUTXOByPKHParams{p, p1}
	return m
}

//Return sets up a mock for Storage.FindUTXOByPKH to return Return's arguments
func (m *mStorageMockFindUTXOByPKH) Return(r UTXOSet, r1 error) *StorageMock {
	m.mock.FindUTXOByPKHFunc = func(p context.Context, p1 btcutil.AddressPubKeyHash) (UTXOSet, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Storage.FindUTXOByPKH method
func (m *mStorageMockFindUTXOByPKH) Set(f func(p context.Context, p1 btcutil.AddressPubKeyHash) (r UTXOSet, r1 error)) *StorageMock {
	m.mock.FindUTXOByPKHFunc = f
	return m.mock
}

//FindUTXOByPKH implements github.com/superstas/gcoin/gcoin/transaction.Storage interface
func (m *StorageMock) FindUTXOByPKH(p context.Context, p1 btcutil.AddressPubKeyHash) (r UTXOSet, r1 error) {
	atomic.AddUint64(&m.FindUTXOByPKHPreCounter, 1)
	defer atomic.AddUint64(&m.FindUTXOByPKHCounter, 1)

	if m.FindUTXOByPKHMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.FindUTXOByPKHMock.mockExpectations, StorageMockFindUTXOByPKHParams{p, p1},
			"Storage.FindUTXOByPKH got unexpected parameters")

		if m.FindUTXOByPKHFunc == nil {

			m.t.Fatal("No results are set for the StorageMock.FindUTXOByPKH")

			return
		}
	}

	if m.FindUTXOByPKHFunc == nil {
		m.t.Fatal("Unexpected call to StorageMock.FindUTXOByPKH")
		return
	}

	return m.FindUTXOByPKHFunc(p, p1)
}

//FindUTXOByPKHMinimockCounter returns a count of StorageMock.FindUTXOByPKHFunc invocations
func (m *StorageMock) FindUTXOByPKHMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.FindUTXOByPKHCounter)
}

//FindUTXOByPKHMinimockPreCounter returns the value of StorageMock.FindUTXOByPKH invocations
func (m *StorageMock) FindUTXOByPKHMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.FindUTXOByPKHPreCounter)
}

type mStorageMockTotalUTXOByPKH struct {
	mock             *StorageMock
	mockExpectations *StorageMockTotalUTXOByPKHParams
}

//StorageMockTotalUTXOByPKHParams represents input parameters of the Storage.TotalUTXOByPKH
type StorageMockTotalUTXOByPKHParams struct {
	p  context.Context
	p1 btcutil.AddressPubKeyHash
}

//Expect sets up expected params for the Storage.TotalUTXOByPKH
func (m *mStorageMockTotalUTXOByPKH) Expect(p context.Context, p1 btcutil.AddressPubKeyHash) *mStorageMockTotalUTXOByPKH {
	m.mockExpectations = &StorageMockTotalUTXOByPKHParams{p, p1}
	return m
}

//Return sets up a mock for Storage.TotalUTXOByPKH to return Return's arguments
func (m *mStorageMockTotalUTXOByPKH) Return(r int64, r1 error) *StorageMock {
	m.mock.TotalUTXOByPKHFunc = func(p context.Context, p1 btcutil.AddressPubKeyHash) (int64, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Storage.TotalUTXOByPKH method
func (m *mStorageMockTotalUTXOByPKH) Set(f func(p context.Context, p1 btcutil.AddressPubKeyHash) (r int64, r1 error)) *StorageMock {
	m.mock.TotalUTXOByPKHFunc = f
	return m.mock
}

//TotalUTXOByPKH implements github.com/superstas/gcoin/gcoin/transaction.Storage interface
func (m *StorageMock) TotalUTXOByPKH(p context.Context, p1 btcutil.AddressPubKeyHash) (r int64, r1 error) {
	atomic.AddUint64(&m.TotalUTXOByPKHPreCounter, 1)
	defer atomic.AddUint64(&m.TotalUTXOByPKHCounter, 1)

	if m.TotalUTXOByPKHMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.TotalUTXOByPKHMock.mockExpectations, StorageMockTotalUTXOByPKHParams{p, p1},
			"Storage.TotalUTXOByPKH got unexpected parameters")

		if m.TotalUTXOByPKHFunc == nil {

			m.t.Fatal("No results are set for the StorageMock.TotalUTXOByPKH")

			return
		}
	}

	if m.TotalUTXOByPKHFunc == nil {
		m.t.Fatal("Unexpected call to StorageMock.TotalUTXOByPKH")
		return
	}

	return m.TotalUTXOByPKHFunc(p, p1)
}

//TotalUTXOByPKHMinimockCounter returns a count of StorageMock.TotalUTXOByPKHFunc invocations
func (m *StorageMock) TotalUTXOByPKHMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.TotalUTXOByPKHCounter)
}

//TotalUTXOByPKHMinimockPreCounter returns the value of StorageMock.TotalUTXOByPKH invocations
func (m *StorageMock) TotalUTXOByPKHMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.TotalUTXOByPKHPreCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *StorageMock) ValidateCallCounters() {

	if m.FindTransactionByIDFunc != nil && atomic.LoadUint64(&m.FindTransactionByIDCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.FindTransactionByID")
	}

	if m.FindUTXOByPKHFunc != nil && atomic.LoadUint64(&m.FindUTXOByPKHCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.FindUTXOByPKH")
	}

	if m.TotalUTXOByPKHFunc != nil && atomic.LoadUint64(&m.TotalUTXOByPKHCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.TotalUTXOByPKH")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *StorageMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *StorageMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *StorageMock) MinimockFinish() {

	if m.FindTransactionByIDFunc != nil && atomic.LoadUint64(&m.FindTransactionByIDCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.FindTransactionByID")
	}

	if m.FindUTXOByPKHFunc != nil && atomic.LoadUint64(&m.FindUTXOByPKHCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.FindUTXOByPKH")
	}

	if m.TotalUTXOByPKHFunc != nil && atomic.LoadUint64(&m.TotalUTXOByPKHCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.TotalUTXOByPKH")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *StorageMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *StorageMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && (m.FindTransactionByIDFunc == nil || atomic.LoadUint64(&m.FindTransactionByIDCounter) > 0)
		ok = ok && (m.FindUTXOByPKHFunc == nil || atomic.LoadUint64(&m.FindUTXOByPKHCounter) > 0)
		ok = ok && (m.TotalUTXOByPKHFunc == nil || atomic.LoadUint64(&m.TotalUTXOByPKHCounter) > 0)

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if m.FindTransactionByIDFunc != nil && atomic.LoadUint64(&m.FindTransactionByIDCounter) == 0 {
				m.t.Error("Expected call to StorageMock.FindTransactionByID")
			}

			if m.FindUTXOByPKHFunc != nil && atomic.LoadUint64(&m.FindUTXOByPKHCounter) == 0 {
				m.t.Error("Expected call to StorageMock.FindUTXOByPKH")
			}

			if m.TotalUTXOByPKHFunc != nil && atomic.LoadUint64(&m.TotalUTXOByPKHCounter) == 0 {
				m.t.Error("Expected call to StorageMock.TotalUTXOByPKH")
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
func (m *StorageMock) AllMocksCalled() bool {

	if m.FindTransactionByIDFunc != nil && atomic.LoadUint64(&m.FindTransactionByIDCounter) == 0 {
		return false
	}

	if m.FindUTXOByPKHFunc != nil && atomic.LoadUint64(&m.FindUTXOByPKHCounter) == 0 {
		return false
	}

	if m.TotalUTXOByPKHFunc != nil && atomic.LoadUint64(&m.TotalUTXOByPKHCounter) == 0 {
		return false
	}

	return true
}
