package database

import (
	"mime/multipart"

	"github.com/PeepoFrog/testProject/internal/model"

	_ "github.com/lib/pq"
)

type TransactionRepository interface {
	LoadFromCSVToPostgre(multipart.File)
	QueryFormer(transactionid string, terminalid string, status string, paymenttype string, datefilter string, phasetosearch string) string
	RunQuery(sqlStatment string) ([]model.Record, error)
}
