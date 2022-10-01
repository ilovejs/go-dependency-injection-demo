package service

import "github.com/google/wire"

// ServiceSet service注入
var ServiceSet = wire.NewSet(
	ProjectSet,
	// other service set...
)
