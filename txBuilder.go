package psl

import (
  "time"
  "errors"
  "math"
)

type TxBuilder struct {
  Params Params
  Tx     Transaction
}

func NewTxBuilder() TxBuilder {
  return TxBuilder{
    Tx: NewTransaction(),
  }
}

func (builder *TxBuilder) EstimateFee() error {
  cborBytes, err := builder.Tx.Body.ToCBOR()
  if err != nil { return err }
  
  maxTxSize := builder.Params.MaxTxSize
  minTxFee := builder.Params.MinTxFee
  txFeePerByte := builder.Params.TxFeePerByte
  builder.Tx.Body.Fee = minTxFee
  
  size := uint(len(cborBytes))
  if size > maxTxSize { return errors.New("tx size exceeds MaxTxSize") }
  
  dryRunFee := minTxFee + (size * txFeePerByte)
  builder.Tx.Body.Fee = dryRunFee
  
  cborBytes, err = builder.Tx.Body.ToCBOR()
  if err != nil { return err }
  
  size = uint(len(cborBytes))
  if size > maxTxSize { return errors.New("tx size exceeds MaxTxSize") }

  finalFee := minTxFee + (size * txFeePerByte)
  if finalFee > math.MaxUint { return errors.New("fee overflow") }
  
  builder.Tx.Body.Fee = finalFee
  return nil
}

func (builder *TxBuilder) Build() error {
  builder.Tx.Body.Network = builder.Params.Network
  builder.Tx.Body.TTL = 3000 // 3 seconds
  builder.Tx.Body.Timestamp = uint(time.Now().UnixMilli())

  err := builder.EstimateFee()
  if err != nil { return err }

  err = builder.Tx.Hash()
  if err != nil { return err }

  return nil
}

func (builder *TxBuilder) Sign(privateKey []byte) error {
  return builder.Tx.Sign(privateKey)
}

func (builder *TxBuilder) Verify() bool {
  return builder.Tx.Verify()
}

func (builder *TxBuilder) AddSimpleOutput(output SimpleOutput) {
  builder.Tx.AddSimpleOutput(output)
}

func (builder *TxBuilder) AddMultiAssetOutput(output MultiAssetOutput) {
  builder.Tx.AddMultiAssetOutput(output)
}

func (builder *TxBuilder) AddMultiAddrOutput(output MultiAddrOutput) {
  builder.Tx.AddMultiAddrOutput(output)
}

func (builder *TxBuilder) AddData(data TxData) {
  builder.Tx.AddData(data)
}

func (builder *TxBuilder) AddRequest(request *Request) {
  builder.Tx.AddRequest(request)
}

func (builder *TxBuilder) AddCertificate(certificate *Certificate) {
  builder.Tx.AddCertificate(certificate)
}
