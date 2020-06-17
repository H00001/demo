//+build wireinject

package controllers

import (
	test "test/internal"

	"github.com/google/wire"
)

func initControllers(app *test.App) *Controllers {
	wire.Build(controllersSet)
	return nil
}
