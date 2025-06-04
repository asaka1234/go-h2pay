package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/spf13/cast"
	"log"
	"strings"
	"time"
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

			if k == "Datetime" {
				t, _ := time.Parse("2006-01-02 03:04:05PM", value)
				value = t.Format("20060102150405")
			}
			//fmt.Printf("%s=>%s\n", k, value)
			sb.WriteString(value)
		} else {
			//fmt.Printf("%s=>%s\n", k, key)
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

// MD5({MerchantCode}{TransactionId}{MemberCode}{Amount}{CurrencyCode}){TransactionDateTime}){ToBankAccountNumber}){SecurityCode}))
func WithdrawSign(params map[string]interface{}, key string) string {

	//参与签名的key
	signKeyList := []string{"MerchantCode", "TransactionID", "MemberCode", "Amount", "CurrencyCode", "TransactionDateTime", "toBankAccountNumber", "SecurityCode"}

	//拼凑字符串
	var sb strings.Builder
	for _, k := range signKeyList {
		if k != "SecurityCode" {
			value := cast.ToString(params[k])

			if k == "TransactionDateTime" {
				t, _ := time.Parse("2006-01-02 03:04:05PM", value)
				value = t.Format("20060102150405")
			}
			//fmt.Printf("%s=>%s\n", k, value)
			sb.WriteString(value)
		} else {
			//fmt.Printf("%s=>%s\n", k, key)
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
