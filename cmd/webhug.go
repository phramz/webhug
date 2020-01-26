package main

import (
	"fmt"
	"github.com/phramz/webhug/internal/config"
	"github.com/phramz/webhug/internal/contract"
	"github.com/phramz/webhug/internal/logger"
	"github.com/spf13/viper"
	"net/http"
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
		if !wh.GetSecurity().IsGranted(wh, r) {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("403 - Forbidden"))

			return
		}

		w.WriteHeader(http.StatusOK)
		wh.GetAction().Dispatch(wh, r, w)
	})
}
