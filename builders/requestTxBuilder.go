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

func (builder *RequestTxBuilder) Build() (PSL.Transaction, error) {
  cborBytes, err := builder.Request.ToCBOR()
  if err != nil { return PSL.NewTransaction(), err }

  requestData := PSL.TxData{}
  requestData.Tag = "validator_request"
  requestData.Data = cborBytes
  requestData.Type = 0

  transaction := PSL.NewTransaction()
  transaction.Body.Data = append(transaction.Body.Data, requestData)
  transaction.Body.Network = builder.Context.Params.Network
  transaction.Body.TTL = builder.Context.Slot + 60
  transaction.Body.Epoch = builder.Context.Epoch
  transaction.Body.Timestamp = uint(time.Now().UnixMilli())

  transaction.Hash()
  txCborBytes, err := transaction.ToCBOR()
  if err != nil { return transaction, err }

  txFee := uint(len(txCborBytes)) * builder.Context.Params.TxFeePerByte
  transaction.Header.Fee = txFee

  return transaction, nil
}
