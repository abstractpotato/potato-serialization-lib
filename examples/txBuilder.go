package main

import "fmt"
// import PSL "github.com/abstractpotato/potato-serialization-lib/psl"
import Builders "github.com/abstractpotato/potato-serialization-lib/builders"

func main() {
  txBuilder := Builders.NewTxBuilder()
  tx, _ := txBuilder.Build()
  fmt.Printf("%+v\n", tx)
}
