package main

import (
	"fmt"

	"github.com/captaincrazybro/go-user-manager/server/routers"

	c "github.com/captaincrazybro/literalutil/console"

	"github.com/gin-gonic/gin"
)

const (
	port int = 8080
)

func main() {
	// creates webserver
	r := gin.Default()

	r.Static("/public", "../public")
	r.StaticFile("/favicon.ico", "../views/favicon.ico")

	// creates routers for "/" and "/api"
	website := r.Group("/")
	api := r.Group("/api")

	routers.HandleWebsite(website)
	routers.HandleAPI(api)

	r.GET("/favicon.ico")

	r.Run(fmt.Sprintf(":%d", port))
	c.Plnf("Webserver started on port %d!", port)
}
