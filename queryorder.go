package icbc_api_sdk_go

type OrderQueryRequest struct {
	MerId      string `json:"mer_id,omitempty"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
	OrderId    string `json:"order_id,omitempty"`
	DealFlag   string `json:"deal_flag,omitempty"`
	IcbcAppid  string `json:"icbc_appid,omitempty"`
	MerPrtclNo string `json:"mer_prtcl_no,omitempty"`
}
type OrderQueryResp struct {
	ResponseBizContent struct {
		ReturnCode            string `json:"return_code"`
		ReturnMsg             string `json:"return_msg"`
		MsgId                 string `json:"msg_id"`
		PayStatus             string `json:"pay_status"`
		CardNo                string `json:"card_no"`
		MerId                 string `json:"mer_id"`
		TotalAmt              string `json:"total_amt"`
		PointAmt              string `json:"point_amt"`
		EcouponAmt            string `json:"ecoupon_amt"`
		MerDiscAmt            string `json:"mer_disc_amt"`
		CouponAmt             string `json:"coupon_amt"`
		BankDiscAmt           string `json:"bank_disc_amt"`
		PaymentAmt            string `json:"payment_amt"`
		OutTradeNo            string `json:"out_trade_no"`
		OrderId               string `json:"order_id"`
		PayTime               string `json:"pay_time"`
		TotalDiscAmt          string `json:"total_disc_amt"`
		Attach                string `json:"attach"`
		ThirdTradeNo          string `json:"third_trade_no"`
		CardFlag              string `json:"card_flag"`
		DecrFlag              string `json:"decr_flag"`
		OpenId                string `json:"open_id"`
		PayType               string `json:"pay_type"`
		AccessType            string `json:"access_type"`
		CardKind              string `json:"card_kind"`
		ThirdPartyReturnCode  string `json:"third_party_return_code"`
		ThirdPartyReturnMsg   string `json:"third_party_return_msg"`
		ThirdPartyCouponAmt   string `json:"third_party_coupon_amt"`
		ThirdPartyDiscountAmt string `json:"third_party_discount_amt"`
		UnionDiscountAmt      string `json:"union_discount_amt"`
		UnionMchtDiscountAmt  string `json:"union_mcht_discount_amt"`
		PromotionDetail       string `json:"promotion_detail"`
		UnionActivityId       string `json:"union_activity_id"`
		UnionActivityNm       string `json:"union_activity_nm"`
		UnionAddnPrintInfo    string `json:"union_addn_print_info"`
		UnionIssAddnData      string `json:"union_iss_addn_data"`
		OrderStatus           string `json:"order_status"`
		BankType              string `json:"bank_type"`
		PayGType              string `json:"pay_g_type"`
		TrxSerno              string `json:"trx_serno"`
		CardTissue            string `json:"card_tissue"`
	} `json:"response_biz_content"`
	Sign string `json:"sign"`
}
