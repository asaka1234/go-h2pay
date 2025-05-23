package go_h2pay

// ----------pre generate-------------------------

type H2PayDepositReq struct {
	Currency  string `json:"Currency" mapstructure:"Currency"`
	Customer  string `json:"Customer" mapstructure:"Customer"`   //Merchant’s customer ID
	Reference string `json:"Reference" mapstructure:"Reference"` //Transaction ID created by Merchant, 必须唯一
	Amount    string `json:"Amount" mapstructure:"Amount"`       //法币是2位精度, VND, IDR,PPTP (THB currency)只允许整数, 不允许浮点数！
	Bank      string `json:"Bank" mapstructure:"Bank"`           //https://docs.google.com/spreadsheets/d/19ylX6-2XNkeke3HdYyT9m5lCV0PLnBmXC2f5SAjIye4/edit?gid=0#gid=0
	Language  string `json:"Language" mapstructure:"Language"`
	ClientIP  string `json:"ClientIP" mapstructure:"ClientIP"`
	//key是sdk帮计算的,所以不需要传
	//Key string `json:"key"` //对请求内容做的一个签名(需要自行运算)
	//商户号和回调地址等.sdk帮赋值
	//Merchant string `json:"Merchant"`
	//FrontURI  string `json:"FrontURI"`  //前端跳转地址
	//BackURI   string `json:"BackURI"`   //后端notify地址 (callback)
	//sdk设置好了
	//Datetime  string `json:"Datetime" mapstructure:"Datetime"`   //YYYY-MM-DD hh:mm:sstt . e.g. 2012-05-01 08:04:00AM

}

type H2PayDepositRsp struct {
	HTMLString string `json:"htmlString"`
}

// 充值的回调.
type H2PayDepositBackReq struct {
	Merchant      string `json:"Merchant" mapstructure:"Merchant"`
	Reference     string `json:"Reference" mapstructure:"Reference"` //Transaction ID created by Merchant, 必须唯一
	Currency      string `json:"Currency" mapstructure:"Currency"`
	Amount        string `json:"Amount" mapstructure:"Amount"`
	Language      string `json:"Language" mapstructure:"Language"`
	Customer      string `json:"Customer" mapstructure:"Customer"`           //Merchant’s customer ID
	Datetime      string `json:"Datetime" mapstructure:"Datetime"`           //YYYY-MM-DD hh:mm:sstt . e.g. 2012-05-01 08:04:00AM
	StatementDate string `json:"StatementDate" mapstructure:"StatementDate"` //Datetime for the transaction processed. In UTC time
	Note          string `json:"Note" mapstructure:"Note"`
	Key           string `json:"Key" mapstructure:"Key"`       //加密的签名
	Status        string `json:"Status" mapstructure:"Status"` //枚举: 000 成功, 001 失败, 006 approved, 007 Rejected,008 Canceled, 009 Pending
	Deposit       string `json:"Deposit" mapstructure:"Deposit"`
	ID            string `json:"ID" mapstructure:"ID"` //三方psp的订单id
	ErrorCode     string `json:"ErrorCode" mapstructure:"ErrorCode"`
}

//--------------------------------------------------------

type H2PayWithdrawReq struct {
	ClientIP            string `json:"ClientIP" mapstructure:"ClientIP"`
	TransactionID       string `json:"TransactionID" mapstructure:"TransactionID"` //商户的订单号
	CurrencyCode        string `json:"CurrencyCode" mapstructure:"CurrencyCode"`   //CNY, THB..
	MemberCode          string `json:"MemberCode" mapstructure:"MemberCode"`       //Merchant’s customer ID
	Amount              string `json:"Amount" mapstructure:"Amount"`
	BankCode            string `json:"BankCode" mapstructure:"BankCode"`
	ToBankAccountName   string `json:"toBankAccountName" mapstructure:"toBankAccountName"`
	ToBankAccountNumber string `json:"toBankAccountNumber" mapstructure:"toBankAccountNumber"`
	//以下都由sdk补充
	//Key          string `json:"Key" mapstructure:"Key"`                   //签名
	//ReturnURI    string `json:"ReturnURI" mapstructure:"ReturnURI"`       //回调地址
	//MerchantCode string `json:"MerchantCode" mapstructure:"MerchantCode"` //psp的商户号
	//sdk给补充
	//TransactionDateTime string `json:"TransactionDateTime" mapstructure:"TransactionDateTime"` //YYYY-MM-DD hh:mm:sstt . e.g. 2012-05-01 08:04:00AM

}

type H2PayWithdrawRsp struct {
	StatusCode string `json:"statusCode"`
	Message    string `json:"message"`
}

//---------------------------------------------

type H2PayWithdrawBackReq struct {
	MerchantCode        string `json:"MerchantCode" mapstructure:"MerchantCode"`
	TransactionID       string `json:"TransactionID" mapstructure:"TransactionID"`
	CurrencyCode        string `json:"CurrencyCode" mapstructure:"CurrencyCode"`
	Amount              string `json:"Amount" mapstructure:"Amount"`
	TransactionDatetime string `json:"TransactionDatetime" mapstructure:"TransactionDatetime"` //YYYY-MM-DD hh:mm:sstt . e.g. 2012-05-01 08:04:00AM
	Key                 string `json:"Key" mapstructure:"Key"`                                 //签名
	Status              string `json:"Status" mapstructure:"Status"`                           //枚举: 000 成功, 001 失败, 006 approved, 007 Rejected,008 Canceled, 009
	Message             string `json:"Message" mapstructure:"Message"`
	MemberCode          string `json:"MemberCode" mapstructure:"MemberCode"`
	ID                  string `json:"ID" mapstructure:"ID"` //Payout ID 三方psp的订单id
}
