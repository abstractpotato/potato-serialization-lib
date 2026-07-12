package psl

type Genesis struct {
  GenesisSeed []byte      `cbor:"2,keyasint"`
  Certificate Certificate `cbor:"0,keyasint"`
  Params      Params      `cbor:"1,keyasint"`
}
