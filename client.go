package go_h2pay

import (
	"github.com/asaka1234/go-h2pay/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	MerchantID string // merchantId
	AccessKey  string // accessKey

	DepositURL         string
	DepositCallbackURL string //sever回调地址
	DepositFeBackURL   string //前端跳转地址

	WithdrawURL         string
	WithdrawCallbackURL string

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, merchantID string, accessKey, depositURL, withdrawURL, depositCallbackURL, depositFeBackURL, withdrawCallbackURL string) *Client {
	return &Client{
		MerchantID:          merchantID,
		AccessKey:           accessKey,
		DepositURL:          depositURL,
		WithdrawURL:         withdrawURL,
		DepositCallbackURL:  depositCallbackURL,
		DepositFeBackURL:    depositFeBackURL,
		WithdrawCallbackURL: withdrawCallbackURL,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
