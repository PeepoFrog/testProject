package handler

import (
	"database/sql"
	"mime/multipart"

	"github.com/PeepoFrog/testProject/iternal/initilisation"
	"github.com/PeepoFrog/testProject/iternal/model"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "database"
// 	port     = 5432
// 	user     = "dima"
// 	password = "pass"
// 	dbname   = "dima"
// )

type Handlers struct {
	db *sql.DB
}

func NewHanlders() *Handlers {
	db := initilisation.CreatePostgreConnection()
	initilisation.CreateIfNotExistTable(db)
	return &Handlers{db: db}

}

type HandlersRepository interface {
	LoadFromCSVToPostgre(multipart.File)
	QueryFormer(transactionid string, terminalid string, status string, paymenttype string, datefilter string, phasetosearch string) string
	RunQuery(sqlStatment string) ([]model.Record, error)
}
