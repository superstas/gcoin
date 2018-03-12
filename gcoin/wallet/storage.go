package wallet

import "context"

// Storage represents a storage for wallet
// It might be a file, memory or network implementation
type Storage interface {
	//Init creates a wallet or reads existed one
	Init(context.Context) (*Wallet, error)
	// Write writes a wallet to an implemented storage
	Write(context.Context, *Wallet) error
	// Close releases allocated resources of the storage
	Close() error
}
