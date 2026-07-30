package psl

import(
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
)

type BlockBody struct {
  ID           uint          `cbor:"0,keyasint" json:"id"`
  Network      uint          `cbor:"1,keyasint" json:"network"`
  PreviousHash string        `cbor:"2,keyasint,omitempty" json:"previousHash,omitempty"`
  VRF          string        `cbor:"3,keyasint,omitempty" json:"vrf,omitempty"`
  Epoch        uint          `cbor:"4,keyasint" json:"epoch"`
  Slot         uint          `cbor:"5,keyasint" json:"slot"`
  Transactions []Transaction `cbor:"6,keyasint,toarray,omitempty" json:"transactions,omitempty"`
  Fees         uint          `cbor:"7,keyasint,omitzero" json:"fees,omitzero"`
  Timestamp    uint          `cbor:"8,keyasint" json:"timestamp"`
  Genesis      *Genesis      `cbor:"9,keyasint,omitempty" json:"genesis,omitempty"`
}

func NewBlockBody() BlockBody {
  return BlockBody{
    Transactions: make([]Transaction, 0),
  }
}

func BlockBodyFromCBOR(cborBytes []byte) (BlockBody, error) {
  var body BlockBody
  err := strictDec.Unmarshal(cborBytes, &body)
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

func (body *BlockBody) AddTx(transaction Transaction) {
  body.Transactions = append(body.Transactions, transaction)
}
