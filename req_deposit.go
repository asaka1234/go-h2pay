package go_h2pay

import (
	"crypto/tls"
	"github.com/asaka1234/go-h2pay/utils"
	"github.com/mitchellh/mapstructure"
	"log"
	"time"
)

// 下单
func (cli *Client) Deposit(req H2PayDepositReq) (*H2PayDepositRsp, error) {

	rawURL := cli.Params.DepositUrl

	loc := time.FixedZone("UTC", 8*3600)

	//先转成map
	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["Merchant"] = cli.Params.MerchantId
	params["FrontURI"] = cli.Params.DepositFeBackUrl
	params["BackURI"] = cli.Params.DepositCallbackUrl
	params["Datetime"] = time.Now().In(loc).Format("2006-01-02 03:04:05PM")

	// Generate signature
	signStr := utils.DepositSign(params, cli.Params.AccessKey)
	params["Key"] = signStr

	// 发送HTTP请求
	resp, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getHeaders()).
		SetFormData(utils.ConvertToStringMap(params)).
		Post(rawURL)

	if err != nil {
		cli.logger.Errorf("请求失败: %s", err.Error())
		return nil, err
	}

	// Log response
	responseStr := string(resp.Body())
	log.Printf("H2PayService#deposit#rsp: %s", responseStr)

	// Build response struct
	rsp := &H2PayDepositRsp{
		HTMLString: responseStr,
	}

	return rsp, nil
}
