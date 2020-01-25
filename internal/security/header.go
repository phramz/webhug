package security

import (
	"fmt"
	"github.com/phramz/webhug/internal/contract"
	"net/http"
)

type header struct {
	key   string
	value string
}

func (hdr *header) IsGranted(wh contract.Webhook, rq *http.Request) bool {
	actual := rq.Header.Get(hdr.key)

	if actual != hdr.value {

		log.Infof("[%s] access denied from %s. Reason: %s", wh.GetName(), rq.RemoteAddr,
			fmt.Sprintf("wrong request header '%s: %s' ", hdr.key, actual))

		return false
	}

	log.Infof("[%s] access granted from %s", wh.GetName(), rq.RemoteAddr)
	return true
}
