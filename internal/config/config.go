package config

import (
	"fmt"

	"github.com/phramz/webhug/internal/contract"
	"github.com/phramz/webhug/internal/logger"
	"github.com/phramz/webhug/internal/webhook"
	"github.com/spf13/viper"
)

var (
	log = logger.Logger()
)

func MustReadConfig() []contract.Webhook {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/webhug/")
	viper.AddConfigPath("$HOME/.webhug")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// fallback to example config
			viper.SetConfigName("config-example")
			if err := viper.ReadInConfig(); err != nil {
				log.Fatal(fmt.Errorf("Fatal error config file: %s \n", err))
			}
		} else {
			log.Fatal(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}

	return parseWebhooks(viper.GetStringMap("webhug.webhooks"))
}

func parseWebhooks(webhooks map[string]interface{}) []contract.Webhook {
	var whs []contract.Webhook

	for name := range webhooks {
		whs = append(whs, webhook.FromConfig(name))
	}

	return whs
}
