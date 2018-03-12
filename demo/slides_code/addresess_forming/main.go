package main

import (
	"fmt"
	"log"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

func main() {
	privKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		log.Fatal(err)
	}

	pubKey := privKey.PubKey().SerializeUncompressed()
	addr, err := btcutil.NewAddressPubKeyHash(btcutil.Hash160(pubKey), &chaincfg.Params{PubKeyHashAddrID: 0x00})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("PubKey: %x\n", privKey.PubKey().SerializeUncompressed())
	fmt.Printf("Address: %s\n", addr.EncodeAddress())
}

//PubKey: 04bc98f0be289684367bd04b262b7ef0fea7bd292cdc978344a578644251ed9fe4c5e81b494318e1aaaa8fc1492250bbe86bf2d18f1b9defef1d40d2fd2d8c84e7
//Address: 1Pwp7sqxYtrTkJExuZmHMfbh9ffbqzk6sH
