package go_h2pay

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// 下单
func (cli *Client) Deposit(req H2PayDepositReq) (*H2PayDepositRsp, error) {

	key := cli.getDepositMD5(req)

	rawURL := cli.DepositURL

	// Build parameter list
	params := []string{
		fmt.Sprintf("Merchant=%s", req.Merchant),
		fmt.Sprintf("Currency=%s", req.Currency),
		fmt.Sprintf("Customer=%s", req.Customer),
		fmt.Sprintf("Reference=%s", req.Reference),
		fmt.Sprintf("Key=%s", key),
		fmt.Sprintf("Amount=%s", req.Amount),
		fmt.Sprintf("Datetime=%s", req.Datetime),
		fmt.Sprintf("FrontURI=%s", req.FrontURI),
		fmt.Sprintf("BackURI=%s", req.BackURI),
		fmt.Sprintf("Bank=%s", req.Bank),
		fmt.Sprintf("Language=%s", req.Language),
		fmt.Sprintf("ClientIP=%s", req.ClientIP),
	}

	// Join parameters with &
	paramStr := strings.Join(params, "&")

	// Log request
	cli.logger.Infof("H2PayService#deposit, url: %s, param: %s", rawURL, paramStr)

	// Make HTTP POST request
	resp, err := http.Post(rawURL,
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(paramStr))
	if err != nil {
		cli.logger.Errorf("请求失败:%s", err.Error())
	}
	defer resp.Body.Close()
	if err != nil {
		log.Printf("H2PayService#deposit error: %v", err)
		return nil, err
	}

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("H2PayService#deposit read response error: %v", err)
		return nil, err
	}

	// Log response
	responseStr := string(body)
	log.Printf("H2PayService#deposit#rsp: %s", responseStr)

	// Build response struct
	rsp := &H2PayDepositRsp{
		HTMLString: responseStr,
	}

	return rsp, nil
}

func (cli *Client) getDepositMD5(req H2PayDepositReq) string {
	// Construct the encoded string
	encodeStr := req.Merchant + req.Reference + req.Customer + req.Amount +
		req.Currency + req.DateTimeMd5 + cli.AccessKey + req.ClientIP

	// Log before MD5
	log.Printf("H2PayService#MD5#deposit#before, s: %s", encodeStr)

	// Generate MD5 hash
	hash := md5.Sum([]byte(encodeStr))
	result := hex.EncodeToString(hash[:])

	// Log after MD5
	log.Printf("H2PayService#MD5#deposit#end, s: %s", result)

	return result
}
