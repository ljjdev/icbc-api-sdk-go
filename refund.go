package icbc_api_sdk_go

type RefundRequest struct {
	OrderId        string `json:"order_id"`
	OuttrxSerialNo string `json:"outtrx_serial_no"`
	RetTotalAmt    string `json:"ret_total_amt"`
	TrnscCcy       string `json:"trnsc_ccy"`
	MerId          string `json:"mer_id"`
	IcbcAppid      string `json:"icbc_appid"`
	MerAcct        string `json:"mer_acct"`
	OutTradeNo     string `json:"out_trade_no"`
	OrderApdInf    string `json:"order_apd_inf"`
	MerPrtclNo     string `json:"mer_prtcl_no"`
	RefundSource   string `json:"refund_source"`
	AcqAddnData    string `json:"acq_addn_data"`
}
type RefundResp struct {
	ResponseBizContent struct {
		ReturnCode                  string `json:"return_code,omitempty"`
		ReturnMsg                   string `json:"return_msg,omitempty"`
		MsgId                       string `json:"msg_id,omitempty"`
		OutTradeNo                  string `json:"out_trade_no,omitempty"`
		OuttrxSerialNo              string `json:"outtrx_serial_no,omitempty"`
		OrderId                     string `json:"order_id,omitempty"`
		CardNo                      string `json:"card_no,omitempty"`
		RejectAmt                   string `json:"reject_amt,omitempty"`
		RealRejectAmt               string `json:"real_reject_amt,omitempty"`
		RejectPoint                 string `json:"reject_point,omitempty"`
		RejectEcoupon               string `json:"reject_ecoupon,omitempty"`
		RejectMerDiscAmt            string `json:"reject_mer_disc_amt,omitempty"`
		RejectBankDiscAmt           string `json:"reject_bank_disc_amt,omitempty"`
		PayType                     string `json:"pay_type,omitempty"`
		SettlementRefundAmt         string `json:"settlement_refund_amt,omitempty"`
		ThirdPartyCouponRefundAmt   string `json:"third_party_coupon_refund_amt,omitempty"`
		ThirdPartyDiscountRefundAmt string `json:"third_party_discount_refund_amt,omitempty"`
		RefundTime                  string `json:"refund_time,omitempty"`
		IntrxSerialNo               string `json:"intrx_serial_no,omitempty"`
		ThirdPartyReturnCode        string `json:"third_party_return_code,omitempty"`
		ThirdPartyReturnMsg         string `json:"third_party_return_msg,omitempty"`
		RejectUnionDiscountamt      string `json:"reject_union_discountamt,omitempty"`
		RejectUnionMchtdiscountamt  string `json:"reject_union_mchtdiscountamt,omitempty"`
		RefundDetail                string `json:"refund_detail,omitempty"`
		UnionActivityId             string `json:"union_activity_id,omitempty"`
		UnionActivityNm             string `json:"union_activity_nm,omitempty"`
		UnionAddnPrintInfo          string `json:"union_addn_print_info,omitempty"`
		UnionIssAddnData            string `json:"union_iss_addn_data,omitempty"`
	} `json:"response_biz_content,omitempty"`
	Sign string `json:"sign,omitempty"`
}
