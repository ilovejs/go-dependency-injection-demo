package service

import (
	"context"

	"DIdemo/dal"
)

func NewProjectService(
	projectDal *dal.ProjectDal,
	questionDal *dal.QuestionDal,
	questionModelDal *dal.QuestionModelDal,
) *ProjectService {
	return &ProjectService{
		ProjectDal:       projectDal,
		QuestionDal:      questionDal,
		QuestionModelDal: questionModelDal,
	}
}

type ProjectService struct {
	ProjectDal       *dal.ProjectDal
	QuestionDal      *dal.QuestionDal
	QuestionModelDal *dal.QuestionModelDal
}

// projectBo *bo.ProjectCreateBo
func (s *ProjectService) Create(ctx context.Context) (int64, error) {
	return 0, nil
}
