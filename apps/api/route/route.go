package route

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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
