package txBuilders

import PSL "github.com/abstractpotato/potato-serialization-lib/psl"

type RequestTxBuilder struct {
  Context Context
  Request Request
}

func (RequestTxBuilder *builder) Build() (PSL.Transaction, error) {
  cborBytes, err := request.ToCBOR()
  if err != nil { return PSL.NewTransaction(), err }

  requestData := PSL.TxData{}
  requestData.Tag = "validator_request"
  requestData.Data = cborBytes
  requestData.Type = 0

  transaction := PSL.NewTransaction()
  transaction.Body.Data = append(transaction.Body.Data, requestData)
  transaction.Hash()
  return transaction, nil
}
