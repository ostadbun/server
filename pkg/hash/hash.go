package hash

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func Hasher(text string) string {

	SecretKey1 := os.Getenv("SECRET_HASHER1")
	SecretKey3 := os.Getenv("SECRET_HASHER3")
	SecretKey5 := os.Getenv("SECRET_HASHER5")

	txt := fmt.Sprintf("(((%s)%s)%s)%s)",
		SecretKey3,
		SecretKey5,
		text,
		SecretKey1,
	)

	return fmt.Sprintf("%x", sha256.Sum256([]byte(txt)))
}
