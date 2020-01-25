package security

import (
	"github.com/phramz/webhug/internal/contract"
	"net/http"
)

type none struct {
}

func (n *none) IsGranted(wh contract.Webhook, rq *http.Request) bool {
	log.Infof("[%s] access granted from %s", wh.GetName(), rq.RemoteAddr)

	return true
}
