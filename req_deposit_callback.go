package go_h2pay

import (
	"encoding/json"
	"errors"
	"github.com/asaka1234/go-h2pay/utils"
	"github.com/mitchellh/mapstructure"
	"log"
)

// 充值回调
func (cli *Client) DepositCallback(req H2PayDepositBackReq, processor func(H2PayDepositBackReq) error) error {
	//验证签名
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	// Verify signature
	flag, err := utils.DepositBackVerify(params, cli.AccessKey)
	if err != nil {
		log.Printf("Signature verification error: %v", err)
		return err
	}
	if !flag {
		//签名校验失败
		reqJson, _ := json.Marshal(req)
		log.Printf("H2Pay back verify fail, req: %s", string(reqJson))
		return errors.New("sign verify error")
	}

	//开始处理
	return processor(req)
}
