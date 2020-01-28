package contract

import (
	"net/http"
)

type Context struct {
	Env     map[string]string
	Webhook *WebhookContext
	Request *RequestContext
}

type WebhookContext struct {
	Name   string
	Format string
}

type RequestContext struct {
	Body       string
	Method     string
	Json       interface{}
	Uri        string
	Host       string
	RemoteAddr string
	Query      string
	Scheme     string
	Username   string
	Password   string
	Header     map[string]string
	Get        map[string][]string
	Cookie     map[string]*http.Cookie
}

type Webhook interface {
	GetName() string
	GetFormat() string
	GetSecurity() Security
	GetAction() Action
}

type Action interface {
	Dispatch(ctx *Context, res http.ResponseWriter)
	HasResponse() bool
}

type Security interface {
	IsGranted(ctx *Context) bool
}
