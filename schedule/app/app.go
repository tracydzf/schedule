package app

import (
	"github.com/google/wire"
	"schedule/common"
)

var Provides = wire.NewSet(New)

func New(i *common.Inject) *App {
	return &App{
		Inject: i,
	}
}

type App struct {
	*common.Inject
}
