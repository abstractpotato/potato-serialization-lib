package psl

import(
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
)

type TxBody struct {
  SimpleOutputs []SimpleOutput `cbor:"0,keyasint,toarray,omitempty" json:"simpleOutputs,omitempty"`
  MultiAssetOutputs []MultiAssetOutput `cbor:"1,keyasint,toarray,omitempty" json:"multiAssetOutputs,omitempty"`
  MultiAddrOutputs []MultiAddrOutput `cbor:"2,keyasint,toarray,omitempty" json:"multiAddrOutputs,omitempty"`
  Data []TxData `cbor:"3,keyasint,toarray,omitempty" json:"data,omitempty"`
  Request Request `cbor:"4,keyasint,omitempty", json:"request,omitempty"`
  Certificate Certificate `cbor:"5,keyasint,omitempty" json:"certificate,omitempty"`
  TTL uint `cbor:"5,keyasint,omitempty" json:"ttl,omitempty"`
  Timestamp uint `cbor:"6,keyasint" json:"timestamp"`
  Network uint `cbor:"7,keyasint" json:"network"`
  Fee uint `cbor:"8,keyasint" json:"fee"`
}

type TxData struct {
  Tag  string `cbor:"0,keyasint"`
  Data []byte `cbor:"1,keyasint"`
  Type uint   `cbor:"2,keyasint"`
}

func NewTxBody() TxBody {
  return TxBody{
    SimpleOutputs: make([]SimpleOutput, 0),
    MultiAssetOutputs: make([]MultiAssetOutput, 0),
    MultiAddrOutputs: make([]MultiAddrOutput, 0),
    Data: make([]TxData, 0),
  }
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

func (body *TxBody) AddSimpleOutput(output SimpleOutput) {
  body.SimpleOutputs = append(body.SimpleOutputs, output)
}

func (body *TxBody) AddMultiAssetOutput(output MultiAssetOutput) {
  body.MultiAssetOutputs = append(body.MultiAssetOutputs, output)
}

func (body *TxBody) AddMultiAddrOutput(output MultiAddrOutput) {
  body.MultiAddrOutputs = append(body.MultiAddrOutputs, output)
}

func (body *TxBody) AddData(data TxData) {
  body.Data = append(body.Data, data)
}

func (body *TxBody) AddRequest(request Request) {
  body.Request = request
}

func (body *TxBody) AddCertificate(certificate Certificate) {
  body.Certificate = certificate
}
