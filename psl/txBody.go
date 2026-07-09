package psl

import(
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
)

type TxBody struct {
  Outputs   TxOutputs `cbor:"0,keyasint" json:"outputs"`
  Data      []TxData  `cbor:"1,keyasint,toarray,omitempty", json:"data"`
  TTL       uint      `cbor:"2,keyasint" json:"ttl"`
  Timestamp uint      `cbor:"3,keyasint" json:"timestamp"`
  Network   uint      `cbor:"4,keyasint" json:"network"`
  Fee       uint      `cbor:"5,keyasint" json:"fee"`
}

type TxData struct {
  Tag  string `cbor:"0,keyasint"`
  Data []byte `cbor:"1,keyasint"`
  Type uint   `cbor:"2,keyasint"`
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
