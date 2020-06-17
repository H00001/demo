package routers

import (
	"test/internal/transports/http/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterGinServer(en *gin.Engine, ctl *controllers.Controllers) {
	en.GET("/", ctl.Esim.Esim)

	en.GET("/demo", ctl.Demo.Demo)

	en.GET("/ping", ctl.Ping.Ping)

	en.POST("/services/node", ctl.ServerNode.Regisite)
	en.DELETE("/services/node", ctl.ServerNode.Cancel)

}
