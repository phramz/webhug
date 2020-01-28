package security

import (
	"crypto/hmac"
	"fmt"
	"github.com/phramz/webhug/internal/contract"
	"github.com/phramz/webhug/pkg/crypto"
	"github.com/phramz/webhug/pkg/tpl"
)

type github struct {
	secret string
}

func (gh *github) IsGranted(ctx *contract.Context) bool {
	secret := tpl.MustRender(gh.secret, ctx)

	actual := []byte(tpl.MustRender(`{{ index .Request.Header "x-hub-signature" }}`, ctx))
	body := []byte(tpl.MustRender(`{{ .Request.Body }}`, ctx))
	expected := []byte(crypto.GithubSign(body, []byte(secret)))

	if !hmac.Equal(actual, expected) {
		reason := fmt.Sprintf("wrong x-hub-signature header. Expected '%s', got '%s'.", expected, actual)
		log.Infof(tpl.MustRender(`[{{ .Webhook.Name }}] access denied from {{ .Request.RemoteAddr }}. Reason %s`, ctx), reason)

		return false
	}

	log.Infof(tpl.MustRender(`[{{ .Webhook.Name }}] access granted from {{ .Request.RemoteAddr }}`, ctx))
	return true
}
