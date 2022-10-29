package handler

import (
	"database/sql"
	"fmt"
	"mime/multipart"

	"github.com/PeepoFrog/testProject/iternal/model"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "dima"
	password = "pass"
	dbname   = "dima"
)

type Handlers struct {
	db *sql.DB
}

func NewHanlders() *Handlers {
	db := createPostgreConnection()
	return &Handlers{db: db}

}

type HandlersRepository interface {
	LoadFromCSVToPostgre(multipart.File)
	QueryFormer(transactionid string, terminalid string, status string, paymenttype string, datefilter string, phasetosearch string) string
	RunQuery(sqlStatment string) ([]model.Record, error)
}

func createPostgreConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

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
