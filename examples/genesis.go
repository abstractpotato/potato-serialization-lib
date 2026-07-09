package main

import "fmt"
import PSL "github.com/abstractpotato/potato-serialization-lib/psl"
import Builders "github.com/abstractpotato/potato-serialization-lib/builders"
import wrapper "github.com/abstractpotato/cardano-signature-wrapper"

// this is an example of a hard coded genesis block
// does not yet include signatures

func main() {
  // genesis validator used to sign the transaction and block
  validator := PSL.NewValidator()
  validator.Addr = "cardano_addr"
  validator.CertificateTx = "genesis"
  validator.Relays = append(validator.Relays, "http://0.0.0.0:5001")
  validator.Relays = append(validator.Relays, "http://0.0.0.0:5002")
  validatorJSON, _ := validator.ToJSON()
  fmt.Printf("Genesis Validator:\n%s\n\n", validatorJSON)

  // initital protocol parameters
  params := PSL.NewParams()
  params.Network = 0
  params.MaxBlockHeaderSize = 1100 // 128 bytes
  params.MaxBlockBodySize = 40000000 // 40 MB or ~200k simple transactions
  params.MaxTxSize = 4000 // 4 KB
  params.MinTxFee = 236500
  params.TxFeePerByte = 430
  params.SlotsPerEpoch = 432000
  params.SlotTimeInMs = 1000
  params.ProtocolVersion = 0

  paramsJSON, _ := params.ToJSON()
  fmt.Printf("Genesis Params:\n%+s\n\n", paramsJSON)

  // convert validator to cbor and format into TxData
  validatorData := PSL.TxData{}
  validatorData.Tag = "genesis_validator"
  validatorCBOR, _ := validator.ToCBOR()
  validatorData.Data = validatorCBOR
  validatorData.Type = 0

  // convert parameters to cbor and format into TxData
  paramsData := PSL.TxData{}
  paramsData.Tag = "genesis_params"
  paramsCBOR, _ := params.ToCBOR()
  paramsData.Data = paramsCBOR
  paramsData.Type = 0

  txBuilder := Builders.NewTxBuilder()
  txBuilder.Tx.Header.Addr = validator.Addr
  txBuilder.AddData(paramsData)
  txBuilder.AddData(validatorData)
  txBuilder.Build()

  // txBodyHex, _ := txBuilder.Tx.Body.ToHex()
  signed, _ := wrapper.Sign(txBuilder.Tx.Header.Hash)
  txBuilder.Tx.Header.Signature = signed

  txJSON, _ := txBuilder.Tx.ToJSON()
  fmt.Printf("Transaction:\n%s\n", string(txJSON))

  txHeaderCBOR, _ := txBuilder.Tx.Header.ToCBOR()
  fmt.Printf("Transaction Header Size: %v bytes\n", len(txHeaderCBOR))
  txBodyCBOR, _ := txBuilder.Tx.Body.ToCBOR()
  fmt.Printf("Transaction Body Size: %v bytes\n", len(txBodyCBOR))
  txCBOR, _ := txBuilder.Tx.ToCBOR()
  fmt.Printf("Transaction Size: %v bytes\n\n", len(txCBOR))

  // add transaction into block body
  block := PSL.NewBlock()
  block.Body.PreviousHash = "potato"
  block.Body.VRF = "genesis"
  block.Body.Transactions = append(block.Body.Transactions, txBuilder.Tx)
  block.Body.Timestamp = txBuilder.Tx.Body.Timestamp
  block.Hash() // generate block hash

  // sign blockCBOR with PaymentKey
  // blockBodyHex, _ := block.Body.ToHex()
  blockSignature, _ := wrapper.Sign(block.Header.Hash)

  witness := PSL.Witness{}
  witness.Addr = validator.Addr
  witness.Signature = blockSignature
  block.AddWitness(witness)

  blockJSON, _ := block.ToJSON()
  fmt.Printf("Genesis Block:\n%+s\n", blockJSON)

  blockHeaderCBOR, _ := block.Header.ToCBOR()
  fmt.Printf("Block Header Size: %v bytes\n", len(blockHeaderCBOR))

  blockBodyCBOR, _ := block.Body.ToCBOR()
  fmt.Printf("Block Body Size: %v bytes\n", len(blockBodyCBOR))

  blockCBOR, _ := block.ToCBOR()
  fmt.Printf("Block Size: %v bytes\n", len(blockCBOR))
}
