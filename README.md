# ICBC API SDK Go

## 项目介绍

ICBC API SDK Go 是工商银行API的Go语言实现，参照了官方icbc-api-sdk-cop的思路进行开发，提供了完整的工商银行API调用功能。

## 开发背景

该项目基于工商银行官方SDK的设计思路，使用Go语言实现，旨在为Go开发者提供便捷、高效的工商银行API调用方式。

## 功能特点

- 支持工商银行API的签名和验签
- 提供HTTP客户端配置优化
- 支持表单构建和提交
- 增强的错误处理和类型安全
- 良好的代码风格和文档

## 当前功能说明

- 当前功能仅根据实际使用过的功能进行开发
- 目前仅支持RSA和RSA2签名算法，不支持加密功能
- 还有很多功能可根据官方API文档进行扩展开发

## 可用功能列表

1. 聚合支付B2C线上消费查询
2. 线上POS退款
3. 线上POS退款查询
4. 线上POS聚合支付非埋名消费下单



## 安装

```bash
go get github.com/ljjdev/icbc-api-sdk-go
```

## 使用示例

### 初始化客户端

```go
client := &icbc_api_sdk_go.DefaultClient{
    APPID:         "your_app_id",
    PrivateKey:    "your_private_key",
    SignType:      "RSA2",
    IcbcPublicKey: "icbc_public_key",
}
```

### 执行请求

```go
request := &icbc_api_sdk_go.ICBCRequest{
    ServiceUrl: "https://api.icbc.com.cn/service",
    BizContent: your_biz_content,
    Method:     "POST",
}

var response YourResponseStruct
result, err := client.Execute(request, "", &response)
if err != nil {
    log.Fatalf("Failed to execute request: %v", err)
}
```

## 项目结构

- `client.go` - 客户端核心实现
- `webutil.go` - Web工具函数
- `icbcmap.go` - 工商银行Map实现
- `sign.go` - 签名和验签实现
- `base.go` - 基础结构体定义

## 开发规范

- 遵循Go语言编码规范
- 统一使用驼峰命名法
- 添加详细的注释
- 增强错误处理和类型安全

## 许可证

MIT
