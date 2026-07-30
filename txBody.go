package psl

import(
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
)

type TxBody struct {
  SimpleOutputs []SimpleOutput `cbor:"0,keyasint,omitempty" json:"simpleOutputs,omitempty"`
  MultiOutputs  []MultiOutput  `cbor:"1,keyasint,omitempty" json:"multiOutputs,omitempty"`
  AirDropOutput *AirDropOutput `cbor:"2,keyasint,omitempty" json:"airdrop,omitempty"`
  Request       *Request       `cbor:"3,keyasint,omitempty" json:"request,omitempty"`
  Certificate   *Certificate   `cbor:"4,keyasint,omitempty" json:"certificate,omitempty"`
  Data          []TxData       `cbor:"5,keyasint,toarray,omitempty" json:"data,omitempty"`
  TTL           uint           `cbor:"6,keyasint,omitempty" json:"ttl,omitempty"`
  Timestamp     uint           `cbor:"7,keyasint" json:"timestamp"`
  Network       uint           `cbor:"8,keyasint" json:"network"`
  Fee           uint           `cbor:"9,keyasint" json:"fee"`
}

type TxData struct {
  Tag  string `cbor:"0,keyasint"`
  Data []byte `cbor:"1,keyasint"`
  Type uint   `cbor:"2,keyasint"`
}

func NewTxBody() TxBody {
  return TxBody{
    SimpleOutputs: make([]SimpleOutput, 0),
    MultiOutputs: make([]MultiOutput, 0),
    Data: make([]TxData, 0),
  }
}

func TxBodyFromCBOR(cborBytes []byte) (TxBody, error) {
  var body TxBody
  err := strictDec.Unmarshal(cborBytes, &body)
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

func (body *TxBody) AddMultiOutput(output MultiOutput) {
  body.MultiOutputs = append(body.MultiOutputs, output)
}

func (body *TxBody) SetAirDropOutput(output *AirDropOutput) {
  body.AirDropOutput = output
}

func (body *TxBody) SetRequest(request *Request) {
  body.Request = request
}

func (body *TxBody) SetCertificate(certificate *Certificate) {
  body.Certificate = certificate
}

func (body *TxBody) AddData(data TxData) {
  body.Data = append(body.Data, data)
}
