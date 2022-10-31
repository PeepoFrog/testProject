package controller

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/PeepoFrog/testProject/internal/database"
	"github.com/PeepoFrog/testProject/internal/model"

	"github.com/gocarina/gocsv"
)

type Controller struct {
	repository database.TransactionRepository
}

func NewController(repository database.TransactionRepository) *Controller {
	return &Controller{repository: repository}
}

// LoadFromCSVToDB godoc
// @Summary      Пошук по параметрам з відповіддю у форматі CSV
// @Tags         uploadfile
// @ID           file.upload
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file          true  "тест файл"
// @Success 200
// @Router  /uploadfile [put]
func (c *Controller) LoadFromCSVToDB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	w.Header().Set("Content-Type", "application/json")

	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(err)

	}
	defer file.Close()
	reader := csv.NewReader(file)
	firstrow := true
	for {
		record, err := reader.Read()
		if err != nil {
			fmt.Println(err)
			break
		}
		if firstrow {
			firstrow = false
			continue
		}
		var rec model.Record
		rec.TransactionId, _ = strconv.Atoi(record[0])
		rec.RequestId, _ = strconv.Atoi(record[1])
		rec.TerminalId, _ = strconv.Atoi(record[2])
		rec.PartnerObjectId, _ = strconv.Atoi(record[3])
		rec.AmountTotal, _ = strconv.ParseFloat(record[4], 64)
		rec.AmountOriginal, _ = strconv.ParseFloat(record[5], 64)
		rec.CommissionPS, _ = strconv.ParseFloat(record[6], 64)
		rec.CommissionClient, _ = strconv.ParseFloat(record[7], 64)
		rec.CommissionProvider, _ = strconv.ParseFloat(record[8], 64)
		//
		t, _ := time.Parse("2006-1-2 15:04:05", record[9])
		rec.DateInput = t.Format("2006-1-2 15:04:05")
		t, _ = time.Parse("2006-1-2 15:04:05", record[10])
		rec.DatePost = t.Format("2006-1-2 15:04:05")
		//
		rec.Status = record[11]
		rec.PaymentType = record[12]
		rec.PaymentNumber = record[13]
		rec.ServiceId, _ = strconv.Atoi(record[14])
		rec.Service = record[15]
		rec.PayeeId = record[16]
		rec.PayeeName = record[17]
		rec.PayeeBankMfo, _ = strconv.Atoi(record[18])
		rec.PayeeBankAccount = record[19]
		rec.PaymentNarrative = record[20]
		c.repository.LoadFromCSVToPostgre(rec)

	}

}

// QrlSearch godoc
// @Summary      Пошук по параметрам
// @Tags         search
// @Accept       json
// @Produce      json
// @Content-Type application/json
// @Param        transactionid    query int    false "Шукати по  transactionid"                                                                           Format(number)
// @Param        terminalid       query string false "Шукати по  terminalid шукати за декількома одночасно можна через кому наприклад 3507,3508,3509...." Format(text)
// @Param        status           query string false "Шукати по  status accepted/declined"                                                                Format(text)
// @Param        paymenttype      query string false "Шукати по  payment type cash/card "                                                                 Format(text)
// @Param        datepost         query string false "Шукати по  date post  рік-місяць-день з,по. Наприклад: 2022-08-18,2022-09-28"                  Format(text)
// @Param        paymentnarrative query string false "Шукати по  Payment narrative"                                                                       Format(text)
// @Success 200
// @Router  /search [get]
func (c *Controller) QrlSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

// QrlSearchToCSV godoc
// @Summary      Пошук по параметрам з відповіддю у форматі CSV
// @Tags         search
// @Accept       json
// @Produce      json
// @Content-Type response/csv
// @Param        transactionid    query int    false "Шукати по  transactionid"                                                                           Format(number)
// @Param        terminalid       query string false "Шукати по  terminalid шукати за декількома одночасно можна через кому наприклад 3507,3508,3509...." Format(text)
// @Param        status           query string false "Шукати по  status accepted/declined"                                                                Format(text)
// @Param        paymenttype      query string false "Шукати по  payment type cash/card "                                                                 Format(text)
// @Param        datepost         query string false "Шукати по  date post  рік-місяць-день з,по. Наприклад: 2022-08-18,2022-09-28"                  Format(text)
// @Param        paymentnarrative query string false "Шукати по  Payment narrative"                                                                       Format(text)
// @Success 200
// @Router  /searchcsv [get]
func (c *Controller) QrlSearchToCSV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "response/csv")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	w.Header().Set("Content-Disposition", "attachment;filename=response.csv")
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
