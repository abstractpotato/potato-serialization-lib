package psl

type Genesis struct {
  Seed        []byte      `cbor:"0,keyasint"`
  Certificate Certificate `cbor:"1,keyasint"`
  Params      Params      `cbor:"2,keyasint"`
}
