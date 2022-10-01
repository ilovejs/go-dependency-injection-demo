package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/wire"

	"DIdemo/internal/schema"
	"DIdemo/internal/service"
)

var ProjectHandlerSet = wire.NewSet(
	wire.Struct(new(ProjectHandler), "*"),
)

type ProjectHandler struct {
	ProjectService service.IProjectService
}

//func (s *ProjectHandler) CreateProject(
//	ctx context.Context,
//	req *schema.CreateProjectRequest,
//) (resp *schema.CreateProjectResponse, err error) {
//	log.Println(ctx, req)
//	return nil, nil
//}

func (s *ProjectHandler) CreateProject(
	w http.ResponseWriter,
	r *http.Request,
) {

	log.Println("creating project")

	var payload *schema.CreateProjectRequest

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("project created: ", payload)

	w.WriteHeader(http.StatusOK)
	//_, _ = w.Write([]byte("project created"))
	_, _ = fmt.Fprintf(w, "Project %+v", payload)
}
