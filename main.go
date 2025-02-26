package main

import (
	"ip_getter/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.GET("/ipv4", utils.GetIPv4)
	r.GET("/ipv6", utils.GetIPv6)
	r.Run(":1234")
}
