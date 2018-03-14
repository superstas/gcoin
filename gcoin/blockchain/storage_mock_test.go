package blockchain

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "Storage" can be found in github.com/superstas/gcoin/gcoin/blockchain
*/
import (
	context "context"
	"sync/atomic"
	"time"

	btcutil "github.com/btcsuite/btcutil"
	"github.com/gojuno/minimock"
	block "github.com/superstas/gcoin/gcoin/block"
	transaction "github.com/superstas/gcoin/gcoin/transaction"

	testify_assert "github.com/stretchr/testify/assert"
)

//StorageMock implements github.com/superstas/gcoin/gcoin/blockchain.Storage
type StorageMock struct {
	t minimock.Tester

	CloseFunc       func(p context.Context) (r error)
	CloseCounter    uint64
	ClosePreCounter uint64
	CloseMock       mStorageMockClose

	FindTransactionByIDFunc       func(p []byte) (r []byte, r1 transaction.Transaction, r2 error)
	FindTransactionByIDCounter    uint64
	FindTransactionByIDPreCounter uint64
	FindTransactionByIDMock       mStorageMockFindTransactionByID

	FindUTXOByPKHFunc       func(p context.Context, p1 btcutil.AddressPubKeyHash) (r transaction.UTXOSet, r1 error)
	FindUTXOByPKHCounter    uint64
	FindUTXOByPKHPreCounter uint64
	FindUTXOByPKHMock       mStorageMockFindUTXOByPKH

	ReadBlockByHashFunc       func(p context.Context, p1 []byte) (r block.Block, r1 error)
	ReadBlockByHashCounter    uint64
	ReadBlockByHashPreCounter uint64
	ReadBlockByHashMock       mStorageMockReadBlockByHash

	ReadGenesisBlockFunc       func(p context.Context) (r block.Block, r1 error)
	ReadGenesisBlockCounter    uint64
	ReadGenesisBlockPreCounter uint64
	ReadGenesisBlockMock       mStorageMockReadGenesisBlock

	ReadLastBlockHashFunc       func(p context.Context) (r []byte, r1 error)
	ReadLastBlockHashCounter    uint64
	ReadLastBlockHashPreCounter uint64
	ReadLastBlockHashMock       mStorageMockReadLastBlockHash

	ReadLastNBlocksFunc       func(p context.Context, p1 int) (r []block.Block, r1 error)
	ReadLastNBlocksCounter    uint64
	ReadLastNBlocksPreCounter uint64
	ReadLastNBlocksMock       mStorageMockReadLastNBlocks

	TotalUTXOByPKHFunc       func(p context.Context, p1 btcutil.AddressPubKeyHash) (r int64, r1 error)
	TotalUTXOByPKHCounter    uint64
	TotalUTXOByPKHPreCounter uint64
	TotalUTXOByPKHMock       mStorageMockTotalUTXOByPKH

	WriteBlockFunc       func(p context.Context, p1 block.Block) (r error)
	WriteBlockCounter    uint64
	WriteBlockPreCounter uint64
	WriteBlockMock       mStorageMockWriteBlock
}

//NewStorageMock returns a mock for github.com/superstas/gcoin/gcoin/blockchain.Storage
func NewStorageMock(t minimock.Tester) *StorageMock {
	m := &StorageMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CloseMock = mStorageMockClose{mock: m}
	m.FindTransactionByIDMock = mStorageMockFindTransactionByID{mock: m}
	m.FindUTXOByPKHMock = mStorageMockFindUTXOByPKH{mock: m}
	m.ReadBlockByHashMock = mStorageMockReadBlockByHash{mock: m}
	m.ReadGenesisBlockMock = mStorageMockReadGenesisBlock{mock: m}
	m.ReadLastBlockHashMock = mStorageMockReadLastBlockHash{mock: m}
	m.ReadLastNBlocksMock = mStorageMockReadLastNBlocks{mock: m}
	m.TotalUTXOByPKHMock = mStorageMockTotalUTXOByPKH{mock: m}
	m.WriteBlockMock = mStorageMockWriteBlock{mock: m}

	return m
}

