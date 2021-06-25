package router

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	Root *gin.Engine
	V1 *gin.RouterGroup
}

var BaseRouter *Router

func InitRouter(router *gin.Engine){
	BaseRouter = &Router{Root: router}
	BaseRouter.V1Group()
}
