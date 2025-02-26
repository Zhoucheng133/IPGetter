package main

import (
	"ip_getter/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/api/ipv4", utils.GetIPv4)
	r.GET("/api/ipv6", utils.GetIPv6)
	r.Run(":1234")
}
