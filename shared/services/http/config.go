package http

type ServiceConfig struct {
	Schema string
	Host   string
	Port   string
}

func NewServiceConfig(schema, host, port string) *ServiceConfig {
	return &ServiceConfig{
		Schema: schema,
		Host:   host,
		Port:   port,
	}
}

func (s *ServiceConfig) Address() string {
	return s.Schema + "://" + s.Host + ":" + s.Port
}
