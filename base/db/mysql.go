package db

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
	"go-leaf/conf"
	"log"
)

var Mysql *gorm.DB

func InitMysql() {
	var err error
	Mysql, err = gorm.Open("mysql", conf.Configure.Mysql.Url)
	if err != nil {
		log.Panicf("Mysql init error: %+v", err)
		return
	}

	Mysql.DB().SetMaxIdleConns(conf.Configure.Mysql.MaxIdleConns)
	Mysql.DB().SetMaxOpenConns(conf.Configure.Mysql.MaxOpenConns)

	log.Println("mysql init success...")

	return
}

func Close() error {
	return Mysql.Close()
}
