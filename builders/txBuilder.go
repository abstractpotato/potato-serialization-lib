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

func (builder *TxBuilder) EstimateFee() error {
  minTxFee := builder.Params.MinTxFee
  txFeePerByte := builder.Params.TxFeePerByte
  builder.Tx.Body.Fee = minTxFee

  cborBytes, err := builder.Tx.BodyToCBOR()
  if err != nil { return err }

  dryRunFee := minTxFee + (uint(len(cborBytes)) * txFeePerByte)
  builder.Tx.Body.Fee = dryRunFee

  cborBytes, err = builder.Tx.BodyToCBOR()
  if err != nil { return err }

  finalFee := minTxFee + (uint(len(cborBytes)) * txFeePerByte)
  builder.Tx.Body.Fee = finalFee
  return nil
}

func (builder *TxBuilder) Build() error {
  builder.Tx.Body.Network = builder.Params.Network
  builder.Tx.Body.TTL = 3000 // 3 seconds
  builder.Tx.Body.Timestamp = uint(time.Now().UnixMilli())
  builder.EstimateFee()
  builder.Tx.Hash()
  return nil
}

func (builder *TxBuilder) AddSimpleOutput(output PSL.SimpleOutput) {
  builder.Tx.AddSimpleOutput(output)
}

func (builder *TxBuilder) AddMultiAssetOutput(output PSL.MultiAssetOutput) {
  builder.Tx.AddMultiAssetOutput(output)
}

func (builder *TxBuilder) AddMultiAddrOutput(output PSL.MultiAddrOutput) {
  builder.Tx.AddMultiAddrOutput(output)
}

func (builder *TxBuilder) AddData(data PSL.TxData) {
  builder.Tx.AddData(data)
}

func (builder *TxBuilder) AddRequest(request PSL.Request) error {
  cborBytes, err := request.ToCBOR()
  if err != nil { return err }

  requestData := PSL.TxData{}
  requestData.Tag = "validator_request"
  requestData.Data = cborBytes
  requestData.Type = 0

  builder.AddData(requestData)
  return nil
}
