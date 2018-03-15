package keys

import (
	"encoding/json"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/pkg/errors"
)

// Keys represents a simple wrapper for generated keys
type Keys struct {
	PrivateKey *btcec.PrivateKey
	PublicKey  *btcec.PublicKey
	AddressPKH *btcutil.AddressPubKeyHash
}

// New creates new keys object
func New() (*Keys, error) {
	privateKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return nil, errors.Wrap(err, "failed to create private key")
	}

	pubKey := privateKey.PubKey().SerializeCompressed()
	addressPKH, err := btcutil.NewAddressPubKeyHash(btcutil.Hash160(pubKey), &chaincfg.Params{})

	if err != nil {
		return nil, errors.Wrap(err, "failed to create address public key hash")
	}

	return &Keys{
		PrivateKey: privateKey,
		PublicKey:  privateKey.PubKey(),
		AddressPKH: addressPKH,
	}, nil
}

// MarshalJSON returns marshaled keys
// This is not the best way to serialize, but this is a prototype :)
func (k *Keys) MarshalJSON() ([]byte, error) {
	key := struct {
		PrivateKey []byte `json:"private_key"`
	}{k.PrivateKey.Serialize()}
	return json.Marshal(&key)
}

// UnmarshalJSON does unmarshal keys
func (k *Keys) UnmarshalJSON(d []byte) error {
	var key struct {
		PrivateKey []byte `json:"private_key"`
	}

	err := json.Unmarshal(d, &key)
	if err != nil {
		return err
	}

	k.PrivateKey, k.PublicKey = btcec.PrivKeyFromBytes(btcec.S256(), key.PrivateKey)

	pubKey := k.PublicKey.SerializeCompressed()
	k.AddressPKH, err = btcutil.NewAddressPubKeyHash(btcutil.Hash160(pubKey), &chaincfg.Params{})

	if err != nil {
		return errors.Wrap(err, "failed to create address public key hash during unmarshal")
	}

	return err
}
