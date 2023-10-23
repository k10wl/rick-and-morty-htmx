package handlers

import (
	"fmt"
	"net/http"

	"rick-and-morty-htmx/internal/pkg/rickandmortyapi"

	"github.com/gin-gonic/gin"
)

func EpisodesPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "episodes", map[string]interface{}{
		"Title": "Episodes",
	})
}

func EpisodesList(ctx *gin.Context) {
	q := ctx.Request.URL.Query()
	e, err := rickandmortyapi.GetEpisodes(q.Encode())
	if err != nil {
		ctx.HTML(http.StatusOK, "episodes-list", map[string]interface{}{
			"Error": err,
		})
		return
	}
	fmt.Printf("e: %v\n", e.Results)

	pages := getPages(ctx, e.Info)
	ctx.HTML(http.StatusOK, "episodes-list", map[string]interface{}{
		"Episodes": e.Results,
		"Pages":    pages,
	})
}
