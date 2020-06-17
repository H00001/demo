package controllers

import (
	"net/http"
	"strconv"
	"test/internal/application"
	"test/internal/infra"
	"test/internal/transports/http/dto"

	"github.com/gin-gonic/gin"
)

type DemoController struct {
	infra *infra.Infra

	UserSvc *application.UserService
}

func (dc *DemoController) Demo(c *gin.Context) {
	username := c.GetString("username")
	info := dc.UserSvc.GetUserInfo(c.Request.Context(), username)

	c.JSON(http.StatusOK, gin.H{
		"data": dto.NewUser(info),
	})
}

type PingController struct {
	infra *infra.Infra
}

func (pc *PingController) Ping(c *gin.Context) {
	errs := pc.infra.HealthCheck()

	if len(errs) > 0 {
		for _, err := range errs {
			infra.NewInfra().Logger.Errorf(err.Error())
		}
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{})
	}
}

type EsimController struct {
	infra *infra.Infra
}

func (ec *EsimController) Esim(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Esim": ec.infra.String(),
	})
}

type ServerNodeController struct {
	infra           *infra.Infra
	registerService *application.ServerRegisteService
}

func (dc *ServerNodeController) Cancel(c *gin.Context) {
	id, _ := strconv.Atoi(c.GetString("server_id"))

	dc.registerService.Cancel(c, id)

	c.JSON(http.StatusOK, gin.H{"status": 0})

}

func (dc *ServerNodeController) Regisite(c *gin.Context) {
	d := dto.ServerNode{}
	if err := c.BindJSON(&d); err != nil {
		c.JSON(http.StatusInternalServerError, dto.CreateFailWithError(400, err.Error()))
		return
	}

	ID := dc.registerService.Regisite(c, dto.NewServerRecode(d))
	c.JSON(http.StatusOK, dto.CreateSucceedWithData(ID))
}
