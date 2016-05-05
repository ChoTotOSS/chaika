package config

type Config struct {
	ConsulHost  string
	ConsulPort  int64
	Port        int64
	GraylogHost string
	GraylogPort int64
}

var config *Config

func InitConfig(cfg *Config) {
	config = cfg
}

func GetConfig() Config {
	return *config
}
