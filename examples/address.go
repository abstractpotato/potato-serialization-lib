package main

import (
  "fmt"
  "github.com/abstractpotato/potato-serialization-lib"
)

func main() {
  privateKey, err := psl.GenerateKeys("")
  if err != nil { panic(err) }
  fmt.Println(len(privateKey))
  
  vkey, err := psl.MakePublicKey(privateKey[:32])
  if err != nil { panic(err) }
  fmt.Println(len(vkey))
  
  addr, err := psl.GenerateEnterpriseAddr(privateKey, true)
  if err != nil { panic(err) }
  fmt.Println(addr)
  fmt.Println(psl.IsValidAddress(addr))
  
  stakeKey, err := psl.GenerateKeys("")
  if err != nil { panic(err) }
  addrB, err := psl.GenerateBaseAddr(privateKey, stakeKey, true)
  fmt.Println(addrB)
  fmt.Println(psl.IsValidAddress(addrB))
}