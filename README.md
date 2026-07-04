# Potato Serialization Lib
This is a library written in Go, for serialization & deserialization of data structures used in StarchPay's Go implementation.

Add to Go project
```
go get github.com/abstractpotato/potato-serialization-lib
```

In Go code
```go
import PSL "github.com/potato-serialization-lib/psl"
```

<!-- ## Documentation -->
<!-- [Click Here](https://starch.one/docs/starch_pay) -->

## Build A Transaction Example
```go
package main

import (
  PSL "github.com/abstractpotato/potato-serialization-lib/psl"
)

func main() {
  output := PSL.TxOutput{}
  output.To = "your_target_cardano_address"
  output.Asset = "3d77d63dfa6033be98021417e08e3368cc80e67f8d7afa196aaa0b3953746172636820546f6b656e"
  output.Amount = 1000

  transaction := PSL.NewTransaction()
  transaction.Outputs = append(transaction.Outputs, output)

  // sign these bytes (not yet included in this module)
  cborBytes := transaction.BodyToCBOR()

  transaction.Hash() // generate the transaction hash

  transaction.ToCBOR() // this would be submitted to the network
}
```

## Core maintainers:
[Abstract Potato](https://github.com/abstractpotato)
