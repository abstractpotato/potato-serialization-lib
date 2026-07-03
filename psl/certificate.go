package psl

import(
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
)

type Certificate struct {
  RequestHash string `cbor: "requestHash"`
  Signature   []byte `cbor: "signature"`
}

func NewCertificate() Certificate {
  return Certificate{}
}

func CertificateFromCBOR(cborBytes []byte) (Certificate, error) {
  var certificate Certificate
  err := cbor.Unmarshal(cborBytes, &certificate)
  if err != nil { return NewCertificate(), err }
  return certificate, nil
}

func CertificateFromHex(hexString string) (Certificate, error) {
  cborBytes, err := hex.DecodeString(hexString)
  if err != nil { return NewCertificate(), err }
  certificate, err :=  CertificateFromCBOR(cborBytes)
  if err != nil { return NewCertificate(), err }
  return certificate, nil
}

func (certificate *Certificate) ToCBOR() ([]byte, error) {
  cborBytes, err := cbor.Marshal(certificate)
  if err != nil { return nil, err }
  return cborBytes, nil
}

func (certificate *Certificate) ToHex() (string, error) {
  cborBytes, err := certificate.ToCBOR()
  if err != nil { return "", err }
  return hex.EncodeToString(cborBytes), nil
}

func (certificate *Certificate) ToJSON() ([]byte, error) {
  jsonBytes, err := json.Marshal(certificate)
  if err != nil { return nil, err }
  return jsonBytes, nil
}
