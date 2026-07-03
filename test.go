package main

import (
  "fmt"
  "potato-serialization-lib/ledger"
)

func main() {
  epochTest()
  blockTest()
}

func epochTest() {
  epoch := ledger.NewEpoch()
  epoch.Hash()
  
  fmt.Printf("Original Epoch:\n %+v\n\n", epoch)
  
  epochHex, _ := epoch.ToHex()
  fmt.Printf("Epoch Hex:\n %s\n\n", epochHex)
  
  // to json 
  epochJSON, _ := epoch.ToJSON()
  fmt.Printf("Epoch JSON:\n %s\n\n", epochJSON)
  
  // Convert struct to CBOR
  epochCBOR, _ := epoch.ToCBOR()
  fmt.Printf("Epoch CBOR:\n %s\n\n", epochCBOR)
  
  // rebuild struct from CBOR
  newEpochCBOR, _ := ledger.EpochFromCBOR(epochCBOR)
  fmt.Printf("Reconstructed Epoch CBOR:\n %+v\n\n", newEpochCBOR)
  
  // rebuild epoch from HEX
  newEpochHex, _ := ledger.EpochFromHex(epochHex)
  fmt.Printf("Reconstructed Epoch HEX:\n %+v\n\n", newEpochHex)
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