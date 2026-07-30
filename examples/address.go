package main

import (
  "fmt"
  "github.com/abstractpotato/potato-serialization-lib"
)

func main() {
  privateKey, err := psl.GenerateKeys("")
  if err != nil { panic(err) }
  fmt.Println(privateKey)
  
  publicKey, err := psl.GetPublicKey(privateKey[:32])
  if err != nil { panic(err) }
  fmt.Println(publicKey)
  
  addr, err := psl.GenerateEnterpriseAddr(privateKey, false)
  if err != nil { panic(err) }
  fmt.Println(addr)
  fmt.Println(psl.IsValidAddress(addr))
  
  stakeKey, err := psl.GenerateKeys("")
  if err != nil { panic(err) }
  addrB, err := psl.GenerateBaseAddr(privateKey, stakeKey, false)
  fmt.Println(addrB)
  fmt.Println(psl.IsValidAddress(addrB))
  
  value, err := psl.AddressBelongsToPubKey(addr, publicKey)
  if err != nil { panic(err) }
  fmt.Println(value)
  
  valueB, err := psl.AddressBelongsToPubKey(addrB, publicKey)
  if err != nil { panic(err) }
  fmt.Println(valueB)
}