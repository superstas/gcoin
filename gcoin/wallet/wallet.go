package wallet

import (
	"sync"

	"encoding/json"

	"github.com/btcsuite/btcutil"
	"github.com/superstas/gcoin/gcoin/keys"
)

// Wallet represents a container with generated keys
type Wallet struct {
	keys []*keys.Keys
	sync.RWMutex
}

// Init creates a wallet with N pre-generated keys
func Init(n int) (*Wallet, error) {
	k := make([]*keys.Keys, n)
	for i := 0; i < n; i++ {
		p, err := keys.New()
		if err != nil {
			return nil, err
		}

		k[i] = p
	}
	return &Wallet{keys: k}, nil
}

// Keys returns all existing keys in the wallet
func (w *Wallet) Keys() []*keys.Keys {
	w.RLock()
	defer w.RUnlock()
	return w.keys
}

// AddKeys adds keys to the wallet
func (w *Wallet) AddKeys(k *keys.Keys) {
	w.Lock()
	defer w.Unlock()
	w.keys = append(w.keys, k)
}

// Addresses returns all addresses for all existed keys
func (w *Wallet) Addresses() []*btcutil.AddressPubKeyHash {
	w.RLock()
	defer w.RUnlock()

	addresses := make([]*btcutil.AddressPubKeyHash, len(w.keys))
	for i, k := range w.keys {
		addresses[i] = k.AddressPKH
	}
	return addresses
}

// MarshalJSON implements Marshaler interface
func (w *Wallet) MarshalJSON() ([]byte, error) {
	k := struct {
		Keys []*keys.Keys `json:"keys"`
	}{w.Keys()}
	return json.Marshal(&k)
}

// UnmarshalJSON implements Unmarshaler interface
func (w *Wallet) UnmarshalJSON(d []byte) error {
	var k struct {
		Keys []*keys.Keys `json:"keys"`
	}

	if err := json.Unmarshal(d, &k); err != nil {
		return err
	}

	w.Lock()
	w.keys = k.Keys
	w.Unlock()

	return nil
}
