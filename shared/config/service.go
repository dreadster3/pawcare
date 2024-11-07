package config

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type ServiceConfig struct {
	Scheme string
	Host   string
	Port   string
}

func InitServiceConfig(root *viper.Viper, serviceName string) {
	prefix := strings.ToUpper(fmt.Sprintf("pawcare_%s", serviceName))

	root.BindEnv("services.profile.scheme", fmt.Sprintf("%s_SCHEME", prefix))
	root.BindEnv("services.profile.host", fmt.Sprintf("%s_HOST", prefix))
	root.BindEnv("services.profile.port", fmt.Sprintf("%s_PORT", prefix))

	root.SetDefault("services.profile.scheme", "http")
	root.SetDefault("services.profile.host", "localhost")
	root.SetDefault("services.profile.port", "8080")
}

func LoadServiceConfig(viper *viper.Viper, serviceName string) (*ServiceConfig, error) {
	var serviceConfig ServiceConfig

	serviceSettings, ok := viper.AllSettings()["services"].(map[string]interface{})[serviceName]
	if !ok {
		return nil, fmt.Errorf("Service %s not found", serviceName)
	}

	jsonString, err := json.Marshal(serviceSettings)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(jsonString, &serviceConfig); err != nil {
		return nil, err
	}

	return &serviceConfig, nil
}

func (s *ServiceConfig) Address() string {
	return fmt.Sprintf("%s://%s:%s", s.Scheme, s.Host, s.Port)
}
