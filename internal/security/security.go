package security

import (
	"fmt"
	"github.com/phramz/webhug/internal/contract"
	"github.com/phramz/webhug/internal/logger"
	"github.com/spf13/viper"
	"strings"
)

const None string = "none"
const Header string = "header"

var (
	log = logger.Logger()
)

func FromConfig(name string) contract.Security {
	securityType := strings.ToLower(viper.GetString(fmt.Sprintf("webhug.webhooks.%s.security.type", name)))
	switch securityType {
	case Header:
		return &header{
			key:   strings.ToLower(viper.GetString(fmt.Sprintf("webhug.webhooks.%s.security.key", name))),
			value: strings.ToLower(viper.GetString(fmt.Sprintf("webhug.webhooks.%s.security.value", name))),
		}
	case None:
		return &none{}
	}

	log.Warningf("unsupported security type for '%s'! Falling back to deny all policy. "+
		"If you dont want to have access control at all please set security.type to 'none'. ",
		fmt.Sprintf("webhug.webhooks.%s.security.type", name))

	return &deny{}
}
