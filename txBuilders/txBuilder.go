package txBuilders

import (
  PSL "github.com/abstractpotato/potato-serialization-lib/psl"
  "time"
)

type TxBuilder struct {
  Context PSL.Context
  Tx      PSL.Transaction
}

func NewTxBuilder() TxBuilder {
  return TxBuilder{
    Tx: PSL.NewTransaction(),
  }
}

func (TxBuilder *builder) Build() (PSL.Transaction, error) {
  transaction := builder.Tx
  transaction.Body.Network = builder.Context.ProtocolParams.Network
  transaction.Body.TTL = builder.Context.Slot + 60
  transaction.Body.Epoch = builder.Context.Epoch
  transaction.Body.Timestamp = time.Now().UnixMilli()

  transaction.Hash()
  txFee := len(transaction.ToCBOR()) * builder.Context.Params.TxFeePerByte
  transaction.Header.Fee = txFee

  return transaction, nil
}

func (TxBuilder *builder) AddOutput(output PSL.TxOutput) {
  builder.Tx.Outputs = append(builder.Tx.Outputs, output)
}

func (TxBuilder *builder) AddData(data PSL.TxData) {
  builder.Tx.Data = append(builder.Tx.Data, data)
}
