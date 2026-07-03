package main

import (
  "fmt"
  "potato-serialization-lib/ledger"
)

func main() {
  blockTest()
}

func blockTest() {
  block := ledger.NewBlock()
  block.Hash()
  fmt.Printf("Original Block:\n %+v\n", block)
  blockHex, _ := block.ToHex()
  
  fmt.Printf("Block Hex:\n %s\n", blockHex)
  
  // Convert struct to CBOR
  blockCBOR, _ := block.ToCBOR()
  fmt.Printf("Block CBOR:\n %s\n", blockCBOR)
  
  // rebuild struct from CBOR
  newBlock, _ := ledger.BlockFromCBOR(blockCBOR)
  fmt.Printf("Reconstructed Block:\n %+v\n", newBlock)
}