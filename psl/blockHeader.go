package psl

import(
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
)

type BlockHeader struct {
  ID      uint    `cbor:"0,keyasint" json:"id"`
  Hash    string  `cbor:"1,keyasint" json:"hash"`
  Witness Witness `cbor:"2,keyasint,toarray" json:"witness"`
}

func NewBlockHeader() BlockHeader {
  return BlockHeader{}
}

func BlockHeaderFromCBOR(cborBytes []byte) (BlockHeader, error) {
  var header BlockHeader
  err := cbor.Unmarshal(cborBytes, &header)
  if err != nil { return BlockHeader{}, err }
  return header, nil
}

func BlockHeaderFromHex(hexString string) (BlockHeader, error) {
  cborBytes, err := hex.DecodeString(hexString)
  if err != nil { return BlockHeader{}, err }
  header, err := BlockHeaderFromCBOR(cborBytes)
  if err != nil { return BlockHeader{}, err }
  return header, nil
}

func (header *BlockHeader) ToCBOR() ([]byte, error) {
  cborBytes, err := cbor.Marshal(header)
  if err != nil { return nil, err}
  return cborBytes, nil
}

func (header *BlockHeader) ToHex() (string, error) {
  cborBytes, err := header.ToCBOR()
  if err != nil { return "", err }
  return hex.EncodeToString(cborBytes), nil
}

func (header *BlockHeader) ToJSON() ([]byte, error) {
  jsonBytes, err := json.Marshal(header)
  if err != nil { return nil, err }
  return jsonBytes, nil
}
