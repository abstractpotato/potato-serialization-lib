package main

import (
  "fmt"
  "potato-serialization-lib/ledger"
)

// this is an example of a hard coded genesis block
// does not yet include signatures

func main() {
  validator := ledger.NewValidator()
  validator.Addr = "test_addr"
  validator.CertificateTx = "genesis"
  validator.Relays = append(validator.Relays, "http://0.0.0.0:5001")
  validator.Relays = append(validator.Relays, "http://0.0.0.0:5002")
  validatorJSON, _ := validator.ToJSON()
  fmt.Printf("Genesis Validator:\n %s\n\n", validatorJSON)

  params := ledger.NewParams()
  params.Network = 0
  params.MaxBlockHeaderSize = 256
  params.MaxBlockBodySize = 2048000000
  params.MaxTxSize = 4096
  params.TxFeePerByte = 430
  params.SlotsPerEpoch = 432000
  params.SlotTimeInMs = 1000
  params.ProtocolVersion = 0
  paramsJSON, _ := params.ToJSON()
  fmt.Printf("Genesis Params:\n %+s\n\n", paramsJSON)

  epoch := ledger.NewEpoch()
  epoch.Header.ID = 0
  epoch.Body.StartTime = 0
  epoch.Body.EndTime = epoch.Body.StartTime + params.SlotsPerEpoch
  epoch.Body.PreviousHash = "0000"
  epoch.Body.Validators = append(epoch.Body.Validators, validator.Addr)
  epoch.Body.ProtocolVersion = 0
  epoch.Hash()
  fmt.Printf("Genesis Epoch:\n %+v\n\n", epoch)

  validatorData := ledger.TxData{}
  validatorData.Tag = "genesis_validator"
  validatorCBOR, _ := validator.ToCBOR()
  validatorData.Data = validatorCBOR
  validatorData.Type = 0

  paramsData := ledger.TxData{}
  paramsData.Tag = "genesis_params"
  paramsCBOR, _ := params.ToCBOR()
  paramsData.Data = paramsCBOR
  paramsData.Type = 0

  transaction := ledger.NewTransaction()
  transaction.Body.Data = append(transaction.Body.Data, validatorData)
  transaction.Body.Data = append(transaction.Body.Data, paramsData)

  txJSON, _ := transaction.ToJSON()
  fmt.Printf("Genesis Transaction:\n %s\n\n", txJSON)

  block := ledger.NewBlock()
  block.Header.ID = 0
  block.Header.Validator = validator.Addr
  // add validator signature as witness
  block.Body.VRF = "genesis"
  block.Body.Epoch = 0
  block.Body.Slot = 0
  block.Body.Transactions = append(block.Body.Transactions, transaction)
  block.Body.Fees = 0
  block.Body.Timestamp = 0
  block.Hash()

  blockJSON, _ := block.ToJSON()
  fmt.Printf("Genesis Block:\n %s\n\n", blockJSON)
}
