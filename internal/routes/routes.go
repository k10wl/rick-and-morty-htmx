package routes

import (
	"rick-and-morty-htmx/internal/handlers"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	r.GET("/", handlers.Home)
	r.GET("/characters", handlers.CharactersPage)
}
