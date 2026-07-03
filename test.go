package main

import (
  "fmt"
  "github.com/abstractpotato/potato-serialization-lib/ledger"
)

func main() {
  paramTest()
  validatorTest()
  epochTest()
  blockTest()
  txTest()
  accountTest()
}

func paramTest() {
  params := ledger.NewParams()

  fmt.Printf("Original Params:\n %+v\n\n", params)

  paramHex, _ := params.ToHex()
  fmt.Printf("Params Hex:\n %s\n\n", paramHex)

  paramJSON, _ := params.ToJSON()
  fmt.Printf("Params JSON:\n %s\n\n", paramJSON)

  paramCBOR, _ := params.ToCBOR()
  fmt.Printf("Params CBOR:\n %s\n\n", paramCBOR)

  newParamsCBOR, _ := ledger.ParamsFromCBOR(paramCBOR)
  fmt.Printf("Reconstructed Params CBOR:\n %+v\n\n", newParamsCBOR)

  newParamsHex, _ := ledger.ParamsFromHex(paramHex)
  fmt.Printf("Reconstructed Params HEX:\n %+v\n\n", newParamsHex)
}

func validatorTest()  {
  validator := ledger.NewValidator()

  fmt.Printf("Original Validator:\n %+v\n\n", validator)

  validatorHex, _ := validator.ToHex()
  fmt.Printf("Validator Hex:\n %s\n\n", validatorHex)

  validatorJSON, _ := validator.ToJSON()
  fmt.Printf("Validator JSON:\n %s\n\n", validatorJSON)

  validatorCBOR, _ := validator.ToCBOR()
  fmt.Printf("Validator CBOR:\n %s\n\n", validatorCBOR)

  newValidatorCBOR, _ := ledger.ValidatorFromCBOR(validatorCBOR)
  fmt.Printf("Reconstructed Validator CBOR:\n %+v\n\n", newValidatorCBOR)

  newValidatorHex, _ := ledger.ValidatorFromHex(validatorHex)
  fmt.Printf("Reconstructed Validator HEX:\n %+v\n\n", newValidatorHex)

}

func epochTest() {
  epoch := ledger.NewEpoch()
  epoch.Hash()

  fmt.Printf("Original Epoch:\n %+v\n\n", epoch)

  epochHex, _ := epoch.ToHex()
  fmt.Printf("Epoch Hex:\n %s\n\n", epochHex)

  // to json
  epochJSON, _ := epoch.ToJSON()
  fmt.Printf("Epoch JSON:\n %s\n\n", epochJSON)

  // Convert struct to CBOR
  epochCBOR, _ := epoch.ToCBOR()
  fmt.Printf("Epoch CBOR:\n %s\n\n", epochCBOR)

  // rebuild struct from CBOR
  newEpochCBOR, _ := ledger.EpochFromCBOR(epochCBOR)
  fmt.Printf("Reconstructed Epoch CBOR:\n %+v\n\n", newEpochCBOR)

  // rebuild epoch from HEX
  newEpochHex, _ := ledger.EpochFromHex(epochHex)
  fmt.Printf("Reconstructed Epoch HEX:\n %+v\n\n", newEpochHex)
}

func blockTest() {
  block := ledger.NewBlock()
  block.Hash()

  fmt.Printf("Original Block:\n %+v\n\n", block)

  blockHex, _ := block.ToHex()
  fmt.Printf("Block Hex:\n %s\n\n", blockHex)

  // to json
  blockJSON, _ := block.ToJSON()
  fmt.Printf("Block JSON:\n %s\n\n", blockJSON)

  // Convert struct to CBOR
  blockCBOR, _ := block.ToCBOR()
  fmt.Printf("Block CBOR:\n %s\n\n", blockCBOR)

  // rebuild struct from CBOR
  newBlockCBOR, _ := ledger.BlockFromCBOR(blockCBOR)
  fmt.Printf("Reconstructed Block CBOR:\n %+v\n\n", newBlockCBOR)

  // rebuild block from HEX
  newBlockHex, _ := ledger.BlockFromHex(blockHex)
  fmt.Printf("Reconstructed Block HEX:\n %+v\n\n", newBlockHex)
}

func txTest() {
  tx := ledger.NewTransaction()
  tx.Hash()

  fmt.Printf("Original Transaction:\n %+v\n\n", tx)

  txHex, _ := tx.ToHex()
  fmt.Printf("Transaction Hex:\n %s\n\n", txHex)

  // to json
  txJSON, _ := tx.ToJSON()
  fmt.Printf("Transaction JSON:\n %s\n\n", txJSON)

  // Convert struct to CBOR
  txCBOR, _ := tx.ToCBOR()
  fmt.Printf("Transaction CBOR:\n %s\n\n", txCBOR)

  // rebuild struct from CBOR
  newTransactionCBOR, _ := ledger.TransactionFromCBOR(txCBOR)
  fmt.Printf("Reconstructed Transaction CBOR:\n %+v\n\n", newTransactionCBOR)

  // rebuild tx from HEX
  newTransactionHex, _ := ledger.TransactionFromHex(txHex)
  fmt.Printf("Reconstructed Transaction HEX:\n %+v\n\n", newTransactionHex)
}

func accountTest() {
  account := ledger.NewAccount()

  fmt.Printf("Original Account:\n %+v\n\n", account)

  paramHex, _ := account.ToHex()
  fmt.Printf("Account Hex:\n %s\n\n", paramHex)

  paramJSON, _ := account.ToJSON()
  fmt.Printf("Account JSON:\n %s\n\n", paramJSON)

  paramCBOR, _ := account.ToCBOR()
  fmt.Printf("Account CBOR:\n %s\n\n", paramCBOR)

  newAccountCBOR, _ := ledger.AccountFromCBOR(paramCBOR)
  fmt.Printf("Reconstructed Account CBOR:\n %+v\n\n", newAccountCBOR)

  newAccountHex, _ := ledger.AccountFromHex(paramHex)
  fmt.Printf("Reconstructed Account HEX:\n %+v\n\n", newAccountHex)
}
