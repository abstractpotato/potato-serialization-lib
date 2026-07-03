package ledger

type Pool struct {
  ID     string   `cbor: "pool_id"`
  Addr   string   `cbor: "addr"`
  Relays []string `cbor: "relays"`
  Status uint     `cbor: "status"`
}