package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/spf13/cast"
	"log"
	"strings"
)

// MD5({Merchant}{Reference}{Customer}{Amount}{Currency}{Datetime}{SecurityCode}{ClientIP})
func DepositSign(params map[string]interface{}, key string) string {

	//参与签名的key
	signKeyList := []string{"Merchant", "Reference", "Customer", "Amount", "Currency", "Datetime", "SecurityCode", "ClientIP"}

	//拼凑字符串
	var sb strings.Builder
	for _, k := range signKeyList {
		if k != "SecurityCode" {
			value := cast.ToString(params[k])
			sb.WriteString(value)
		} else {
			sb.WriteString(key)
		}
	}
	signStr := sb.String()

	// Log before MD5
	log.Printf("H2PayService#MD5#deposit#before, s: %s", signStr)

	// Generate MD5 hash
	hash := md5.Sum([]byte(signStr))
	result := hex.EncodeToString(hash[:])

	// Log after MD5
	log.Printf("H2PayService#MD5#deposit#end, s: %s", result)

	return result
}

//---------------

// MD5({Merchant}{Reference}{Customer}{Amount}{Currency}{Status}{SecurityCode})
func DepositBackSign(params map[string]interface{}, key string) string {

	//参与签名的key
	signKeyList := []string{"Merchant", "Reference", "Customer", "Amount", "Currency", "Status", "SecurityCode"}

	//拼凑字符串
	var sb strings.Builder
	for _, k := range signKeyList {
		if k != "SecurityCode" {
			value := cast.ToString(params[k])
			sb.WriteString(value)
		} else {
			sb.WriteString(key)
		}
	}
	signStr := sb.String()

	// Log before MD5
	log.Printf("H2PayService#MD5#deposit#before, s: %s", signStr)

	// Generate MD5 hash
	hash := md5.Sum([]byte(signStr))
	result := hex.EncodeToString(hash[:])

	// Log after MD5
	log.Printf("H2PayService#MD5#deposit#end, s: %s", result)

	return result
}

func DepositBackVerify(params map[string]interface{}, signKey string) (bool, error) {
	// Check if signature exists in params
	signature, exists := params["Key"]
	if !exists {
		return false, nil
	}

	// Remove signature from params for verification
	delete(params, "Key")

	// Generate current signature
	currentSignature := DepositBackSign(params, signKey)

	// Compare signatures
	return signature.(string) == currentSignature, nil
}

// MD5({Merchant}{Reference}{Customer}{Amount}{Currency}{Datetime}{SecurityCode}{ClientIP})
func WithdrawSign(params map[string]interface{}, key string) string {

	//参与签名的key
	signKeyList := []string{"MerchantCode", "TransactionID", "MemberCode", "Amount", "CurrencyCode", "TransactionDatetime", "toBankAccountNumber", "SecurityCode"}

	//拼凑字符串
	var sb strings.Builder
	for _, k := range signKeyList {
		if k != "SecurityCode" {
			value := cast.ToString(params[k])
			sb.WriteString(value)
		} else {
			sb.WriteString(key)
		}
	}
	signStr := sb.String()

	// Log before MD5
	log.Printf("H2PayService#MD5#deposit#before, s: %s", signStr)

	// Generate MD5 hash
	hash := md5.Sum([]byte(signStr))
	result := hex.EncodeToString(hash[:])

	// Log after MD5
	log.Printf("H2PayService#MD5#deposit#end, s: %s", result)

	return result
}

//---------------

// MD5({MerchantCode}{TransactionID}{MemberCode}{Amount}{CurrencyCode}{Status}{SecurityCode}
func WithdrawBackSign(params map[string]interface{}, key string) string {

	//参与签名的key
	signKeyList := []string{"MerchantCode", "TransactionID", "MemberCode", "Amount", "CurrencyCode", "Status", "SecurityCode"}

	//拼凑字符串
	var sb strings.Builder
	for _, k := range signKeyList {
		if k != "SecurityCode" {
			value := cast.ToString(params[k])
			sb.WriteString(value)
		} else {
			sb.WriteString(key)
		}
	}
	signStr := sb.String()

	// Log before MD5
	log.Printf("H2PayService#MD5#deposit#before, s: %s", signStr)

	// Generate MD5 hash
	hash := md5.Sum([]byte(signStr))
	result := hex.EncodeToString(hash[:])

	// Log after MD5
	log.Printf("H2PayService#MD5#deposit#end, s: %s", result)

	return result
}

func WithdrawBackVerify(params map[string]interface{}, signKey string) (bool, error) {
	// Check if signature exists in params
	signature, exists := params["Key"]
	if !exists {
		return false, nil
	}

	// Remove signature from params for verification
	delete(params, "Key")

	// Generate current signature
	currentSignature := WithdrawBackSign(params, signKey)

	// Compare signatures
	return signature.(string) == currentSignature, nil
}
