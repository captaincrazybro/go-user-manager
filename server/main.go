package main

import (
	"fmt"

	"github.com/gin-gonic/contrib/static"

	"github.com/captaincrazybro/go-user-manager/server/routers"

	"github.com/gin-gonic/gin"
)

const (
	port int = 8080
)

func main() {
	// creates webserver
	r := gin.Default()

	r.Static("/public", "public")
	r.StaticFile("/favicon.ico", "resources/favicon.ico")
	r.Use(static.Serve("/", static.LocalFile("views", true)))

	// creates routers for "/" and "/api"
	website := r.Group("/")
	api := r.Group("/api")

	routers.HandleWebsite(website)
	routers.HandleAPI(api)

	r.Run(fmt.Sprintf(":%d", port))
}
