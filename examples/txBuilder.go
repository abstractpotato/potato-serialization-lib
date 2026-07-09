package main

import "fmt"
import PSL "github.com/abstractpotato/potato-serialization-lib/psl"
import Builders "github.com/abstractpotato/potato-serialization-lib/builders"
import wrapper "github.com/abstractpotato/cardano-signature-wrapper"

// these transactions do not have signatures yet

func main() {
  // sample data
  params := PSL.NewParams()
  params.Network = 0
  params.MaxTxSize = 4000
  params.MinTxFee = 222310
  params.TxFeePerByte = 430

  // // simple 1 receiver 1 asset transaction
  createBasicTx(params)

  // 1 receiver multiple asset transaction
  createMultiAssetTx(params)

  // 1 asset multiple receivers transaction
  createMultiAddrTx(params)

  // validator registration
  createRequestTx(params)
}

func createBasicTx(params PSL.Params) {
  txBuilder := Builders.NewTxBuilder()
  txBuilder.Params = params

  output := PSL.SimpleOutput{}
  output.To = "target_cardano_addr"
  output.Asset = "policy_id+asset_name"
  output.Amount = 10000

  txBuilder.AddSimpleOutput(output)
  txBuilder.Build()

  signature, err := wrapper.Sign(txBuilder.Tx.Header.Hash)
  if err != nil { fmt.Println(err) }
  txBuilder.Tx.Header.Signature = signature

  txJSON, _ := txBuilder.Tx.ToJSON()
  fmt.Printf("Simple Transaction:\n%s\n", string(txJSON))
  txHeaderCBOR, _ := txBuilder.Tx.Header.ToCBOR()
  fmt.Printf("Transaction Header Size: %v bytes\n", len(txHeaderCBOR))
  txBodyCBOR, _ := txBuilder.Tx.Body.ToCBOR()
  fmt.Printf("Transaction Body Size: %v bytes\n", len(txBodyCBOR))
  txCBOR, _ := txBuilder.Tx.ToCBOR()
  fmt.Printf("Transaction Size: %v bytes\n\n", len(txCBOR))
}

func createMultiAssetTx(params PSL.Params) {
  txBuilder := Builders.NewTxBuilder()
  txBuilder.Params = params

  asset0 := PSL.AssetOutput{}
  asset0.Asset = "policy_id+asset_name"
  asset0.Amount = 1000

  asset1 := PSL.AssetOutput{}
  asset1.Asset = "policy_id+asset_name"
  asset1.Amount = 1000

  outputs := PSL.NewMultiAssetOutput()
  outputs.AddAssetOutput(asset0)
  outputs.AddAssetOutput(asset1)

  txBuilder.AddMultiAssetOutput(outputs)
  txBuilder.Build()

  signature, err := wrapper.Sign(txBuilder.Tx.Header.Hash)
  if err != nil { fmt.Println(err) }
  txBuilder.Tx.Header.Signature = signature

  txJSON, _ := txBuilder.Tx.ToJSON()
  fmt.Printf("Multi-Asset Transaction:\n%s\n", string(txJSON))
  txHeaderCBOR, _ := txBuilder.Tx.Header.ToCBOR()
  fmt.Printf("Transaction Header Size: %v bytes\n", len(txHeaderCBOR))
  txBodyCBOR, _ := txBuilder.Tx.Body.ToCBOR()
  fmt.Printf("Transaction Body Size: %v bytes\n", len(txBodyCBOR))
  txCBOR, _ := txBuilder.Tx.ToCBOR()
  fmt.Printf("Transaction Size: %v bytes\n\n", len(txCBOR))
}

func createMultiAddrTx(params PSL.Params) {
  txBuilder := Builders.NewTxBuilder()
  txBuilder.Params = params

  addr0 := PSL.AddrOutput{}
  addr0.Addr = "target_cardano_addr"
  addr0.Amount = 1000

  outputs := PSL.NewMultiAddrOutput()
  outputs.AddAddrOutput(addr0)
  outputs.AddAddrOutput(addr0)

  txBuilder.AddMultiAddrOutput(outputs)
  txBuilder.Build()

  signature, err := wrapper.Sign(txBuilder.Tx.Header.Hash)
  if err != nil { fmt.Println(err) }
  txBuilder.Tx.Header.Signature = signature

  txJSON, _ := txBuilder.Tx.ToJSON()
  fmt.Printf("Multi-Addr Transaction:\n%s\n", string(txJSON))
  txHeaderCBOR, _ := txBuilder.Tx.Header.ToCBOR()
  fmt.Printf("Transaction Header Size: %v bytes\n", len(txHeaderCBOR))
  txBodyCBOR, _ := txBuilder.Tx.Body.ToCBOR()
  fmt.Printf("Transaction Body Size: %v bytes\n", len(txBodyCBOR))
  txCBOR, _ := txBuilder.Tx.ToCBOR()
  fmt.Printf("Transaction Size: %v bytes\n\n", len(txCBOR))
}

func createRequestTx(params PSL.Params) {
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

  signature, err := wrapper.Sign(txBuilder.Tx.Header.Hash)
  if err != nil { fmt.Println(err) }
  txBuilder.Tx.Header.Signature = signature

  txJSON, _ := txBuilder.Tx.ToJSON()
  fmt.Printf("Request Transaction:\n%s\n", string(txJSON))
  txHeaderCBOR, _ := txBuilder.Tx.Header.ToCBOR()
  fmt.Printf("Transaction Header Size: %v bytes\n", len(txHeaderCBOR))
  txBodyCBOR, _ := txBuilder.Tx.Body.ToCBOR()
  fmt.Printf("Transaction Body Size: %v bytes\n", len(txBodyCBOR))
  txCBOR, _ := txBuilder.Tx.ToCBOR()
  fmt.Printf("Transaction Size: %v bytes\n\n", len(txCBOR))
}
