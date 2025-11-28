package icbc_api_sdk_go

type ShowPayUIRequest struct {
	Attach        string `json:"attach,omitempty"`
	Body          string `json:"body,omitempty"`
	ExpireTime    string `json:"expireTime,omitempty"`
	IcbcAppid     string `json:"icbc_appid,omitempty"`
	MerAcct       string `json:"mer_acct,omitempty"`
	MerId         string `json:"mer_id,omitempty"`
	MerPrtclNo    string `json:"mer_prtcl_no,omitempty"`
	NotifyType    string `json:"notify_type,omitempty"`
	NotifyUrl     string `json:"notify_url,omitempty"`
	OpenId        string `json:"openId,omitempty"`
	OrderAmt      string `json:"order_amt,omitempty"`
	OrderApdInf   string `json:"order_apd_inf,omitempty"`
	OutTradeNo    string `json:"out_trade_no,omitempty"`
	PayLimit      string `json:"pay_limit,omitempty"`
	ResultType    string `json:"result_type,omitempty"`
	ReturnUrl     string `json:"return_url,omitempty"`
	Saledepname   string `json:"saledepname,omitempty"`
	ShopAppid     string `json:"shop_appid,omitempty"`
	Subject       string `json:"subject,omitempty"`
	Detail        string `json:"detail,omitempty"`
	CustName      string `json:"cust_name,omitempty"`
	CustCertType  string `json:"cust_cert_type,omitempty"`
	CustCertNo    string `json:"cust_cert_no,omitempty"`
	GoodsTag      string `json:"goods_tag,omitempty"`
	StartDatetime string `json:"start_datetime,omitempty"`
}
