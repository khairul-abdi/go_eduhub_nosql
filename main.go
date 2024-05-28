package main

import (
	"go_eduhub_nosql/api"
	db "go_eduhub_nosql/db"
	repo "go_eduhub_nosql/repository"

	_ "embed"

	_ "github.com/lib/pq"
)

func main() {
	client, ctx, cancel, err := db.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	defer db.Close(client, ctx, cancel)

	// err = db.SQLExecute(client)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	studentRepo := repo.NewStudentRepo(ctx, client)

	//routing
	mainAPI := api.NewAPI(studentRepo)
	mainAPI.Start()
}
