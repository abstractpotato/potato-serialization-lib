package main

import (
  "os"
  "fmt"
  "time"
  PSL "github.com/abstractpotato/potato-serialization-lib/psl"
  Builders "github.com/abstractpotato/potato-serialization-lib/builders"
)

func loadPrivateKey() ([]byte, error) {
  privateKey, err := os.ReadFile(".env/skey")
  if err != nil { return nil, err }
  return privateKey[:96], nil
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
  privateKey, err := loadPrivateKey()
  if err != nil { panic(err) }

  params := loadParams()

  tx := createBasicTx(params, privateKey)

  blockBuilder := Builders.NewBlockBuilder()
  blockBuilder.Params = params
  blockBuilder.Seed = []byte("bonepool")
  blockBuilder.AddTx(tx)
  blockBuilder.Build()

  start := time.Now()
  err = blockBuilder.Sign(privateKey)
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
