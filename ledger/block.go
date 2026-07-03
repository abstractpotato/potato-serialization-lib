package ledger

import(
  "fmt"
  "crypto/sha256"
  "github.com/fxamacker/cbor/v2"
)

type Block struct {
  Header BlockHeader `cbor: "header"`
  Body   BlockBody   `cbor: "body"`
}

type BlockHeader struct {
  ID        uint   `cbor: "id"`
  Hash      string `cbor: "hash"`
  Pool      string `cbor: "pool"`
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

func (block *Block) ToCBOR() ([]byte, error) {
  cborBytes, err := cbor.Marshal(block)
  if err != nil { return []byte, err }
  return cborBytes, nil
}
 
func (block *Block) BodyToCBOR() ([]byte, error) {
  cborBytes, err := cbor.Marshal(block.Body)
  if err != nil { return []byte, err }
  return cborBytes, nil
}

func (block *Block) Hash() error {
  cborBytes, err := block.BodyToCBOR()
  if err != nil { return err }
  block.Header.Hash = fmt.Sprintf("%x", sha256.Sum256(cborBytes))
  return nil
}