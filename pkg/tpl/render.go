package tpl

import (
	"bytes"
	"github.com/patrickmn/go-cache"
	"github.com/phramz/webhug/internal/contract"
	"github.com/phramz/webhug/pkg/crypto"
	"text/template"
)

var tplCache = cache.New(-1, -1)

func MustRender(tpl string, ctx *contract.Context) string {
	name := crypto.Sha512Hex(tpl)

	cached, exists := tplCache.Get(name)
	if !exists {
		cached = template.Must(template.New(name).Parse(tpl))
		_ = tplCache.Add(name, cached, -1)
	}

	td := cached.(*template.Template)
	var out bytes.Buffer
	err := td.Execute(&out, ctx)
	if err != nil {
		panic(err)
	}

	return out.String()
}

func MustRenderAll(tpls []string, ctx *contract.Context) []string {
	var out []string

	for _, v := range tpls {
		out = append(out, MustRender(v, ctx))
	}

	return out
}
