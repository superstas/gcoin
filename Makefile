all: fmt imports generate test install

generate: 
	go generate 

install: 
	go install ./cmd/gcoind/...
	go install ./cmd/gcoin-cli/...

test:
	go test -cover ./gcoin/...

imports:
	find ./ -type f -name '*.go' -not -path "./vendor*" -exec goimports -w {} +

fmt:
	find ./ -type f -name '*.go' -not -path "./vendor*" -exec gofmt -w -s {} +

install_deps:
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/gojuno/minimock/cmd/minimock

