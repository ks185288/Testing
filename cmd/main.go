package main

import (
	"Testing/src/router"

	"github.com/gin-gonic/gin"
)

func main() {
	handler := router.Handler{}
	server := gin.Default()
	server.POST("/message", handler.HandleMessage)
	server.Run(":4000")

}
