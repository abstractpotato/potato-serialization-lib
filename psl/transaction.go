package psl

import(
  "fmt"
  "crypto/sha256"
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
)

type Transaction struct {
  Header TxHeader `cbor: "header"`
  Body   TxBody   `cbor: "body"`
}

type TxHeader struct {
  Hash      string `cbor: "hash"`
  Sender    string `cbor: "sender"`
  Signature []byte `cbor: "signature"`
}

type TxBody struct {
  Outputs   TxOutputs `cbor: "outputs"`
  Data      []TxData    `cbor: "data"`
  TTL       uint        `cbor: "ttl"`
  Timestamp uint        `cbor: "timestamp"`
  Network   uint        `cbor: "network"`
  Fee       uint        `cbor: "fee"`
}

type TxData struct {
  Tag  string `cbor: "tag"`
  Data []byte `cbor: "data"`
  Type uint   `cbor: "type"`
}

func NewTransaction() Transaction {
  return Transaction{
    Header: TxHeader{},
    Body: TxBody{
      Outputs: NewTxOutputs(),
      Data: make([]TxData, 0),
    },
  }
}

func TransactionFromCBOR(cborBytes []byte) (Transaction, error) {
  var transaction Transaction
  err := cbor.Unmarshal(cborBytes, &transaction)
  if err != nil { return NewTransaction(), err }
  return transaction, nil
}

func TransactionFromHex(hexString string) (Transaction, error) {
  cborBytes, err := hex.DecodeString(hexString)
  if err != nil { return NewTransaction(), err }
  transaction, err := TransactionFromCBOR(cborBytes)
  if err != nil { return NewTransaction(), err }
  return transaction, nil
}

func (transaction *Transaction) ToCBOR() ([]byte, error) {
  cborBytes, err := cbor.Marshal(transaction)
  if err != nil { return nil, err}
  return cborBytes, nil
}

func (transaction *Transaction) ToHex() (string, error) {
  cborBytes, err := transaction.ToCBOR()
  if err != nil { return "", err }
  return hex.EncodeToString(cborBytes), nil
}

func (transaction *Transaction) ToJSON() ([]byte, error) {
  jsonBytes, err := json.Marshal(transaction)
  if err != nil { return nil, err }
  return jsonBytes, nil
}

func (transaction *Transaction) BodyToCBOR() ([]byte, error) {
  cborBytes, err := cbor.Marshal(transaction.Body)
  if err != nil { return nil, err }
  return cborBytes, nil
}

func (transaction *Transaction) Hash() error {
  cborBytes, err := transaction.BodyToCBOR()
  if err != nil { return err }
  transaction.Header.Hash = fmt.Sprintf("%x", sha256.Sum256(cborBytes))
  return nil
}

func (transaction *Transaction) AddSimpleOutput(output SimpleOutput) {
  transaction.Body.Outputs.AddSimpleOutput(output)
}

func (transaction *Transaction) AddMultiAssetOutput(output MultiAssetOutput) {
  transaction.Body.Outputs.AddMultiAssetOutput(output)
}

func (transaction *Transaction) AddMultiAddrOutput(output MultiAddrOutput) {
  transaction.Body.Outputs.AddMultiAddrOutput(output)
}

func (transaction *Transaction) AddData(data TxData) {
  transaction.Body.Data = append(transaction.Body.Data, data)
}
