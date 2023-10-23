package handlers

import (
	"net/http"

	"rick-and-morty-htmx/internal/pkg/rickandmortyapi"

	"github.com/gin-gonic/gin"
)

func LocationsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "locations-page", map[string]interface{}{
		"Title": "Locations",
	})
}

func LocationList(c *gin.Context) {
	q := c.Request.URL.Query()
	locations, err := rickandmortyapi.GetLocations(q.Encode())
	if err != nil {
		c.HTML(http.StatusOK, "locations-data", map[string]interface{}{
			"Locations": []string{},
		})
		return
	}

	pages := getPages(c, locations.Info)

	c.HTML(http.StatusOK, "locations-data", map[string]interface{}{
		"Locations": locations.Results,
		"Pages":     pages,
	})
}
