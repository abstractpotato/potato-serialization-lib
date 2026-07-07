package psl

type TxOutputs struct {
  SimpleOutputs     []SimpleOutput     `cbor: "simpleOutputs",omitempty`
  MultiAssetOutputs []MultiAssetOutput `cbor: "multiAssetOutputs",omitempty`
  MultiAddrOutputs  []MultiAddrOutput  `cbor: "multiAddrOutputs",omitempty`
}

type SimpleOutput struct {
  To     string `cbor: "to"`
  Asset  string `cbor: "asset"`
  Amount uint   `cbor: "amount"`
}

type MultiAssetOutput struct {
  To     string        `cbor: "to"`
  Assets []AssetOutput `cbor: "assets"`
}

type AssetOutput struct {
  ID     string `cbor: "id"`
  Amount uint   `cbor: "amount"`
}

type MultiAddrOutput struct {
  Asset  string       `cbor: "asset"`
  Addrs  []AddrOutput `cbor: "addr"`
}

type AddrOutput struct {
  Addr   string `cbor: "addr"`
  Amount uint   `cbor: "amount"`
}

func NewTxOutputs() TxOutputs {
  return TxOutputs{
    SimpleOutputs: make([]SimpleOutput, 0),
    MultiAssetOutputs: make([]MultiAssetOutput, 0),
    MultiAddrOutputs: make([]MultiAddrOutput, 0),
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
