package psl

import(
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
)

type BlockBody struct {
  PreviousHash string        `cbor:"0,keyasint" json:"previousHash"`
  VRF          string        `cbor:"1,keyasint" json:"vrf"`
  Epoch        uint          `cbor:"2,keyasint" json:"epoch"`
  Slot         uint          `cbor:"3,keyasint" json:"slot"`
  Transactions []Transaction `cbor:"4,keyasint,toarray,omitempty" json:"transactions,omitempty"`
  Fees         uint          `cbor:"5,keyasint,omitempty" json:"fees"`
  Timestamp    uint          `cbor:"6,keyasint" json:"timestamp"`
}

func NewBlockBody() BlockBody {
  return BlockBody{
    Transactions: make([]Transaction, 0),
  }
}

func BlockBodyFromCBOR(cborBytes []byte) (BlockBody, error) {
  var body BlockBody
  err := cbor.Unmarshal(cborBytes, &body)
  if err != nil { return BlockBody{}, err }
  return body, nil
}

func BlockBodyFromHex(hexString string) (BlockBody, error) {
  cborBytes, err := hex.DecodeString(hexString)
  if err != nil { return BlockBody{}, err }
  body, err := BlockBodyFromCBOR(cborBytes)
  if err != nil { return BlockBody{}, err }
  return body, nil
}

func (body *BlockBody) ToCBOR() ([]byte, error) {
  cborBytes, err := cbor.Marshal(body)
  if err != nil { return nil, err}
  return cborBytes, nil
}

func (body *BlockBody) ToHex() (string, error) {
  cborBytes, err := body.ToCBOR()
  if err != nil { return "", err }
  return hex.EncodeToString(cborBytes), nil
}

func (body *BlockBody) ToJSON() ([]byte, error) {
  jsonBytes, err := json.Marshal(body)
  if err != nil { return nil, err }
  return jsonBytes, nil
}
