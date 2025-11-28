package icbc_api_sdk_go

import (
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

func buildForm(baseUrl string, params *IcbcMap) string {
	var sb strings.Builder
	sb.WriteString("<form name=\"auto_submit_form\" method=\"post\" action=\"")
	sb.WriteString(baseUrl)
	sb.WriteString("\">")
	sb.WriteString(buildHiddenFields(params))
	sb.WriteString("<input type=\"submit\" value=\"立刻提交\" style=\"display:none\" >\n")
	sb.WriteString("</form>\n")
	sb.WriteString("<script>document.forms[0].submit();</script>")
	return sb.String()
}

func buildHiddenFields(params *IcbcMap) string {
	var sb strings.Builder
	for key, value := range params.data {
		if value != "" {
			sb.WriteString(buildHiddenFieldsWithKV(key, value))
		}
	}
	return sb.String()
}

// buildHiddenFieldsWithKV 构建隐藏域字符串
// @param key 键
// @param value 值
// @return hiddenField 隐藏域字符串
func buildHiddenFieldsWithKV(key, value string) string {
	var sb strings.Builder
	sb.WriteString("<input type=\"hidden\" name=\"")
	sb.WriteString(key)
	sb.WriteString("\" value=\"")
	sb.WriteString(strings.ReplaceAll(value, "\"", "&quot;"))
	sb.WriteString("\" >\n")
	return sb.String()
}

// BuildGetUrl 构建GET请求URL,添加到URL查询参数
// @param strUrl 请求URL
// @param icbcMap 参数对象
// @return getUrl GET请求URL
func BuildGetUrl(strUrl string, icbcMap *IcbcMap) (string, error) {
	u, err := URL.Parse(strUrl)
	if err != nil {
		return "", err
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

// BuildOrderedSignStr  构建有序参数签名字符串
// @param params 参数对象
// @param path 请求路径
// @return signStr 有序参数签名字符串
func BuildOrderedSignStr(params *IcbcMap, path string) string {
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