type mStorageMockClose struct {
	mock             *StorageMock
	mockExpectations *StorageMockCloseParams
}

//StorageMockCloseParams represents input parameters of the Storage.Close
type StorageMockCloseParams struct {
	p context.Context
}

//Expect sets up expected params for the Storage.Close
func (m *mStorageMockClose) Expect(p context.Context) *mStorageMockClose {
	m.mockExpectations = &StorageMockCloseParams{p}
	return m
}

//Return sets up a mock for Storage.Close to return Return's arguments
func (m *mStorageMockClose) Return(r error) *StorageMock {
	m.mock.CloseFunc = func(p context.Context) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Storage.Close method
func (m *mStorageMockClose) Set(f func(p context.Context) (r error)) *StorageMock {
	m.mock.CloseFunc = f
	return m.mock
}

//Close implements github.com/superstas/gcoin/gcoin/blockchain.Storage interface
func (m *StorageMock) Close(p context.Context) (r error) {
	atomic.AddUint64(&m.ClosePreCounter, 1)
	defer atomic.AddUint64(&m.CloseCounter, 1)

	if m.CloseMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.CloseMock.mockExpectations, StorageMockCloseParams{p},
			"Storage.Close got unexpected parameters")

		if m.CloseFunc == nil {

			m.t.Fatal("No results are set for the StorageMock.Close")

			return
		}
	}

	if m.CloseFunc == nil {
		m.t.Fatal("Unexpected call to StorageMock.Close")
		return
	}

	return m.CloseFunc(p)
}

//CloseMinimockCounter returns a count of StorageMock.CloseFunc invocations
func (m *StorageMock) CloseMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.CloseCounter)
}

//CloseMinimockPreCounter returns the value of StorageMock.Close invocations
func (m *StorageMock) CloseMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.ClosePreCounter)
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
func (m *mStorageMockFindTransactionByID) Return(r []byte, r1 transaction.Transaction, r2 error) *StorageMock {
	m.mock.FindTransactionByIDFunc = func(p []byte) ([]byte, transaction.Transaction, error) {
		return r, r1, r2
	}
	return m.mock
}

//Set uses given function f as a mock of Storage.FindTransactionByID method
func (m *mStorageMockFindTransactionByID) Set(f func(p []byte) (r []byte, r1 transaction.Transaction, r2 error)) *StorageMock {
	m.mock.FindTransactionByIDFunc = f
	return m.mock
}

