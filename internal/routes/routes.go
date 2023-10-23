package routes

import (
	"rick-and-morty-htmx/internal/handlers"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	r.GET("/", handlers.Home)
	r.GET("/characters", handlers.CharactersPage)
	r.GET("/locations", handlers.LocationsPage)
	r.GET("/locations/list", handlers.LocationList)
	r.GET("/episodes", handlers.EpisodesPage)
	r.GET("/episodes/list", handlers.EpisodesList)
}
