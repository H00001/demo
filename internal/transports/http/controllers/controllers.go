package controllers

import (
	test "test/internal"
	"test/internal/application"

	"github.com/google/wire"
)

type Controllers struct {
	App *test.App

	Ping *PingController

	Esim *EsimController

	Demo *DemoController

	ServerNode *ServerNodeController
}

//nolint:deadcode,varcheck,unused
var controllersSet = wire.NewSet(
	wire.Struct(new(Controllers), "*"),
	providePingController,
	provideEsimController,
	provideDemoController,
	provideServerController,
)

func NewControllers(app *test.App) *Controllers {
	controllers := initControllers(app)
	return controllers
}

func providePingController(app *test.App) *PingController {
	pingController := &PingController{}
	pingController.infra = app.Infra
	return pingController
}

func provideEsimController(app *test.App) *EsimController {
	esimController := &EsimController{}
	esimController.infra = app.Infra
	return esimController
}

func provideDemoController(app *test.App) *DemoController {
	userSvc := application.NewUserSvc(app.Infra)

	demoController := &DemoController{}
	demoController.infra = app.Infra
	demoController.UserSvc = userSvc

	return demoController
}

func provideServerController(app *test.App) *ServerNodeController {
	serverSvc := application.NewServerRegisteService(app.Infra)

	c := &ServerNodeController{}
	c.infra = app.Infra
	c.registerService = serverSvc

	return c
}
