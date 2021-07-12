package routers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Users     []string `form:"users" json:"users" binding:"required"`
	Passwords []string `form:"passwords" json:"passwords" binding:"required"`
}

func HandleAPI(r *gin.RouterGroup) {
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong!")
	})

	r.GET("", func(ctx *gin.Context) {
		contents, _ := ioutil.ReadFile(FILENAME)
		ctx.String(http.StatusOK, string(contents))
	})

	r.POST("", func(ctx *gin.Context) {
		var login Login
		ctx.ShouldBindJSON(&login)
		fmt.Println(login)
		b, _ := json.Marshal(login)
		os.WriteFile(FILENAME, []byte(b), 0666)
	})
}
