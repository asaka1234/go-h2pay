package go_h2pay

import (
	"fmt"
	"testing"
)

//--------------------------------------------

func TestWithdraw(t *testing.T) {

	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &H2PayInitParams{MERCHANT_ID, ACCESS_KEY, DEPOSIT_URL, DEPOSIT_CALLBACK_URL, DEPOSIT_FE_BACK_URL, WITHDRAW_URL, WITHDRAW_CALLBACK_URL})

	//发请求
	resp, err := cli.Withdraw(GenWithdrawRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenWithdrawRequestDemo() H2PayWithdrawReq {

	return H2PayWithdrawReq{
		CurrencyCode:        "MYR",
		MemberCode:          "220099",     //uid
		TransactionID:       "1609032335", //outNo
		Amount:              "1.00",
		BankCode:            "MBB", //必须要是存在的,不然报错
		ToBankAccountName:   "cy",
		ToBankAccountNumber: "72972972927",
		ClientIP:            "18.162.184.178",
	}
}
