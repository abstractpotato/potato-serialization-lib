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

func main() {
  // sample param data
  params := PSL.NewParams()
  params.Network = 0
  params.MaxTxSize = 4000
  params.TxFeePerByte = 430
  params.MinTxFee = params.TxFeePerByte * 175 // signature size

  privateKey, err := loadPrivateKey()
  if err != nil { panic(err) }

  // simple 1 receiver 1 asset transaction
  createBasicTx(params, privateKey)

  // 1 receiver multiple asset transaction
  createMultiAssetTx(params, privateKey)

  // 1 asset multiple receivers transaction
  createMultiAddrTx(params, privateKey)

  // validator registration
  createRequestTx(params, privateKey)
}

func createBasicTx(params PSL.Params, privateKey []byte) {
  txBuilder := Builders.NewTxBuilder()
  txBuilder.Params = params

  output := PSL.SimpleOutput{}
  output.To = "target_cardano_addr"
  output.Asset = "policy_id+asset_name"
  output.Amount = 10000

  txBuilder.AddSimpleOutput(output)
  txBuilder.Build()
  start := time.Now()
  err := txBuilder.Sign(privateKey)
  if err != nil { fmt.Println(err) }
  fmt.Printf("Signature took %s\n", time.Since(start))

  txJSON, _ := txBuilder.Tx.ToJSON()
  fmt.Printf("Simple Transaction:\n%s\n", string(txJSON))
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

func createMultiAssetTx(params PSL.Params, privateKey []byte) {
  txBuilder := Builders.NewTxBuilder()
  txBuilder.Params = params

  asset0 := PSL.AssetOutput{}
  asset0.Asset = "policy_id+asset_name"
  asset0.Amount = 1000

  asset1 := PSL.AssetOutput{}
  asset1.Asset = "policy_id+asset_name"
  asset1.Amount = 1000

  output := PSL.NewMultiAssetOutput()
  output.Add(asset0)
  output.Add(asset1)

  txBuilder.AddMultiAssetOutput(output)
  txBuilder.Build()

  start := time.Now()
  err := txBuilder.Sign(privateKey)
  if err != nil { fmt.Println(err) }
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

func createMultiAddrTx(params PSL.Params, privateKey []byte) {
  txBuilder := Builders.NewTxBuilder()
  txBuilder.Params = params

  addr0 := PSL.AddrOutput{}
  addr0.Addr = "target_cardano_addr"
  addr0.Amount = 1000

  outputs := PSL.NewMultiAddrOutput()
  outputs.Add(addr0)
  outputs.Add(addr0)

  txBuilder.AddMultiAddrOutput(outputs)
  txBuilder.Build()

  start := time.Now()
  err := txBuilder.Sign(privateKey)
  if err != nil { fmt.Println(err) }
  fmt.Printf("Signature took %s\n", time.Since(start))


  txJSON, _ := txBuilder.Tx.ToJSON()
  fmt.Printf("Multi-Addr Transaction:\n%s\n", string(txJSON))
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

func createRequestTx(params PSL.Params, privateKey []byte) {
  txBuilder := Builders.NewTxBuilder()
  txBuilder.Params = params

  request := PSL.NewRequest()
  request.Ticker = "bone"
  request.Url = "https://bonepool.com"
  request.Addr = "rewards_addr"
  request.Relays = append(request.Relays, "0.0.0.0:5001")
  request.Relays = append(request.Relays, "0.0.0.0:5002")
  txBuilder.AddRequest(request)

  txBuilder.Build()
  start := time.Now()
  err := txBuilder.Sign(privateKey)
  if err != nil { fmt.Println(err) }
  fmt.Printf("Signature took %s\n", time.Since(start))

  txJSON, _ := txBuilder.Tx.ToJSON()
  fmt.Printf("Request Transaction:\n%s\n", string(txJSON))
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
