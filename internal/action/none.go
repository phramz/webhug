package action

import (
	"github.com/phramz/webhug/internal/contract"
	"github.com/phramz/webhug/pkg/tpl"
	"net/http"
)

type none struct {
	response bool
}

func (n *none) Dispatch(ctx *contract.Context, res http.ResponseWriter) (bool, error) {
	log.Infof(tpl.MustRender(`[{{ .Webhook.Name }}] running action none`, ctx))

	return true, nil
}

func (n *none) HasResponse() bool {
	return n.response
}
