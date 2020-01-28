package security

import (
	"github.com/phramz/webhug/internal/contract"
	"github.com/phramz/webhug/pkg/tpl"
)

type none struct {
}

func (n *none) IsGranted(ctx *contract.Context) bool {
	log.Infof(tpl.MustRender(`[{{ .Webhook.Name }}] access granted from {{ .Request.RemoteAddr }}`, ctx))

	return true
}
