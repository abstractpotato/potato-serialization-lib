package main

import (
  "os"
  "fmt"
  "time"
  PSL "github.com/abstractpotato/potato-serialization-lib/psl"
)

func loadPrivateKey() ([]byte, error) {
  privateKey, err := os.ReadFile(".env/skey")
  if err != nil { return nil, err }
  return privateKey[:96], nil
}

func main() {
  privateKey, err := loadPrivateKey()
  if err != nil { panic(err) }

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

  // initial node certificate
  cert := PSL.NewCertificate()
  cert.RequestTx = "genesis"
  cert.RewardAddr = "genesis"
  cert.AddRelay("0.0.0.0:5001")
  cert.AddRelay("0.0.0.0:5002")
  cert.Status = 1

  genesis := PSL.Genesis{}
  genesis.Seed = []byte("bonepool")
  genesis.Certificate = cert
  genesis.Params = params

  block := PSL.NewBlock()
  block.Body.Genesis = &genesis
  block.Body.Timestamp = uint(time.Now().UnixMilli())
  block.Hash()

  start := time.Now()
  err = block.Sign(privateKey)
  if err != nil { panic(err) }
  fmt.Printf("Signature took %s\n", time.Since(start))

  blockJSON, _ := block.ToJSON()
  fmt.Printf("Genesis Block Demo:\n%s\n", blockJSON)

  blockHeaderBOR, _ := block.Header.ToCBOR()
  fmt.Printf("Block Header Size: %v bytes\n", len(blockHeaderBOR))
  
  blockBodyCBOR, _ := block.Body.ToCBOR()
  fmt.Printf("Block Body Size: %v bytes\n", len(blockBodyCBOR))

  start = time.Now()
  fmt.Printf("Block Verification: %v\n", block.Verify())
  fmt.Printf("Verification took %s\n\n", time.Since(start))
}
