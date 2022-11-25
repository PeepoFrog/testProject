package postgre

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/PeepoFrog/testProject/internal/model"
)

type Postgre struct {
	db *sql.DB
}

func NewPostgre() *Postgre {
	db := CreatePostgreConnection()
	CreateIfNotExistTable(db)
	return &Postgre{db: db}

}
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
	b, err := ioutil.ReadFile("./migrations/createTable.sql")
	if err != nil {
		fmt.Println(err)
	}
	s := string(b)
	res, err := db.Exec(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s, res)

}

func (h *Postgre) LoadFromCSVToPostgre(rec model.Record, statment string) string {
	var ret string
	switch statment {
	case "BEGIN":
		fmt.Println("BEGIN")
		h.db.QueryRow(statment)
		return ret
	case "COMMIT":
		fmt.Println("COMMIT")
		h.db.QueryRow(statment)
		return ret
	default:
		fmt.Println("INSERT")
		sqlStatment := `INSERT INTO transactions (
			TransactionId,	
			RequestId,
			TerminalId,
			PartnerObjectId,
			AmountTotal,
			AmountOriginal,
			CommissionPS,
			CommissionClient,
			CommissionProvider,
			DateInput,
			DatePost,
			Status,
			PaymentType,
			PaymentNumber,
			ServiceId,
			Service,
			PayeeId,
			PayeeName,
			PayeeBankMfo,
			PayeeBankAccount,
			PaymentNarrative   
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21 ) RETURNING *`
		h.db.QueryRow(sqlStatment,
			rec.TransactionId,
			rec.RequestId,
			rec.TerminalId,
			rec.PartnerObjectId,
			rec.AmountTotal,
			rec.AmountOriginal,
			rec.CommissionPS,
			rec.CommissionClient,
			rec.CommissionProvider,
			rec.DateInput,
			rec.DatePost,
			rec.Status,
			rec.PaymentType,
			rec.PaymentNumber,
			rec.ServiceId,
			rec.Service,
			rec.PayeeId,
			rec.PayeeName,
			rec.PayeeBankMfo,
			rec.PayeeBankAccount,
			rec.PaymentNarrative).Scan(ret)
		return ret
	}
}

func (h *Postgre) QueryFormer(transactionid string, terminalid string, status string, paymenttype string, datepost string, paymentnarrative string) string {
	sqlStatment := "Select * from transactions WHERE "
	if transactionid != "" {
		sqlStatment = string(sqlStatment + " transactionid=" + transactionid + " AND")
	}
	if terminalid != "" {
		s := strings.Split(terminalid, ",")
		for _, element := range s {
			sqlStatment = string(sqlStatment + " terminalid=" + element + " OR")
		}
		sqlStatment = strings.TrimSuffix(sqlStatment, "OR")
		sqlStatment = string(sqlStatment + " AND")
	}
	if status != "" {
		sqlStatment = string(sqlStatment + " status='" + status + "' AND")
	}
	if paymenttype != "" {
		sqlStatment = string(sqlStatment + " paymenttype='" + paymenttype + "' AND")
	}
	if paymentnarrative != "" {
		sqlStatment = string(sqlStatment + " paymentnarrative like '%" + paymentnarrative + "%' AND")
	}
	if datepost != "" {
		s := strings.Split(datepost, ",")
		sqlStatment = string(sqlStatment + " datepost BETWEEN '" + s[0] + "' AND '" + s[1] + "'")
	}
	sqlStatment = strings.TrimSuffix(sqlStatment, "AND")
	sqlStatment = strings.TrimSuffix(sqlStatment, "WHERE ")
	fmt.Println(sqlStatment)
	return sqlStatment
}
func (h *Postgre) RunQuery(sqlStatment string) ([]model.Record, error) {
	var response model.Record
	var arresponse []model.Record
	rows, err := h.db.Query(sqlStatment)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
	}
	for rows.Next() {
		err = rows.Scan(
			&response.TransactionId,
			&response.RequestId,
			&response.TerminalId,
			&response.PartnerObjectId,
			&response.AmountTotal,
			&response.AmountOriginal,
			&response.CommissionPS,
			&response.CommissionClient,
			&response.CommissionProvider,
			&response.DateInput,
			&response.DatePost,
			&response.Status,
			&response.PaymentType,
			&response.PaymentNumber,
			&response.ServiceId,
			&response.Service,
			&response.PayeeId,
			&response.PayeeName,
			&response.PayeeBankMfo,
			&response.PayeeBankAccount,
			&response.PaymentNarrative)
		if err != nil {
			log.Printf("Unable to scan the row. %v", err)
		}
		arresponse = append(arresponse, response)
	}
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return arresponse, nil
	case nil:
		return arresponse, nil
	default:
		fmt.Printf("Unable to scan the row. %v", err)
	}
	return arresponse, err
}
