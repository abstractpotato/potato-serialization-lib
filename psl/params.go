package ledger

import(
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
)

type Params struct {
  Network            uint `cbor: "network"`
  MaxBlockHeaderSize uint `cbor: "maxBlockHeaderSize"`
  MaxBlockBodySize   uint `cbor: "maxBlockBodySize"`
  MaxTxSize          uint `cbor: "maxTxSize"`
  TxFeePerByte       uint `cbor: "txFeePerByte"`
  SlotsPerEpoch      uint `cbor: "slotsPerEpoch"`
  SlotTimeInMs       uint `cbor: "slotTimeInMs"`
  ProtocolVersion    uint `cbor: "protocolVersion"`
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
