package routers

import (
<<<<<<< HEAD
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
=======
	"io/ioutil"
	"net/http"

	//"strings"
>>>>>>> 9929aac8263374ffdc8ec0691ce1e109b7f26827

	"github.com/gin-gonic/gin"
)

type Login struct {
<<<<<<< HEAD
	Users[]		string `form:"users" json:"users" binding:"required"`
	Passwords[]	string `form:"passwords" json:"passwords" binding:"required"`
=======
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
>>>>>>> 9929aac8263374ffdc8ec0691ce1e109b7f26827
}

func HandleAPI(r *gin.RouterGroup) {
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong!")
	})

<<<<<<< HEAD
	r.GET("", func (ctx *gin.Context) {
		contents, _ := ioutil.ReadFile(FILENAME)
		ctx.String(http.StatusOK, string(contents))
=======
	r.GET("", func(ctx *gin.Context) {
		contents := getData()
		ctx.String(http.StatusOK, contents)
>>>>>>> 9929aac8263374ffdc8ec0691ce1e109b7f26827
	})

<<<<<<< HEAD
	r.POST("", func (ctx *gin.Context) {
		var login Login
		ctx.ShouldBindJSON(&login)
		fmt.Println(login)
		b, _ := json.Marshal(login)
		os.WriteFile(FILENAME, []byte(b), 0666)
	})
}
=======
func getData() (parsed string) {
	data, _ := ioutil.ReadFile(FILENAME)
	parsed = string(data)
	return parsed
}
>>>>>>> 9929aac8263374ffdc8ec0691ce1e109b7f26827
