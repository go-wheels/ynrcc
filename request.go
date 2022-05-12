package ynrcc

type request struct {
	Request any `json:"request"`
}

type RequestCommon struct {
	Service  string `json:"service"`
	TranCode string `json:"tranCode"`
	SeqNo    string `json:"seqNo"`
	TxnTime  string `json:"txnTime"`
	MerID    string `json:"merId"`
	TemID    string `json:"temId"`
}

type QueryTrxStateRequest struct {
	RequestCommon
	TradeNo string `json:"tradeNo"`
}

type MPCreateTradeRequest struct {
	RequestCommon
	TradeNo      string `json:"tradeNo"`
	TradeChannel string `json:"tradeChannel"`
	BusinessType string `json:"businessType"`
	TotalAmt     string `json:"totalAmt"`
	TotalNum     string `json:"totalNum"`
	OrderDesc    string `json:"orderDesc"`
	OnlineFlag   string `json:"onlineFlag"`
	EventFlag    string `json:"eventFlag"`
	Ccy          string `json:"ccy"`
	SubOpenID    string `json:"subOpenId"`
	SubAppID     string `json:"subAppId"`
	NotifyURL    string `json:"notifyUrl"`
}

type PayResultMerNotifyRequest struct {
	RequestCommon
	Code        string `json:"code"`
	Msg         string `json:"msg"`
	TradeNo     string `json:"tradeNo"`
	State       string `json:"state"`
	OrderID     string `json:"orderId"`
	TradeTime   string `json:"tradeTime"`
	TotalAmt    string `json:"totalAmt"`
	PaidAmt     string `json:"paidAmt"`
	DiscountAmt int    `json:"discountAmt"`
}
