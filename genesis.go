package psl

type Genesis struct {
  Seed        []byte      `cbor:"0,keyasint"`
  Params      Params      `cbor:"1,keyasint"`
  Certificate Certificate `cbor:"2,keyasint"`
}
