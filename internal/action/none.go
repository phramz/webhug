package action

import (
	"github.com/phramz/webhug/internal/contract"
	"net/http"
)

type none struct {
	response bool
}

func (n *none) Dispatch(wh contract.Webhook, rq *http.Request, res http.ResponseWriter) {
	log.Infof("[%s] running action none", wh.GetName())
}

func (n *none) HasResponse() bool {
	return n.response
}
