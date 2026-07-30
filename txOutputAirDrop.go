package psl

type AirDropOutput struct {
  From   string   `cbor:"0,keyasint" json:"from"`
  Asset  string   `cbor:"1,keyasint" json:"asset"`
  To     []string `cbor:"2,keyasint" json:"to"`
  Amount uint     `cbor:"3,keyasint" json:"amount"`
}

func NewAirDropOutput() AirDropOutput {
  return AirDropOutput{
    To: make([]string, 0),
  }
}

func (output *AirDropOutput) SetSender(addr string) {
  output.From = addr
}

func (output *AirDropOutput) AddAddr(addr string) {
  output.To = append(output.To, addr)
}