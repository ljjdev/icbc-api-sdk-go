package icbc_api_sdk_go

import (
	"fmt"
	"strconv"
	"time"
)

type IcbcMap struct {
	data map[string]string
}

func NewIcbcMap() *IcbcMap {
	return &IcbcMap{
		data: make(map[string]string),
	}
}

// Put 添加键值对到IcbcMap
// @param key 键
// @param value 值
func (m *IcbcMap) Put(key string, value interface{}) {
	if key == "" {
		return
	}
	var strValue string
	switch v := value.(type) {
	case string:
		strValue = v
	case int:
		strValue = strconv.Itoa(v)
	case int64:
		strValue = strconv.FormatInt(v, 10)
	case float64:
		strValue = strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		strValue = strconv.FormatBool(v)
	case time.Time:
		strValue = FormatTime(v)
	default:
		strValue = fmt.Sprintf("%v", value)
	}
	if strValue != "" {
		m.data[key] = strValue
	}
}

// PutString 添加字符串键值对
// @param key 键
// @param value 字符串值
func (m *IcbcMap) PutString(key, value string) {
	if key != "" && value != "" {
		m.data[key] = value
	}
}

// PutInt 添加整数键值对
// @param key 键
// @param value 整数值
func (m *IcbcMap) PutInt(key string, value int) {
	if key != "" {
		m.data[key] = strconv.Itoa(value)
	}
}

// PutInt64 添加长整数键值对
// @param key 键
// @param value 长整数值
func (m *IcbcMap) PutInt64(key string, value int64) {
	if key != "" {
		m.data[key] = strconv.FormatInt(value, 10)
	}
}

// PutFloat64 添加浮点数键值对
// @param key 键
// @param value 浮点数值
func (m *IcbcMap) PutFloat64(key string, value float64) {
	if key != "" {
		m.data[key] = strconv.FormatFloat(value, 'f', -1, 64)
	}
}

// PutBool 添加布尔值键值对
// @param key 键
// @param value 布尔值
func (m *IcbcMap) PutBool(key string, value bool) {
	if key != "" {
		m.data[key] = strconv.FormatBool(value)
	}
}

// PutTime 添加时间键值对
// @param key 键
// @param value 时间值
func (m *IcbcMap) PutTime(key string, value time.Time) {
	if key != "" {
		m.data[key] = FormatTime(value)
	}
}

// Get 获取字符串值
// @param key 键
// @return string 值
func (m *IcbcMap) Get(key string) string {
	return m.data[key]
}

// GetInt 获取整数值
// @param key 键
// @return int 值
// @return error 转换错误
func (m *IcbcMap) GetInt(key string) (int, error) {
	value, exists := m.data[key]
	if !exists {
		return 0, fmt.Errorf("key %s not found", key)
	}
	return strconv.Atoi(value)
}

// GetInt64 获取长整数值
// @param key 键
// @return int64 值
// @return error 转换错误
func (m *IcbcMap) GetInt64(key string) (int64, error) {
	value, exists := m.data[key]
	if !exists {
		return 0, fmt.Errorf("key %s not found", key)
	}
	return strconv.ParseInt(value, 10, 64)
}

// GetFloat64 获取浮点数值
// @param key 键
// @return float64 值
// @return error 转换错误
func (m *IcbcMap) GetFloat64(key string) (float64, error) {
	value, exists := m.data[key]
	if !exists {
		return 0, fmt.Errorf("key %s not found", key)
	}
	return strconv.ParseFloat(value, 64)
}

// GetBool 获取布尔值
// @param key 键
// @return bool 值
// @return error 转换错误
func (m *IcbcMap) GetBool(key string) (bool, error) {
	value, exists := m.data[key]
	if !exists {
		return false, fmt.Errorf("key %s not found", key)
	}
	return strconv.ParseBool(value)
}

// String 方法，方便打印整个 map 的内容
func (m *IcbcMap) String() string {
	return fmt.Sprintf("%v", m.data)
}
