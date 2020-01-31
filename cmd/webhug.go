package main

import (
	"fmt"
	"github.com/phramz/webhug/internal/config"
	"github.com/phramz/webhug/internal/contract"
	"github.com/phramz/webhug/internal/logger"
	"github.com/phramz/webhug/pkg/tpl"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

var (
	log = logger.Logger()
)

func main() {
	log.Infoln("reading config ...")

	for _, wh := range config.MustReadConfig() {
		path := fmt.Sprintf("/%s/", wh.GetName())
		log.Infof("setting up webhook '%s' at path '%s'", wh.GetName(), path)

		handle(path, wh)
	}

	listen := viper.GetString("webhug.listen")
	log.Infof("🤗 webhug listening on %s ...", listen)
	log.Fatal(http.ListenAndServe(listen, nil))
}

func handle(path string, wh contract.Webhook) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		ctx := tpl.DefaultContext(os.Environ())
		ctx = tpl.AddRequestContext(ctx, r)
		ctx = tpl.AddWebhookContext(ctx, wh)

		if !wh.GetSecurity().IsGranted(ctx) {
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte("403 - Forbidden"))

			return
		}

		success, _ := wh.GetAction().Dispatch(ctx, w)

		if success || wh.GetAction().HasResponse() {
			w.WriteHeader(http.StatusOK)
			return
		}

		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write([]byte("422 - Unprocessable Entity"))
		return
	})
}
