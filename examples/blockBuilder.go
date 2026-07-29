package main

import (
  "fmt"
  "time"
  "encoding/hex"
  PSL "github.com/abstractpotato/potato-serialization-lib"
  Builders "github.com/abstractpotato/potato-serialization-lib/builders"
)

const skey = "c0e5981efee192773da5a3542b28da40b48638eff0bf5495dc016f4ecc0c55534b0853da95378d4ecbf184920b1dec5747212915977718b5094ef0c45ee0cfb0a8f448cbb86544765fa7ae7a0ef604768c10054de52498d59ba00995ca6ec66696bcefe574605f16a8166e3219a1a012fc04c6f1929003917f9f805784930784"

func GetPrivateKey() []byte {
  privateKey, err := hex.DecodeString(skey)
  if err != nil { panic(err) }
  return privateKey[:96]
}


func loadParams() PSL.Params {
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
  return params
}

func createBasicTx(params PSL.Params, privateKey []byte) PSL.Transaction {
  txBuilder := Builders.NewTxBuilder()
  txBuilder.Params = params

  output := PSL.SimpleOutput{}
  output.To = "target_cardano_addr"
  output.Asset = "policy_id+asset_name"
  output.Amount = 10000

  txBuilder.AddSimpleOutput(output)
  txBuilder.Build()
  err := txBuilder.Sign(privateKey)
  if err != nil { panic(err) }
  return txBuilder.Tx
}

func main () {
  privateKey := GetPrivateKey()

  params := loadParams()

  tx := createBasicTx(params, privateKey)

  blockBuilder := Builders.NewBlockBuilder()
  blockBuilder.Params = params
  blockBuilder.Seed = []byte("bonepool")
  blockBuilder.AddTx(tx)
  blockBuilder.Build()

  start := time.Now()
  err := blockBuilder.Sign(privateKey)
  if err != nil { panic(err) }
  fmt.Printf("Signature took %s\n", time.Since(start))

  blockJSON, _ := blockBuilder.Block.ToJSON()
  fmt.Printf("Block :\n%s\n", blockJSON)

  blockCBOR, _ := blockBuilder.Block.ToCBOR()
  fmt.Printf("Block Size: %v bytes\n", len(blockCBOR))

  blockHeaderCBOR, _ := blockBuilder.Block.Header.ToCBOR()
  fmt.Printf("Block Header Size: %v bytes\n", len(blockHeaderCBOR))

  blockBodyCBOR, _ := blockBuilder.Block.Body.ToCBOR()
  fmt.Printf("Block Body Size: %v bytes\n", len(blockBodyCBOR))

  start = time.Now()
  fmt.Printf("Block Verification: %v\n", blockBuilder.Verify())
  fmt.Printf("Verification took %s\n\n", time.Since(start))
}
