package go_h2pay

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"github.com/asaka1234/go-h2pay/utils"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
	"time"
)

// 提现
func (cli *Client) Withdraw(req H2PayWithdrawReq) (*H2PayWithdrawRsp, error) {
	rawURL := cli.Params.WithdrawUrl

	//----------------------判断bank code的正确性------------------
	_, ok := lo.Find(WithdrawBankCodes, func(i H2PayBankCode) bool {
		return i.Code == req.BankCode && i.Currency == req.CurrencyCode
	})
	if !ok {
		return nil, fmt.Errorf("bank code %s error", req.BankCode)
	}
	//---------------------------------------------------------

	loc := time.FixedZone("UTC", 8*3600)

	//先转成map
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	//以此确保amount是2位精度!
	amount := decimal.NewFromFloat(cast.ToFloat64(params["Amount"])) //转为decimal
	params["Amount"] = amount.StringFixed(2)

	params["ReturnURI"] = cli.Params.WithdrawBackUrl
	params["MerchantCode"] = cli.Params.MerchantId
	params["TransactionDateTime"] = time.Now().In(loc).Format("2006-01-02 03:04:05PM")

	// Generate signature
	signStr := utils.WithdrawSign(params, cli.Params.AccessKey)
	params["Key"] = signStr

	// 发送HTTP请求
	resp, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getHeaders()).
		SetFormData(utils.ConvertToStringMap(params)).
		SetDebug(cli.debugMode).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp))
	cli.logger.Infof("PSPResty#h2pay#withdraw->%s", string(restLog))

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

	responseStr := string(resp.Body())

	// Parse XML response (implement your parseXml function)
	result, err := cli.parseXml(responseStr)
	if err != nil {
		cli.logger.Infof("H2PayService#withdraw parse error: %v", err)
		return nil, err
	}

	return result, nil
}

//----------------------------------------------------

type Payout struct {
	XMLName    xml.Name `xml:"Payout"`
	StatusCode string   `xml:"statusCode"`
	Message    string   `xml:"message"`
}

func (cli *Client) parseXml(xmlString string) (*H2PayWithdrawRsp, error) {
	// Parse the XML
	var payout Payout
	err := xml.Unmarshal([]byte(xmlString), &payout)
	if err != nil {
		cli.logger.Infof("parseXml#error, xml: %s, error: %v", xmlString, err)
		return nil, err
	}

	// Create the response
	rsp := &H2PayWithdrawRsp{
		StatusCode: payout.StatusCode,
		Message:    payout.Message,
	}

	return rsp, nil
}
