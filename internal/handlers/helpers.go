package handlers

import (
	"strconv"

	"rick-and-morty-htmx/internal/pkg/rickandmortyapi"

	"github.com/gin-gonic/gin"
)

func getPages(c *gin.Context, i rickandmortyapi.Info) map[string]int {
	Pages := map[string]int{
		"Prev": 0,
		"Next": 0,
	}
	p, exists := c.GetQuery("page")
	if exists {
		pi, err := strconv.Atoi(p)
		if err == nil {
			Pages["Prev"] = pi - 1
			if i.Next != "" {
				Pages["Next"] = pi + 1
			}
		}
	} else {
		Pages["Prev"] = 0
		if i.Next != "" {
			Pages["Next"] = 2
		}
	}
	return Pages
}
