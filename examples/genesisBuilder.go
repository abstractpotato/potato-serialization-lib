package main

import (
  "os"
  // "fmt"
  // "time"
  PSL "github.com/abstractpotato/potato-serialization-lib/psl"
)

func loadPrivateKey() ([]byte, error) {
  privateKey, err := os.ReadFile(".env/skey")
  if err != nil { return nil, err }
  return privateKey[:96], nil
}

func main() {
  // privateKey, err := loadPrivateKey()
  // if err != nil { panic(err) }

  // initital protocol parameters
  params := PSL.NewParams()
  params.Network = 0
  params.MaxBlockHeaderSize = 1100 // 128 bytes
  params.MaxBlockBodySize = 40000000 // 40 MB or ~200k simple transactions
  params.MaxTxSize = 4000 // 4 KB
  params.TxFeePerByte = 430
  params.MinTxFee = params.TxFeePerByte * 175 // signature size
  params.SlotsPerEpoch = 432000
  params.SlotTimeInMs = 1000
  params.ProtocolVersion = 0

  // add cert
  // add genesis seed

  genesis := PSL.Genesis{}
  genesis.Params = params

  block := PSL.NewBlock()
  block.Body.Genesis = genesis

  // sign & verify
}
