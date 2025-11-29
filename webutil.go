package icbc_api_sdk_go

import (
	"fmt"
	URL "net/url"
	"strings"
	"time"

	"github.com/igrmk/treemap/v2"
)

var ApiParamNames = []string{
	"app_id",
	"sign_type",
	"msg_id",
	"sign",
	"charset",
	"format",
	"encrypt_type",
	"timestamp",
}

// FormatTime 格式化时间为ICBC API所需的格式
// @param t 时间对象
// @return string 格式化后的时间字符串
func FormatTime(t time.Time) string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return t.In(loc).Format("2006-01-02 15:04:05")
}

// GetCurrentTime 获取当前时间的格式化字符串
// @return string 格式化后的当前时间字符串
func GetCurrentTime() string {
	return FormatTime(time.Now())
}

// BuildForm 构建POST表单页面
//
// 参数:
//   - baseUrl: 表单提交的目标URL
//   - params: 表单参数
//
// 返回值:
//   - string: 构建好的HTML表单
func BuildForm(baseUrl string, params *IcbcMap) string {
	var sb strings.Builder
	sb.WriteString("<form name=\"auto_submit_form\" method=\"post\" action=\"")
	sb.WriteString(baseUrl)
	sb.WriteString("\">")
	sb.WriteString(BuildHiddenFields(params))
	sb.WriteString("<input type=\"submit\" value=\"立刻提交\" style=\"display:none\" >\n")
	sb.WriteString("</form>\n")
	sb.WriteString("<script>document.forms[0].submit();</script>")
	return sb.String()
}

// BuildHiddenFields 构建隐藏域字段
//
// 参数:
//   - params: 表单参数
//
// 返回值:
//   - string: 构建好的隐藏域字段
func BuildHiddenFields(params *IcbcMap) string {
	if params == nil {
		return ""
	}

	var sb strings.Builder
	for key, value := range params.data {
		if value != "" {
			sb.WriteString(BuildHiddenFieldsWithKV(key, value))
		}
	}
	return sb.String()
}

// BuildHiddenFieldsWithKV 构建隐藏域字符串
//
// 参数:
//   - key: 字段名
//   - value: 字段值
//
// 返回值:
//   - string: 构建好的隐藏域字符串
func BuildHiddenFieldsWithKV(key, value string) string {
	var sb strings.Builder
	sb.WriteString("<input type=\"hidden\" name=\"")
	sb.WriteString(key)
	sb.WriteString("\" value=\"")
	sb.WriteString(strings.ReplaceAll(value, "\"", "&quot;"))
	sb.WriteString("\" >\n")
	return sb.String()
}

// BuildGetUrl 构建GET请求URL,添加到URL查询参数
//
// 参数:
//   - strUrl: 请求URL
//   - icbcMap: 参数对象
//
// 返回值:
//   - string: 构建好的GET请求URL
//   - error: 错误信息
func BuildGetUrl(strUrl string, icbcMap *IcbcMap) (string, error) {
	if strUrl == "" {
		return "", fmt.Errorf("url cannot be empty")
	}
	if icbcMap == nil {
		return strUrl, nil
	}

	u, err := URL.Parse(strUrl)
	if err != nil {
		return "", fmt.Errorf("failed to parse url: %w", err)
	}
	q := u.Query()
	for k, v := range icbcMap.data {
		if q.Has(k) {
			q.Set(k, v)
		} else {
			q.Add(k, v)
		}
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}

// BuildOrderedSignStr 构建有序参数签名字符串
//
// 参数:
//   - params: 参数对象
//   - path: 请求路径
//
// 返回值:
//   - string: 有序参数签名字符串
func BuildOrderedSignStr(params *IcbcMap, path string) string {
	if params == nil {
		return path
	}

	tr := treemap.New[string, string]()
	for k, v := range params.data {
		tr.Set(k, v)
	}

	// 拼接有序参数
	var queryParts []string
	for it := tr.Iterator(); it.Valid(); it.Next() {
		if it.Key() != "" && it.Value() != "" {
			queryParts = append(queryParts, it.Key()+"="+it.Value())
		}
	}

	var sb strings.Builder
	sb.WriteString(path)
	sb.WriteString("?")
	sb.WriteString(strings.Join(queryParts, "&"))
	return sb.String()
}
