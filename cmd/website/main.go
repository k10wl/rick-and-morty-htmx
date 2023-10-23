package main

import (
	"rick-and-morty-htmx/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("internal/views/**/*.html")
	r.Static("/assets", "./assets")
	routes.ApplyRoutes(r)

	r.Run(":8080")
}
