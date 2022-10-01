package dal

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"DIdemo/entity"
)

func NewQuestionDal(db *gorm.DB) *QuestionDal {
	return &QuestionDal{
		DB: db,
	}
}

type QuestionDal struct {
	DB *gorm.DB
}

func (dal *QuestionDal) Create(ctx context.Context, item *entity.Question) error {
	result := dal.DB.Create(item)
	return errors.WithStack(result.Error)
}
