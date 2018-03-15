//go:generate protoc -I ./gcoin/network/cli --go_out=plugins=grpc:./gcoin/network/cli ./gcoin/network/cli/cli.proto
//go:generate protoc -I ./gcoin/network/message --go_out=plugins=grpc:./gcoin/network/message ./gcoin/network/message/message.proto

package gcoin
