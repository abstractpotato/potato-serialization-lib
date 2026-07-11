package psl

type Witness struct {
  PublicKey []byte `cbor:"0,keyasint"`
  Signature []byte `cbor:"1,keyasint"`
}
