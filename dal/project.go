package dal

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"DIdemo/entity"
)

type IProjectDal interface {
	Create(ctx context.Context, item *entity.Project) (err error)
}

type ProjectDal struct {
	DB *gorm.DB
}

// wire.Struct方法是wire提供的构造器，"*"代表为所有字段注入值，在这里可以用"DB"代替
// wire.Bind方法把接口和实现绑定起来

var ProjectSet = wire.NewSet(
	wire.Struct(new(ProjectDal), "*"),
	wire.Bind(new(IProjectDal), new(*ProjectDal)),
)

func (dal *ProjectDal) Create(ctx context.Context, item *entity.Project) error {
	result := dal.DB.Create(item)
	return errors.WithStack(result.Error)
}
