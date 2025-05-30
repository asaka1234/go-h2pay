package go_h2pay

import (
	"github.com/asaka1234/go-h2pay/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Params H2PayInitParams

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, params H2PayInitParams) *Client {
	return &Client{
		Params: params,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