//FindTransactionByID implements github.com/superstas/gcoin/gcoin/blockchain.Storage interface
func (m *StorageMock) FindTransactionByID(p []byte) (r []byte, r1 transaction.Transaction, r2 error) {
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
func (m *mStorageMockFindUTXOByPKH) Return(r transaction.UTXOSet, r1 error) *StorageMock {
	m.mock.FindUTXOByPKHFunc = func(p context.Context, p1 btcutil.AddressPubKeyHash) (transaction.UTXOSet, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Storage.FindUTXOByPKH method
func (m *mStorageMockFindUTXOByPKH) Set(f func(p context.Context, p1 btcutil.AddressPubKeyHash) (r transaction.UTXOSet, r1 error)) *StorageMock {
	m.mock.FindUTXOByPKHFunc = f
	return m.mock
}

//FindUTXOByPKH implements github.com/superstas/gcoin/gcoin/blockchain.Storage interface
func (m *StorageMock) FindUTXOByPKH(p context.Context, p1 btcutil.AddressPubKeyHash) (r transaction.UTXOSet, r1 error) {
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

type mStorageMockReadBlockByHash struct {
	mock             *StorageMock
	mockExpectations *StorageMockReadBlockByHashParams
}

//StorageMockReadBlockByHashParams represents input parameters of the Storage.ReadBlockByHash
type StorageMockReadBlockByHashParams struct {
	p  context.Context
	p1 []byte
}

//Expect sets up expected params for the Storage.ReadBlockByHash
func (m *mStorageMockReadBlockByHash) Expect(p context.Context, p1 []byte) *mStorageMockReadBlockByHash {
	m.mockExpectations = &StorageMockReadBlockByHashParams{p, p1}
	return m
}

//Return sets up a mock for Storage.ReadBlockByHash to return Return's arguments
func (m *mStorageMockReadBlockByHash) Return(r block.Block, r1 error) *StorageMock {
	m.mock.ReadBlockByHashFunc = func(p context.Context, p1 []byte) (block.Block, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Storage.ReadBlockByHash method
func (m *mStorageMockReadBlockByHash) Set(f func(p context.Context, p1 []byte) (r block.Block, r1 error)) *StorageMock {
	m.mock.ReadBlockByHashFunc = f
	return m.mock
}

//ReadBlockByHash implements github.com/superstas/gcoin/gcoin/blockchain.Storage interface
func (m *StorageMock) ReadBlockByHash(p context.Context, p1 []byte) (r block.Block, r1 error) {
	atomic.AddUint64(&m.ReadBlockByHashPreCounter, 1)
	defer atomic.AddUint64(&m.ReadBlockByHashCounter, 1)

	if m.ReadBlockByHashMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.ReadBlockByHashMock.mockExpectations, StorageMockReadBlockByHashParams{p, p1},
			"Storage.ReadBlockByHash got unexpected parameters")

		if m.ReadBlockByHashFunc == nil {

			m.t.Fatal("No results are set for the StorageMock.ReadBlockByHash")

			return
		}
	}

	if m.ReadBlockByHashFunc == nil {
		m.t.Fatal("Unexpected call to StorageMock.ReadBlockByHash")
		return
	}

	return m.ReadBlockByHashFunc(p, p1)
}

//ReadBlockByHashMinimockCounter returns a count of StorageMock.ReadBlockByHashFunc invocations
func (m *StorageMock) ReadBlockByHashMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.ReadBlockByHashCounter)
}

//ReadBlockByHashMinimockPreCounter returns the value of StorageMock.ReadBlockByHash invocations
func (m *StorageMock) ReadBlockByHashMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.ReadBlockByHashPreCounter)
}

type mStorageMockReadGenesisBlock struct {
	mock             *StorageMock
	mockExpectations *StorageMockReadGenesisBlockParams
}

//StorageMockReadGenesisBlockParams represents input parameters of the Storage.ReadGenesisBlock
type StorageMockReadGenesisBlockParams struct {
	p context.Context
}

//Expect sets up expected params for the Storage.ReadGenesisBlock
func (m *mStorageMockReadGenesisBlock) Expect(p context.Context) *mStorageMockReadGenesisBlock {
	m.mockExpectations = &StorageMockReadGenesisBlockParams{p}
	return m
}

//Return sets up a mock for Storage.ReadGenesisBlock to return Return's arguments
func (m *mStorageMockReadGenesisBlock) Return(r block.Block, r1 error) *StorageMock {
	m.mock.ReadGenesisBlockFunc = func(p context.Context) (block.Block, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Storage.ReadGenesisBlock method
func (m *mStorageMockReadGenesisBlock) Set(f func(p context.Context) (r block.Block, r1 error)) *StorageMock {
	m.mock.ReadGenesisBlockFunc = f
	return m.mock
}

//ReadGenesisBlock implements github.com/superstas/gcoin/gcoin/blockchain.Storage interface
func (m *StorageMock) ReadGenesisBlock(p context.Context) (r block.Block, r1 error) {
	atomic.AddUint64(&m.ReadGenesisBlockPreCounter, 1)
	defer atomic.AddUint64(&m.ReadGenesisBlockCounter, 1)

	if m.ReadGenesisBlockMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.ReadGenesisBlockMock.mockExpectations, StorageMockReadGenesisBlockParams{p},
			"Storage.ReadGenesisBlock got unexpected parameters")

		if m.ReadGenesisBlockFunc == nil {

			m.t.Fatal("No results are set for the StorageMock.ReadGenesisBlock")

			return
		}
	}

	if m.ReadGenesisBlockFunc == nil {
		m.t.Fatal("Unexpected call to StorageMock.ReadGenesisBlock")
		return
	}

	return m.ReadGenesisBlockFunc(p)
}

