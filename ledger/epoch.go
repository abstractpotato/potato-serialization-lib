package ledger

type Epoch struct {
  Header EpochHeader `cbor: "header"`
  Body   EpochBody   `cbor: "body"`
} 

type EpochHeader struct {
  ID   uint   `cbor: "id"`
  Hash string `cbor: "hash"`
}

type EpochBody struct {
  PreviousHash    string   `cbor: "previousHash"`
  Pools           []string `cbor: "pools"`
  StartTime       uint     `cbor: "startTime"`
  EndTime         uint     `cbor: "endTime"`
  ProtocolVersion uint     `cbor: "protocolVersion"`
}

