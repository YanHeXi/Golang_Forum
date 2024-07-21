package go_demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func func1(ctx *gin.Context) {}
func func4(ctx *gin.Context) {}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	shopGroup := r.Group("/shop")
	shopGroup.Use(func1)
	{
		shopGroup.GET("/index", func4)
	}

	err := r.Run()
	if err != nil {
		return
	}
}
