package controllers

import (
	"github.com/google/wire"
	test "test/internal"
	"test/internal/application"
)

type Controllers struct {
	App *test.App

	ServerNode *TcpNodeController
}

var controllersSet = wire.NewSet(
	wire.Struct(new(Controllers), "*"),
	providePingController,
)

func NewControllers(app *test.App) *Controllers {
	controllers := initControllers(app)
	return controllers
}

func providePingController(app *test.App) *TcpNodeController {
	serverSvc := application.NewServerRegisteService(app.Infra)

	c := &TcpNodeController{}
	c.infra = app.Infra
	c.registerService = serverSvc
	return c
}
