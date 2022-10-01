//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"gorm.io/gorm"

	"DIdemo/dal"
	"DIdemo/handler"
	"DIdemo/internal/service"
)

func BuildInjector(db *gorm.DB) (*Injector, error) {
	wire.Build(
		//NewInjector,
		InjectorSet,

		// handler
		handler.ProjectHandlerSet,

		// services
		service.ServiceSet,

		//dal
		dal.DalSet,

		// db
		//common.InitGormDB,
	)

	return new(Injector), nil
}
