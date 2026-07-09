package psl

import(
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
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

func (body *TxBody) ToCBOR() ([]byte, error) {
  cborBytes, err := cbor.Marshal(body)
  if err != nil { return nil, err}
  return cborBytes, nil
}

func (body *TxBody) ToHex() (string, error) {
  cborBytes, err := body.ToCBOR()
  if err != nil { return "", err }
  return hex.EncodeToString(cborBytes), nil
}

func (body *TxBody) ToJSON() ([]byte, error) {
  jsonBytes, err := json.Marshal(body)
  if err != nil { return nil, err }
  return jsonBytes, nil
}
