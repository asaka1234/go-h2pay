package go_h2pay

import (
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

//--------------------------------------------

func TestDeposit(t *testing.T) {

	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, MERCHANT_ID, ACCESS_KEY, DEPOSIT_URL, WITHDRAW_URL, DEPOSIT_CALLBACK_URL, DEPOSIT_FE_BACK_URL, WITHDRAW_CALLBACK_URL)

	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenDepositRequestDemo() H2PayDepositReq {

	return H2PayDepositReq{
		//Merchant:  MERCHANT_ID,
		Currency:  "MYR",
		Customer:  "220099",     //uid
		Reference: "1609032335", //outNo
		//Key         string `json:"key"`
		Amount: "1.00",
		//FrontURI:    "https://usercenter.cptinternational.com/",
		//BackURI:     "https://usercenter.cptinternational.com/",
		Bank:     "MCC", //必须要是存在的,不然报错
		Language: "en-us",
		ClientIP: "128.199.171.73",
	}
}
