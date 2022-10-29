package handler

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/PeepoFrog/testProject/iternal/model"
)

func (h *Handlers) LoadFromCSVToPostgre(f multipart.File) {
	defer f.Close()
	reader := csv.NewReader(f)
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
		var ret string
		sqlStatment := `INSERT INTO exampledima (
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
		h.db.QueryRow(sqlStatment, rec.TransactionId, rec.RequestId, rec.TerminalId, rec.PartnerObjectId, rec.AmountTotal, rec.AmountOriginal, rec.CommissionPS, rec.CommissionClient, rec.CommissionProvider, rec.DateInput, rec.DatePost, rec.Status, rec.PaymentType, rec.PaymentNumber, rec.ServiceId, rec.Service, rec.PayeeId, rec.PayeeName, rec.PayeeBankMfo, rec.PayeeBankAccount, rec.PaymentNarrative).Scan(ret)

	}
}
