package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
	// 1.利用rsa加解密
	// 生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	// 通过私钥生成公钥
	publicKey := privateKey.PublicKey

	modulesBytes := base64.StdEncoding.EncodeToString(privateKey.N.Bytes())
	privateExponentBytes := base64.StdEncoding.EncodeToString(privateKey.D.Bytes())
	fmt.Println("公钥值：", modulesBytes)
	fmt.Println("私钥值:", privateExponentBytes)
	fmt.Println("公钥幂值：", publicKey.E)

	// 利用公钥来进行加密
	bytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&publicKey,
		[]byte("hello emily"),
		nil,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("公钥加密后的文本：", bytes)

	// 利用私钥进行解密
	decryptBytes, err := privateKey.Decrypt(nil, bytes, &rsa.OAEPOptions{
		Hash: crypto.SHA256,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("私钥解密后的文档：", string(decryptBytes))

	// 2. 利用rsa签名
	msg := []byte("hello, Emily")
	msgHash := sha256.New()
	_, err = msgHash.Write(msg)
	if err != nil {
		panic(err)
	}
	msgHashSum := msgHash.Sum(nil)
	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		panic(err)
	}

	// 签名验证
	err = rsa.VerifyPSS(&publicKey, crypto.SHA256, msgHashSum, signature, nil)
	if err != nil {
		fmt.Println("签名验证失败：", err)
		return
	}

	fmt.Println("签名验证成功")
}
