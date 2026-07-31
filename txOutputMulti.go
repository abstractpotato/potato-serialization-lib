package psl

type MultiOutput struct {
  From   string `cbor:"0,keyasint" json:"from"`
  To     string        `cbor:"0,keyasint" json:"to"`
  Assets []AssetOutput `cbor:"1,keyasint,toarray" json:"asset"`
}

type AssetOutput struct {
  Asset  string `cbor:"0,keyasint" json:"asset"`
  Amount uint   `cbor:"1,keyasint" json:"amount"`
}

func NewMultiOutput() MultiOutput {
  return MultiOutput{
    Assets: make([]AssetOutput, 0),
  }
}

func (outputs *MultiOutput) SetSender(addr string) {
  outputs.From = addr
}

func (outputs *MultiOutput) Add(output AssetOutput) {
  outputs.Assets = append(outputs.Assets, output)
}
