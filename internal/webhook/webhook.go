package webhook

import (
	"fmt"
	"github.com/phramz/webhug/internal/action"
	"github.com/phramz/webhug/internal/contract"
	"github.com/phramz/webhug/internal/logger"
	"github.com/phramz/webhug/internal/security"
	"github.com/spf13/viper"
)

var (
	log = logger.Logger()
)

func FromConfig(name string) contract.Webhook {
	wh := &webhook{
		name:     name,
		format:   viper.GetString(fmt.Sprintf("webhug.webhooks.%s.format", name)),
		security: security.FromConfig(name),
		action:   action.FromConfig(name),
	}

	return wh
}

type webhook struct {
	name     string
	format   string
	security contract.Security
	action   contract.Action
}

func (wh *webhook) GetName() string {
	return wh.name
}

func (wh *webhook) GetFormat() string {
	return wh.format
}

func (wh *webhook) GetSecurity() contract.Security {
	return wh.security
}

func (wh *webhook) GetAction() contract.Action {
	return wh.action
}
