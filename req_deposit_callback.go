package go_h2pay

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
)

// 充值回调
func (cli *Client) DepositCallback(req H2PayDepositBackReq, processor func(H2PayDepositBackReq) error) error {
	//验证签名
	key := cli.getDepositBackMD5(req)
	if key != req.Key {
		cli.logger.Warnf("H2PayBackService#depositBack#verify,req:%+v,key:%s", req, key)
		return errors.New("key is wrong")
	}

	//开始处理
	return processor(req)
}

func (cli *Client) getDepositBackMD5(req H2PayDepositBackReq) string {
	// Create the concatenated string in the exact same order as Java
	encodeStr := req.Merchant + req.Reference + req.Customer + req.Amount +
		req.Currency + req.Status + cli.AccessKey

	// Generate MD5 hash
	hash := md5.Sum([]byte(encodeStr))

	// Convert to hex string
	return hex.EncodeToString(hash[:])
}
