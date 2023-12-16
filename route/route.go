package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRoute(e *gin.Engine) {
	e.LoadHTMLGlob("templates/*")
	e.GET("/test", Index)
}

type test struct {
	Name string
	Age  int
}

func Index(c *gin.Context) {
	var list []*test

	list = append(list, &test{
		Name: "test1",
		Age:  1,
	})
	list = append(list, &test{
		Name: "test2",
		Age:  2,
	})
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "CRMS",
		"msg":   &list[0].Name,
	})
}
