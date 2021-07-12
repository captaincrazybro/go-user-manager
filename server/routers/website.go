package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func HandleWebsite(r *gin.RouterGroup) {
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong!")
	})
}
