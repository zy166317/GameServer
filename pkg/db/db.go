package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var gormDB *gorm.DB

// 拼接地址
func getConnStr(user string, pass string, protocol string, addr string, db string) string {
	return fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", user, pass, protocol, addr, db)
}

// InitDB 初始化mysql
func InitDB(my *Mysql) {
	connStr := getConnStr(my.User, my.Password, my.Protocol, my.Address, my.DBName)
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		panic(err)
	}
	myDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	err = myDB.Ping()
	if err != nil {
		panic(err)
	}
	db.Debug()
	myDB.SetConnMaxLifetime(time.Duration(my.ConnMaxLifeTime) * time.Second)
	myDB.SetMaxOpenConns(my.MaxOpenConnNum)
	myDB.SetMaxIdleConns(my.MaxIdleConnNum)
}
