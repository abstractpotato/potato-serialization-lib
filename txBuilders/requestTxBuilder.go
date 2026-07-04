package txBuilders

import (
  PSL "github.com/abstractpotato/potato-serialization-lib/psl"
  "time"
)

type RequestTxBuilder struct {
  Context PSL.Context
  Request PSL.Request
}

func NewRequestTxBuilder() RequestTxBuilder {
  return RequestTxBuilder{}
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
  transaction.Body.Network = builder.Context.ProtocolParams.Network
  transaction.Body.TTL = builder.Context.Slot + 60
  transaction.Body.Epoch = builder.Context.Epoch
  transaction.Body.Timestamp = time.Now().UnixMilli()

  transaction.Hash()
  txFee := len(transaction.ToCBOR()) * builder.Context.Params.TxFeePerByte
  transaction.Header.Fee = txFee

  return transaction, nil
}
