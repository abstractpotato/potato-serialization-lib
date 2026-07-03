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
  
  fmt.Printf("Original Block:\n %+v\n\n", block)
  
  blockHex, _ := block.ToHex()
  fmt.Printf("Block Hex:\n %s\n\n", blockHex)
  
  // to json 
  blockJSON, _ := block.ToJSON()
  fmt.Printf("Block JSON:\n %s\n\n", blockJSON)
  
  // Convert struct to CBOR
  blockCBOR, _ := block.ToCBOR()
  fmt.Printf("Block CBOR:\n %s\n\n", blockCBOR)
  
  // rebuild struct from CBOR
  newBlockCBOR, _ := ledger.BlockFromCBOR(blockCBOR)
  fmt.Printf("Reconstructed Block CBOR:\n %+v\n\n", newBlockCBOR)
  
  // rebuild block from HEX
  newBlockHex, _ := ledger.BlockFromHex(blockHex)
  fmt.Printf("Reconstructed Block HEX:\n %+v\n\n", newBlockHex)
}