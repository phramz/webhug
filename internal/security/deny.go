package security

import (
	"github.com/phramz/webhug/internal/contract"
	"net/http"
)

type deny struct {
}

func (n *deny) IsGranted(wh contract.Webhook, rq *http.Request) bool {
	log.Infof("[%s] access denied from %s", wh.GetName(), rq.RemoteAddr)
	return false
}
