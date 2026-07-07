package psl

import(
  "fmt"
  "crypto/sha256"
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
)

type Block struct {
  Header BlockHeader `cbor: "header"`
  Body   BlockBody   `cbor: "body"`
}

type BlockHeader struct {
  ID        uint      `cbor: "id"`
  Hash      string    `cbor: "hash"`
  Witnesses []Witness `cbor: "witness"`
  Bytes     []byte    `cbor: "bytes"`
}

type Witness struct {
  Addr      string `cbor: "validator"`
  Signature []byte `cbor: "signature"`
}

type BlockBody struct {
  PreviousHash string        `cbor: "previousHash"`
  VRF          string        `cbor: "vrf"`
  Epoch        uint          `cbor: "epoch"`
  Slot         uint          `cbor: "slot"`
  Transactions []Transaction `cbor: "transactions"`
  Fees         uint          `cbor: "fees"`
  Timestamp    uint          `cbor: "timestamp"`
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

func (block *Block) HeaderToCBOR() ([]byte, error) {
  cborBytes, err := cbor.Marshal(block.Header)
  if err != nil { return nil, err }
  return cborBytes, nil
}

func (block *Block) BodyToCBOR() ([]byte, error) {
  cborBytes, err := cbor.Marshal(block.Body)
  if err != nil { return nil, err }
  return cborBytes, nil
}

func (block *Block) Hash() error {
  cborBytes, err := block.BodyToCBOR()
  if err != nil { return err }
  block.Header.Hash = fmt.Sprintf("%x", sha256.Sum256(cborBytes))
  return nil
}

func (block *Block) AddWitness(witness Witness) {
  block.Header.Witnesses = append(block.Header.Witnesses, witness)
}
