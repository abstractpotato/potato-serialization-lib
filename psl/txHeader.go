package psl

import(
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
)

type TxHeader struct {
  Hash      string `cbor: "hash"`
  Addr      string `cbor: "sender"`
  Signature []byte `cbor: "signature"`
  Key       []byte `cbor: "key",omitempty`
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
