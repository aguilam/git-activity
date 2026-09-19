package config

type Config struct {
	Services []ServiceConfig `yaml:"services"`
}

type ServiceConfig struct {
	Username string  `yaml:"username"`
	Token    *string `yaml:"token"`
	Url      *string `yaml:"url"`
	Service  string  `yaml:"service"`
}