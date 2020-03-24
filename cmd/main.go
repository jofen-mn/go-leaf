package main

import (
	"github.com/gin-gonic/gin"
	"go-leaf/base/db"
	"go-leaf/conf"
	"go-leaf/ids"
	"log"
)

func main() {
	log.Println("server is running...")
	conf.ConfInit()
	log.Printf("%+v", conf.Configure)
	db.InitMysql()

	server := gin.New()
	server.Use(gin.Recovery())

	baseGroup := server.Group("/idsproducer/v1")
	ids.NewIds(baseGroup)

	server.Run("0.0.0.0:" + conf.Configure.Application.Port)
}
