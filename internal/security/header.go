package security

import (
	"fmt"
	"github.com/phramz/webhug/internal/contract"
	"github.com/phramz/webhug/pkg/tpl"
	"net/textproto"
)

type header struct {
	key   string
	value string
}

func (hdr *header) IsGranted(ctx *contract.Context) bool {
	expectedVal := tpl.MustRender(hdr.value, ctx)
	expectedHdr := tpl.MustRender(hdr.key, ctx)
	actual := tpl.MustRender(fmt.Sprintf(`{{ index .Request.Header "%s" }}`, textproto.CanonicalMIMEHeaderKey(expectedHdr)), ctx)

	if actual != expectedVal {
		reason := fmt.Sprintf("wrong request header '%s: %s' ", expectedHdr, actual)
		log.Infof(tpl.MustRender(`[{{ .Webhook.Name }}] access denied from {{ .Request.RemoteAddr }}. Reason %s`, ctx), reason)

		return false
	}

	log.Infof(tpl.MustRender(`[{{ .Webhook.Name }}] access granted from {{ .Request.RemoteAddr }}`, ctx))
	return true
}
