package app

import "DIdemo/handler"

func NewInjector(handler *handler.ProjectHandler) *Injector {
	return &Injector{
		ProjectHandler: handler,
	}
}

type Injector struct {
	ProjectHandler *handler.ProjectHandler
	// components，others...
}
