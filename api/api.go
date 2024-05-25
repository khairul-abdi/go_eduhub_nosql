package api

import (
	"fmt"
	repo "go_eduhub/repository"
	"net/http"
)

type API struct {
	studentRepo repo.StudentRepository
	courseRepo  repo.CourseRepository
	mux         *http.ServeMux
}

func NewAPI(studentRepo repo.StudentRepository, courseRepo repo.CourseRepository) API {
	mux := http.NewServeMux()
	api := API{
		studentRepo,
		courseRepo,
		mux,
	}

	mux.Handle("/student/get-all", api.Get(http.HandlerFunc(api.FetchAllStudent)))
	mux.Handle("/student/get", api.Get(http.HandlerFunc(api.FetchStudentByID)))
	mux.Handle("/student/add", api.Post(http.HandlerFunc(api.Storestudent)))

	mux.Handle("/course/get-all", api.Get(http.HandlerFunc(api.FetchAllCourse)))
	mux.Handle("/course/get", api.Get(http.HandlerFunc(api.FetchCourseByID)))
	mux.Handle("/course/add", api.Post(http.HandlerFunc(api.StoreCourse)))
	mux.Handle("/course/update", api.Put(http.HandlerFunc(api.UpdateCourse)))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}
