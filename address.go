package psl

import (
  "fmt"
  "bytes"
	"github.com/btcsuite/btcutil/bech32"
	"golang.org/x/crypto/blake2b"
)

const (
	TypeBase       = 0 // payment key + stake key
	TypeEnterprise = 6 // payment key only
	// (other types exist but are omitted for brevity)
)

type ParsedAddress struct {
	HRP          string
	NetworkMain  bool   // true = mainnet
	Type         int    // high nibble
	PaymentHash  []byte // 28 bytes (always present for the types we care about)
	StakeHash    []byte // 28 bytes (nil for enterprise)
	RawPayload   []byte // the full decoded bytes including header
}

func hashPubKey(pubKey []byte) ([]byte, error) {
	if len(pubKey) != 32 {
    return nil, fmt.Errorf("public key must be exactly 32 bytes. ", len(pubKey))
	}
	h, err := blake2b.New(28, nil)
	if err != nil {
		return nil, err
	}
	h.Write(pubKey)
	return h.Sum(nil), nil
}

func PubKeyToEnterpriseAddress(pubKey []byte, mainnet bool) (string, error) {
  keyHash, err := hashPubKey(pubKey)
  if err != nil { return "", err }

	var header byte
	var hrp string
	if mainnet {
		header = 0x61
    hrp = "spud"
	} else {
		header = 0x60
    hrp = "spud_test"
	}

  payload := make([]byte, 1+len(keyHash))
	payload[0] = header
	copy(payload[1:], keyHash)

	converted, err := bech32.ConvertBits(payload, 8, 5, true)
	if err != nil { return "", err }

	return bech32.Encode(hrp, converted)
}

func PubKeysToBaseAddress(paymentPub, stakePub []byte, mainnet bool) (string, error) {
	paymentHash, err := hashPubKey(paymentPub)
	if err != nil {
		return "", err
	}
	stakeHash, err := hashPubKey(stakePub)
	if err != nil { return "", err }

	var header byte
	var hrp string
	if mainnet {
		header = 0x01
    hrp = "spud"
	} else {
		header = 0x00
    hrp = "spud_test"
	}

	payload := make([]byte, 1+28+28)
	payload[0] = header
	copy(payload[1:29], paymentHash)
	copy(payload[29:], stakeHash)

	converted, err := bech32.ConvertBits(payload, 8, 5, true)
	if err != nil { return "", err }
	return bech32.Encode(hrp, converted)
}

func ParseAndValidate(addr string) (*ParsedAddress, error) {
	hrp, data5, err := bech32.Decode(addr)
	if err != nil {
		return nil, fmt.Errorf("bech32 decode: %w", err)
	}
  if hrp != "spud" && hrp != "spud_test" {
		return nil, fmt.Errorf("unsupported HRP %q (expected addr or addr_test)", hrp)
	}

	payload, err := bech32.ConvertBits(data5, 5, 8, false)
	if err != nil {
		return nil, fmt.Errorf("convert bits: %w", err)
	}
	if len(payload) < 1 {
		return nil, fmt.Errorf("empty payload")
	}

	header := payload[0]
	addrType := int(header >> 4)
	networkBit := header & 0x0F
	isMainnet := networkBit == 1

	// Enforce network ↔ HRP consistency
  if (hrp == "spud" && !isMainnet) || (hrp == "spud_test" && isMainnet) {
		return nil, fmt.Errorf("HRP/network mismatch")
	}

	pa := &ParsedAddress{
		HRP:         hrp,
		NetworkMain: isMainnet,
		Type:        addrType,
		RawPayload:  payload,
	}

	switch addrType {
	case TypeEnterprise: // 0x6x
		if len(payload) != 29 {
			return nil, fmt.Errorf("enterprise address must be 29 bytes, got %d", len(payload))
		}
		pa.PaymentHash = payload[1:29]
	case TypeBase: // 0x0x
		if len(payload) != 57 {
			return nil, fmt.Errorf("base address must be 57 bytes, got %d", len(payload))
		}
		pa.PaymentHash = payload[1:29]
		pa.StakeHash = payload[29:57]
	default:
		return nil, fmt.Errorf("unsupported address type %d (only base=0 and enterprise=6 implemented)", addrType)
	}

	return pa, nil
}

func VerifyAgainstPubKeys(addr string, paymentPub, stakePub []byte) (bool, error) {
	pa, err := ParseAndValidate(addr)
	if err != nil { return false, err }

	if paymentPub != nil {
		want, err := hashPubKey(paymentPub)
		if err != nil { return false, err }
		if !bytes.Equal(pa.PaymentHash, want) { return false, nil }
	}

	if stakePub != nil {
		if pa.Type != TypeBase {
			return false, fmt.Errorf("stake key supplied but address is not a base address")
		}
		want, err := hashPubKey(stakePub)
    if err != nil { return false, err }
		if !bytes.Equal(pa.StakeHash, want) { return false, nil }
	}

	return true, nil
}