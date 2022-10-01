package dal

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"DIdemo/entity"
)

type IQuestionDal interface {
	Create(ctx context.Context, item *entity.Question) error
}

type QuestionDal struct {
	DB *gorm.DB
}

var QuestionDalSet = wire.NewSet(
	wire.Struct(new(QuestionDal), "*"),
	wire.Bind(new(IQuestionDal), new(*QuestionDal)),
)

func (dal *QuestionDal) Create(ctx context.Context, item *entity.Question) error {
	result := dal.DB.WithContext(ctx).Create(item)
	return errors.WithStack(result.Error)
}
