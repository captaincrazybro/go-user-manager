package routers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"log"

	"github.com/gin-gonic/gin"
)

type LoginBody struct {
	Username string ` json:"username" binding:"required"`
	Password string ` json:"password" binding:"required"`
}

type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Permission int    `json:"permission"`
}

type Users []User

func HandleAPI(r *gin.RouterGroup) {

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong!")
	})

	r.GET("", func(ctx *gin.Context) {
		users, err := GetUsers()
		if err != nil {log.Println(err); return}
		ctx.JSON(http.StatusOK, users)
	})

	r.POST("", func(ctx *gin.Context) {
		var users Users
		ctx.ShouldBindJSON(&users)
		fmt.Println(users)
		b, _ := json.Marshal(users)
		os.WriteFile(FILENAME, []byte(b), 0666)
	})

	r.POST("/login", func(ctx *gin.Context) {
		var body LoginBody
		err := ctx.ShouldBind(&body)
		if err != nil {
			ctx.JSON(http.StatusOK, MakeErrRes(err))
			return
		}

		users, err := GetUsers()
		if err != nil {
			ctx.JSON(http.StatusOK, MakeErrRes(err))
			return
		}

		user := users.GetUser(body.Username)
		if user == nil {
			ctx.JSON(http.StatusOK, MakeErrRes("Username does not exist!"))
			return
		}

		if user.Password != body.Password {
			ctx.JSON(http.StatusOK, MakeErrRes("Incorrect password!"))
			return
		}

		// handle session
		ctx.JSON(http.StatusOK, users)
	})
}

func MakeErrRes(v interface{}) gin.H {
	return gin.H{
		"type":    "error",
		"message": fmt.Sprintf("%s", v),
	}
}

func GetUsers() (Users, error) {
	file, err := os.Open(FILENAME)
	if os.IsNotExist(err) {
		file, err = os.Create(FILENAME)
		if err != nil {
			return nil, err
		}

		_, err := file.Write([]byte("[]"))
		if err != nil {
			return nil, err
		}
		return []User{}, nil
	}

	bts, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var users Users
	err = json.Unmarshal(bts, &users)

	return users, err
}

func (us Users) GetUser(username string) *User {
	for _, v := range us {
		if v.Username == username {
			return &v
		}
	}

	return nil
}
