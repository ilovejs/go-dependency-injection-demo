package app

import (
	"github.com/google/wire"

	"DIdemo/handler"
)

//func NewInjector(handler *handler.ProjectHandler) *Injector {
//	return &Injector{
//		ProjectHandler: handler,
//	}
//}

var InjectorSet = wire.NewSet(
	wire.Struct(new(Injector), "*"),
)

type Injector struct {
	ProjectHandler *handler.ProjectHandler
}
