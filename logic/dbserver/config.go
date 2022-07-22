package dbserver

import "HSLX/pkg/db"

type Configuration struct {
	Mysql db.Mysql
	Redis db.Redis
}

func NewConfiguration() *Configuration {
	return &Configuration{}
}
