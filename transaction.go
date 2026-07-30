package psl

import(
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
  "golang.org/x/crypto/blake2b"
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
  err := strictDec.Unmarshal(cborBytes, &transaction)
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

func (transaction *Transaction) HashToBytes() ([]byte, error) {
  return hex.DecodeString(transaction.Header.Hash)
}

func (transaction *Transaction) AddSimpleOutput(output SimpleOutput) {
  transaction.Body.AddSimpleOutput(output)
}

func (transaction *Transaction) AddMultiOutput(output MultiOutput) {
  transaction.Body.AddMultiOutput(output)
}

func (transaction *Transaction) SetAirDropOutput(output *AirDropOutput) {
  transaction.Body.SetAirDropOutput(output)
}

func (transaction *Transaction) SetRequest(request *Request) {
  transaction.Body.Request = request
}

func (transaction *Transaction) SetCertificate(certificate *Certificate) {
  transaction.Body.Certificate = certificate
}

func (transaction *Transaction) AddData(data TxData) {
  transaction.Body.AddData(data)
}

func (transaction *Transaction) Sign(privateKey []byte) error {
  err := transaction.Hash()
  if err != nil { return err }

  hashBytes, err := transaction.HashToBytes()
  if err != nil { return err }

  signature, err := Sign(privateKey, hashBytes)
  if err != nil { return err }

  publicKey, err := GetPublicKey(privateKey[:32])
  if err != nil { return err }

  witness := Witness{
    PublicKey: publicKey,
    Signature: signature,
  }

  transaction.Header.Witness = witness
  return nil
}

func (transaction *Transaction) Verify() bool {
  vkey := transaction.Header.Witness.PublicKey
  sig := transaction.Header.Witness.Signature
  if len(vkey) == 0 || len(sig) == 0 { return false }
  
  cborBytes, err := transaction.Body.ToCBOR()
  if err != nil { return false }
  
  // check simpleOutputs
  valid_simple := transaction.verifySimpleOutputAddr(vkey)
  if !valid_simple { return false }
  
  // // check multiOutputs
  // valid_multi := transaction.verifyMultiOutputAddr(vkey)
  // if !valid_multi { return false }
  // 
  // // check airdrop
  // valid_airdrop := transaction.verifyAirDrop(vkey)
  // if !valid_airdrop { return false }
  
  hashBytes := blake2b.Sum256(cborBytes)
  return Verify(vkey, sig, hashBytes[:])
}

func (transaction *Transaction) verifySimpleOutputAddr(publicKey []byte) bool {
  if len(transaction.Body.SimpleOutputs) == 0 { return true }
  
  for _, output := range transaction.Body.SimpleOutputs {
    valid_sender, err := AddrBelongsToPubKey(output.From, publicKey)
    if err != nil { panic(err) }
    if err != nil || !valid_sender { return false }
    if !IsValidAddress(output.To) { return false }
  }
  
  return true
}

func (transaction *Transaction) verifyMultiOutputAddr(publicKey []byte) bool {
  if len(transaction.Body.MultiOutputs) == 0 { return true }
  
  for _, output := range transaction.Body.MultiOutputs {
    valid_sender, err := AddrBelongsToPubKey(output.From, publicKey)
    if err != nil || !valid_sender { return false }
    if !IsValidAddress(output.To) { return false }  
  }
  
  return true
}

func (transaction *Transaction) verifyAirDrop(publicKey []byte) bool {
  if transaction.Body.AirDropOutput == nil { return true }
  if len(transaction.Body.AirDropOutput.To) == 0 { return false }
  
  sender_addr := transaction.Body.AirDropOutput.From
  valid_sender, err := AddrBelongsToPubKey(sender_addr, publicKey)
  if err != nil || !valid_sender { return false }
  
  for _, addr := range transaction.Body.AirDropOutput.To {
    if !IsValidAddress(addr) { return false }
  }
  
  return true
}