package utils

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIPv4(c *gin.Context) {
	resp, err := http.Get("https://4.ipw.cn")
	if err != nil {
		c.JSON(400, gin.H{
			"ok":      false,
			"content": err,
		})
		return
	}
	body, _ := io.ReadAll(resp.Body)
	c.JSON(200, gin.H{
		"ok":      true,
		"content": string(body),
	})
}

func GetIPv6(c *gin.Context) {
	resp, err := http.Get("https://6.ipw.cn")
	if err != nil {
		c.JSON(400, gin.H{
			"ok":      false,
			"content": err,
		})
		return
	}
	body, _ := io.ReadAll(resp.Body)
	c.JSON(200, gin.H{
		"ok":      true,
		"content": string(body),
	})
}
