package icbc_api_sdk_go

type QueryRefundRequest struct {
	MerId          string `json:"mer_id,omitempty"`
	OutTradeNo     string `json:"out_trade_no,omitempty"`
	OrderId        string `json:"order_id,omitempty"`
	OuttrxSerialNo string `json:"outtrx_serial_no,omitempty"`
	MerPrtclNo     string `json:"mer_prtcl_no,omitempty"`
}
type QueryRefundResponse struct {
	ReturnCode                  string `json:"return_code"`
	ReturnMsg                   string `json:"return_msg"`
	PayStatus                   string `json:"pay_status"`
	MsgId                       string `json:"msg_id"`
	OutTradeNo                  string `json:"out_trade_no"`
	OrderId                     string `json:"order_id"`
	OuttrxSerialNo              string `json:"outtrx_serial_no"`
	RealRejectAmt               string `json:"real_reject_amt"`
	RejectAmt                   string `json:"reject_amt"`
	RejectPoint                 string `json:"reject_point"`
	RejectEcoupon               string `json:"reject_ecoupon"`
	CardNo                      string `json:"card_no"`
	RejectMerDiscAmt            string `json:"reject_mer_disc_amt"`
	RejectBankDiscAmt           string `json:"reject_bank_disc_amt"`
	PayType                     string `json:"pay_type"`
	IntrxSerialNo               string `json:"intrx_serial_no"`
	RefundTime                  string `json:"refund_time"`
	RejectUnionDiscountamt      string `json:"reject_union_discountamt"`
	RejectUnionMchtdiscountamt  string `json:"reject_union_mchtdiscountamt"`
	RefundDetail                string `json:"refund_detail"`
	SettlementRefundFee         string `json:"settlement_refund_fee"`
	ThirdPartyDiscountRefundAmt string `json:"third_party_discount_refund_amt"`
	ThirdPartyCouponRefundAmt   string `json:"third_party_coupon_refund_amt"`
	UnionActivityId             string `json:"union_activity_id"`
	UnionActivityNm             string `json:"union_activity_nm"`
	UnionAddnPrintInfo          string `json:"union_addn_print_info"`
	UnionIssAddnData            string `json:"union_iss_addn_data"`
}
