package go_h2pay

import (
	"crypto/tls"
	"fmt"
	"github.com/asaka1234/go-h2pay/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
	"time"
)

// 下单
func (cli *Client) Deposit(req H2PayDepositReq) (*H2PayDepositRsp, error) {

	rawURL := cli.Params.DepositUrl

	//----------------------判断bank code的正确性------------------
	_, ok := lo.Find(DepositBankCodes, func(i H2PayBankCode) bool {
		return i.Code == req.Bank
	})
	if !ok {
		return nil, fmt.Errorf("bank code %s error", req.Bank)
	}
	//---------------------------------------------------------

	loc := time.FixedZone("UTC", 8*3600)

	//先转成map
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	//以此确保amount是2位精度!
	amount := decimal.NewFromFloat(cast.ToFloat64(params["Amount"])) //转为decimal
	params["Amount"] = amount.StringFixed(2)

	params["Merchant"] = cli.Params.MerchantId
	params["FrontURI"] = cli.Params.DepositFeBackUrl
	params["BackURI"] = cli.Params.DepositBackUrl
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
		SetDebug(cli.debugMode).
		Post(rawURL)

	if err != nil {
		cli.logger.Errorf("请求失败: %s", err.Error())
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode())
	}

	if resp.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%v, body:%s", resp.Error(), resp.Body())
	}

	// Log response
	responseStr := string(resp.Body())

	// Build response struct
	rsp := &H2PayDepositRsp{
		HTMLString: responseStr,
	}

	return rsp, nil
}
