package repository

import (
	"database/sql"
	"go_eduhub/model"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
}

type studentRepoImpl struct {
	db *sql.DB
}

func NewStudentRepo(db *sql.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {

	rows, err := s.db.Query("SELECT * FROM public.students")
	if err != nil {
		panic(err)
	}

	var listStudents []model.Student

	for rows.Next() {
		var student model.Student

		err := rows.Scan(&student.ID, &student.Name, &student.Address, &student.Class)
		if err != nil {
			return nil, err
		}

		listStudents = append(listStudents, student)
	}

	return listStudents, nil
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	row := s.db.QueryRow("SELECT id, name, address, class FROM students WHERE id = $1", id)

	var student model.Student
	err := row.Scan(&student.ID, &student.Name, &student.Address, &student.Class)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}

		return nil, err
	}

	return &student, nil
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	_, err := s.db.Exec(`INSERT INTO public.students
	("name", address, "class")
	VALUES($1, $2, $3)`, student.Name, student.Address, student.Class)

	if err != nil {
		return err
	}

	return nil
}
