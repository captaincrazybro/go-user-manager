package routers

import (
	"net/http"
	"io/ioutil"
	//"strings"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User		string `form:"user" json:"user" binding:"required"`
	Password	string `form:"password" json:"password" binding:"required"`
}

func HandleAPI(r *gin.RouterGroup) {
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong!")
	})

	r.GET("", func (ctx *gin.Context) {
		contents := getData()
		ctx.String(http.StatusOK, contents)
	})
}

func getData() (parsed string) {
	data, _ := ioutil.ReadFile(FILENAME)
	parsed = string(data)
	return parsed
}