package wallet

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const (
	defaultFilename   = "wallet.dat"
	defaultKeysAmount = 5
)

// fileStorage represents a file storage of a wallet
type fileStorage struct {
	f *os.File
}

// NewFileStorage creates new Storage.
// It creates new filename or uses existed one
func NewFileStorage(filename string) Storage {
	if filename == "" {
		filename = defaultFilename
	}

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return &fileStorage{f}
}

//Init creates a wallet or reads existed one
func (f *fileStorage) Init(ctx context.Context) (*Wallet, error) {
	data, err := ioutil.ReadAll(f.f)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return Init(defaultKeysAmount)
	}

	var w Wallet
	err = json.Unmarshal(data, &w)
	if err != nil {
		return nil, err
	}

	if len(w.Keys()) == 0 {
		return Init(defaultKeysAmount)
	}

	return &w, nil
}

// Write writes a wallet data to a file
// It's not the best way to store private keys :)
// You can read more about it here: https://stackoverflow.com/questions/21322182/how-to-store-ecdsa-private-key-in-go
func (f *fileStorage) Write(ctx context.Context, w *Wallet) error {
	data, err := json.Marshal(w)
	if err != nil {
		return err
	}

	// since we have only wallet.AddKeys behavior, we'll rewrite the file every time
	_, err = f.f.WriteAt(data, 0)
	return err
}

// Close closes a file
func (f *fileStorage) Close() error {
	return f.f.Close()
}
