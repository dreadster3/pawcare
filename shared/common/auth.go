package common

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func JWTKeyFactory(viper viper.Viper) func(token *jwt.Token) (interface{}, error) {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString(JWTSecretKey)), nil
	}
}
