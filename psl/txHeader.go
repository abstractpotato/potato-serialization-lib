package psl

import(
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
)

type TxHeader struct {
  Hash      string `cbor:"0,keyasint" json:"hash"`
  Addr      string `cbor:"1,keyasint" json:"addr"`
  Signature []byte `cbor:"2,keyasint" json:"signature"`
  Key       []byte `cbor:"3,keyasint,omitempty" json:"key"`
}

func TxHeaderFromCBOR(cborBytes []byte) (TxHeader, error) {
  var header TxHeader
  err := cbor.Unmarshal(cborBytes, &header)
  if err != nil { return TxHeader{}, err }
  return header, nil
}

func TxHeaderFromHex(hexString string) (TxHeader, error) {
  cborBytes, err := hex.DecodeString(hexString)
  if err != nil { return TxHeader{}, err }
  header, err := TxHeaderFromCBOR(cborBytes)
  if err != nil { return TxHeader{}, err }
  return header, nil
}

func (header *TxHeader) ToCBOR() ([]byte, error) {
  cborBytes, err := cbor.Marshal(header)
  if err != nil { return nil, err}
  return cborBytes, nil
}

func (header *TxHeader) ToHex() (string, error) {
  cborBytes, err := header.ToCBOR()
  if err != nil { return "", err }
  return hex.EncodeToString(cborBytes), nil
}

func (header *TxHeader) ToJSON() ([]byte, error) {
  jsonBytes, err := json.Marshal(header)
  if err != nil { return nil, err }
  return jsonBytes, nil
}
