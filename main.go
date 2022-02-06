package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"visits/databases/redis"
	"visits/handlers/home"
)

func main() {
	redis.Connect()
	err := redis.Set("visits", "0")
	if err != nil {
		log.Fatal(err.Error())
	}
	r := gin.Default()
	r.GET("/", home.Handler)
	r.GET("/stop", home.StopServer)
	log.Fatal(r.Run("0.0.0.0:8081"))
}
