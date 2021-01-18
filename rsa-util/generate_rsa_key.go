package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

//generateKeyPair: 生成key对
func generateKeyPair(bites int) (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bites)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return privateKey, &privateKey.PublicKey
}

func main() {
	privateKey, publicKey := generateKeyPair(2048)

	fmt.Printf("Private key: %v\n", privateKey)
	fmt.Printf("Public key: %v\n", publicKey)
}
