package security

import (
	"github.com/phramz/webhug/internal/contract"
	"github.com/phramz/webhug/pkg/tpl"
)

type deny struct {
}

func (n *deny) IsGranted(ctx *contract.Context) bool {
	log.Infof(tpl.MustRender(`[{{ .Webhook.Name }}] access denied from {{ .Request.RemoteAddr }}`, ctx))
	return false
}
