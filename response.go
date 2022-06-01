package ynrcc

import "fmt"

const (
	ReturnCodeSuccess     = "000000"
	ReturnCodeInvalidSign = "100036"
)

type response struct {
	Response any `json:"response"`
}

type ResponseCommon struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"subCode"`
	SubMsg  string `json:"subMsg"`
}

func (r ResponseCommon) Err() error {
	if r.Code != ReturnCodeSuccess {
		return fmt.Errorf("%s (code: %s)", r.Msg, r.Code)
	}
	if r.SubCode != ReturnCodeSuccess {
		return fmt.Errorf("%s (subCode: %s)", r.SubMsg, r.SubCode)
	}
	return nil
}

type CloseOrderResponse struct {
	ResponseCommon
	State string `json:"state"`
}

type QueryTrxStateResponse struct {
	ResponseCommon
	TradeNo      string `json:"tradeNo"`
	State        string `json:"state"`
	OrderID      string `json:"orderId"`
	BuyerAnonyID string `json:"buyerAnonyId"`
	TotalAmt     string `json:"totalAmt"`
	PaidAmt      string `json:"paidAmt"`
	DiscountAmt  string `json:"discountAmt"`
	PayTime      string `json:"payTime"`
}

type MPCreateTradeResponse struct {
	ResponseCommon
	TradeNo   string `json:"tradeNo"`
	State     string `json:"state"`
	OrderID   string `json:"orderId"`
	TradeTime string `json:"tradeTime"`
	WxPayData struct {
		TimeStamp string `json:"timeStamp"`
		Package   string `json:"package"`
		PaySign   string `json:"paySign"`
		AppID     string `json:"appId"`
		SignType  string `json:"signType"`
		NonceStr  string `json:"nonceStr"`
	} `json:"wxPayData"`
}

type PayResultMerNotifyResponse struct {
	ResponseCommon
}
