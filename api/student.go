package api

import (
	"encoding/json"
	"go_eduhub_nosql/model"
	"net/http"
)

func (api *API) FetchAllStudent(w http.ResponseWriter, r *http.Request) {
	student, err := api.studentRepo.FetchAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

func (api *API) FetchStudentByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	student, err := api.studentRepo.FetchByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

func (api *API) Storestudent(w http.ResponseWriter, r *http.Request) {
	var student model.Student

	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	err = api.studentRepo.Store(&student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

func (api *API) Updatestudent(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var student model.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	err = api.studentRepo.Update(id, &student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}
