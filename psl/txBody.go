package psl

import(
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
)

type TxBody struct {
  Outputs   TxOutputs `cbor: "outputs"`
  Data      []TxData  `cbor: "data"`
  TTL       uint      `cbor: "ttl"`
  Timestamp uint      `cbor: "timestamp"`
  Network   uint      `cbor: "network"`
  Fee       uint      `cbor: "fee"`
}

type TxData struct {
  Tag  string `cbor: "tag"`
  Data []byte `cbor: "data"`
  Type uint   `cbor: "type"`
}

func TxBodyFromCBOR(cborBytes []byte) (TxBody, error) {
  var body TxBody
  err := cbor.Unmarshal(cborBytes, &body)
  if err != nil { return TxBody{}, err }
  return body, nil
}

func TxBodyFromHex(hexString string) (TxBody, error) {
  cborBytes, err := hex.DecodeString(hexString)
  if err != nil { return TxBody{}, err }
  body, err := TxBodyFromCBOR(cborBytes)
  if err != nil { return TxBody{}, err }
  return body, nil
}
