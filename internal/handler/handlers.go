package handler

import (
	"database/sql"
	"mime/multipart"

	"github.com/PeepoFrog/testProject/internal/initilisation"
	"github.com/PeepoFrog/testProject/internal/model"

	_ "github.com/lib/pq"
)

type Handlers struct {
	db *sql.DB
}

func NewHanlders() *Handlers {
	db := initilisation.CreatePostgreConnection()
	initilisation.CreateIfNotExistTable(db)
	return &Handlers{db: db}

}

type TransactionRepository interface {
	LoadFromCSVToPostgre(multipart.File)
	QueryFormer(transactionid string, terminalid string, status string, paymenttype string, datefilter string, phasetosearch string) string
	RunQuery(sqlStatment string) ([]model.Record, error)
}
