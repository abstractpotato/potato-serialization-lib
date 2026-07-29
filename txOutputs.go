package psl

type SimpleOutput struct {
  From   string `cbor:"0,keyasint" json:"from"`
  To     string `cbor:"1,keyasint" json:"to"`
  Asset  string `cbor:"2,keyasint" json:"asset"`
  Amount uint   `cbor:"3,keyasint" json:"amount"`
}

type MultiAssetOutput struct {
  From   string        `cbor:"0,keyasint" json:"from"`
  To     string        `cbor:"1,keyasint" json:"to"`
  Assets []AssetOutput `cbor:"2,keyasint,toarray" json:"asset"`
}

type AssetOutput struct {
  Asset  string `cbor:"0,keyasint" json:"asset"`
  Amount uint   `cbor:"1,keyasint" json:"amount"`
}

type MultiAddrOutput struct {
  From   string       `cbor:"0,keyasint" json:"from"`
  To     []AddrOutput `cbor:"1,keyasint,toarray" json:"to"`
  Asset  string       `cbor:"2,keyasint" json:"asset"`
}

type AddrOutput struct {
  Addr   string `cbor:"0,keyasint" json:"addr"`
  Amount uint   `cbor:"1,keyasint" json:"amount"`
}

func NewMultiAssetOutput() MultiAssetOutput {
  return MultiAssetOutput{
    Assets: make([]AssetOutput, 0),
  }
}

func NewMultiAddrOutput() MultiAddrOutput {
  return MultiAddrOutput{
    To: make([]AddrOutput, 0),
  }
}

func (outputs *MultiAssetOutput) Add(output AssetOutput) {
  outputs.Assets = append(outputs.Assets, output)
}

func (outputs *MultiAddrOutput) Add(output AddrOutput) {
  outputs.To = append(outputs.To, output)
}
