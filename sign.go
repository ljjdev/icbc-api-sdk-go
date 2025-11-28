package icbc_api_sdk_go

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"strings"
)

const (
	PrivateKeyPrefix = "-----BEGIN PRIVATE KEY-----"
	PrivateKeySuffix = "-----END PRIVATE KEY-----"
	PublicKeyPrefix  = "-----BEGIN PUBLIC KEY-----"
	PublicKeySuffix  = "-----END PUBLIC KEY-----"
)

// SignWithSHA256RSA 使用 SHA-256 哈希算法和 RSA 私钥对数据进行签名
//
// 参数:
//   - data: 待签名的数据
//   - privateKey:  Base64 编码的 RSA 私钥字符串
//
// 返回值:
//   - string: 签名后的 Base64 编码字符串
//   - error: 签名过程中出现的错误
func SignWithSHA256RSA(data, privateKey string) (string, error) {
	privateKey = strings.TrimSpace(privateKey)
	//判断私钥是否包含开始结束标记
	if !strings.HasPrefix(privateKey, PrivateKeyPrefix) {
		privateKey = PrivateKeyPrefix + "\n" + privateKey

	}
	if !strings.HasSuffix(privateKey, PrivateKeySuffix) {
		privateKey = privateKey + "\n" + PrivateKeySuffix
	}
	//解析PEM私钥
	pemBlock, _ := pem.Decode([]byte(privateKey))
	if pemBlock == nil {
		return "", fmt.Errorf("failed to decode PEM block containing the private key")
	}
	//解析 PEM 格式的私钥块
	privateKeyInterface, err := x509.ParsePKCS8PrivateKey(pemBlock.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse PKCS8 private key: %w", err)
	}
	// 类型断言，确保是 RSA 私钥
	pk, ok := privateKeyInterface.(*rsa.PrivateKey)
	if !ok {
		return "", fmt.Errorf("private key is not RSA")
	}
	// 对数据进行 SHA-256 哈希
	digest := sha256.Sum256([]byte(data))
	// 使用 RSA-PKCS1v15 签名算法对哈希值进行签名
	signature, signErr := rsa.SignPKCS1v15(rand.Reader, pk, crypto.SHA256, digest[:])
	// 检查签名是否成功
	if signErr != nil {
		return "", fmt.Errorf("could not sign message:%w", signErr)
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

// VerifySHA1RSA 使用 SHA-1 哈希算法和 RSA 公钥验证数据的签名
// 参数:
//   - data: 待验证的数据 ,如 响应工行响应结果 response_biz_content的值
//
// {
// "response_biz_content":{
// "return_code":0,
// "return_msg":"success",
// "class_id":"your class id",
// "class_name:"your class name"
// },
// "sign":"abcd"
// }
// 那么待验签的内容就是
// {
// "return_code":0,
// "return_msg":"success",
// "class_id":"your class id",
// "class_name:"your class name"
// }
//   - signature: Base64 编码的签名字符串
//   - publicKey: Base64 编码的 RSA 公钥字符串
//
// 返回值:
//   - bool: 如果签名验证成功则返回 true，否则返回 false
func VerifySHA1RSA(data, signature, publicKey string) (bool, error) {
	publicKey = strings.TrimSpace(publicKey)
	if !strings.HasPrefix(publicKey, PublicKeyPrefix) {
		publicKey = PublicKeyPrefix + "\n" + publicKey
	}
	//判断公钥是否包含开始结束标记
	if !strings.HasSuffix(publicKey, PublicKeySuffix) {
		publicKey = publicKey + "\n" + PublicKeySuffix
	}
	//解析PEM公钥
	pemBlock, _ := pem.Decode([]byte(publicKey))
	if pemBlock == nil {
		return false, fmt.Errorf("failed to decode PEM block containing the private key")
	}
	//解析 PEM 格式的私钥块
	pkcs1PublicKey, err := x509.ParsePKIXPublicKey(pemBlock.Bytes)
	if err != nil {
		return false, fmt.Errorf("failed to parse PKCS1 public key: %w", err)
	}
	pk, ok := pkcs1PublicKey.(*rsa.PublicKey)
	if !ok {
		return false, fmt.Errorf("public key is not RSA")
	}
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, fmt.Errorf("failed to decode signature: %w", err)
	}
	// 对数据进行 SHA1 哈希
	digest := sha1.Sum([]byte(data))
	err = rsa.VerifyPKCS1v15(pk, crypto.SHA1, digest[:], signatureBytes)
	if err != nil {
		return false, fmt.Errorf("could not verify signature:%w", err)
	}
	return true, nil
}
