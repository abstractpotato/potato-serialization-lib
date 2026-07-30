package psl

import (
  "time"
)

type BlockBuilder struct {
  Seed    []byte
  Params  Params
  Block   Block
}

func NewBlockBuilder() BlockBuilder {
  return BlockBuilder{
    Block: NewBlock(),
  }
}

func (builder *BlockBuilder) CalcFees() {
  builder.Block.Body.Fees = 0
  for _, tx := range builder.Block.Body.Transactions {
    builder.Block.Body.Fees += tx.Body.Fee
  }
}

func (builder *BlockBuilder) Build() {
  builder.Block.Body.Network = builder.Params.Network
  builder.Block.Body.Timestamp = uint(time.Now().UnixMilli())
  builder.CalcFees()
  builder.Block.Hash()
}

func (builder *BlockBuilder) AddTx(transaction Transaction) {
  builder.Block.Body.AddTx(transaction)
}

func (builder *BlockBuilder) SetGenesis(genesis *Genesis) {
  builder.Block.SetGenesis(genesis)
}

func (builder *BlockBuilder) Sign(privateKey []byte) error {
  return builder.Block.Sign(privateKey)
}

func (builder *BlockBuilder) Verify() bool {
  return builder.Block.Verify()
}
