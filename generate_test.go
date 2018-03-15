// +build test

//go:generate minimock -i ./gcoin/transaction.Storage -o ./gcoin/transaction
//go:generate minimock -i ./gcoin/blockchain.Storage -o ./gcoin/blockchain
//go:generate minimock -i ./gcoin/network/message.MessageService_MessageClient,./gcoin/network/message.MessageService_MessageServer -o ./gcoin/network/message/mocks -s _mock.go
package gcoin
