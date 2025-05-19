package go_h2pay

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
)

// 充值回调
func (cli *Client) WithdrawCallback(req H2PayWithdrawBackReq, processor func(H2PayWithdrawBackReq) error) error {
	//验证签名
	key := cli.getWithdrawBackMD5(req)
	if key != req.Key {
		cli.logger.Warnf("H2PayBackService#depositBack#verify,req:%+v,key:%s", req, key)
		return errors.New("key is wrong")
	}

	//开始处理
	return processor(req)
}

func (cli *Client) getWithdrawBackMD5(req H2PayWithdrawBackReq) string {
	// Construct the encoded string in the exact same order as Java
	encodeStr := req.MerchantCode + req.TransactionID + req.MemberCode +
		req.Amount + req.CurrencyCode + req.Status + cli.AccessKey

	// Generate MD5 hash
	hash := md5.Sum([]byte(encodeStr))

	// Return hex-encoded string
	return hex.EncodeToString(hash[:])
}
