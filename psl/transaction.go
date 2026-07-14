package psl

import(
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
  "golang.org/x/crypto/blake2b"
  cardano "github.com/abstractpotato/potato-serialization-lib/cardano"
)

type Transaction struct {
  Header TxHeader `cbor:"0,keyasint" json:"header"`
  Body   TxBody   `cbor:"1,keyasint" json:"body"`
}

func NewTransaction() Transaction {
  return Transaction{
    Header: NewTxHeader(),
    Body: NewTxBody(),
  }
}

func TransactionFromCBOR(cborBytes []byte) (Transaction, error) {
  var transaction Transaction
  err := cbor.Unmarshal(cborBytes, &transaction)
  if err != nil { return NewTransaction(), err }
  return transaction, nil
}

func TransactionFromHex(hexString string) (Transaction, error) {
  cborBytes, err := hex.DecodeString(hexString)
  if err != nil { return NewTransaction(), err }
  transaction, err := TransactionFromCBOR(cborBytes)
  if err != nil { return NewTransaction(), err }
  return transaction, nil
}

func (transaction *Transaction) ToCBOR() ([]byte, error) {
  cborBytes, err := cbor.Marshal(transaction)
  if err != nil { return nil, err}
  return cborBytes, nil
}

func (transaction *Transaction) ToHex() (string, error) {
  cborBytes, err := transaction.ToCBOR()
  if err != nil { return "", err }
  return hex.EncodeToString(cborBytes), nil
}

func (transaction *Transaction) ToJSON() ([]byte, error) {
  jsonBytes, err := json.Marshal(transaction)
  if err != nil { return nil, err }
  return jsonBytes, nil
}

func (transaction *Transaction) Hash() error {
  cborBytes, err := transaction.Body.ToCBOR()
  if err != nil { return err }
  hashBytes := blake2b.Sum256(cborBytes)
  transaction.Header.Hash = hex.EncodeToString(hashBytes[:])
  return nil
}

func (transaction *Transaction) HashToBytes() []byte {
  return []byte(transaction.Header.Hash)
}

func (transaction *Transaction) AddWitness(witness Witness) {
  transaction.Header.AddWitness(witness)
}

func (transaction *Transaction) AddSimpleOutput(output SimpleOutput) {
  transaction.Body.AddSimpleOutput(output)
}

func (transaction *Transaction) AddMultiAssetOutput(output MultiAssetOutput) {
  transaction.Body.AddMultiAssetOutput(output)
}

func (transaction *Transaction) AddMultiAddrOutput(output MultiAddrOutput) {
  transaction.Body.AddMultiAddrOutput(output)
}

func (transaction *Transaction) AddData(data TxData) {
  transaction.Body.AddData(data)
}

func (transaction *Transaction) AddRequest(request *Request) {
  transaction.Body.AddRequest(request)
}

func (transaction *Transaction) AddCertificate(certificate *Certificate) {
  transaction.Body.AddCertificate(certificate)
}

func (transaction *Transaction) Sign(privateKey []byte) error {
  hashBytes := transaction.HashToBytes()
  signature, err := cardano.Sign(privateKey, hashBytes)
  if err != nil { return err }

  publicKey, err := cardano.MakePublicKey(privateKey[:32])
  if err != nil { return err }

  witness := Witness{
    PublicKey: publicKey,
    Signature: signature,
  }

  transaction.Header.AddWitness(witness)
  return nil
}

func (transaction *Transaction) Verify() bool {
  transaction.Hash()
  witness := transaction.Header.Witnesses[0]
  vkey := witness.PublicKey
  sig := witness.Signature
  hashBytes := transaction.HashToBytes()
  return cardano.Verify(vkey, sig, hashBytes)
}
