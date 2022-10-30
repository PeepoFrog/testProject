create table IF NOT EXISTS transactions(
TransactionId      int,
	RequestId          int,
	TerminalId         int,
	PartnerObjectId    int,
	AmountTotal        float8 ,
	AmountOriginal     float8 ,
	CommissionPS       float8 ,
	CommissionClient   float8 ,
	CommissionProvider float8 ,
	DateInput          TIMESTAMP,
	DatePost           TIMESTAMP,
	Status             VARCHAR,
	PaymentType        VARCHAR,
	PaymentNumber      VARCHAR,
	ServiceId          int,
	Service            VARCHAR,
	PayeeId            VARCHAR,
	PayeeName          VARCHAR,
	PayeeBankMfo       int,
	PayeeBankAccount   VARCHAR,
	PaymentNarrative   VARCHAR

);