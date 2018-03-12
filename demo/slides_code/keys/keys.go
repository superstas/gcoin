package main

import (
	"fmt"
	"log"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

func main() {
	privKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		log.Fatal(err)
	}

	d := chainhash.DoubleHashB([]byte("super secret data"))
	signature, err := privKey.Sign(d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("PubKey: %x\n", privKey.PubKey().SerializeUncompressed())
	fmt.Printf("Signature: %x\n", signature.Serialize())
	fmt.Printf("Is sign valid: %v\n", signature.Verify(d, privKey.PubKey()))
}

//PubKey: 04144837e3cb253214f095744930f35c6f3b1afacc46278ae5e6f528535f1180c4d22f0a63430877bed6e026ced53c285f0392b1ccafae5cae5df1db246e548f38
//Signature: 304402206b22f3e5cc357a17f18539e273235f8ad548ff5c44c9e8c3ae66dde7323d456602203aa4a9cae21c0b6b2caaecf3e4d98135952ce2957df6775a66b040e4fef763c8
//Is sign valid: true
