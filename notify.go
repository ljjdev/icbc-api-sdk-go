package icbc_api_sdk_go

type Notify struct {
	ReturnCode   string `json:"return_code"`
	ReturnMsg    string `json:"return_msg"`
	MsgId        string `json:"msg_id"`
	CardNo       string `json:"card_no"`
	MerId        string `json:"mer_id"`
	TotalAmt     string `json:"total_amt"`
	PointAmt     string `json:"point_amt"`
	EcouponAmt   string `json:"ecoupon_amt"`
	MerDiscAmt   string `json:"mer_disc_amt"`
	CouponAmt    string `json:"coupon_amt"`
	BankDiscAmt  string `json:"bank_disc_amt"`
	PaymentAmt   string `json:"payment_amt"`
	OutTradeNo   string `json:"out_trade_no"`
	OrderId      string `json:"order_id"`
	PayTime      string `json:"pay_time"`
	TotalDiscAmt string `json:"total_disc_amt"`
	Attach       string `json:"attach"`
	ThirdTradeNo string `json:"third_trade_no"`
	CardFlag     string `json:"card_flag"`
	DecrFlag     string `json:"decr_flag"`
	OpenId       string `json:"open_id"`
	PayType      string `json:"pay_type"`
	AccessType   string `json:"access_type"`
	CardKind     string `json:"card_kind"`
	BankType     string `json:"bank_type"`
}
