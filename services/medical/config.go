package main

import "github.com/spf13/viper"

func SetupConfig() *viper.Viper {
	viper := viper.New()

	viper.SetEnvPrefix("pawcare")

	viper.BindEnv("port")
	viper.BindEnv("log_level")
	viper.BindEnv("db_host")
	viper.BindEnv("db_port")
	viper.BindEnv("db_name")
	viper.BindEnv("db_user")
	viper.BindEnv("db_password")
	viper.BindEnv("jwt_secret")

	viper.SetDefault("port", "8080")
	viper.SetDefault("log_level", "INFO")
	viper.SetDefault("db_host", "localhost")
	viper.SetDefault("db_port", "27017")
	viper.SetDefault("db_name", "medical")
	viper.SetDefault("db_user", "user")
	viper.SetDefault("db_password", "password")

	viper.AddConfigPath("$HOME/.config/pawcare")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	return viper
}
