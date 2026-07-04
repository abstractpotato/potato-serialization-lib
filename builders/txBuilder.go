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

func (builder *TxBuilder) Build() (PSL.Transaction, error) {
  transaction := builder.Tx
  transaction.Body.Network = builder.Context.Params.Network
  transaction.Body.TTL = builder.Context.Slot + 60
  transaction.Body.Epoch = builder.Context.Epoch
  transaction.Body.Timestamp = uint(time.Now().UnixMilli())

  transaction.Hash()
  cborBytes, err := transaction.ToCBOR()
  if err != nil { return transaction, err }

  txFee := uint(len(cborBytes)) * builder.Context.Params.TxFeePerByte
  transaction.Header.Fee = txFee

  return transaction, nil
}

func (builder *TxBuilder) AddOutput(output PSL.TxOutput) {
  builder.Tx.Body.Outputs = append(builder.Tx.Body.Outputs, output)
}

func (builder *TxBuilder) AddData(data PSL.TxData) {
  builder.Tx.Body.Data = append(builder.Tx.Body.Data, data)
}
