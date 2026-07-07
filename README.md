# Potato Serialization Lib
This is a library written in Go, for serialization & deserialization of data structures used in StarchPay's Go implementation.

Add to Go project
```
go get github.com/abstractpotato/potato-serialization-lib
```

In Go code
```go
import "github.com/abstractpotato/potato-serialization-lib/psl"
import "github.com/abstractpotato/potato-serialization-lib/builders"
```

Run Examples
```
go run examples/<example_name.go>
```

## Build A Transaction Example
```go
package main

import (
  PSL "github.com/abstractpotato/potato-serialization-lib/psl"
  Builders "github.com/abstractpotato/potato-serialization-lib/builders"
)

func main() {
  output := PSL.SimpleOutput()
  output.To = "targer_cardano_address"
  output.Asset = "policy_id+asset_name"
  output.Amount = 1000

  txBuilder := Builders.NewTxBuilder()
  txBuilder.Params = params // this assumes you already have network params
  txBuilder.AddSimpleOutput(output)
  txBuilder.EstimateFee()
  txBuilder.Build()

  transaction := txBuilder.Tx

  bodyCBOR, _ := transaction.BodyToCBOR()

  // sign bodyCBOR (not yet in this module)

  transaction.AddSignature(signature)

  txCBOR, _ := transaction.ToCBOR()

  // submit txCBOR to the network
}
```

## Core maintainers:
[Abstract Potato](https://github.com/abstractpotato)
