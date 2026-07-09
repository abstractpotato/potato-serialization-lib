package psl

import(
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
)

type Params struct {
  Network            uint `cbor:"0,keyasint" json:"network"`
  MaxBlockHeaderSize uint `cbor:"1,keyasint" json:"maxBlockHeaderSize"`
  MaxBlockBodySize   uint `cbor:"2,keyasint" json:"maxBlockBodySize"`
  MaxTxSize          uint `cbor:"3,keyasint" json:"maxTxSize"`
  MinTxFee           uint `cbor:"4,keyasint" json:"minTxFee"`
  TxFeePerByte       uint `cbor:"5,keyasint" json:"txFeePerByte"`
  SlotsPerEpoch      uint `cbor:"6,keyasint" json:"slotsPerEpoch"`
  SlotTimeInMs       uint `cbor:"7,keyasint" json:"slotTimeInMs"`
  ProtocolVersion    uint `cbor:"8,keyasint" json:"protocolVersion"`
}

func NewParams() Params {
  return Params{}
}

func ParamsFromCBOR(cborBytes []byte) (Params, error) {
  var params Params
  err := cbor.Unmarshal(cborBytes, &params)
  if err != nil { return NewParams(), err }
  return params, nil
}

func ParamsFromHex(hexString string) (Params, error) {
  cborBytes, err := hex.DecodeString(hexString)
  if err != nil { return NewParams(), err }
  params, err :=  ParamsFromCBOR(cborBytes)
  if err != nil { return NewParams(), err }
  return params, nil
}

func (params *Params) ToCBOR() ([]byte, error) {
  cborBytes, err := cbor.Marshal(params)
  if err != nil { return nil, err }
  return cborBytes, nil
}

func (params *Params) ToHex() (string, error) {
  cborBytes, err := params.ToCBOR()
  if err != nil { return "", err }
  return hex.EncodeToString(cborBytes), nil
}

func (params *Params) ToJSON() ([]byte, error) {
  jsonBytes, err := json.Marshal(params)
  if err != nil { return nil, err }
  return jsonBytes, nil
}
