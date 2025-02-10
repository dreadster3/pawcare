package config

import (
	"github.com/spf13/viper"
)

const (
	HTTPPortKey  = "HTTP_PORT"
	GRPCPortKey  = "GRPC_PORT"
	DBHostKey    = "DB_HOST"
	DBPortKey    = "DB_PORT"
	DBUserKey    = "DB_USER"
	DBPassKey    = "DB_PASS"
	JWTSecretKey = "JWT_SECRET"
)

func InitConfig() viper.Viper {
	viper := viper.New()

	viper.SetEnvPrefix("PAWCARE")
	viper.AddConfigPath("$HOME/.config/pawcare")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.BindEnv(HTTPPortKey)
	viper.BindEnv(GRPCPortKey)
	viper.BindEnv(DBHostKey)
	viper.BindEnv(DBPortKey)
	viper.BindEnv(DBUserKey)
	viper.BindEnv(DBPassKey)
	viper.BindEnv(JWTSecretKey)

	viper.SetDefault(HTTPPortKey, "8080")
	viper.SetDefault(GRPCPortKey, "8081")
	viper.SetDefault(DBHostKey, "localhost")
	viper.SetDefault(DBPortKey, "27017")
	viper.SetDefault(DBUserKey, "root")
	viper.SetDefault(DBPassKey, "root")

	return *viper
}
