package go_h2pay

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// 提现
func (cli *Client) Withdraw(req H2PayWithdrawReq) (*H2PayWithdrawRsp, error) {
	key := cli.getWithdrawMD5(req)

	rawURL := cli.WithdrawURL

	// Build parameter list in the exact same order as Java
	params := []string{
		fmt.Sprintf("Key=%s", key),
		fmt.Sprintf("ClientIP=%s", req.ClientIP),
		fmt.Sprintf("ReturnURI=%s", req.ReturnURI),
		fmt.Sprintf("MerchantCode=%s", req.MerchantCode),
		fmt.Sprintf("TransactionID=%s", req.TransactionId),
		fmt.Sprintf("CurrencyCode=%s", req.CurrencyCode),
		fmt.Sprintf("MemberCode=%s", req.MemberCode),
		fmt.Sprintf("Amount=%s", req.Amount),
		fmt.Sprintf("TransactionDateTime=%s", req.TransactionDateTime),
		fmt.Sprintf("BankCode=%s", req.BankCode),
		fmt.Sprintf("toBankAccountName=%s", req.ToBankAccountName),
		fmt.Sprintf("toBankAccountNumber=%s", req.ToBankAccountNumber),
	}

	// Join parameters with &
	paramStr := strings.Join(params, "&")

	// Log request
	cli.logger.Infof("H2PayService#withdraw, url: %s, param: %s", rawURL, paramStr)

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: time.Second * 30, // Set appropriate timeout
	}

	// Create POST request
	resp, err := client.Post(
		rawURL,
		"application/x-www-form-urlencoded",
		strings.NewReader(paramStr),
	)
	if err != nil {
		cli.logger.Infof("H2PayService#withdraw error: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cli.logger.Infof("H2PayService#withdraw read response error: %v", err)
		return nil, err
	}

	responseStr := string(body)
	cli.logger.Infof("H2PayService#withdraw#rsp: %s", responseStr)

	// Parse XML response (implement your parseXml function)
	result, err := cli.parseXml(responseStr)
	if err != nil {
		cli.logger.Infof("H2PayService#withdraw parse error: %v", err)
		return nil, err
	}

	return result, nil
}

func (cli *Client) getWithdrawMD5(req H2PayWithdrawReq) string {
	// Construct the encoded string in the exact same order as Java
	encodeStr := req.MerchantCode + req.TransactionId + req.MemberCode +
		req.Amount + req.CurrencyCode + req.DateTimeMd5 +
		req.ToBankAccountNumber + cli.AccessKey

	// Log before MD5 (matches Java log format)
	cli.logger.Infof("H2PayService#MD5#withdraw#before, s: %s", encodeStr)

	// Generate MD5 hash
	hash := md5.Sum([]byte(encodeStr))
	result := hex.EncodeToString(hash[:])

	// Log after MD5 (matches Java log format)
	cli.logger.Infof("H2PayService#MD5#withdraw#end, s: %s", result)

	return result
}

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
