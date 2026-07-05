package main

import (
	"fmt"

	Builders "github.com/abstractpotato/potato-serialization-lib/builders"
	PSL "github.com/abstractpotato/potato-serialization-lib/psl"
)

func main() {
	txBuilder := Builders.NewTxBuilder()
	txBuilder.Context = PSL.NewContext()
	txBuilder.Context.Params = PSL.NewParams()
	txBuilder.Context.Params.Network = 0
	txBuilder.Context.Params.TxFeePerByte = 430
	txBuilder.Context.Epoch = 1
	txBuilder.Context.Slot = 42
	txBuilder.AddOutput(PSL.TxOutput{
		To:     "target_addr",
		Asset:  "starch",
		Amount: 1000,
	})
	tx, _ := txBuilder.Build()
	fmt.Printf("%+v\n", tx)
}
