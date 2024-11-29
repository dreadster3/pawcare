package main

import (
	"fmt"

	"github.com/dreadster3/pawcare/services/account/config"
	"github.com/spf13/viper"
)

func BuildConnectionString(viper viper.Viper) string {
	user := viper.GetString(config.DBUserKey)
	password := viper.GetString(config.DBPassKey)
	host := viper.GetString(config.DBHostKey)
	port := viper.GetString(config.DBPortKey)

	return fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)
}
