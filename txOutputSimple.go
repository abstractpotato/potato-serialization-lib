package psl

type SimpleOutput struct {
  From   string `cbor:"0,keyasint" json:"from"`
  To     string `cbor:"1,keyasint,omitempty" json:"to,omitempty"`
  Asset  string `cbor:"2,keyasint,omitempty" json:"asset,omitempty"`
  Amount uint   `cbor:"3,keyasint,omitempty" json:"amount,omitempty"`
}