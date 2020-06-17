package component_test

import (
	"context"
	"os"
	test "test/internal"
	"test/internal/infra"
	"test/internal/transports/http"
	"testing"

	"github.com/jukylin/esim/container"
	"github.com/jukylin/esim/grpc"
	_grpc "google.golang.org/grpc"
)

func TestMain(m *testing.M) {
	appOptions := test.AppOptions{}
	app := test.NewApp(appOptions.WithConfPath("../../../../conf/"))

	setUp(app)

	code := m.Run()

	tearDown(app)

	os.Exit(code)
}

func provideStubsGrpcClient(esim *container.Esim) *grpc.Client {
	clientOptional := grpc.ClientOptionals{}
	clientOptions := grpc.NewClientOptions(
		clientOptional.WithLogger(esim.Logger),
		clientOptional.WithConf(esim.Conf),
		clientOptional.WithDialOptions(_grpc.WithUnaryInterceptor(
			grpc.ClientStubs(func(ctx context.Context, method string, req,
				reply interface{}, cc *_grpc.ClientConn, invoker _grpc.UnaryInvoker,
				opts ..._grpc.CallOption) error {
				esim.Logger.Infof(method)
				err := invoker(ctx, method, req, reply, cc, opts...)
				return err
			}),
		)),
	)

	grpcClient := grpc.NewClient(clientOptions)

	return grpcClient
}

func setUp(app *test.App) {
	app.Infra = infra.NewStubsInfra(provideStubsGrpcClient(app.Esim))

	app.RegisterTran(http.NewGinServer(app))

	app.Start()

	errs := app.Infra.HealthCheck()
	if len(errs) > 0 {
		for _, err := range errs {
			app.Logger.Errorf(err.Error())
		}
	}
}

func tearDown(app *test.App) {
	app.Infra.Close()
}
