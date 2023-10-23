package handlers

import (
	"net/http"
	"strconv"

	"rick-and-morty-htmx/internal/pkg/rickandmortyapi"

	"github.com/gin-gonic/gin"
)

func CharactersPage(ctx *gin.Context) {
	q := ctx.Request.URL.Query()

	c, _ := rickandmortyapi.GetCharacters(q.Encode())

	Pages := map[string]int{
		"Prev": 0,
		"Next": 0,
	}
	p, exists := ctx.GetQuery("page")
	if exists {
		pi, err := strconv.Atoi(p)
		if err == nil {
			Pages["Prev"] = pi - 1
			if c.Info.Next != "" {
				Pages["Next"] = pi + 1
			}
		}
	} else {
		Pages["Prev"] = 0
		if c.Info.Next != "" {
			Pages["Next"] = 2
		}
	}

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
		"Pages":      Pages,
		"Params":     q,
	})
}
