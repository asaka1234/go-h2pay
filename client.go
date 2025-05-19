package go_h2pay

import (
	"github.com/asaka1234/go-h2pay/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	MerchantID string // merchantId
	AccessKey  string // accessKey

	DepositURL  string
	WithdrawURL string

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, merchantID string, accessKey, depositURL, withdrawURL string) *Client {
	return &Client{
		MerchantID:  merchantID,
		AccessKey:   accessKey,
		DepositURL:  depositURL,
		WithdrawURL: withdrawURL,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
