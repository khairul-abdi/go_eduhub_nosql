package api

import (
	"fmt"
	repo "go_eduhub_nosql/repository"
	"net/http"
)

type API struct {
	studentRepo repo.StudentRepository
	mux         *http.ServeMux
}

func NewAPI(studentRepo repo.StudentRepository) API {
	mux := http.NewServeMux()
	api := API{
		studentRepo,
		mux,
	}

	mux.Handle("/student/get-all", api.Get(http.HandlerFunc(api.FetchAllStudent)))
	mux.Handle("/student/get", api.Get(http.HandlerFunc(api.FetchStudentByID)))
	mux.Handle("/student/add", api.Post(http.HandlerFunc(api.Storestudent)))
	mux.Handle("/student/update", api.Put(http.HandlerFunc(api.Updatestudent)))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}
