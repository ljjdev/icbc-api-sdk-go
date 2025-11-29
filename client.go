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

// BuildPostForm UI类客户端 构建POST表单页面
//
// 参数:
//   - request: 请求对象
//
// 返回值:
//   - string: 构建好的POST表单
//   - error: 错误信息
func (c *UiIcbcClient) BuildPostForm(request *ICBCRequest) (string, error) {
	// 验证请求对象
	if request == nil {
		return "", fmt.Errorf("request cannot be nil")
	}
	if request.ServiceUrl == "" {
		return "", fmt.Errorf("service url cannot be empty")
	}

	//把 bizContent 转换成 Icbc map
	params, err := c.PrepareParams(request, "")
	if err != nil {
		return "", fmt.Errorf("failed to prepare params: %w", err)
	}

	queryParams := c.BuildUrlQueryParams(params)
	// 拼接到service url 作为查询参数
	buildGetUrl, err := BuildGetUrl(request.ServiceUrl, queryParams)
	if err != nil {
		return "", fmt.Errorf("failed to build get url: %w", err)
	}

	bodyParams := c.BuildBodyParams(params)
	return BuildForm(buildGetUrl, bodyParams), nil
}

// BuildUrlQueryParams 构建URL查询参数
//
// 参数:
//   - params: 参数对象
//
// 返回值:
//   - *IcbcMap: URL查询参数
func (c *UiIcbcClient) BuildUrlQueryParams(params *IcbcMap) *IcbcMap {
	if params == nil {
		return NewIcbcMap()
	}

	m := NewIcbcMap()
	for k, v := range params.data {
		if slices.Contains(ApiParamNames, k) {
			m.Put(k, v)
		}
	}
	return m
}

// BuildBodyParams 构建请求体参数
//
// 参数:
//   - params: 参数对象
//
// 返回值:
//   - *IcbcMap: 请求体参数
func (c *UiIcbcClient) BuildBodyParams(params *IcbcMap) *IcbcMap {
	if params == nil {
		return NewIcbcMap()
	}

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
	// 验证请求对象
	if request == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if request.ServiceUrl == "" {
		return nil, fmt.Errorf("service url cannot be empty")
	}

	//msgId 不传 默认用 UUID V7 生成
	if msgId == "" {
		uuidV7, err := uuid.NewV7()
		if err != nil {
			return nil, fmt.Errorf("failed to generate uuid: %w", err)
		}
		msgId = strings.ReplaceAll(uuidV7.String(), "-", "")
	}

	// 构建业务内容
	bizContentStr, err := c.BuildBizContentStr(request)
	if err != nil {
		return nil, fmt.Errorf("failed to build biz content: %w", err)
	}

	// 构建请求参数
	params := NewIcbcMap()
	params.Put("app_id", c.APPID)
	params.Put("sign_type", c.SignType)
	params.Put("msg_id", msgId)
	params.Put("biz_content", bizContentStr)
	params.Put("charset", "UTF-8")
	params.Put("format", "json")
	params.Put("timestamp", GetCurrentTime())

	// 解析URL获取路径
	u, err := URL.Parse(request.ServiceUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse service url: %w", err)
	}

	// 构建签名字符串并签名
	a := BuildOrderedSignStr(params, u.Path)
	signStr, err := SignWithSHA256RSA(a, c.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign request: %w", err)
	}

	params.Put("sign", signStr)
	return params, nil
}

// BuildBizContentStr 构建业务内容字符串
//
// 参数:
//   - request: 请求对象
//
// 返回值:
//   - string: 业务内容JSON字符串
//   - error: 错误信息
func (c *DefaultClient) BuildBizContentStr(request *ICBCRequest) (string, error) {
	if request == nil {
		return "", fmt.Errorf("request cannot be nil")
	}

	if request.BizContent == nil {
		return "", nil
	}

	bizContentStr, marshalErr := json.Marshal(request.BizContent)
	if marshalErr != nil {
		return "", fmt.Errorf("failed to marshal biz content: %w", marshalErr)
	}

	return string(bizContentStr), nil
}

// Execute 执行请求
//
// 参数:
//   - request: 请求对象
//   - msgId: 消息ID
//   - res: 响应对象指针
//
// 返回值:
//   - any: 响应对象
//   - error: 错误信息
func (c *DefaultClient) Execute(request *ICBCRequest, msgId string, res any) (any, error) {
	// 验证请求对象
	if request == nil {
		return "", fmt.Errorf("request cannot be nil")
	}
	if request.ServiceUrl == "" {
		return "", fmt.Errorf("service url cannot be empty")
	}

	// 验证响应对象
	if res == nil {
		return "", fmt.Errorf("response object cannot be nil")
	}

	// 准备请求参数
	params, err := c.PrepareParams(request, msgId)
	if err != nil {
		return "", fmt.Errorf("failed to prepare params: %w", err)
	}

	// 构建表单数据
	formData := URL.Values{}
	for k, v := range params.data {
		formData.Set(k, v)
	}
	encodedData := formData.Encode()

	// 创建HTTP请求
	req, err := http.NewRequest("POST", request.ServiceUrl, strings.NewReader(encodedData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("APIGW-VERSION", version)

	// 使用自定义HTTP客户端或默认客户端
	httpClient := c.HTTPClient
	if httpClient == nil {
		// 默认HTTP客户端配置，优化超时和连接池
		httpClient = &http.Client{
			Timeout: time.Second * 30, // 延长超时时间
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 20,
				IdleConnTimeout:     90 * time.Second,
			},
		}
	}

	// 发送HTTP请求
	response, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}

	// 确保响应体被关闭
	defer func() {
		if closeErr := response.Body.Close(); closeErr != nil {
			fmt.Printf("warning: failed to close response body: %v\n", closeErr)
		}
	}()

	// 检查HTTP状态码
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d, response: %s", response.StatusCode, response.Status)
	}

	// 读取响应体
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// 验证响应是否为有效JSON
	isValid := json.Valid(body)
	if !isValid {
		return "", fmt.Errorf("invalid response json body: %s", string(body))
	}

	// 解析ICBC响应
	var icbcResponse IcbcResponse
	err = json.Unmarshal(body, &icbcResponse)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal icbc response: %w", err)
	}

	// 验证签名
	rawBizContent := string(icbcResponse.ResponseBizContent)
	sign := icbcResponse.Sign
	pass, err := VerifySHA1RSA(rawBizContent, sign, c.IcbcPublicKey)
	if err != nil {
		return "", fmt.Errorf("failed to verify signature: %w", err)
	}
	if !pass {
		return "", fmt.Errorf("signature verification failed")
	}

	// 解析响应体到目标对象
	err = json.Unmarshal(body, res)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response to target type: %w", err)
	}

	return res, nil
}
