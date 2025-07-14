package go_h2pay

import (
	"fmt"
	"testing"
)

//--------------------------------------------

func TestDepositCallback(t *testing.T) {

	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &H2PayInitParams{MERCHANT_ID, ACCESS_KEY, DEPOSIT_URL, DEPOSIT_CALLBACK_URL, DEPOSIT_FE_BACK_URL, WITHDRAW_URL, WITHDRAW_CALLBACK_URL})

	//发请求
	err := cli.DepositCallback(GenDepositCallbackRequestDemo(), AAAaa)
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
}

func AAAaa(H2PayDepositBackReq) error {
	return nil
}

func GenDepositCallbackRequestDemo() H2PayDepositBackReq {

	//{"Merchant":"T0326","Reference":"202507141125100167","Currency":"VND","Amount":"574530.00","Language":"en-us","Customer":"820003522","Datetime":"2025-07-14 04:25:10PM","StatementDate":"2025-07-14 08:27:17AM","Note":"",
	//	"Key":"6720EADCC84AF6346AB195512D53E623","Status":"000","Deposit":"","ID":"046633348","ErrorCode":""}

	return H2PayDepositBackReq{
		Currency:      "VND",
		Customer:      "820003522",          //uid
		Reference:     "202507141125100167", //outNo
		Amount:        "574530.00",
		Merchant:      "T0326", //必须要是存在的,不然报错
		Language:      "en-us",
		Datetime:      "2025-07-14 04:25:10PM",
		StatementDate: "2025-07-14 08:27:17AM",
		Status:        "000",
		ID:            "046633348",
		Key:           "6720EADCC84AF6346AB195512D53E623",
	}
}
