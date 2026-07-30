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

func main() {
  // sample param data
  params := psl.NewParams()
  params.Network = 0
  params.MaxTxSize = 4000
  params.TxFeePerByte = 430
  params.MinTxFee = params.TxFeePerByte * 173 // signature size
  
  privateKey := GetPrivateKey()
  addr, err := psl.GenerateEnterpriseAddr(privateKey, true)
  if err != nil { panic(err) }
  
  txBuilder := psl.NewTxBuilder()
  txBuilder.Params = params

  asset0 := psl.AssetOutput{}
  asset0.Asset = ""
  asset0.Amount = 1000

  asset1 := psl.AssetOutput{}
  asset1.Asset = "3d77d63dfa6033be98021417e08e3368cc80e67f8d7afa196aaa0b3953746172636820546f6b656e"
  asset1.Amount = 1000

  output := psl.NewMultiOutput()
  output.From = "3d77d63dfa6033be98021417e08e3368cc80e67f8d7afa196aaa0b3953746172636820546f6b656e"
  output.To = addr
  output.Add(asset0)
  output.Add(asset1)
  
  txBuilder.AddMultiOutput(output)
  txBuilder.Build()

  start := time.Now()
  err = txBuilder.Sign(privateKey)
  if err != nil { panic(err) }
  fmt.Printf("Signature took %s\n", time.Since(start))

  txJSON, _ := txBuilder.Tx.ToJSON()
  fmt.Printf("Multi-Asset Transaction:\n%s\n", string(txJSON))
  txHeaderCBOR, _ := txBuilder.Tx.Header.ToCBOR()
  fmt.Printf("Transaction Header Size: %v bytes\n", len(txHeaderCBOR))
  txBodyCBOR, _ := txBuilder.Tx.Body.ToCBOR()
  fmt.Printf("Transaction Body Size: %v bytes\n", len(txBodyCBOR))
  txCBOR, _ := txBuilder.Tx.ToCBOR()
  fmt.Printf("Transaction Size: %v bytes\n", len(txCBOR))

  start = time.Now()
  fmt.Printf("Transaction Verification: %v\n", txBuilder.Verify())
  fmt.Printf("Verification took %s\n\n", time.Since(start))
}