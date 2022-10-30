package initilisation

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"time"
)

func CreatePostgreConnection() *sql.DB {
	psqlInfo := "host=database port=5432 user=dima " +
		"password=pass dbname=dima sslmode=disable"

	db, err := sql.Open("postgres", psqlInfo)

	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()

	takes := 0
	for err != nil {
		fmt.Println("failed to connect waiting 4 seconds and trying again")
		time.Sleep(4 * time.Second)

		db, err = sql.Open("postgres", psqlInfo)
		fmt.Println(err)
		err = db.Ping()
		fmt.Println(err)

		takes++
		fmt.Println(takes, " take out of 8")
		if takes > 7 {
			break
		}
	}

	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	return db
}

func CreateIfNotExistTable(db *sql.DB) {
	b, err := ioutil.ReadFile("../../migration/createTable.sql")
	if err != nil {
		fmt.Println(err)
	}
	s := string(b)
	fmt.Println(s)
	res, err := db.Exec(s)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s, res)

}
