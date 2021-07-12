package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/contrib/static"

	"github.com/captaincrazybro/go-user-manager/server/routers"

	"github.com/gin-gonic/gin"
)

const (
	port int = 8080
)

func main() {
	// check if "users.txt" exists, otherwise create it
	_, err := os.Stat(routers.UsersFile)
	if os.IsNotExist(err) {
		os.Create(routers.UsersFile)
	}

	// creates webserver
	r := gin.Default()

	r.Static("/public", "public")
	r.StaticFile("/favicon.ico", "resources/favicon.ico")
	r.Use(static.Serve("/", static.LocalFile("views", true)))

	// creates routers for "/" and "/api"
	website := r.Group("/")
	api := r.Group("/api/")

	routers.HandleWebsite(website)
	routers.HandleAPI(api)

	r.Run(fmt.Sprintf(":%d", port))
}
