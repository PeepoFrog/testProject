package controller

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
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
	c.repository.LoadFromCSVToPostgre(model.Record{}, "BEGIN")
	var arrarr []string
	var errstring string

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
		var aerr error
		rec.TransactionId, aerr = strconv.Atoi(record[0])
		if aerr != nil {
			errstring = "error with converting TransactionID with " + record[0] + " data"
			arrarr = append(arrarr, errstring)
			log.Println(aerr)
		}
		rec.RequestId, aerr = strconv.Atoi(record[1])
		if aerr != nil {
			errstring = "error with converting RequestID with " + record[1] + " data"
			arrarr = append(arrarr, errstring)
			log.Println(aerr)
		}
		rec.TerminalId, aerr = strconv.Atoi(record[2])
		if aerr != nil {
			errstring = "error with converting TerminalID with " + record[2] + " data"
			arrarr = append(arrarr, errstring)
			log.Println(aerr)
		}
		rec.PartnerObjectId, aerr = strconv.Atoi(record[3])
		if aerr != nil {
			errstring = "error with converting PartnerObjectID with " + record[3] + " data"
			arrarr = append(arrarr, errstring)
			log.Println(aerr)
		}
		rec.AmountTotal, aerr = strconv.ParseFloat(record[4], 64)
		if aerr != nil {
			errstring = "error with converting AmountTotal with " + record[4] + " data"
			arrarr = append(arrarr, errstring)
			log.Println(aerr)
		}
		rec.AmountOriginal, aerr = strconv.ParseFloat(record[5], 64)
		if aerr != nil {
			errstring = "error with converting AmountOriginal with " + record[5] + " data"
			arrarr = append(arrarr, errstring)
			log.Println(aerr)
		}
		rec.CommissionPS, aerr = strconv.ParseFloat(record[6], 64)
		if aerr != nil {
			errstring = "error with converting CommisionPS with " + record[6] + " data"
			arrarr = append(arrarr, errstring)
			log.Println(aerr)
		}
		rec.CommissionClient, aerr = strconv.ParseFloat(record[7], 64)
		if err != nil {
			errstring = "error with converting CommisionClient with " + record[7] + " data"
			arrarr = append(arrarr, errstring)
			log.Println(err)
		}
		rec.CommissionProvider, aerr = strconv.ParseFloat(record[8], 64)
		if aerr != nil {
			errstring = "error with converting CommisionProvider with " + record[8] + " data"
			arrarr = append(arrarr, errstring)
			log.Println(aerr)
		}
		t, aerr := time.Parse("2006-1-2 15:04:05", record[9])
		if err != nil {
			errstring = "cannot parse data Wrong Date type " + record[9]
			arrarr = append(arrarr, errstring)
			log.Println(err)
		}
		rec.DateInput = t.Format("2006-1-2 15:04:05")
		t, aerr = time.Parse("2006-1-2 15:04:05", record[10])
		if aerr != nil {
			errstring = "cannot parse data wrong Date type" + record[10]
			arrarr = append(arrarr, errstring)
			log.Println(aerr)
		}
		rec.DatePost = t.Format("2006-1-2 15:04:05")
		rec.Status = record[11]
		rec.PaymentType = record[12]
		rec.PaymentNumber = record[13]
		rec.ServiceId, aerr = strconv.Atoi(record[14])
		if aerr != nil {
			errstring = "error with converting ServiceID with " + record[14] + " data"
			arrarr = append(arrarr, errstring)
			log.Println(aerr)
		}
		rec.Service = record[15]
		rec.PayeeId = record[16]
		rec.PayeeName = record[17]
		rec.PayeeBankMfo, aerr = strconv.Atoi(record[18])
		if aerr != nil {
			errstring = "error with converting PayeeBankMfo with " + record[18] + " data"
			arrarr = append(arrarr, errstring)
			log.Println(aerr)
		}
		rec.PayeeBankAccount = record[19]
		rec.PaymentNarrative = record[20]
		c.repository.LoadFromCSVToPostgre(rec, "")
	}
	fmt.Println(arrarr)
	if arrarr == nil {
		c.repository.LoadFromCSVToPostgre(model.Record{}, "COMMIT")

	} else {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(arrarr)
		err = json.NewEncoder(w).Encode(arrarr)
		if err != nil {
			log.Println(err)
		}
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
		fmt.Println(err)
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		fmt.Println(err)
	}

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
	w.Header().Set("Content-Type", "text/csv")
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
		fmt.Println(err)
	}
	err = gocsv.Marshal(qres, w)
	if err != nil {
		fmt.Println(err)
	}
}
