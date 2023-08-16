package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type DatabaseConfig struct {
	DbHost   string
	DbName   string
	User     string
	Password string
}

func (cfg *DatabaseConfig) URI() (uri string) {
	if !strings.Contains(viper.GetString("stage"), "local") {
		uri = fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority",
			cfg.User,
			cfg.Password,
			cfg.DbHost,
			cfg.DbName,
		)

	} else {
		uri = fmt.Sprintf("mongodb://%s:%s@%s/%s",
			cfg.User,
			cfg.Password,
			cfg.DbHost,
			cfg.DbName,
		)

	}
	return
}
