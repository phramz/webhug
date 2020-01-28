package crypto

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

func GithubSign(payload []byte, secret []byte) string {
	mac := hmac.New(sha1.New, secret)
	mac.Write(payload)

	return fmt.Sprintf("sha1=%x", mac.Sum(nil))
}

func Sha512Hex(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}
