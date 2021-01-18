package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

// readKeyFromFile：从文件中读取数据
func readKeyFromFile(filename string) []byte {
	key, _ := ioutil.ReadFile(filename)
	return key
}

// exportPEMStrToPrivKey：从pem文件导出privateKey
func exportPEMStrToPrivKey(priv []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(priv)
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return key
}

// exportPEMStrToPubKey：从pem文件导出publicKey
func exportPEMStrToPubKey(pub []byte) *rsa.PublicKey {
	block, _ := pem.Decode(pub)
	key, _ := x509.ParsePKCS1PublicKey(block.Bytes)
	return key
}

func main() {

	bytes := readKeyFromFile("privkey.pem")
	privKey := exportPEMStrToPrivKey(bytes)

	fmt.Printf("Private Key: %v\n", privKey)
}
