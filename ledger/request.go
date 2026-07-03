package ledger

import(
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
)

type Request struct {
  Ticker string   `cbor: "ticker"`
  Url    string   `cbor: "url"`
  Addr   string   `cbor: "addr"`
  Relays []string `cbor: "relays"`
}

func NewRequest() Request {
  return Request{}
}

func RequestFromCBOR(cborBytes []byte) (Request, error) {
  var request Request
  err := cbor.Unmarshal(cborBytes, &request)
  if err != nil { return NewRequest(), err }
  return request, nil
}

func RequestFromHex(hexString string) (Request, error) {
  cborBytes, err := hex.DecodeString(hexString)
  if err != nil { return NewRequest(), err }
  request, err :=  RequestFromCBOR(cborBytes)
  if err != nil { return NewRequest(), err }
  return request, nil
}

func (request *Request) ToCBOR() ([]byte, error) {
  cborBytes, err := cbor.Marshal(request)
  if err != nil { return nil, err }
  return cborBytes, nil
}

func (request *Request) ToHex() (string, error) {
  cborBytes, err := request.ToCBOR()
  if err != nil { return "", err }
  return hex.EncodeToString(cborBytes), nil
}

func (request *Request) ToJSON() ([]byte, error) {
  jsonBytes, err := json.Marshal(request)
  if err != nil { return nil, err }
  return jsonBytes, nil
}
