package go_h2pay

import (
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

//--------------------------------------------

func TestDeposit(t *testing.T) {

	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &H2PayInitParams{MERCHANT_ID, ACCESS_KEY, DEPOSIT_URL, DEPOSIT_CALLBACK_URL, DEPOSIT_FE_BACK_URL, WITHDRAW_URL, WITHDRAW_CALLBACK_URL})

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
		Currency:  "MYR",
		Customer:  "220099",      //uid
		Reference: "16090323351", //outNo
		Amount:    "12.0",
		Bank:      "MBB", //必须要是存在的,不然报错
		Language:  "en-us",
		ClientIP:  "128.199.171.73",
	}
}
