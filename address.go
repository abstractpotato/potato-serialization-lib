package psl

import (
  "bytes"
  "crypto/rand"
  "github.com/echovl/cardano-go"
  "github.com/echovl/cardano-go/crypto"
)

func GenerateKeys(password string) ([]byte, error) {
  entropy := make([]byte, 32)
  if _, err := rand.Read(entropy); err != nil { return nil, err }

  rootXPrv := crypto.NewXPrvKeyFromEntropy(entropy, password)
  
  return rootXPrv, nil
}

func GenerateEnterpriseAddr(privateKey []byte, mainnet bool) (string, error) {
  pk := crypto.XPrvKey(privateKey)

  pub := pk.PubKey()

  keyHash, err := pub.Hash() // blake2b-224
  if err != nil { return "", err }

  paymentCred := cardano.StakeCredential{
    Type:    cardano.KeyCredential,
    KeyHash: keyHash,
  }

  if mainnet {
    addr, err := cardano.NewEnterpriseAddress(cardano.Mainnet, paymentCred)
    return addr.Bech32(), err
  }
  
  addr, err := cardano.NewEnterpriseAddress(cardano.Testnet, paymentCred)
  return addr.Bech32(), err
}

func GenerateBaseAddr(privateKey, stakeKey []byte, mainnet bool) (string, error){
  pk := crypto.XPrvKey(privateKey)
  sk := crypto.XPrvKey(stakeKey)
  
  paymentCred, err := cardano.NewKeyCredential(pk.PubKey())
  if err != nil { return "", err }
  
  stakeCred, err := cardano.NewKeyCredential(sk.PubKey())
  if err != nil { return "", err }
  
  if mainnet {
    addr, err := cardano.NewBaseAddress(cardano.Mainnet, paymentCred, stakeCred)
    return addr.Bech32(), err
  }
  
  addr, err := cardano.NewBaseAddress(cardano.Testnet, paymentCred, stakeCred)
  return addr.Bech32(), err
}

func IsValidAddress(addrStr string) bool {
  _, err := cardano.NewAddress(addrStr)
  return err == nil
}

func AddrBelongsToPubKey(addrStr string, publicKey []byte) (bool, error) {
  pk := crypto.PubKey(publicKey)
  
  addr, err := cardano.NewAddress(addrStr)
  if err != nil { return false, err }
  if addr.Payment.Type != cardano.KeyCredential { return false, nil }
  
  expected, err := pk.Hash()
  if err != nil { return false, err }
  
  return bytes.Equal(addr.Payment.KeyHash, expected), nil
}