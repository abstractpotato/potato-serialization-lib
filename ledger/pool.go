package ledger

import(
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
)

type Pool struct {
  ID     string   `cbor: "id"`
  Addr   string   `cbor: "addr"`
  Relays []string `cbor: "relays"`
  Status uint     `cbor: "status"`
}

func NewPool() Pool {
  return Pool{}
}

func PoolFromCBOR(cborBytes []byte) (Pool, error) {
  var pool Pool
  err := cbor.Unmarshal(cborBytes, &pool)
  if err != nil { return NewPool(), err }
  return pool, nil
}

func PoolFromHex(hexString string) (Pool, error) {
  cborBytes, err := hex.DecodeString(hexString)
  if err != nil { return NewPool(), err }
  pool, err :=  PoolFromCBOR(cborBytes)
  if err != nil { return NewPool(), err }
  return pool, nil
}

func (pool *Pool) ToCBOR() ([]byte, error) {
  cborBytes, err := cbor.Marshal(pool)
  if err != nil { return nil, err }
  return cborBytes, nil
}

func (pool *Pool) ToHex() (string, error) {
  cborBytes, err := pool.ToCBOR()
  if err != nil { return "", err }
  return hex.EncodeToString(cborBytes), nil
}

func (pool *Pool) ToJSON() ([]byte, error) {
  jsonBytes, err := json.Marshal(pool)
  if err != nil { return nil, err }
  return jsonBytes, nil
}