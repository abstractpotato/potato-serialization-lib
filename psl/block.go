package psl

import(
  "fmt"
  "crypto/sha256"
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
)

type Block struct {
  Header BlockHeader `cbor:"0,keyasint"`
  Body   BlockBody   `cbor:"1,keyasint"`
}

func NewBlock() Block {
  return Block{
    Header: BlockHeader{
      Witnesses: make([]Witness, 0),
    },
    Body: BlockBody{
      Transactions: make([]Transaction, 0),
    },
  }
}

func BlockFromCBOR(cborBytes []byte) (Block, error) {
  var block Block
  err := cbor.Unmarshal(cborBytes, &block)
  if err != nil { return NewBlock(), err }
  return block, nil
}

func BlockFromHex(hexString string) (Block, error) {
  cborBytes, err := hex.DecodeString(hexString)
  if err != nil { return NewBlock(), err }
  block, err :=  BlockFromCBOR(cborBytes)
  if err != nil { return NewBlock(), err }
  return block, nil
}

func (block *Block) ToCBOR() ([]byte, error) {
  cborBytes, err := cbor.Marshal(block)
  if err != nil { return nil, err }
  return cborBytes, nil
}

func (block *Block) ToHex() (string, error) {
  cborBytes, err := block.ToCBOR()
  if err != nil { return "", err }
  return hex.EncodeToString(cborBytes), nil
}

func (block *Block) ToJSON() ([]byte, error) {
  jsonBytes, err := json.Marshal(block)
  if err != nil { return nil, err }
  return jsonBytes, nil
}

func (block *Block) Hash() error {
  cborBytes, err := block.Body.ToCBOR()
  if err != nil { return err }
  block.Header.Hash = fmt.Sprintf("%x", sha256.Sum256(cborBytes))
  return nil
}

func (block *Block) AddWitness(witness []byte) {
  block.Header.Witnesses = append(block.Header.Witnesses, witness)
}
