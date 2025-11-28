package icbc_api_sdk_go

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	URL "net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

const version = "v2_20190522"

// DefaultClient 默认客户端
type DefaultClient struct {
	APPID         string
	PrivateKey    string
	SignType      string
	IcbcPublicKey string
	HTTPClient    *http.Client // 允许自定义HTTP客户端
}

// UiIcbcClient 页面类客户端
type UiIcbcClient struct {
	DefaultClient
}

// buildPostForm UI类客户端 构建POST表单页面
// @param request 请求对象
// @return form 构建好的POST表单
func (c *UiIcbcClient) buildPostForm(request *ICBCRequest) (string, error) {
	//把 bizContent 转换成 Icbc map
	params, err := c.PrepareParams(request, "")
	if err != nil {
		return "", err
	}
	queryParams := c.buildUrlQueryParams(params)
	// 拼接到service url 作为查询参数
	buildGetUrl, err := BuildGetUrl(request.ServiceUrl, queryParams)
	if err != nil {
		return "", err
	}
	return buildForm(buildGetUrl, c.buildBodyParams(params)), nil
}

func (c *UiIcbcClient) buildUrlQueryParams(params *IcbcMap) *IcbcMap {
	m := NewIcbcMap()
	for k, v := range params.data {
		if slices.Contains(ApiParamNames, k) {
			m.Put(k, v)
		}
	}
	return m
}
func (c *UiIcbcClient) buildBodyParams(params *IcbcMap) *IcbcMap {
	m := NewIcbcMap()
	for k, v := range params.data {
		if !slices.Contains(ApiParamNames, k) && v != "" {
			m.Put(k, v)
		}
	}
	return m
}

// PrepareParams 准备请求参数
//
// 参数:
//   - request: 请求对象
//   - msgId: 消息ID
//
// 返回值:
//   - *IcbcMap: 准备好的参数
//   - error: 错误信息
func (c *DefaultClient) PrepareParams(request *ICBCRequest, msgId string) (*IcbcMap, error) {
	//msgId 不传 默认用 UUID V7 生成
	if msgId == "" {
		uuidV7, _ := uuid.NewV7()
		msgId = strings.ReplaceAll(uuidV7.String(), "-", "")
	}
	bizContentStr, err := c.buildBizContentStr(request)
	if err != nil {
		return nil, err
	}
	params := NewIcbcMap()
	params.Put("app_id", c.APPID)
	params.Put("sign_type", c.SignType)
	params.Put("msg_id", msgId)
	params.Put("biz_content", bizContentStr)
	params.Put("charset", "UTF-8")
	params.Put("format", "json")
	params.Put("timestamp", GetCurrentTime())
	u, err := URL.Parse(request.ServiceUrl)
	if err != nil {
		return nil, err
	}
	a := BuildOrderedSignStr(params, u.Path)
	// 签名
	signStr, err := SignWithSHA256RSA(a, c.PrivateKey)
	if err != nil {
		return nil, err
	}
	params.Put("sign", signStr)
	return params, nil
}

// buildBizContentStr 构建业务内容字符串
// @param request 请求对象
// @return bizContentStr 业务内容JSON字符串
// @return err 错误信息
func (c *DefaultClient) buildBizContentStr(request *ICBCRequest) (string, error) {
	if request.BizContent == nil {
		return "", nil
	}
	bizContentStr, marshalErr := json.Marshal(request.BizContent)
	if marshalErr != nil {
		return "", fmt.Errorf("failed to marshal biz content: %w", marshalErr)
	}
	return string(bizContentStr), nil
}

// execute 执行请求
// @param request 请求对象
// @param msgId 消息ID
// @param res 响应对象
// @return res 响应对象
// @return err 错误信息
func (c *DefaultClient) execute(request *ICBCRequest, msgId string, res any) (any, error) {
	params, err := c.PrepareParams(request, msgId)
	if err != nil {
		return "", err
	}
	formData := URL.Values{}
	for k, v := range params.data {
		formData.Set(k, v)
	}
	encodedData := formData.Encode()
	req, err := http.NewRequest("POST", request.ServiceUrl, strings.NewReader(encodedData))
	if err != nil {
		// 如果创建请求失败（例如URL格式错误），直接返回错误
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("APIGW-VERSION", version)

	// 使用自定义HTTP客户端或默认客户端
	httpClient := c.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: time.Second * 10,
		}
	}

	response, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	//判断 http 状态码是否正常
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}
	//body 读取  要记得关闭
	defer func() {
		if err := response.Body.Close(); err != nil {
			fmt.Printf("warning: failed to close response body: %v\n", err)
		}
	}()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	//bodyStr := string(body)
	//验证签名
	isValid := json.Valid(body)
	if !isValid {
		return "", fmt.Errorf("invalid response json body: %s", string(body))
	}
	var icbcResponse IcbcResponse
	err = json.Unmarshal(body, &icbcResponse)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal icbc response: %w", err)
	}
	rawBizContent := string(icbcResponse.ResponseBizContent)
	sign := icbcResponse.Sign
	pass, err := VerifySHA1RSA(rawBizContent, sign, c.IcbcPublicKey)
	if err != nil {
		return "", fmt.Errorf("failed to verify signature: %w", err)
	}
	if !pass {
		return "", fmt.Errorf("signature verification failed")
	}
	// 解析响应体
	err = json.Unmarshal(body, res)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response to target type: %w", err)
	}
	return res, nil
}
