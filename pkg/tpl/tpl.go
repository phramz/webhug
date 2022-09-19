package tpl

import (
	"encoding/json"
	"io"
	"net/http"
	"net/textproto"
	"net/url"
	"strings"

	"github.com/phramz/webhug/internal/contract"
	"github.com/phramz/webhug/internal/logger"
)

var (
	log = logger.Logger()
)

func DefaultContext(environ []string) *contract.Context {
	context := &contract.Context{
		Env: parseEnvironment(environ),
		Request: &contract.RequestContext{
			Header: make(map[string]string),
			Get:    make(map[string][]string),
			Cookie: make(map[string]*http.Cookie),
		},
		Webhook: &contract.WebhookContext{},
	}

	return context
}

func AddRequestContext(context *contract.Context, rq *http.Request) *contract.Context {
	body, err := io.ReadAll(rq.Body)
	if nil == err {
		context.Request.Body = string(body)
		var j interface{}

		ctype := rq.Header.Get("content-type")
		if strings.Contains(ctype, "json") {
			err = json.Unmarshal(body, &j)
			if err != nil {
				log.Warningf("error parsing json: %s", err)
			}
		}

		context.Request.Json = j
	}

	context.Request.Method = rq.Method
	context.Request.Uri = rq.RequestURI
	context.Request.Host = rq.Host
	context.Request.RemoteAddr = rq.RemoteAddr
	context.Request.Query = rq.URL.RawQuery
	context.Request.Scheme = rq.URL.Scheme
	context.Request.Username = rq.URL.User.Username()
	context.Request.Password, _ = rq.URL.User.Password()
	context.Request.Header = parseHeader(rq.Header)
	context.Request.Get = parseGet(rq.URL.Query())
	for _, cookie := range rq.Cookies() {
		context.Request.Cookie[cookie.Name] = cookie
	}

	return context
}

func AddWebhookContext(context *contract.Context, wh contract.Webhook) *contract.Context {
	context.Webhook.Name = wh.GetName()
	context.Webhook.Format = wh.GetFormat()

	return context
}

func parseGet(vals url.Values) map[string][]string {
	return vals
}

func parseHeader(header http.Header) map[string]string {
	hdrVars := make(map[string]string)

	for name, hdrs := range header {
		for _, hdr := range hdrs {
			hdrVars[strings.ToLower(name)] = hdr
			hdrVars[textproto.CanonicalMIMEHeaderKey(name)] = hdr
		}
	}

	return hdrVars
}

func parseEnvironment(environ []string) map[string]string {
	envVars := make(map[string]string)

	for _, pair := range environ {
		envVar := strings.SplitAfterN(pair, "=", 2)
		envVars[strings.TrimRight(envVar[0], "=")] = envVar[1]
	}

	return envVars
}
