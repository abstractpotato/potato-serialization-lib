package main

import (
  "fmt"
  "github.com/abstractpotato/potato-serialization-lib"
)

func main() {
  privateKey, err := psl.GenerateKeys("")
  if err != nil { panic(err) }
  fmt.Println(len(privateKey))
  
  vkey, err := psl.GetPublicKey(privateKey[:32])
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
  
  publicKey, err := psl.GetPublicKey(privateKey[:32])
  if err != nil { panic(err) }
  
  value, err := psl.AddressBelongsToPubKey(addr, publicKey)
  if err != nil { panic(err) }
  fmt.Println(value)
  
  valueB, err := psl.AddressBelongsToPubKey(addrB, publicKey)
  if err != nil { panic(err) }
  fmt.Println(valueB)
}