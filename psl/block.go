package psl

import(
  // "fmt"
  // "crypto/sha256"
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
  "golang.org/x/crypto/blake2b"
  cardano "github.com/abstractpotato/potato-serialization-lib/cardano"
)

type Block struct {
  Header BlockHeader `cbor:"0,keyasint" json:"header"`
  Body   BlockBody   `cbor:"1,keyasint" json:"body"`
}

func NewBlock() Block {
  return Block{
    Header: NewBlockHeader(),
    Body: NewBlockBody(),
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
  hashBytes := blake2b.Sum256(cborBytes)
  block.Header.Hash = hex.EncodeToString(hashBytes[:])
  return nil
}

func (block *Block) HashToBytes() []byte {
  return []byte(block.Header.Hash)
}

func (block *Block) Sign(privateKey []byte) error {
  hashBytes := block.HashToBytes()
  signature, err := cardano.Sign(privateKey, hashBytes)
  if err != nil { return err }

  publicKey, err := cardano.MakePublicKey(privateKey[:32])
  if err != nil { return err }

  witness := Witness{
    PublicKey: publicKey,
    Signature: signature,
  }

  block.Header.Witness = witness
  return nil
}

func (block *Block) Verify() bool {
  block.Hash() // hash internally for externally loaded blocks
  witness := block.Header.Witness
  vkey := witness.PublicKey
  sig := witness.Signature
  hashBytes := block.HashToBytes()
  return cardano.Verify(vkey, sig, hashBytes)
}
