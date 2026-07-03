package ledger

type Parameters struct {
  Network            uint `cbor: "network"`
  MaxBlockHeaderSize uint `cbor: "maxBlockHeaderSize"`
  MaxBlockBodySize   uint `cbor: "maxBlockBodySize"`
  MaxTxSize          uint `cbor: "maxTxSize"`
  TxFeePerByte       uint `cbor: "txFeePerByte"`
  SlotsPerEpoch      uint `cbor: "slotsPerEpoch"`
  SlotTimeInMs       uint `cbor: "slotTimeInMs"`
  ProtocolVersion    uint `cbor: "protocolVersion"`
}