package bip32

// credit to https://github.com/blinklabs-io/bursa

import (
	"errors"
	"crypto/sha512"
	"filippo.io/edwards25519"
	"crypto/ed25519"
)

func Sign(skey, message []byte) ([]byte, error) {
  if len(skey) != 96 {
    return nil, errors.New("skey must be 96 bytes")
  }

  kL := skey[:32]
  kR := skey[32:64]
	publicKey, err := MakePublicKey(kL)
  if err != nil { return nil, err }

  nonceHash := sha512.New()
	nonceHash.Write(kR)
	nonceHash.Write(message)
  nonceDigest := nonceHash.Sum(nil)

  r, err := edwards25519.NewScalar().SetUniformBytes(nonceDigest)
  if err != nil { return nil, err }

	R := edwards25519.NewIdentityPoint().ScalarBaseMult(r)

  challengeHash := sha512.New()
	challengeHash.Write(R.Bytes())
	challengeHash.Write(publicKey)
	challengeHash.Write(message)
	challengeDigest := challengeHash.Sum(nil)

  h, err := edwards25519.NewScalar().SetUniformBytes(challengeDigest)
  if err != nil { return nil, err }

  padded := make([]byte, 64)
	copy(padded[:32], kL)
	s, err := edwards25519.NewScalar().SetUniformBytes(padded)
  if err != nil { return nil, err }

  S := edwards25519.NewScalar().MultiplyAdd(h, s, r)
  signature := make([]byte, 64)
	copy(signature[:32], R.Bytes())
	copy(signature[32:], S.Bytes())
  return signature, nil
}

func MakePublicKey(kL []byte) ([]byte, error) {
  padded := make([]byte, 64)
  copy(padded[:32], kL)

	s, err := edwards25519.NewScalar().SetUniformBytes(padded)
  if err != nil { return nil, err }

  p := edwards25519.NewIdentityPoint().ScalarBaseMult(s)
  return p.Bytes(), nil
}

func Verify(vkey, signature, message []byte) bool {
	publicKey := ed25519.PublicKey(vkey)
	return ed25519.Verify(publicKey, message, signature)
}
