package dal

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"DIdemo/entity"
)

func NewProjectDal(db *gorm.DB) *ProjectDal {
	return &ProjectDal{
		DB: db,
	}
}

type ProjectDal struct {
	DB *gorm.DB
}

func (dal *ProjectDal) Create(ctx context.Context, item *entity.Project) error {
	result := dal.DB.Create(item)
	return errors.WithStack(result.Error)
}
