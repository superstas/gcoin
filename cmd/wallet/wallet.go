package main

import (
	"context"
	"log"
	"os"

	"github.com/superstas/gcoin/gcoin/wallet"
)

func main() {
	l := log.New(os.Stdout, "[wallet]: ", log.Lshortfile)
	s := wallet.NewFileStorage("alice_wallet.dat")
	w, err := s.Init(context.Background())
	if err != nil {
		l.Fatal(err)
	}

	for _, k := range w.Keys() {
		l.Println(k.AddressPKH.EncodeAddress())
	}
}
