package http

import (
	"context"
	"net/http"
	"strings"
	test "test/internal"
	"test/internal/transports/http/controllers"
	"test/internal/transports/http/routers"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jukylin/esim/log"
	middleware "github.com/jukylin/esim/middle-ware"
)

type GinServer struct {
	en *gin.Engine

	addr string

	logger log.Logger

	server *http.Server

	app *test.App
}

func NewGinServer(app *test.App) *GinServer {
	httpport := app.Conf.GetString("httpport")

	in := strings.Index(httpport, ":")
	if in < 0 {
		httpport = ":" + httpport
	}

	if app.Conf.GetString("runmode") != "pro" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	en := gin.Default()

	if app.Conf.GetBool("http_trace") {
		en.Use(middleware.GinTracer(app.Tracer))
	}

	if app.Conf.GetBool("http_metrics") {
		en.Use(middleware.GinMonitor())
	}

	en.Use(middleware.GinTracerID())

	server := &GinServer{
		en:     en,
		addr:   httpport,
		logger: app.Logger,
		app:    app,
	}

	return server
}

func (gs *GinServer) Start() {
	routers.RegisterGinServer(gs.en, controllers.NewControllers(gs.app))

	server := &http.Server{Addr: gs.addr, Handler: gs.en}
	gs.server = server
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				gs.logger.Fatalf("start http server err %s", err.Error())
			}
			return
		}
	}()
}

func (gs *GinServer) GracefulShutDown() {
	ctx, cannel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cannel()
	if err := gs.server.Shutdown(ctx); err != nil {
		gs.logger.Errorf("stop http server error %s", err.Error())
	}
}
