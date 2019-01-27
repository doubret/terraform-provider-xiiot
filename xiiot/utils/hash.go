package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(val interface{}) string {
	sha := sha256.Sum256([]byte(val.(string)))
	return hex.EncodeToString(sha[:])
}
