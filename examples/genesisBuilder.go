package main

import (
  "fmt"
  "time"
  "encoding/hex"
  PSL "github.com/abstractpotato/potato-serialization-lib"
)

const skey = "c0e5981efee192773da5a3542b28da40b48638eff0bf5495dc016f4ecc0c55534b0853da95378d4ecbf184920b1dec5747212915977718b5094ef0c45ee0cfb0a8f448cbb86544765fa7ae7a0ef604768c10054de52498d59ba00995ca6ec66696bcefe574605f16a8166e3219a1a012fc04c6f1929003917f9f805784930784"

func GetPrivateKey() []byte {
  privateKey, err := hex.DecodeString(skey)
  if err != nil { panic(err) }
  return privateKey[:96]
}

func main() {
  privateKey := GetPrivateKey()

  // initital protocol parameters
  params := PSL.NewParams()
  params.Network = 0
  params.MaxBlockHeaderSize = 1100 // 128 bytes
  params.MaxBlockBodySize = 4000000 // 4 MB or ~15k simple transactions
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
  err := block.Sign(privateKey)
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
