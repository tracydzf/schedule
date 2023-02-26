//go:build wireinject
// +build wireinject

package bootstrap

import (
	"github.com/google/wire"
	"schedule/common"
	"schedule/schedule/app"
)

func App(value *common.Config) (*app.App, error) {
	wire.Build(
		wire.Struct(new(common.Inject), "*"),
		Provides,
		app.Provides,
	)
	return &app.App{}, nil
}
