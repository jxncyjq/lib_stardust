package crypto

import (
	"crypto/sha1"
	"encoding/hex"
)

func generateSHA1(input string) string {
	hash := sha1.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
