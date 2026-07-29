package psl

type Witness struct {
  PublicKey []byte `cbor:"0,keyasint" json:"publicKey"`
  Signature []byte `cbor:"1,keyasint" json:"signature"`
}
