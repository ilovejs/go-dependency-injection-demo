package dal

import "github.com/google/wire"

// DalSet dal注入, used by BuildInjector()

var DalSet = wire.NewSet(
	ProjectSet,
	QuestionDalSet,
	QuestionModelDalSet,
)
