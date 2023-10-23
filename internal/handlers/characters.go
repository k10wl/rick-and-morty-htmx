package handlers

import (
	"net/http"

	"rick-and-morty-htmx/internal/pkg/rickandmortyapi"

	"github.com/gin-gonic/gin"
)

func CharactersPage(ctx *gin.Context) {
	q := ctx.Request.URL.Query()

	c, _ := rickandmortyapi.GetCharacters(q.Encode())

	pages := getPages(ctx, c.Info)

	if ok := q.Has("name"); !ok {
		q.Set("name", "")
	}
	if ok := q.Has("species"); !ok {
		q.Set("species", "")
	}
	if ok := q.Has("gender"); !ok {
		q.Set("gender", "")
	}
	if ok := q.Has("status"); !ok {
		q.Set("status", "")
	}

	ctx.HTML(http.StatusOK, "characters", map[string]interface{}{
		"Title":      "Characters",
		"Characters": c.Results,
		"Info":       c.Info,
		"Pages":      pages,
		"Params":     q,
	})
}
