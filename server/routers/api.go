package routers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type LoginBody struct {
	Username string ` json:"username" binding:"required"`
	Password string ` json:"password" binding:"required"`
}

type SessionBody struct {
	SessionID int    `json:"sessionId"`
	Username  string `json:"username"`
}

type LogoutBody struct {
	SessionID int `json:"sessionId"`
}

type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Permission int    `json:"permission"`
}

type Users []User

type Session struct {
	Id   int    `json:"id"`
	User string `json:"username"`
}

type Sessions []Session

func HandleAPI(r *gin.RouterGroup) {

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong!")
	})

	r.GET("", func(ctx *gin.Context) {
		users, err := GetUsers()
		if err != nil {
			log.Println(err)
			return
		}
		ctx.JSON(http.StatusOK, users)
	})

	r.POST("", func(ctx *gin.Context) {
		var users Users
		ctx.ShouldBindJSON(&users)
		fmt.Println(users)
		b, _ := json.Marshal(users)
		os.WriteFile(UsersFile, []byte(b), 0666)
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

		// create session
		sessions, err := GetSessions()
		if err != nil {
			ctx.JSON(http.StatusOK, MakeErrRes(err))
			return
		}

		id, err := sessions.createSession(user.Username)
		if err != nil {
			ctx.JSON(http.StatusOK, MakeErrRes(err))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"type":    "success",
			"message": "Successfully logged in!",
			"json": gin.H{
				"sessionId": id,
			},
		})
	})

	r.POST("/logout", func(ctx *gin.Context) {
		var body LogoutBody
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			ctx.JSON(http.StatusOK, MakeErrRes(err))
		}

		sessions, err := GetSessions()
		if err != nil {
			ctx.JSON(http.StatusOK, MakeErrRes(err))
		}

		err = sessions.deleteSession(body.SessionID)
		if err != nil {
			ctx.JSON(http.StatusOK, MakeErrRes(err))
		}

		ctx.JSON(http.StatusOK, gin.H{
			"type":    "success",
			"message": "Successfully logged out user!",
		})
	})

	r.POST("/validate-session", func(ctx *gin.Context) {
		var body SessionBody
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			ctx.JSON(http.StatusOK, MakeErrRes(err))
			return
		}

		sessions, err := GetSessions()
		if err != nil {
			ctx.JSON(http.StatusOK, MakeErrRes(err))
			return
		}

		ses := sessions.GetSessionFromId(body.SessionID)
		if ses == nil {
			ctx.JSON(http.StatusOK, gin.H{"type": "success", "message": "Session not found!", "validated": false})
			return
		}

		if ses.User != body.Username {
			ctx.JSON(http.StatusOK, gin.H{"type": "success", "message": "Username doesn't match session user!", "validated": false})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"type": "success", "message": "Successfully validated session!", "validated": true})
	})
}

func MakeErrRes(v interface{}) gin.H {
	return gin.H{
		"type":    "error",
		"message": fmt.Sprintf("%s", v),
	}
}

func GetSessions() (Sessions, error) {
	bts, err := ioutil.ReadFile(SessionsFile)
	if os.IsNotExist(err) {
		file, err := os.Create(SessionsFile)
		defer file.Close()
		if err != nil {
			return nil, err
		}
		_, err = file.Write([]byte("[]"))
		if err != nil {
			return nil, err
		}
		return Sessions{}, nil
	}

	if err != nil {
		return nil, err
	}

	var sessions Sessions
	err = json.Unmarshal(bts, &sessions)

	return sessions, err
}

func (sz Sessions) GetSessionFromId(sessionId int) *Session {
	for _, v := range sz {
		if v.Id == sessionId {
			return &v
		}
	}

	return nil
}

func (sz Sessions) GetSessionFromUser(username string) *Session {
	for _, v := range sz {
		if v.User == username {
			return &v
		}
	}

	return nil
}

func (sz Sessions) createSession(username string) (int, error) {
	id := sz.generateSessionId(username)

	newSession := Session{
		Id:   id,
		User: username,
	}

	sz = append(sz, newSession)

	jsonString, err := json.MarshalIndent(sz, "", "    ")
	if err != nil {
		return 0, err
	}

	err = ioutil.WriteFile(SessionsFile, jsonString, 0644)
	if err != nil {
		return 0, err
	}

	return id, err
}

func (sz Sessions) generateSessionId(username string) int {
	// generates the id (an int with 6 digits, ex. 123876)
	id := rand.Intn(999999-100000) - 100000

	if sz.GetSessionFromId(id) != nil {
		return sz.generateSessionId(username)
	} else {
		return id
	}

}

func (sz Sessions) DeleteAll() error {
	jsonString, err := json.MarshalIndent(Sessions{}, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(SessionsFile, jsonString, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (sz Sessions) deleteSession(sessionId int) error {
	if sz.GetSessionFromId(sessionId) == nil {
		return nil
	}

	for i, v := range sz {
		if v.Id == sessionId {
			sz = RemoveIndex(sz, i)
		}
	}

	jsonString, err := json.MarshalIndent(sz, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(SessionsFile, jsonString, 0644)
	if err != nil {
		return err
	}

	return nil
}

func RemoveIndex(s Sessions, index int) Sessions {
	ret := make(Sessions, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func GetUsers() (Users, error) {
	file, err := os.Open(UsersFile)
	if os.IsNotExist(err) {
		file, err = os.Create(UsersFile)
		defer file.Close()
		if err != nil {
			return nil, err
		}

		_, err = file.Write([]byte("[]"))
		if err != nil {
			return nil, err
		}
		return Users{}, nil
	}
	defer file.Close()

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
