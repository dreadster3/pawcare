package common

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	HTTPPortKey              = "HTTP_PORT"
	GRPCPortKey              = "GRPC_PORT"
	DBHostKey                = "DB_HOST"
	DBPortKey                = "DB_PORT"
	DBUserKey                = "DB_USER"
	DBPassKey                = "DB_PASS"
	DBConnectionStringKey    = "DB_CONNECTION_STRING"
	JWTSecretKey             = "JWT_SECRET"
	IntrospectionEndpointKey = "INTROSPECTION_ENDPOINT"
	ClientIdKey              = "CLIENT_ID"
	ClientSecretKey          = "CLIENT_SECRET"
)

func SetConnectionStringConfig(viper *viper.Viper) {
	user := viper.GetString(DBUserKey)
	password := viper.GetString(DBPassKey)
	host := viper.GetString(DBHostKey)
	port := viper.GetString(DBPortKey)

	dbConnectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)

	viper.Set(DBConnectionStringKey, dbConnectionString)
}
