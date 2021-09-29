package util

import (
	"crypto/sha512"
	"encoding/hex"
)

func EncodeSHA512(value string) string {
	m := sha512.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
