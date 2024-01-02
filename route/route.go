package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func NewRoute(e *gin.Engine) {
	e.LoadHTMLGlob("templates/*")
	e.Static("/static", "./static")
	e.NoRoute(Index)
}

func Index(c *gin.Context) {
	htmlBytes, _ := os.ReadFile("./templates/index.html")
	c.Data(http.StatusOK, "text/html; charset=utf-8", htmlBytes)
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, authorization, token")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
