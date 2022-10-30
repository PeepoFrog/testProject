package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PeepoFrog/testProject/internal/handler"

	"github.com/gocarina/gocsv"
)

type Controller struct {
	repository handler.TransactionRepository
}

func NewControllerRepository(repository handler.TransactionRepository) *Controller {
	return &Controller{repository: repository}
}
func (c *Controller) LoadFromCSVToDB(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(header)
	c.repository.LoadFromCSVToPostgre(file)
}

func (c *Controller) QrlSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	transactionid := r.URL.Query().Get("transactionid")
	terminalid := r.URL.Query().Get("terminalid")
	status := r.URL.Query().Get("status")
	paymenttype := r.URL.Query().Get("paymenttype")
	datefilter := r.URL.Query().Get("datepost")
	phasetosearch := r.URL.Query().Get("paymentnarrative")
	statment := c.repository.QueryFormer(transactionid, terminalid, status, paymenttype, datefilter, phasetosearch)
	res, err := c.repository.RunQuery(statment)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(res)

}
func (c *Controller) QrlSearchToCSV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "response/csv")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	w.Header().Set("Content-Disposition", "attachment;filename=tests.csv")
	transactionid := r.URL.Query().Get("transactionid")
	terminalid := r.URL.Query().Get("terminalid")
	status := r.URL.Query().Get("status")
	paymenttype := r.URL.Query().Get("paymenttype")
	datefilter := r.URL.Query().Get("datepost")
	phasetosearch := r.URL.Query().Get("paymentnarrative")
	statment := c.repository.QueryFormer(transactionid, terminalid, status, paymenttype, datefilter, phasetosearch)
	qres, err := c.repository.RunQuery(statment)

	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	gocsv.Marshal(qres, w)

}
