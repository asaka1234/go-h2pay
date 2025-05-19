package go_h2pay

// ----------pre generate-------------------------

type H2PayDepositReq struct {
	Merchant    string `json:"merchant"`
	Currency    string `json:"currency"`
	Customer    string `json:"customer"`
	Reference   string `json:"reference"`
	Key         string `json:"key"`
	Amount      string `json:"amount"`
	Datetime    string `json:"datetime"`
	FrontURI    string `json:"frontURI"`
	BackURI     string `json:"backURI"`
	Bank        string `json:"bank"`
	Language    string `json:"language"`
	ClientIP    string `json:"clientIP"`
	DateTimeMd5 string `json:"dateTimeMd5"`
}

type H2PayDepositRsp struct {
	HTMLString string `json:"htmlString"`
}

//--------------------------------------------------------

type H2PayWithdrawReq struct {
	Key                 string `json:"key"`
	ClientIP            string `json:"clientIP"`
	ReturnURI           string `json:"returnURI"`
	MerchantCode        string `json:"merchantCode"`
	TransactionId       string `json:"transactionId"`
	CurrencyCode        string `json:"currencyCode"`
	MemberCode          string `json:"memberCode"`
	Amount              string `json:"amount"`
	TransactionDateTime string `json:"transactionDateTime"`
	BankCode            string `json:"bankCode"`
	ToBankAccountName   string `json:"toBankAccountName"`
	ToBankAccountNumber string `json:"toBankAccountNumber"`
	CustomerName        string `json:"customerName"`
	CustomerEmail       string `json:"customerEmail"`
	CustomerMobile      string `json:"customerMobile"`
	IFSC                string `json:"IFSC"`
	DateTimeMd5         string `json:"dateTimeMd5"`
}

type H2PayWithdrawRsp struct {
	StatusCode string `json:"statusCode"`
	Message    string `json:"message"`
}

//---------------------------------------------

type H2PayDepositBackReq struct {
	Merchant      string `json:"Merchant"`
	Reference     string `json:"Reference"`
	Currency      string `json:"Currency"`
	Amount        string `json:"Amount"`
	Language      string `json:"Language"`
	Customer      string `json:"Customer"`
	Datetime      string `json:"Datetime"`
	StatementDate string `json:"StatementDate"`
	Note          string `json:"Note"`
	Key           string `json:"Key"`
	Status        string `json:"Status"`
	Deposit       string `json:"Deposit"`
	ID            string `json:"ID"`
	ErrorCode     string `json:"ErrorCode"`
}

type H2PayWithdrawBackReq struct {
	MerchantCode        string `json:"MerchantCode"`
	TransactionID       string `json:"TransactionID"`
	CurrencyCode        string `json:"CurrencyCode"`
	Amount              string `json:"Amount"`
	TransactionDatetime string `json:"TransactionDatetime"`
	Key                 string `json:"Key"`
	Status              string `json:"Status"`
	Message             string `json:"Message"`
	MemberCode          string `json:"MemberCode"`
	ID                  string `json:"ID"`
	PayoutFee           string `json:"PayoutFee"`
}
