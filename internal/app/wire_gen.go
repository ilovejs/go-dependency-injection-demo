// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"DIdemo/dal"
	"DIdemo/handler"
	"DIdemo/service"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func BuildInjector(db *gorm.DB) (*Injector, error) {
	projectDal := dal.NewProjectDal(db)
	questionDal := dal.NewQuestionDal(db)
	questionModelDal := dal.NewQuestionModelDal(db)
	projectService := service.NewProjectService(projectDal, questionDal, questionModelDal)
	projectHandler := handler.NewProjectHandler(projectService)
	injector := NewInjector(projectHandler)
	return injector, nil
}