//ReadGenesisBlockMinimockCounter returns a count of StorageMock.ReadGenesisBlockFunc invocations
func (m *StorageMock) ReadGenesisBlockMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.ReadGenesisBlockCounter)
}

//ReadGenesisBlockMinimockPreCounter returns the value of StorageMock.ReadGenesisBlock invocations
func (m *StorageMock) ReadGenesisBlockMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.ReadGenesisBlockPreCounter)
}

type mStorageMockReadLastBlockHash struct {
	mock             *StorageMock
	mockExpectations *StorageMockReadLastBlockHashParams
}

//StorageMockReadLastBlockHashParams represents input parameters of the Storage.ReadLastBlockHash
type StorageMockReadLastBlockHashParams struct {
	p context.Context
}

//Expect sets up expected params for the Storage.ReadLastBlockHash
func (m *mStorageMockReadLastBlockHash) Expect(p context.Context) *mStorageMockReadLastBlockHash {
	m.mockExpectations = &StorageMockReadLastBlockHashParams{p}
	return m
}

//Return sets up a mock for Storage.ReadLastBlockHash to return Return's arguments
func (m *mStorageMockReadLastBlockHash) Return(r []byte, r1 error) *StorageMock {
	m.mock.ReadLastBlockHashFunc = func(p context.Context) ([]byte, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Storage.ReadLastBlockHash method
func (m *mStorageMockReadLastBlockHash) Set(f func(p context.Context) (r []byte, r1 error)) *StorageMock {
	m.mock.ReadLastBlockHashFunc = f
	return m.mock
}

//ReadLastBlockHash implements github.com/superstas/gcoin/gcoin/blockchain.Storage interface
func (m *StorageMock) ReadLastBlockHash(p context.Context) (r []byte, r1 error) {
	atomic.AddUint64(&m.ReadLastBlockHashPreCounter, 1)
	defer atomic.AddUint64(&m.ReadLastBlockHashCounter, 1)

	if m.ReadLastBlockHashMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.ReadLastBlockHashMock.mockExpectations, StorageMockReadLastBlockHashParams{p},
			"Storage.ReadLastBlockHash got unexpected parameters")

		if m.ReadLastBlockHashFunc == nil {

			m.t.Fatal("No results are set for the StorageMock.ReadLastBlockHash")

			return
		}
	}

	if m.ReadLastBlockHashFunc == nil {
		m.t.Fatal("Unexpected call to StorageMock.ReadLastBlockHash")
		return
	}

	return m.ReadLastBlockHashFunc(p)
}

//ReadLastBlockHashMinimockCounter returns a count of StorageMock.ReadLastBlockHashFunc invocations
func (m *StorageMock) ReadLastBlockHashMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.ReadLastBlockHashCounter)
}

//ReadLastBlockHashMinimockPreCounter returns the value of StorageMock.ReadLastBlockHash invocations
func (m *StorageMock) ReadLastBlockHashMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.ReadLastBlockHashPreCounter)
}

type mStorageMockReadLastNBlocks struct {
	mock             *StorageMock
	mockExpectations *StorageMockReadLastNBlocksParams
}

//StorageMockReadLastNBlocksParams represents input parameters of the Storage.ReadLastNBlocks
type StorageMockReadLastNBlocksParams struct {
	p  context.Context
	p1 int
}

