package service

import (
	"context"
	"log"

	"github.com/google/wire"

	"DIdemo/dal"
	"DIdemo/internal/schema"
)

type IProjectService interface {
	Create(ctx context.Context, payload *schema.CreateProjectRequest) (int64, error)
}

type ProjectService struct {
	ProjectDal       *dal.ProjectDal
	QuestionDal      *dal.QuestionDal
	QuestionModelDal *dal.QuestionModelDal
}

var ProjectSet = wire.NewSet(
	wire.Struct(new(ProjectService), "*"),
	wire.Bind(new(IProjectService), new(*ProjectService)))

// Create do nnt have projectBo *bo.ProjectCreateBo in original article.
func (s *ProjectService) Create(
	ctx context.Context,
	payload *schema.CreateProjectRequest,
) (int64, error) {
	log.Println("project created: ", payload)
	return 0, nil
}
