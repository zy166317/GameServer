package main

import (
	"HSLX/logic/dbserver"
	"HSLX/pkg/db"
	"HSLX/pkg/redis"
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func main() {
	var conf string
	flag.StringVar(&conf, "config", "config/conf.yml", "config file path")
	flag.Parse()
	file, err := ioutil.ReadFile(conf)
	if err != nil {
		log.Fatalf("read config failed, %s", err)
	}
	configuration := dbserver.NewConfiguration()
	err = yaml.Unmarshal(file, configuration)
	if err != nil {
		log.Fatalf("parse config failed, %s", err)
	}
	//初始化MySQL
	db.InitDB(&configuration.Mysql)
	//初始化Redis
	redis.InitRedis(&configuration.Redis)
}
