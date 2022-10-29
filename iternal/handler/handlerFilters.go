package handler

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/PeepoFrog/testProject/iternal/model"
)

func (h *Handlers) QueryFormer(transactionid string, terminalid string, status string, paymenttype string, datepost string, paymentnarrative string) string {
	sqlStatment := "Select * from exampledima WHERE "
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
func (h *Handlers) RunQuery(sqlStatment string) ([]model.Record, error) {
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
