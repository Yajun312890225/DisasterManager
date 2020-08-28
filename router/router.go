package router

import (
	"DisasterManager/handler"
	"DisasterManager/pkg/ws"
	"os"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter 配置路由
func InitRouter() *gin.Engine {
	gin.SetMode("debug")
	r := gin.Default()
	m := ws.GetMelody()
	// 配置swagger
	swagURL := ginSwagger.URL(os.Getenv("SWAGGER_URL"))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swagURL))

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})
	m.HandleConnect(handler.Connect)
	m.HandleDisconnect(handler.DisConnect)
	m.HandleMessage(handler.Message)

	r.POST("/register", handler.Register)

	return r
}
