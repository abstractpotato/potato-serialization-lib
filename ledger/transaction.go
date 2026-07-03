package ledger

type Transaction struct {
  Header TxHeader `cbor: "header"`
  Body   TxBody   `cbor: "body"`
}

type TxHeader struct {
  Hash      string `cbor: "hash"`
  Sender    string `cbor: "sender"`
  Signature []byte `cbor: "signature"`
  Fee       uint   `cbor: "fee"`
}

type TxBody struct {
  Outputs   []TxOutput `cbor: "outputs"`
  Data      []TxData   `cbor: "data"`
  Epoch     uint       `cbor: "epoch"`
  TTL       uint       `cbor: "ttl"`
  Timestamp uint       `cbor: "timestamp"`
  Network   uint       `cbor: "network"`
}

type TxOutput struct {
  To     string `cbor: "to"`
  Asset  string `cbor: "asset"`
  Amount uint   `cbor: "amount"`
}

type TxData struct {
  Tag  string `cbor: "tag"`
  Data []byte `cbor: "data"`
  Type uint   `cbor: "type"`
}
