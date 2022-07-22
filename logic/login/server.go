package login

import (
	"HSLX/pkg/db"
	"HSLX/pkg/redis"
	"log"
)

type Login struct {
	cfg *Configuration
}

func NewServer(cfg *Configuration) *Login {
	return &Login{cfg: cfg}
}

func (l *Login) init() {
	db.InitDB(l.cfg.Mysql)
	redis.InitRedis(l.cfg.Redis)
}

func (l *Login) Run() error {
	log.Fatalf("login is running")
	l.init()

	return nil
}
