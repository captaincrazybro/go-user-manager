package main

import (
	"fmt"

	c "github.com/captaincrazybro/literalutil/console"
	"github.com/gin-gonic/gin"
)

const (
	port int = 80
)

func main() {
	c.Plnf("Webserver started on port %d!", port)

	r := gin.Default()

	r.Run(fmt.Sprintf(":%d", port))
}
