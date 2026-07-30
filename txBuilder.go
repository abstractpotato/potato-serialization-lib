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

func (builder *TxBuilder) SetSender(addr string) {
  builder.Tx.SetSender(addr)
}

func (builder *TxBuilder) SetSimpleOutput(output *SimpleOutput) {
  builder.Tx.Body.SimpleOutput = output
}

func (builder *TxBuilder) SetMultiAssetOutput(output *MultiAssetOutput) {
  builder.Tx.Body.MultiAssetOutput = output
}

func (builder *TxBuilder) SetMultiAddrOutput(output *MultiAddrOutput) {
  builder.Tx.Body.MultiAddrOutput = output
}

func (builder *TxBuilder) SetRequest(request *Request) {
  builder.Tx.SetRequest(request)
}

func (builder *TxBuilder) AddCertificate(certificate *Certificate) {
  builder.Tx.SetCertificate(certificate)
}

func (builder *TxBuilder) AddData(data TxData) {
  builder.Tx.AddData(data)
}