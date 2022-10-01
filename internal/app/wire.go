//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"gorm.io/gorm"

	"DIdemo/dal"
	"DIdemo/handler"
	"DIdemo/service"
)

func BuildInjector(db *gorm.DB) (*Injector, error) {
	wire.Build(
		NewInjector,

		// handler
		handler.NewProjectHandler,

		// services
		service.NewProjectService,
		// 更多service...

		//dal
		dal.NewProjectDal,
		dal.NewQuestionDal,
		dal.NewQuestionModelDal,
		// 更多dal...

		// db
		//common.InitGormDB,
		// other components...
	)

	return new(Injector), nil
}
