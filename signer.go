package signer

import (
  "os/exec"
  "strings"
  "fmt"
)

type Verified struct {
  Verified bool   `cbor: "verified"`
  Addr     string `cbor: "addr"`
  Message  string `cbor: "message"`
}

func Verify(signature, key []byte) (Verified, error) {
  strSignature := string(signature)
  strKey := string(key)
  cmd := exec.Command("python3", "python/verify.py", strSignature, strKey)
  out, err := cmd.Output()
  if err != nil { return Verified{}, err }
  return ParseVerified(string(out)), nil
}

func ParseVerified(text string) Verified {
  verified := Verified{ Verified: true }
  s := strings.Split(text, "\n")
  if s[0] == "false" { verified.Verified = false; return verified }
  verified.Addr = s[1]
  verified.Message = s[2]
  return verified
}

func Sign(data string) ([]byte, error) {
  fmt.Println("v1")
  cmd := exec.Command("python3", "python/sign.py", data)
  return cmd.Output()
}
