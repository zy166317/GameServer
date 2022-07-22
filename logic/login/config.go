package login

import "HSLX/pkg/db"

type Configuration struct {
	ListenAddr string    `yaml:"listenAddr"`
	SignKey    string    `yaml:"signKey"`
	Mysql      *db.Mysql `yaml:"mysql"`
	Redis      *db.Redis `yaml:"redis"`
}

func NewConfiguration() *Configuration {
	return &Configuration{}
}
