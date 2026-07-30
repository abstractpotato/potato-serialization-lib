package main

import (
  "fmt"
  "time"
  "github.com/abstractpotato/potato-serialization-lib"
)

const asset = "3d77d63dfa6033be98021417e08e3368cc80e67f8d7afa196aaa0b3953746172636820546f6b656e"

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
  
  // simple 1 receiver 1 asset transaction
  createBasicTx(params, privateKey, addr)

  // // 1 receiver multiple asset transaction
  createMultiAssetTx(params, privateKey, addr)
  // 
  // // 1 asset multiple receivers transaction
  createMultiAddrTx(params, privateKey, addr)
  // 
  // // validator registration
  createRequestTx(params, privateKey, addr)
}

func createBasicTx(params psl.Params, privateKey []byte, addr string) {
  txBuilder := psl.NewTxBuilder()
  txBuilder.Params = params

  output := psl.SimpleOutput{}
  output.To = addr
  output.Asset = asset
  output.Amount = 10000
  
  txBuilder.SetSender(addr)
  txBuilder.SetSimpleOutput(&output)
  txBuilder.Build()
  start := time.Now()
  err := txBuilder.Sign(privateKey)
  if err != nil { panic(err) }
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

func createMultiAssetTx(params psl.Params, privateKey []byte, addr string) {
  txBuilder := psl.NewTxBuilder()
  txBuilder.Params = params

  asset0 := psl.AssetOutput{}
  asset0.Asset = asset
  asset0.Amount = 1000

  asset1 := psl.AssetOutput{}
  asset1.Asset = asset
  asset1.Amount = 1000

  output := psl.NewMultiAssetOutput()
  output.To = addr
  output.Add(asset0)
  output.Add(asset1)
  
  txBuilder.SetSender(addr)
  txBuilder.SetMultiAssetOutput(&output)
  txBuilder.Build()

  start := time.Now()
  err := txBuilder.Sign(privateKey)
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

func createMultiAddrTx(params psl.Params, privateKey []byte, addr string) {
  txBuilder := psl.NewTxBuilder()
  txBuilder.Params = params

  addr0 := psl.AddrOutput{}
  addr0.Addr = addr
  addr0.Amount = 1000

  output := psl.NewMultiAddrOutput()
  output.Asset = asset
  output.Add(addr0)
  output.Add(addr0)
  
  txBuilder.SetSender(addr)
  txBuilder.SetMultiAddrOutput(&output)
  txBuilder.Build()

  start := time.Now()
  err := txBuilder.Sign(privateKey)
  if err != nil { panic(err) }
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

func createRequestTx(params psl.Params, privateKey []byte, addr string) {
  txBuilder := psl.NewTxBuilder()
  txBuilder.Params = params

  request := psl.NewRequest()
  request.Ticker = "bone"
  request.Url = "https://bonepool.com"
  request.Addr = "rewards_addr"
  request.Relays = append(request.Relays, "0.0.0.0:5001")
  request.Relays = append(request.Relays, "0.0.0.0:5002")
  txBuilder.SetRequest(&request)
  
  txBuilder.SetSender(addr)
  
  txBuilder.Build()
  start := time.Now()
  err := txBuilder.Sign(privateKey)
  if err != nil { panic(err) }
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
