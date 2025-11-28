package icbc_api_sdk_go

import "encoding/json"

type ICBCRequest struct {
	ServiceUrl  string
	BizContent  interface{}
	ExtraParams map[string]string
	Method      string
}

type IcbcResponse struct {
	ResponseBizContent json.RawMessage `json:"response_biz_content"` // 使用 json.RawMessage 来保留原始的JSON字符串
	Sign               string          `json:"sign"`
}
