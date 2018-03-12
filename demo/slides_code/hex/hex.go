package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
)

func main() {
	blockHash, _ := hex.DecodeString("0001AA")
	targetHash, _ := hex.DecodeString("0001FF")

	blockHashInt := new(big.Int).SetBytes(blockHash)
	targetHashInt := new(big.Int).SetBytes(targetHash)

	fmt.Printf("BlockHashInt: %s\n", blockHashInt.String())
	fmt.Printf("TargetHashInt: %s\n", targetHashInt.String())
	fmt.Printf("BlockHashInt <= TargetHashInt: %v\n", blockHashInt.Cmp(targetHashInt) <= 0)
}
