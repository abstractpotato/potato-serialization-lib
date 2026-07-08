package main

import "fmt"
import PSL "github.com/abstractpotato/potato-serialization-lib/psl"
import Builders "github.com/abstractpotato/potato-serialization-lib/builders"

// these transactions do not have signatures yet

func main() {
  // sample data
  params := PSL.NewParams()
  params.Network = 0
  params.MaxTxSize = 4000
  params.MinTxFee = 4300
  params.TxFeePerByte = 430

  // validator registration
  // createRequestTx(params)

  // // simple 1 receiver 1 asset transaction
  createBasicTx(params)
  //
  // // 1 receiver multiple asset transaction
  // createMultiAssetTx(params)
  //
  // // 1 asset multiple receivers transaction
  // createMultiAddrTx(params)
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

  // txBodyHex, err := txBuilder.Tx.BodyToHex()
  // fmt.Println(err)
  // signature, err := wrapper.Sign(txBodyHex)
  // fmt.Println(err)

  // txBuilder.Tx.Header.Signature = signature

  fmt.Printf("%+v\n", txBuilder.Tx)
}

func createMultiAssetTx(params PSL.Params) {
  txBuilder := Builders.NewTxBuilder()
  txBuilder.Params = params

  asset0 := PSL.AssetOutput{}
  asset0.ID = "policy_id+asset_name"
  asset0.Amount = 1000

  asset1 := PSL.AssetOutput{}
  asset1.ID = "policy_id+asset_name"
  asset1.Amount = 1000

  outputs := PSL.NewMultiAssetOutput()
  outputs.AddAssetOutput(asset0)
  outputs.AddAssetOutput(asset1)

  txBuilder.AddMultiAssetOutput(outputs)
  txBuilder.Build()

  fmt.Printf("%+v\n", txBuilder.Tx)
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

  fmt.Printf("%+v\n", txBuilder.Tx)
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

  fmt.Printf("%+v\n", txBuilder.Tx)
}
