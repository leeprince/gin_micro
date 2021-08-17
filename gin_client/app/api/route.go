package api

import (
	v1 "gin_client/app/api/v1"
	"gin_client/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	routes.RegisterRouter(Routes)
}

func Routes(g *gin.Engine) {
	// 健康检查
	g.GET("/ping", Ping)
	
	hello := g.Group("v1")
	{
		hello.GET("/getUser", v1.GetUser)
	}
}