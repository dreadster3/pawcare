package config

import (
	"github.com/dreadster3/pawcare/shared/common"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	PetServiceHostKey = "PET_SERVICE_HOST"
)

func InitConfig() viper.Viper {
	viper := viper.New()

	viper.SetEnvPrefix("PAWCARE")
	viper.AddConfigPath("$HOME/.config/pawcare")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.BindEnv(common.HTTPPortKey)
	viper.BindEnv(common.GRPCPortKey)
	viper.BindEnv(common.DBHostKey)
	viper.BindEnv(common.DBPortKey)
	viper.BindEnv(common.DBUserKey)
	viper.BindEnv(common.DBPassKey)
	viper.BindEnv(common.CertsEndpointKey)
	viper.BindEnv(common.InsecureSkipVerifyKey)
	viper.BindEnv(PetServiceHostKey)

	viper.SetDefault(common.HTTPPortKey, "8080")
	viper.SetDefault(common.GRPCPortKey, "8081")
	viper.SetDefault(common.DBHostKey, "localhost")
	viper.SetDefault(common.DBPortKey, "27017")
	viper.SetDefault(common.DBUserKey, "root")
	viper.SetDefault(common.DBPassKey, "root")
	viper.SetDefault(common.InsecureSkipVerifyKey, false)

	pflag.String("http-port", "", "HTTP Port")
	pflag.String("grpc-port", "", "GRPC Port")
	pflag.Parse()

	viper.BindPFlag(common.HTTPPortKey, pflag.Lookup("http-port"))
	viper.BindPFlag(common.GRPCPortKey, pflag.Lookup("grpc-port"))

	common.SetConnectionStringConfig(viper)

	return *viper
}
