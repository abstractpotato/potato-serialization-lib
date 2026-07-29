package main

import (
  "fmt"
  "encoding/hex"
  "github.com/abstractpotato/potato-serialization-lib"
)

const skey = "c0e5981efee192773da5a3542b28da40b48638eff0bf5495dc016f4ecc0c55534b0853da95378d4ecbf184920b1dec5747212915977718b5094ef0c45ee0cfb0a8f448cbb86544765fa7ae7a0ef604768c10054de52498d59ba00995ca6ec66696bcefe574605f16a8166e3219a1a012fc04c6f1929003917f9f805784930784"

func GetPrivateKey() []byte {
  privateKey, err := hex.DecodeString(skey)
  if err != nil { panic(err) }
  return privateKey[:96]
}

func main() {
  privateKey := GetPrivateKey()
  publicKey, err := psl.MakePublicKey(privateKey[:32])
  if err != nil { panic(err) }
  
  addr_enterprise, err := psl.PubKeyToEnterpriseAddress(publicKey, true)
  if err != nil { panic(err) }
  
  fmt.Println(addr_enterprise)
  
  addr_base, err := psl.PubKeysToBaseAddress(publicKey, publicKey, true)
  if err != nil { panic(err) }
  
  fmt.Println(addr_base)
  
  valid, err := psl.VerifyAgainstPubKeys(addr_enterprise, publicKey, nil)
  if err != nil { panic(err) }
  
  fmt.Println(valid)
}