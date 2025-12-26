package hash

import (
	"crypto/sha256"
	"fmt"
)

func Hasher(text string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(text)))
}
