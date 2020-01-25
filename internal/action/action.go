package action

import (
	"fmt"
	"github.com/phramz/webhug/internal/contract"
	"github.com/phramz/webhug/internal/logger"
	"github.com/spf13/viper"
	"strings"
)

const None string = "none"
const Shell string = "shell"

var (
	log = logger.Logger()
)

func FromConfig(name string) contract.Action {
	actionType := strings.ToLower(viper.GetString(fmt.Sprintf("webhug.webhooks.%s.action.type", name)))
	switch actionType {
	case Shell:
		return &shell{
			cmd:      viper.GetString(fmt.Sprintf("webhug.webhooks.%s.action.cmd", name)),
			args:     viper.GetStringSlice(fmt.Sprintf("webhug.webhooks.%s.action.args", name)),
			env:      viper.GetStringSlice(fmt.Sprintf("webhug.webhooks.%s.action.env", name)),
			response: viper.GetBool(fmt.Sprintf("webhug.webhooks.%s.action.response", name)),
		}
	case None:
		return &none{}
	}

	log.Warningf("unsupported action type for '%s'! Falling back to none. ",
		fmt.Sprintf("webhug.webhooks.%s.action.type", name))

	return &none{}
}