//Expect sets up expected params for the Storage.ReadLastNBlocks
func (m *mStorageMockReadLastNBlocks) Expect(p context.Context, p1 int) *mStorageMockReadLastNBlocks {
	m.mockExpectations = &StorageMockReadLastNBlocksParams{p, p1}
	return m
}

//Return sets up a mock for Storage.ReadLastNBlocks to return Return's arguments
func (m *mStorageMockReadLastNBlocks) Return(r []block.Block, r1 error) *StorageMock {
	m.mock.ReadLastNBlocksFunc = func(p context.Context, p1 int) ([]block.Block, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Storage.ReadLastNBlocks method
func (m *mStorageMockReadLastNBlocks) Set(f func(p context.Context, p1 int) (r []block.Block, r1 error)) *StorageMock {
	m.mock.ReadLastNBlocksFunc = f
	return m.mock
}

//ReadLastNBlocks implements github.com/superstas/gcoin/gcoin/blockchain.Storage interface
func (m *StorageMock) ReadLastNBlocks(p context.Context, p1 int) (r []block.Block, r1 error) {
	atomic.AddUint64(&m.ReadLastNBlocksPreCounter, 1)
	defer atomic.AddUint64(&m.ReadLastNBlocksCounter, 1)

	if m.ReadLastNBlocksMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.ReadLastNBlocksMock.mockExpectations, StorageMockReadLastNBlocksParams{p, p1},
			"Storage.ReadLastNBlocks got unexpected parameters")

		if m.ReadLastNBlocksFunc == nil {

			m.t.Fatal("No results are set for the StorageMock.ReadLastNBlocks")

			return
		}
	}

	if m.ReadLastNBlocksFunc == nil {
		m.t.Fatal("Unexpected call to StorageMock.ReadLastNBlocks")
		return
	}

	return m.ReadLastNBlocksFunc(p, p1)
}

//ReadLastNBlocksMinimockCounter returns a count of StorageMock.ReadLastNBlocksFunc invocations
func (m *StorageMock) ReadLastNBlocksMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.ReadLastNBlocksCounter)
}

//ReadLastNBlocksMinimockPreCounter returns the value of StorageMock.ReadLastNBlocks invocations
func (m *StorageMock) ReadLastNBlocksMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.ReadLastNBlocksPreCounter)
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

//TotalUTXOByPKH implements github.com/superstas/gcoin/gcoin/blockchain.Storage interface
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

type mStorageMockWriteBlock struct {
	mock             *StorageMock
	mockExpectations *StorageMockWriteBlockParams
}

//StorageMockWriteBlockParams represents input parameters of the Storage.WriteBlock
type StorageMockWriteBlockParams struct {
	p  context.Context
	p1 block.Block
}

//Expect sets up expected params for the Storage.WriteBlock
func (m *mStorageMockWriteBlock) Expect(p context.Context, p1 block.Block) *mStorageMockWriteBlock {
	m.mockExpectations = &StorageMockWriteBlockParams{p, p1}
	return m
}

//Return sets up a mock for Storage.WriteBlock to return Return's arguments
func (m *mStorageMockWriteBlock) Return(r error) *StorageMock {
	m.mock.WriteBlockFunc = func(p context.Context, p1 block.Block) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Storage.WriteBlock method
func (m *mStorageMockWriteBlock) Set(f func(p context.Context, p1 block.Block) (r error)) *StorageMock {
	m.mock.WriteBlockFunc = f
	return m.mock
}

//WriteBlock implements github.com/superstas/gcoin/gcoin/blockchain.Storage interface
func (m *StorageMock) WriteBlock(p context.Context, p1 block.Block) (r error) {
	atomic.AddUint64(&m.WriteBlockPreCounter, 1)
	defer atomic.AddUint64(&m.WriteBlockCounter, 1)

	if m.WriteBlockMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.WriteBlockMock.mockExpectations, StorageMockWriteBlockParams{p, p1},
			"Storage.WriteBlock got unexpected parameters")

		if m.WriteBlockFunc == nil {

			m.t.Fatal("No results are set for the StorageMock.WriteBlock")

			return
		}
	}

	if m.WriteBlockFunc == nil {
		m.t.Fatal("Unexpected call to StorageMock.WriteBlock")
		return
	}

	return m.WriteBlockFunc(p, p1)
}

