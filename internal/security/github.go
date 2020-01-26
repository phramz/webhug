package security

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"github.com/phramz/webhug/internal/contract"
	"io/ioutil"
	"net/http"
)

type github struct {
	secret string
}

func (gh *github) IsGranted(wh contract.Webhook, rq *http.Request) bool {
	actual := []byte(rq.Header.Get("x-hub-signature"))
	body, _ := ioutil.ReadAll(rq.Body)
	expected := []byte(githubSign(body, []byte(gh.secret)))

	if !hmac.Equal(actual, expected) {
		log.Infof("[%s] access denied from %s. Reason: %s", wh.GetName(), rq.RemoteAddr,
			fmt.Sprintf("wrong x-hub-signature header expected '%s', got '%s'.", expected, actual))

		return false
	}

	log.Infof("[%s] access granted from %s", wh.GetName(), rq.RemoteAddr)
	return true
}

func githubSign(payload []byte, secret []byte) string {
	mac := hmac.New(sha1.New, secret)
	mac.Write(payload)

	return fmt.Sprintf("sha1=%x", mac.Sum(nil))
}
