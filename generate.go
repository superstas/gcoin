package gcoin

//go:generate protoc -I ./gcoin/network/cli --go_out=plugins=grpc:./gcoin/network/cli ./gcoin/network/cli/cli.proto
//go:generate protoc -I ./gcoin/network/message --go_out=plugins=grpc:./gcoin/network/message ./gcoin/network/message/message.proto
//go:generate minimock -i ./gcoin/transaction.Storage -o ./gcoin/transaction
//go:generate minimock -i ./gcoin/network/message.MessageService_MessageClient,./gcoin/network/message.MessageService_MessageServer -o ./gcoin/network/message/mocks -s _mock.go
