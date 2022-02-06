package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"visits/databases/redis"
)

func Handler(c *gin.Context) {
	str, ge := redis.Get("visits")
	if ge != nil {
		c.String(http.StatusBadRequest, ge.Error())
		c.Abort()
	}
	c.String(http.StatusOK, str)
	i, _ := strconv.Atoi(str)
	se := redis.Set("visits", strconv.Itoa(i+1))
	if se != nil {
		c.String(http.StatusBadRequest, se.Error())
	}
	c.Abort()
}

func StopServer(c *gin.Context) {
	os.Exit(0)
}
