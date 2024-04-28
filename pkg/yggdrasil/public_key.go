package yggdrasil

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func MustParseKey(s string) PublicKey {
	if key, err := ParseKey(s); err != nil {
		panic(err)
	} else {
		return key
	}
}

func ParseKey(s string) (PublicKey, error) {
	key, err := hex.DecodeString(s)
	switch {
	case err != nil:
		return nil, err
	case len(key) < 32:
		return nil, fmt.Errorf("not enough octets in string (got %d)", len(key))
	}
	return key, nil
}

// PublicKey represents Yggdrasil network node public key.
type PublicKey []byte

// String returns the string representation of the PublicKey.
//
// The String method encodes the PublicKey to a hexadecimal string.
func (key PublicKey) String() string {
	return hex.EncodeToString(key)
}

// Equal checks if the current PublicKey is equal to another PublicKey.
func (key PublicKey) Equal(other PublicKey) bool {
	return key.String() == other.String()
}

// IPv6Address transforms the PublicKey into an network IPv6 address.
func (key PublicKey) IPv6Address() string {
	buf := make([]byte, 32)
	copy(buf[:], key[:])
	for i := range buf {
		buf[i] = ^buf[i]
	}
	leadOnesCount := countLeadOnes(buf)
	stripBitsCount := leadOnesCount + 1

	var addressBytes [16]byte
	buf = buf[stripBitsCount/8:]
	for i := 0; i < 14; i++ {
		buf[i] = buf[i] << (stripBitsCount % 8)
		buf[i] |= buf[i+1] >> (8 - stripBitsCount%8)
		addressBytes[i+2] = buf[i]
	}

	addressBytes[0] = 0x2
	addressBytes[1] = byte(leadOnesCount)
	return bytesToIPv6(addressBytes)
}

func countLeadOnes(bytes []byte) uint {
	var leadOnesCount uint
	for _, b := range bytes {
		for i := 7; i >= 0; i-- {
			if b&(1<<i) == 0 {
				return leadOnesCount
			}
			leadOnesCount++
		}
	}
	return leadOnesCount
}

func bytesToIPv6(bytes [16]byte) string {
	var hextects [8]string
	for i := 0; i < 8; i++ {
		var firstOctet string
		if bytes[i*2] > 0 {
			firstOctet = hex.EncodeToString(bytes[i*2 : i*2+1])
			firstOctet = stripZero(firstOctet)
		}
		secondOctet := hex.EncodeToString(bytes[i*2+1 : i*2+2])
		if firstOctet == "" {
			secondOctet = stripZero(secondOctet)
		}
		hextects[i] = firstOctet + secondOctet
	}

	result := strings.Builder{}
	zeroesSkipped := false
	for i := 0; i < 8; {
		if hextects[i] == "0" && !zeroesSkipped {
			for hextects[i] == "0" {
				i++
			}
			zeroesSkipped = true
			result.WriteRune(':')
			continue
		}

		result.WriteString(hextects[i])
		if i < 7 {
			result.WriteRune(':')
		}
		i++
	}
	return result.String()
}

func stripZero(s string) string {
	if s[0] == '0' {
		return s[1:]
	} else {
		return s
	}
}
