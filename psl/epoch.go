package psl

import(
  "fmt"
  "crypto/sha256"
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
)

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
  Validators      []string `cbor: "pools"`
  StartTime       uint     `cbor: "startTime"`
  EndTime         uint     `cbor: "endTime"`
  ProtocolVersion uint     `cbor: "protocolVersion"`
}

func NewEpoch() Epoch {
  return Epoch{
    Header: EpochHeader{},
    Body: EpochBody{
      Validators: make([]string, 0),
    },
  }
}

func EpochFromCBOR(cborBytes []byte) (Epoch, error) {
  var epoch Epoch
  err := cbor.Unmarshal(cborBytes, &epoch)
  if err != nil { return NewEpoch(), err }
  return epoch, nil
}

func EpochFromHex(hexString string) (Epoch, error) {
  cborBytes, err := hex.DecodeString(hexString)
  if err != nil { return NewEpoch(), err }
  epoch, err := EpochFromCBOR(cborBytes)
  if err != nil { return NewEpoch(), err }
  return epoch, nil
}

func (epoch *Epoch) ToCBOR() ([]byte, error) {
  cborBytes, err := cbor.Marshal(epoch)
  if err != nil { return nil, err}
  return cborBytes, nil
}

func (epoch *Epoch) ToHex() (string, error) {
  cborBytes, err := epoch.ToCBOR()
  if err != nil { return "", err }
  return hex.EncodeToString(cborBytes), nil
}

func (epoch *Epoch) ToJSON() ([]byte, error) {
  jsonBytes, err := json.Marshal(epoch)
  if err != nil { return nil, err }
  return jsonBytes, nil
}

func (epoch *Epoch) BodyToCBOR() ([]byte, error) {
  cborBytes, err := cbor.Marshal(epoch.Body)
  if err != nil { return nil, err }
  return cborBytes, nil
}

func (epoch *Epoch) Hash() error {
  cborBytes, err := epoch.BodyToCBOR()
  if err != nil { return err }
  epoch.Header.Hash = fmt.Sprintf("%x", sha256.Sum256(cborBytes))
  return nil
}
