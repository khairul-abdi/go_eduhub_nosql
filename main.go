package main

import (
	"database/sql"
	"fmt"
	"go_eduhub/api"
	"go_eduhub/model"
	repo "go_eduhub/repository"
	"log"

	_ "embed"

	_ "github.com/lib/pq"
)

func Connect(creds *model.Credential) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", creds.Host, creds.Username, creds.Password, creds.DatabaseName, creds.Port)

	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func SQLExecute(db *sql.DB) error {
	//create table students
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS students (id SERIAL PRIMARY KEY, name VARCHAR(255), address VARCHAR(255), class VARCHAR(255))")
	if err != nil {
		return err
	}

	//create table courses
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS courses (id SERIAL PRIMARY KEY, name VARCHAR(255), schedule VARCHAR(255), attendance int)")
	if err != nil {
		return err
	}

	return nil
}

func Reset(db *sql.DB, table string) error {
	_, err := db.Exec("TRUNCATE " + table)
	if err != nil {
		return err
	}

	_, err = db.Exec("ALTER SEQUENCE " + table + "_id_seq RESTART WITH 1")
	if err != nil {
		return err
	}

	return nil
}

func main() {
	dbCredential := model.Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "postgres",
		DatabaseName: "kampusmerdeka",
		Port:         5432,
	}
	dbConn, err := Connect(&dbCredential)
	if err != nil {
		log.Fatal(err)
	}

	err = SQLExecute(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	defer dbConn.Close()

	studentRepo := repo.NewStudentRepo(dbConn)
	courseRepo := repo.NewCourseRepo(dbConn)

	mainAPI := api.NewAPI(studentRepo, courseRepo)
	mainAPI.Start()
}
