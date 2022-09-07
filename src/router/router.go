package router

import "github.com/gin-gonic/gin"

type router interface {
	HandleMessage(ctx *gin.Context)
}
