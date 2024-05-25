package repository

import (
	"database/sql"
	"go_eduhub/model"
)

type CourseRepository interface {
	FetchAll() ([]model.Course, error)
	FetchByID(id int) (*model.Course, error)
	Store(g *model.Course) error
	Update(id int, g *model.Course) error
}

type courseRepoImpl struct {
	db *sql.DB
}

func NewCourseRepo(db *sql.DB) *courseRepoImpl {
	return &courseRepoImpl{db}
}

func (g *courseRepoImpl) FetchAll() ([]model.Course, error) {
	rows, err := g.db.Query("SELECT * FROM courses")
	if err != nil {
		panic(err)
	}

	var listCourse []model.Course

	for rows.Next() {
		var course model.Course

		err := rows.Scan(&course.ID, &course.Name, &course.Schedule, &course.Attendance)
		if err != nil {
			return nil, err
		}

		listCourse = append(listCourse, course)
	}

	return listCourse, nil
}

func (g *courseRepoImpl) FetchByID(id int) (*model.Course, error) {
	row := g.db.QueryRow("SELECT id, name, schedule, attendance FROM courses WHERE id = $1", id)

	var course model.Course
	err := row.Scan(&course.ID, &course.Name, &course.Schedule, &course.Attendance)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}

		return nil, err
	}

	return &course, nil
}

func (g *courseRepoImpl) Store(course *model.Course) error {
	_, err := g.db.Exec(`INSERT INTO public.courses
	(name, schedule, attendance)
	VALUES($1, $2, $3)`, course.Name, course.Schedule, course.Attendance)

	if err != nil {
		return err
	}

	return nil
}

func (g *courseRepoImpl) Update(id int, course *model.Course) error {
	_, err := g.db.Exec(`UPDATE public.courses SET 
							name=$2,
							schedule=$3, 
							attendance=$4
						WHERE id =$1`,
		id,
		course.Name,
		course.Schedule,
		course.Attendance)

	if err != nil {
		return err
	}

	return nil
}
