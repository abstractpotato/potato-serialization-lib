package main

import (
  "fmt"
  "time"
  "github.com/abstractpotato/potato-serialization-lib"
)

func GetPrivateKey() []byte {
  skey, err := psl.GenerateKeys("")
  if err != nil { panic(err) }
  return skey
}

func loadParams() psl.Params {
  // initital protocol parameters
  params := psl.NewParams()
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

func createBasicTx(params psl.Params, privateKey []byte) psl.Transaction {
  txBuilder := psl.NewTxBuilder()
  txBuilder.Params = params

  output := psl.SimpleOutput{}
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

  blockBuilder := psl.NewBlockBuilder()
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
