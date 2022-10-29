package init

import (
	"database/sql"
	"fmt"
)

func createPostgreConnection() *sql.DB {
	psqlInfo := "host=database port=5432 user=dima " +
		"password=pass dbname=dima sslmode=disable"

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db
}
