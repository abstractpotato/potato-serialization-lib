package txBuilders

import (
  PSL "github.com/abstractpotato/potato-serialization-lib/psl"
  "time"
)

type TxBuilder struct {
  Params PSL.Params
  Tx     PSL.Transaction
}

func NewTxBuilder() TxBuilder {
  return TxBuilder{
    Tx: PSL.NewTransaction(),
  }
}

func EstimateFee(params *PSL.Params, tx *PSL.Transaction) {
  minTxFee := params.MintTxFee
  txFeePerByte := params.TxFeePerByte
  tx.Body.Fee = minTxFee
  dryRunFee = minTxFee + (len(tx.BodyToCBOR()) * txFeePerByte)
  tx.Body.Fee = dryRunFee
  finalFee = minTxFee + (len(tx.BodyToCBOR()) * txFeePerByte)
  tx.Body.Fee = finalFee
}

func (builder *TxBuilder) Build() (PSL.Transaction, error) {
  transaction := builder.Tx
  transaction.Body.Network = builder.Params.Network
  transaction.Body.TTL = uint(time.Now().UnixMilli()) + 60
  transaction.Body.Timestamp = uint(time.Now().UnixMilli())

  EstimateFee(builder.Params, builder.Tx)
  
  transaction.Hash()
  return transaction, nil
}

func (builder *TxBuilder) AddOutput(output PSL.TxOutput) {
  builder.Tx.Body.Outputs = append(builder.Tx.Body.Outputs, output)
}

func (builder *TxBuilder) AddData(data PSL.TxData) {
  builder.Tx.Body.Data = append(builder.Tx.Body.Data, data)
}
