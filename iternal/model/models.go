package model

type Record struct {
	TransactionId      int
	RequestId          int
	TerminalId         int
	PartnerObjectId    int
	AmountTotal        float64
	AmountOriginal     float64
	CommissionPS       float64
	CommissionClient   float64
	CommissionProvider float64
	DateInput          string
	DatePost           string
	Status             string
	PaymentType        string
	PaymentNumber      string
	ServiceId          int
	Service            string
	PayeeId            string
	PayeeName          string
	PayeeBankMfo       int
	PayeeBankAccount   string
	PaymentNarrative   string
}
