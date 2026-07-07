package main

import "fmt"
import PSL "github.com/abstractpotato/potato-serialization-lib/psl"
import Builders "github.com/abstractpotato/potato-serialization-lib/builders"

func main() {
  // sample data
  params := PSL.NewParams()
  params.Network = 0
  params.MaxTxSize = 4000
  params.MinTxFee = 4300
  params.TxFeePerByte = 430

  txBuilder := Builders.NewTxBuilder()
  txBuilder.Params = params

  request := PSL.NewRequest()
  request.Ticker = "bone"
  request.Url = "https://bonepool.com"
  request.Addr = "addr1qytdq4cjldj7lruyq5ppm7wzg6z7tk95j8njqay3g60f8rq7k3cvvlkegt9wv8ar4knyl4vkj63w8e5a8rzm6fqsx6hjk4gzvt3"
  request.Relays = append(request.Relays, "0.0.0.0:5001")
  request.Relays = append(request.Relays, "0.0.0.0:5002")

  txBuilder.AddRequest(request)

  txBuilder.Build()
  fmt.Printf("%+v\n", txBuilder.Tx)
  cborBytes, _ := txBuilder.Tx.BodyToCBOR()
  fmt.Println(len(cborBytes))
}
