package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home", map[string]interface{}{
		"Title": "Tech explorer",
	})
}
