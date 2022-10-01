package dal

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"DIdemo/entity"
)

func NewQuestionModelDal(db *gorm.DB) *QuestionModelDal {
	return &QuestionModelDal{
		DB: db,
	}
}

type QuestionModelDal struct {
	DB *gorm.DB
}

func (dal *QuestionModelDal) Create(ctx context.Context, item *entity.QuestionModel) error {
	result := dal.DB.Create(item)
	return errors.WithStack(result.Error)
}
