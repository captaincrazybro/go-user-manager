package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User		string `form:"user" json:"user"`
	Password	string `form:"password" json:"password"`
}

func HandleWebsite(r *gin.RouterGroup) {
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong!")
	})

	r.POST("/api/login", func (ctx *gin.Context) {
		var json Login
		err := ctx.ShouldBindJSON(&json)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.String(http.StatusOK, json.User + "\n" + json.Password)
		}
	})
}