//WriteBlockMinimockCounter returns a count of StorageMock.WriteBlockFunc invocations
func (m *StorageMock) WriteBlockMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.WriteBlockCounter)
}

//WriteBlockMinimockPreCounter returns the value of StorageMock.WriteBlock invocations
func (m *StorageMock) WriteBlockMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.WriteBlockPreCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *StorageMock) ValidateCallCounters() {

	if m.CloseFunc != nil && atomic.LoadUint64(&m.CloseCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.Close")
	}

	if m.FindTransactionByIDFunc != nil && atomic.LoadUint64(&m.FindTransactionByIDCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.FindTransactionByID")
	}

	if m.FindUTXOByPKHFunc != nil && atomic.LoadUint64(&m.FindUTXOByPKHCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.FindUTXOByPKH")
	}

	if m.ReadBlockByHashFunc != nil && atomic.LoadUint64(&m.ReadBlockByHashCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.ReadBlockByHash")
	}

	if m.ReadGenesisBlockFunc != nil && atomic.LoadUint64(&m.ReadGenesisBlockCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.ReadGenesisBlock")
	}

	if m.ReadLastBlockHashFunc != nil && atomic.LoadUint64(&m.ReadLastBlockHashCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.ReadLastBlockHash")
	}

	if m.ReadLastNBlocksFunc != nil && atomic.LoadUint64(&m.ReadLastNBlocksCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.ReadLastNBlocks")
	}

	if m.TotalUTXOByPKHFunc != nil && atomic.LoadUint64(&m.TotalUTXOByPKHCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.TotalUTXOByPKH")
	}

	if m.WriteBlockFunc != nil && atomic.LoadUint64(&m.WriteBlockCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.WriteBlock")
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

	if m.CloseFunc != nil && atomic.LoadUint64(&m.CloseCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.Close")
	}

	if m.FindTransactionByIDFunc != nil && atomic.LoadUint64(&m.FindTransactionByIDCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.FindTransactionByID")
	}

	if m.FindUTXOByPKHFunc != nil && atomic.LoadUint64(&m.FindUTXOByPKHCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.FindUTXOByPKH")
	}

	if m.ReadBlockByHashFunc != nil && atomic.LoadUint64(&m.ReadBlockByHashCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.ReadBlockByHash")
	}

	if m.ReadGenesisBlockFunc != nil && atomic.LoadUint64(&m.ReadGenesisBlockCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.ReadGenesisBlock")
	}

	if m.ReadLastBlockHashFunc != nil && atomic.LoadUint64(&m.ReadLastBlockHashCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.ReadLastBlockHash")
	}

	if m.ReadLastNBlocksFunc != nil && atomic.LoadUint64(&m.ReadLastNBlocksCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.ReadLastNBlocks")
	}

	if m.TotalUTXOByPKHFunc != nil && atomic.LoadUint64(&m.TotalUTXOByPKHCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.TotalUTXOByPKH")
	}

	if m.WriteBlockFunc != nil && atomic.LoadUint64(&m.WriteBlockCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.WriteBlock")
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
		ok = ok && (m.CloseFunc == nil || atomic.LoadUint64(&m.CloseCounter) > 0)
		ok = ok && (m.FindTransactionByIDFunc == nil || atomic.LoadUint64(&m.FindTransactionByIDCounter) > 0)
		ok = ok && (m.FindUTXOByPKHFunc == nil || atomic.LoadUint64(&m.FindUTXOByPKHCounter) > 0)
		ok = ok && (m.ReadBlockByHashFunc == nil || atomic.LoadUint64(&m.ReadBlockByHashCounter) > 0)
		ok = ok && (m.ReadGenesisBlockFunc == nil || atomic.LoadUint64(&m.ReadGenesisBlockCounter) > 0)
		ok = ok && (m.ReadLastBlockHashFunc == nil || atomic.LoadUint64(&m.ReadLastBlockHashCounter) > 0)
		ok = ok && (m.ReadLastNBlocksFunc == nil || atomic.LoadUint64(&m.ReadLastNBlocksCounter) > 0)
		ok = ok && (m.TotalUTXOByPKHFunc == nil || atomic.LoadUint64(&m.TotalUTXOByPKHCounter) > 0)
		ok = ok && (m.WriteBlockFunc == nil || atomic.LoadUint64(&m.WriteBlockCounter) > 0)

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if m.CloseFunc != nil && atomic.LoadUint64(&m.CloseCounter) == 0 {
				m.t.Error("Expected call to StorageMock.Close")
			}

			if m.FindTransactionByIDFunc != nil && atomic.LoadUint64(&m.FindTransactionByIDCounter) == 0 {
				m.t.Error("Expected call to StorageMock.FindTransactionByID")
			}

			if m.FindUTXOByPKHFunc != nil && atomic.LoadUint64(&m.FindUTXOByPKHCounter) == 0 {
				m.t.Error("Expected call to StorageMock.FindUTXOByPKH")
			}

			if m.ReadBlockByHashFunc != nil && atomic.LoadUint64(&m.ReadBlockByHashCounter) == 0 {
				m.t.Error("Expected call to StorageMock.ReadBlockByHash")
			}

			if m.ReadGenesisBlockFunc != nil && atomic.LoadUint64(&m.ReadGenesisBlockCounter) == 0 {
				m.t.Error("Expected call to StorageMock.ReadGenesisBlock")
			}

			if m.ReadLastBlockHashFunc != nil && atomic.LoadUint64(&m.ReadLastBlockHashCounter) == 0 {
				m.t.Error("Expected call to StorageMock.ReadLastBlockHash")
			}

			if m.ReadLastNBlocksFunc != nil && atomic.LoadUint64(&m.ReadLastNBlocksCounter) == 0 {
				m.t.Error("Expected call to StorageMock.ReadLastNBlocks")
			}

			if m.TotalUTXOByPKHFunc != nil && atomic.LoadUint64(&m.TotalUTXOByPKHCounter) == 0 {
				m.t.Error("Expected call to StorageMock.TotalUTXOByPKH")
			}

			if m.WriteBlockFunc != nil && atomic.LoadUint64(&m.WriteBlockCounter) == 0 {
				m.t.Error("Expected call to StorageMock.WriteBlock")
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

	if m.CloseFunc != nil && atomic.LoadUint64(&m.CloseCounter) == 0 {
		return false
	}

	if m.FindTransactionByIDFunc != nil && atomic.LoadUint64(&m.FindTransactionByIDCounter) == 0 {
		return false
	}

	if m.FindUTXOByPKHFunc != nil && atomic.LoadUint64(&m.FindUTXOByPKHCounter) == 0 {
		return false
	}

	if m.ReadBlockByHashFunc != nil && atomic.LoadUint64(&m.ReadBlockByHashCounter) == 0 {
		return false
	}

	if m.ReadGenesisBlockFunc != nil && atomic.LoadUint64(&m.ReadGenesisBlockCounter) == 0 {
		return false
	}

	if m.ReadLastBlockHashFunc != nil && atomic.LoadUint64(&m.ReadLastBlockHashCounter) == 0 {
		return false
	}

	if m.ReadLastNBlocksFunc != nil && atomic.LoadUint64(&m.ReadLastNBlocksCounter) == 0 {
		return false
	}

	if m.TotalUTXOByPKHFunc != nil && atomic.LoadUint64(&m.TotalUTXOByPKHCounter) == 0 {
		return false
	}

	if m.WriteBlockFunc != nil && atomic.LoadUint64(&m.WriteBlockCounter) == 0 {
		return false
	}

	return true
}
