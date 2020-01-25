package contract

import (
	"net/http"
)

type Webhook interface {
	GetName() string
	GetFormat() string
	GetSecurity() Security
	GetAction() Action
}

type Action interface {
	Dispatch(wh Webhook, rq *http.Request, res http.ResponseWriter)
	HasResponse() bool
}

type Security interface {
	IsGranted(wh Webhook, rq *http.Request) bool
}
