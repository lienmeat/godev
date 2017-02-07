package main

import (
	"companyname.com/exampleapp/controller"
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	server := gin.Default()
	controller.SetRoutes(server)
	server.Run(fmt.Sprintf("0.0.0.0:%s", "8080"))
}