package psl

type TxOutputs struct {
  SimpleOutputs     []SimpleOutput     `cbor:"0,keyasint,toarray,omitempty"`
  MultiAssetOutputs []MultiAssetOutput `cbor:"1,keyasint,toarray,omitempty"`
  MultiAddrOutputs  []MultiAddrOutput  `cbor:"2,keyasint,toarray,omitempty"`
}

type SimpleOutput struct {
  To     string `cbor:"0,keyasint"`
  Asset  string `cbor:"1,keyasint"`
  Amount uint   `cbor:"2,keyasint"`
}

type MultiAssetOutput struct {
  To     string        `cbor:"0,keyasint"`
  Assets []AssetOutput `cbor:"1,keyasint"`
}

type AssetOutput struct {
  ID     string `cbor:"0,keyasint"`
  Amount uint   `cbor:"1,keyasint"`
}

type MultiAddrOutput struct {
  Asset  string       `cbor:"0,keyasint"`
  Addrs  []AddrOutput `cbor:"1,keyasint"`
}

type AddrOutput struct {
  Addr   string `cbor:"0,keyasint"`
  Amount uint   `cbor:"1,keyasint"`
}

func NewTxOutputs() TxOutputs {
  return TxOutputs{
    SimpleOutputs: make([]SimpleOutput, 0),
    MultiAssetOutputs: make([]MultiAssetOutput, 0),
    MultiAddrOutputs: make([]MultiAddrOutput, 0),
  }
}

func NewMultiAssetOutput() MultiAssetOutput {
  return MultiAssetOutput{
    Assets: make([]AssetOutput, 0),
  }
}

func NewMultiAddrOutput() MultiAddrOutput {
  return MultiAddrOutput{
    Addrs: make([]AddrOutput, 0),
  }
}

func (outputs *TxOutputs) AddSimpleOutput(output SimpleOutput) {
  outputs.SimpleOutputs = append(outputs.SimpleOutputs, output)
}

func (outputs *TxOutputs) AddMultiAssetOutput(output MultiAssetOutput) {
  outputs.MultiAssetOutputs = append(outputs.MultiAssetOutputs, output)
}

func (outputs *TxOutputs) AddMultiAddrOutput(output MultiAddrOutput) {
  outputs.MultiAddrOutputs = append(outputs.MultiAddrOutputs, output)
}

func (outputs *MultiAssetOutput) AddAssetOutput(asset AssetOutput) {
  outputs.Assets = append(outputs.Assets, asset)
}

func (outputs *MultiAddrOutput) AddAddrOutput(addr AddrOutput) {
  outputs.Addrs = append(outputs.Addrs, addr)
}
