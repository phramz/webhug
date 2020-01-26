package crypto

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

func GithubSign(payload []byte, secret []byte) string {
	mac := hmac.New(sha1.New, secret)
	mac.Write(payload)

	return fmt.Sprintf("sha1=%x", mac.Sum(nil))
}
