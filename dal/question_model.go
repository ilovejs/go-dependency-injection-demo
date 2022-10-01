package dal

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"DIdemo/entity"
)

type IQuestionModelDal interface {
	Create(ctx context.Context, item *entity.QuestionModel) error
}

type QuestionModelDal struct {
	DB *gorm.DB
}

var QuestionModelDalSet = wire.NewSet(
	wire.Struct(new(QuestionModelDal), "*"),
	wire.Bind(new(IQuestionModelDal), new(*QuestionModelDal)),
)

func (dal *QuestionModelDal) Create(ctx context.Context, item *entity.QuestionModel) error {
	result := dal.DB.WithContext(ctx).Create(item)
	return errors.WithStack(result.Error)
}
