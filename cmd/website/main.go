package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("internal/views/**/*.html")
	r.Static("/assets", "./assets")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home", map[string]interface{}{
			"Title": "Tech explorer",
		})
	})

	r.Run(":8080")
}
