package handler

import (
	"context"
	"log"

	"DIdemo/schema"
	"DIdemo/service"
)

func NewProjectHandler(srv *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		ProjectService: srv,
	}
}

type ProjectHandler struct {
	ProjectService *service.ProjectService
}

func (s *ProjectHandler) CreateProject(
	ctx context.Context,
	req *schema.CreateProjectRequest,
) (resp *schema.CreateProjectResponse, err error) {

	log.Println(ctx, req)

	return nil, nil
}
