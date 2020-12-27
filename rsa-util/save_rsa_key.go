package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

//generateKeyPair: 生成key对
func GenerateKeyPair(bites int) (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bites)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return privateKey, &privateKey.PublicKey
}

// exportPubKeyAsPEMStr：以pem格式导出public key
func exportPubKeyAsPEMStr(pubkey *rsa.PublicKey) string {
	pubKeyPem := string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(pubkey),
		}))
	return pubKeyPem
}

// exportPrivKeyAsPemStr：以pem格式导出private key
func exportPrivKeyAsPemStr(privKey *rsa.PrivateKey) string {
	return string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privKey),
		}))
}

func saveKeyToFile(keyPem, filename string) {
	pemBytes := []byte(keyPem)
	ioutil.WriteFile(filename, pemBytes, 0400)
}

func main() {
	privateKey, publicKey := GenerateKeyPair(2048)

	fmt.Printf("Private Key: %v\n", privateKey)
	fmt.Printf("Public Key: %v\n", publicKey)

	privKeyAsPemStr := exportPrivKeyAsPemStr(privateKey)
	pubKeyAsPEMStr := exportPubKeyAsPEMStr(publicKey)

	fmt.Println(privKeyAsPemStr)
	fmt.Println(pubKeyAsPEMStr)

	saveKeyToFile(privKeyAsPemStr, "privKey.pem")
	saveKeyToFile(pubKeyAsPEMStr, "pubkey.pem")
}
