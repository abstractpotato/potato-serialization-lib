package ledger

type Block struct {
  Header BlockHeader `cbor: "header"`
  Body   BlockBody   `cbor: "body"`
}

type BlockHeader struct {
  ID        uint   `cbor: "id"`
  Hash      string `cbor: "hash"`
  Pool      string `cbor: "pool"`
  Signature []byte `cbor: "signature"`
}

type BlockBody struct {
  PreviousHash string        `cbor: "previousHash"`
  VRF          string        `cbor: "vrf"`
  Epoch        uint          `cbor: "epoch"`
  Slot         uint          `cbor: "slot"`
  Transactions []Transaction `cbor: "transactions"`
  Fees         uint          `cbor: "fees"`
  Timestamp    uint          `cbor: "timestamp"`
}